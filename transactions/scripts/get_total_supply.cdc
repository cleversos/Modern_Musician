import ModernMusicianNFT from "../../contracts/ModernMusicianNFT.cdc"

pub fun main(): UInt64 {
    return ModernMusicianNFT.totalSupply
}
