import NonFungibleToken from "../contracts/NonFungibleToken.cdc"
import ModernMusicianNFT from "../contracts/ModernMusicianNFT.cdc"

// This transaction is what an account would run
// to set itself up to receive NFTs

transaction {

    prepare(signer: AuthAccount) {
        // Return early if the account already has a collection
        if signer.borrow<&ModernMusicianNFT.Collection>(from: ModernMusicianNFT.CollectionStoragePath) != nil {
            return
        }

        // Create a new empty collection
        let collection <- ModernMusicianNFT.createEmptyCollection()

        // save it to the account
        signer.save(<-collection, to: ModernMusicianNFT.CollectionStoragePath)

        // create a public capability for the collection
        signer.link<&{NonFungibleToken.CollectionPublic, ModernMusicianNFT.ModernMusicianNFTCollectionPublic}>(
            ModernMusicianNFT.CollectionPublicPath,
            target: ModernMusicianNFT.CollectionStoragePath
        )
    }
}
