// This is an example implementation of a Flow Non-Fungible Token
// It is not part of the official standard but it assumed to be
// very similar to how many NFTs would implement the core functionality.

import NonFungibleToken from 0x631e88ae7f1d7c20
import MetadataViews from 0x631e88ae7f1d7c20
import FungibleToken from 0x9a0766d93b6608b7
import FlowToken from 0x7e60df042a9c0868

pub contract ModernMusicianNFT: NonFungibleToken {

    pub struct UpcomingInfo {
        pub let name: String
        pub let description: String
        pub let type: UInt8
        pub let preOrderer: Address
        init(
            name: String,
            description: String,
            type: UInt8
        ) {
            self.name = name
            self.description = description
            self.type = type
            self.preOrderer = ""
        }
    }

    pub var waitingList: [UpcomingInfo]
    pub var totalSupply: UInt64

    pub event ContractInitialized()
    pub event Withdraw(id: UInt64, from: Address?)
    pub event Deposit(id: UInt64, to: Address?)

    pub let CollectionStoragePath: StoragePath
    pub let CollectionPublicPath: PublicPath
    pub let MinterStoragePath: StoragePath

    pub resource NFT: NonFungibleToken.INFT, MetadataViews.Resolver {
        pub let id: UInt64

        pub let name: String
        pub let description: String
        pub let thumbnail: String
        pub let preOrderer: Address

        init(
            id: UInt64,
            name: String,
            description: String,
            thumbnail: String,
            preOrderer: Address,
        ) {
            self.id = id
            self.name = name
            self.description = description
            self.thumbnail = thumbnail
            self.preOrderer = preOrderer
        }
    
        pub fun getViews(): [Type] {
            return [
                Type<MetadataViews.Display>()
            ]
        }

        pub fun resolveView(_ view: Type): AnyStruct? {
            switch view {
                case Type<MetadataViews.Display>():
                    return MetadataViews.Display(
                        name: self.name,
                        description: self.description,
                        thumbnail: MetadataViews.HTTPFile(
                            url: self.thumbnail
                        )
                    )
            }

            return nil
        }
    }

    pub resource interface ModernMusicianNFTCollectionPublic {
        pub fun deposit(token: @NonFungibleToken.NFT)
        pub fun getIDs(): [UInt64]
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT
        pub fun borrowModernMusicianNFT(id: UInt64): &ModernMusicianNFT.NFT? {
            post {
                (result == nil) || (result?.id == id):
                    "Cannot borrow ModernMusicianNFT reference: the ID of the returned reference is incorrect"
            }
        }
    }

    pub resource Collection: ModernMusicianNFTCollectionPublic, NonFungibleToken.Provider, NonFungibleToken.Receiver, NonFungibleToken.CollectionPublic, MetadataViews.ResolverCollection {
        // dictionary of NFT conforming tokens
        // NFT is a resource type with an `UInt64` ID field
        pub var ownedNFTs: @{UInt64: NonFungibleToken.NFT}

        init () {
            self.ownedNFTs <- {}
        }

        // buy bu an NFT from Contract Addcount
        pub fun buy(receiver_id: Address, nftID: UInt64, from: @FungibleToken.Vault) {
            let collectionRef = ModernMusicianNFT.account
                .getCapability(ModernMusicianNFT.CollectionPublicPath)
                .borrow<&{NonFungibleToken.CollectionPublic}>()
                ?? panic("Could not borrow capability from public collection")

            // Borrow a reference to a specific NFT in the collection
            let buyingNFT = collectionRef.borrowNFT(id: id) ?? panic("Could not find NFT in the marketplace")

            if(buyingNFT.preOrderer != "" && buyingNFT.preOrderer != receiver_id){
                panic("This NFT has already been sold")
            } elseif (buyingNFT.preOrderer != "" && buyingNFT.preOrderer != receiver_id){
                if(buyingNFT.type == 0 && from.balance < 450.0){
                    panic("You have to deposit more than 450.0 Flow")
                }
            } elseif (buyingNFT.preOrderer == ""){
                if(buyingNFT.type == 0 && from.balance < 500.0){
                    panic("You have to deposit more than 500.0 Flow")
                } elseif(buyingNFT.type == 1 && from.balance < 5.0){
                    panic("You have to deposit more than 5.0 Flow")
                }
            }

            let marketplaceCollection: &ModernMusicianNFT.Collection = ModernMusicianNFT.account.borrow<&ModernMusicianNFT.Collection>(from: NonFungibleToken.CollectionStoragePath)
                ?? panic("Could not borrow capability from storage collection")

            let token <-marketplaceCollection.withdraw(id: nftID) ?? panic("Failed to Withdraw NFT")

            // borrow a public reference to the receivers collection
            let depositRef = getAccount(receiver_id)
                .getCapability(ModernMusicianNFT.CollectionPublicPath)
                .borrow<&{NonFungibleToken.CollectionPublic}>()
                ?? panic("Could not borrow a reference to the receiver's collection")

            // Deposit the NFT in the recipient's collection
            depositRef.deposit(token: <-token)
        }

        // withdraw removes an NFT from the collection and moves it to the caller
        pub fun withdraw(withdrawID: UInt64): @NonFungibleToken.NFT {
            let token <- self.ownedNFTs.remove(key: withdrawID) ?? panic("missing NFT")

            emit Withdraw(id: token.id, from: self.owner?.address)

            return <-token
        }

        // deposit takes a NFT and adds it to the collections dictionary
        // and adds the ID to the id array
        pub fun deposit(token: @NonFungibleToken.NFT) {
            let token <- token as! @ModernMusicianNFT.NFT

            let id: UInt64 = token.id

            // add the new token to the dictionary which removes the old one
            let oldToken <- self.ownedNFTs[id] <- token

            emit Deposit(id: id, to: self.owner?.address)

            destroy oldToken
        }

        // getIDs returns an array of the IDs that are in the collection
        pub fun getIDs(): [UInt64] {
            return self.ownedNFTs.keys
        }

        // borrowNFT gets a reference to an NFT in the collection
        // so that the caller can read its metadata and call its methods
        pub fun borrowNFT(id: UInt64): &NonFungibleToken.NFT {
            return &self.ownedNFTs[id] as &NonFungibleToken.NFT
        }
 
        pub fun borrowModernMusicianNFT(id: UInt64): &ModernMusicianNFT.NFT? {
            if self.ownedNFTs[id] != nil {
                // Create an authorized reference to allow downcasting
                let ref = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT
                return ref as! &ModernMusicianNFT.NFT
            }

            return nil
        }

        pub fun borrowViewResolver(id: UInt64): &AnyResource{MetadataViews.Resolver} {
            let nft = &self.ownedNFTs[id] as auth &NonFungibleToken.NFT
            let ModernMusicianNFT = nft as! &ModernMusicianNFT.NFT
            return ModernMusicianNFT as &AnyResource{MetadataViews.Resolver}
        }

        destroy() {
            destroy self.ownedNFTs
        }
    }

    // public function that anyone can call to create a new empty collection
    pub fun createEmptyCollection(): @NonFungibleToken.Collection {
        return <- create Collection()
    }

    // Resource that an admin or something similar would own to be
    // able to mint new NFTs
    //
    pub resource NFTMinter {

        // mintNFT mints a new NFT with a new ID
        // and deposit it in the recipients collection using their collection reference
        pub fun mintNFT(
            recipient: &{NonFungibleToken.CollectionPublic},
            name: String,
            description: String,
            thumbnail: String,
            waitListIndex: UInt64
        ) {
            if(waitingList.length <= waitListIndex){
                panic("NFT is not exist in waiting list")
            }

            // create a new NFT
            var newNFT <- create NFT(
                id: ModernMusicianNFT.totalSupply,
                name: name,
                description: description,
                thumbnail: thumbnail,
                preOrderer: waitingList[waitListIndex].preOrderer
            )

            // deposit it in the recipient's account using their reference
            recipient.deposit(token: <-newNFT)

            self.waitingList.remove(at: waitListIndex);

            ModernMusicianNFT.totalSupply = ModernMusicianNFT.totalSupply + UInt64(1)
        }

        pub fun addUpcomingInfo(info: UpcomingInfo){
            self.waitingList.append(info);
        }

        pub fun removeUpcomingInfo(index: UInt64){
            self.waitingList.remove(at: index);
        }

    }

    init() {
        // Initialize the total supply
        self.totalSupply = 0

        // Set the named paths
        self.CollectionStoragePath = /storage/ModernMusicianNFTCollection
        self.CollectionPublicPath = /public/ModernMusicianNFTCollection
        self.MinterStoragePath = /storage/ModernMusicianNFTMinter

        // Create a Collection resource and save it to storage
        let collection <- create Collection()
        self.account.save(<-collection, to: self.CollectionStoragePath)

        // create a public capability for the collection
        self.account.link<&ModernMusicianNFT.Collection{NonFungibleToken.CollectionPublic, ModernMusicianNFT.ModernMusicianNFTCollectionPublic}>(
            self.CollectionPublicPath,
            target: self.CollectionStoragePath
        )

        // Create a Minter resource and save it to storage
        let minter <- create NFTMinter()
        self.account.save(<-minter, to: self.MinterStoragePath)

        emit ContractInitialized()
    }
}
