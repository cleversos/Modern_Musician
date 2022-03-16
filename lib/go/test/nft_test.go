package test

import (
	"testing"

	"github.com/onflow/cadence"
	jsoncdc "github.com/onflow/cadence/encoding/json"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-nft/lib/go/contracts"
	"github.com/onflow/flow-nft/lib/go/templates"
)

func TestNFTDeployment(t *testing.T) {
	b := newBlockchain()

	nftAddress := deploy(t, b, "NonFungibleToken", contracts.NonFungibleToken())
	metadataAddress := deploy(t, b, "MetadataViews", contracts.MetadataViews())

	_ = deploy(
		t, b, 
		"ModernMusicianNFT", 
		contracts.ModernMusicianNFT(nftAddress, metadataAddress),
	)
}

func TestCreateNFT(t *testing.T) {
	b := newBlockchain()

	nftAddress := deploy(t, b, "NonFungibleToken", contracts.NonFungibleToken())
	metadataAddress := deploy(t, b, "MetadataViews", contracts.MetadataViews())

	accountKeys := test.AccountKeyGenerator()

	ModernMusicianNFTAccountKey, ModernMusicianNFTSigner := accountKeys.NewWithSigner()
	ModernMusicianNFTAddress := deploy(
		t, b, 
		"ModernMusicianNFT", 
		contracts.ModernMusicianNFT(nftAddress, metadataAddress), 
		ModernMusicianNFTAccountKey,
	)

	script := templates.GenerateGetTotalSupplyScript(nftAddress, ModernMusicianNFTAddress)
	supply := executeScriptAndCheck(t, b, script, nil)
	assert.Equal(t, cadence.NewUInt64(0), supply)

	script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
	length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(ModernMusicianNFTAddress))})
	assert.Equal(t, cadence.NewInt(0), length)

	t.Run("Should be able to mint a token", func(t *testing.T) {

		script := templates.GenerateMintNFTScript(nftAddress, ModernMusicianNFTAddress)

		tx := createTxWithTemplateAndAuthorizer(b, script, ModernMusicianNFTAddress)
		
		tx.AddArgument(cadence.NewAddress(ModernMusicianNFTAddress))
		tx.AddArgument(cadence.String("Example NFT 0"))
		tx.AddArgument(cadence.String("This is an example NFT"))
		tx.AddArgument(cadence.String("example.jpeg"))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				ModernMusicianNFTAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				ModernMusicianNFTSigner,
			},
			false,
		)

		script = templates.GenerateBorrowNFTScript(nftAddress, ModernMusicianNFTAddress)
		executeScriptAndCheck(
			t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(ModernMusicianNFTAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, ModernMusicianNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(1), supply)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(ModernMusicianNFTAddress))})
		assert.Equal(t, cadence.NewInt(1), length)
	})

	t.Run("Shouldn't be able to borrow a reference to an NFT that doesn't exist", func(t *testing.T) {
		script = templates.GenerateBorrowNFTScript(nftAddress, ModernMusicianNFTAddress)

		result, err := b.ExecuteScript(
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(ModernMusicianNFTAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(5)),
			},
		)
		require.NoError(t, err)

		assert.True(t, result.Reverted())
	})
}

func TestTransferNFT(t *testing.T) {
	b := newBlockchain()

	nftAddress := deploy(t, b, "NonFungibleToken", contracts.NonFungibleToken())
	metadataAddress := deploy(t, b, "MetadataViews", contracts.MetadataViews())

	accountKeys := test.AccountKeyGenerator()

	ModernMusicianNFTAccountKey, ModernMusicianNFTSigner := accountKeys.NewWithSigner()
	ModernMusicianNFTAddress := deploy(
		t, b, 
		"ModernMusicianNFT", 
		contracts.ModernMusicianNFT(nftAddress, metadataAddress), 
		ModernMusicianNFTAccountKey,
	)

	joshAccountKey, joshSigner := accountKeys.NewWithSigner()
	joshAddress, err := b.CreateAccount([]*flow.AccountKey{joshAccountKey}, nil)
	require.NoError(t, err)

	script := templates.GenerateMintNFTScript(nftAddress, ModernMusicianNFTAddress)

	tx := flow.NewTransaction().
		SetScript(script).
		SetGasLimit(100).
		SetProposalKey(
			b.ServiceKey().Address,
			b.ServiceKey().Index,
			b.ServiceKey().SequenceNumber,
		).
		SetPayer(b.ServiceKey().Address).
		AddAuthorizer(ModernMusicianNFTAddress)

	tx.AddArgument(cadence.NewAddress(ModernMusicianNFTAddress))
	tx.AddArgument(cadence.String("Example NFT 0"))
	tx.AddArgument(cadence.String("This is an example NFT"))
	tx.AddArgument(cadence.String("example.jpeg"))

	signAndSubmit(
		t, b, tx,
		[]flow.Address{
			b.ServiceKey().Address,
			ModernMusicianNFTAddress,
		},
		[]crypto.Signer{
			b.ServiceKey().Signer(),
			ModernMusicianNFTSigner,
		},
		false,
	)

	// create a new Collection
	t.Run("Should be able to create a new empty NFT Collection", func(t *testing.T) {

		script := templates.GenerateSetupAccountScript(nftAddress, ModernMusicianNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				joshSigner,
			},
			false,
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)
	})

	t.Run("Shouldn't be able to withdraw an NFT that doesn't exist in a collection", func(t *testing.T) {

		script := templates.GenerateTransferNFTScript(nftAddress, ModernMusicianNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, ModernMusicianNFTAddress)

		tx.AddArgument(cadence.NewAddress(joshAddress))
		tx.AddArgument(cadence.NewUInt64(3))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				ModernMusicianNFTAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				ModernMusicianNFTSigner,
			},
			true,
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(ModernMusicianNFTAddress))})
		assert.Equal(t, cadence.NewInt(1), length)
	})

	// transfer an NFT
	t.Run("Should be able to withdraw an NFT and deposit to another accounts collection", func(t *testing.T) {
		script := templates.GenerateTransferNFTScript(nftAddress, ModernMusicianNFTAddress)
		tx := createTxWithTemplateAndAuthorizer(b, script, ModernMusicianNFTAddress)

		tx.AddArgument(cadence.NewAddress(joshAddress))
		tx.AddArgument(cadence.NewUInt64(0))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				ModernMusicianNFTAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				ModernMusicianNFTSigner,
			},
			false,
		)

		script = templates.GenerateBorrowNFTScript(nftAddress, ModernMusicianNFTAddress)
		executeScriptAndCheck(
			t, b,
			script,
			[][]byte{
				jsoncdc.MustEncode(cadence.NewAddress(joshAddress)),
				jsoncdc.MustEncode(cadence.NewUInt64(0)),
			},
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(1), length)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(ModernMusicianNFTAddress))})
		assert.Equal(t, cadence.NewInt(0), length)
	})

	// transfer an NFT
	t.Run("Should be able to withdraw an NFT and destroy it, not reducing the supply", func(t *testing.T) {

		script := templates.GenerateDestroyNFTScript(nftAddress, ModernMusicianNFTAddress)

		tx := createTxWithTemplateAndAuthorizer(b, script, joshAddress)

		tx.AddArgument(cadence.NewUInt64(0))

		signAndSubmit(
			t, b, tx,
			[]flow.Address{
				b.ServiceKey().Address,
				joshAddress,
			},
			[]crypto.Signer{
				b.ServiceKey().Signer(),
				joshSigner,
			},
			false,
		)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length := executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(joshAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetCollectionLengthScript(nftAddress, ModernMusicianNFTAddress)
		length = executeScriptAndCheck(t, b, script, [][]byte{jsoncdc.MustEncode(cadence.NewAddress(ModernMusicianNFTAddress))})
		assert.Equal(t, cadence.NewInt(0), length)

		script = templates.GenerateGetTotalSupplyScript(nftAddress, ModernMusicianNFTAddress)
		supply := executeScriptAndCheck(t, b, script, nil)
		assert.Equal(t, cadence.NewUInt64(1), supply)
	})
}
