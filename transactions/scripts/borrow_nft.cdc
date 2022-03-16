import NonFungibleToken from "../../contracts/NonFungibleToken.cdc"
import ModernMusicianNFT from "../../contracts/ModernMusicianNFT.cdc"

// This script borrows an NFT from a collection
pub fun main(address: Address, id: UInt64) {
    let account = getAccount(address)

    let collectionRef = account
        .getCapability(ModernMusicianNFT.CollectionPublicPath)
        .borrow<&{NonFungibleToken.CollectionPublic}>()
        ?? panic("Could not borrow capability from public collection")

    // Borrow a reference to a specific NFT in the collection
    let _ = collectionRef.borrowNFT(id: id)
}
