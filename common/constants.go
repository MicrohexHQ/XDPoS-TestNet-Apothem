package common

import (
	"math/big"
)

const (
	RewardMasterPercent        = 90
	RewardVoterPercent         = 0
	RewardFoundationPercent    = 10
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 800
	EpocBlockOpening           = 850
	EpocBlockRandomize         = 900
	MaxMasternodes             = 18
	MaxMasternodesV2           = 108
	LimitPenaltyEpoch          = 4
	BlocksPerYear              = uint64(15768000)
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 2500
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1
)

var TIP2019Block = big.NewInt(1)
var TIPSigning = big.NewInt(3000000)
var TIPRandomize = big.NewInt(3464000)
var BlackListHFNumber = uint64(9349100)
var TIPIncreaseMasternodes = big.NewInt(1199999) // example 10 milions
var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var XRC21IssuerSMCTestNet = HexToAddress("0x7081C72c9DC44686C7B7EAB1d338EA137Fa9f0D3")
var XRC21IssuerSMC = HexToAddress("0x8c0faeb5C6bEd2129b8674F262Fd45c4e9468bee")
var XRC21GasPrice = big.NewInt(DefaultMinGasPrice)
var Blacklist = map[Address]bool{
	HexToAddress("0x5248bfb72fd4f234e062d3e9bb76f08643004fcd"): true,
	HexToAddress("0x5ac26105b35ea8935be382863a70281ec7a985e9"): true,
	HexToAddress("0x09c4f991a41e7ca0645d7dfbfee160b55e562ea4"): true,
	HexToAddress("0xb3157bbc5b401a45d6f60b106728bb82ebaa585b"): true,
	HexToAddress("0x741277a8952128d5c2ffe0550f5001e4c8247674"): true,
	HexToAddress("0x10ba49c1caa97d74b22b3e74493032b180cebe01"): true,
	HexToAddress("0x07048d51d9e6179578a6e3b9ee28cdc183b865e4"): true,
	HexToAddress("0x4b899001d73c7b4ec404a771d37d9be13b8983de"): true,
	HexToAddress("0x85cb320a9007f26b7652c19a2a65db1da2d0016f"): true,
	HexToAddress("0x06869dbd0e3a2ea37ddef832e20fa005c6f0ca39"): true,
	HexToAddress("0x82e48bc7e2c93d89125428578fb405947764ad7c"): true,
	HexToAddress("0x1f9a78534d61732367cbb43fc6c89266af67c989"): true,
	HexToAddress("0x7c3b1fa91df55ff7af0cad9e0399384dc5c6641b"): true,
	HexToAddress("0x5888dc1ceb0ff632713486b9418e59743af0fd20"): true,
	HexToAddress("0xa512fa1c735fc3cc635624d591dd9ea1ce339ca5"): true,
	HexToAddress("0x0832517654c7b7e36b1ef45d76de70326b09e2c7"): true,
	HexToAddress("0xca14e3c4c78bafb60819a78ff6e6f0f709d2aea7"): true,
	HexToAddress("0x652ce195a23035114849f7642b0e06647d13e57a"): true,
	HexToAddress("0x29a79f00f16900999d61b6e171e44596af4fb5ae"): true,
	HexToAddress("0xf9fd1c2b0af0d91b0b6754e55639e3f8478dd04a"): true,
	HexToAddress("0xb835710c9901d5fe940ef1b99ed918902e293e35"): true,
	HexToAddress("0x04dd29ce5c253377a9a3796103ea0d9a9e514153"): true,
	HexToAddress("0x2b4b56846eaf05c1fd762b5e1ac802efd0ab871c"): true,
	HexToAddress("0x1d1f909f6600b23ce05004f5500ab98564717996"): true,
	HexToAddress("0x0dfdcebf80006dc9ab7aae8c216b51c6b6759e86"): true,
	HexToAddress("0x2b373890a28e5e46197fbc04f303bbfdd344056f"): true,
	HexToAddress("0xa8a3ef3dc5d8e36aee76f3671ec501ec31e28254"): true,
	HexToAddress("0x4f3d18136fe2b5665c29bdaf74591fc6625ef427"): true,
	HexToAddress("0x175d728b0e0f1facb5822a2e0c03bde93596e324"): true,
	HexToAddress("0xd575c2611984fcd79513b80ab94f59dc5bab4916"): true,
	HexToAddress("0x0579337873c97c4ba051310236ea847f5be41bc0"): true,
	HexToAddress("0xed12a519cc15b286920fc15fd86106b3e6a16218"): true,
	HexToAddress("0x492d26d852a0a0a2982bb40ec86fe394488c419e"): true,
	HexToAddress("0xce5c7635d02dc4e1d6b46c256cae6323be294a32"): true,
	HexToAddress("0x8b94db158b5e78a6c032c7e7c9423dec62c8b11c"): true,
	HexToAddress("0x0e7c48c085b6b0aa7ca6e4cbcc8b9a92dc270eb4"): true,
	HexToAddress("0x206e6508462033ef8425edc6c10789d241d49acb"): true,
	HexToAddress("0x7710e7b7682f26cb5a1202e1cff094fbf7777758"): true,
	HexToAddress("0xcb06f949313b46bbf53b8e6b2868a0c260ff9385"): true,
	HexToAddress("0xf884e43533f61dc2997c0e19a6eff33481920c00"): true,
	HexToAddress("0x8b635ef2e4c8fe21fc2bda027eb5f371d6aa2fc1"): true,
	HexToAddress("0x10f01a27cf9b29d02ce53497312b96037357a361"): true,
	HexToAddress("0x693dd49b0ed70f162d733cf20b6c43dc2a2b4d95"): true,
	HexToAddress("0xe0bec72d1c2a7a7fb0532cdfac44ebab9f6f41ee"): true,
	HexToAddress("0xc8793633a537938cb49cdbbffd45428f10e45b64"): true,
	HexToAddress("0x0d07a6cbbe9fa5c4f154e5623bfe47fb4d857d8e"): true,
	HexToAddress("0xd4080b289da95f70a586610c38268d8d4cf1e4c4"): true,
	HexToAddress("0x8bcfb0caf41f0aa1b548cae76dcdd02e33866a1b"): true,
	HexToAddress("0xabfef22b92366d3074676e77ea911ccaabfb64c1"): true,
	HexToAddress("0xcc4df7a32faf3efba32c9688def5ccf9fefe443d"): true,
	HexToAddress("0x7ec1e48a582475f5f2b7448a86c4ea7a26ea36f8"): true,
	HexToAddress("0xe3de67289080f63b0c2612844256a25bb99ac0ad"): true,
	HexToAddress("0x3ba623300cf9e48729039b3c9e0dee9b785d636e"): true,
	HexToAddress("0x402f2cfc9c8942f5e7a12c70c625d07a5d52fe29"): true,
	HexToAddress("0xd62358d42afbde095a4ca868581d85f9adcc3d61"): true,
	HexToAddress("0x3969f86acb733526cd61e3c6e3b4660589f32bc6"): true,
	HexToAddress("0x67615413d7cdadb2c435a946aec713a9a9794d39"): true,
	HexToAddress("0xfe685f43acc62f92ab01a8da80d76455d39d3cb3"): true,
	HexToAddress("0x3538a544021c07869c16b764424c5987409cba48"): true,
	HexToAddress("0xe187cf86c2274b1f16e8225a7da9a75aba4f1f5f"): true,
}
