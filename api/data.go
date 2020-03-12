package api

import (
	"github.com/asuleymanov/bitshares-go/types"
)

//Config structure for the GetConfig function.
type Config struct {
	Symbol                                  string         `json:"GRAPHENE_SYMBOL"`
	AddressPrefix                           string         `json:"GRAPHENE_ADDRESS_PREFIX"`
	MinAccountNameLength                    uint64         `json:"GRAPHENE_MIN_ACCOUNT_NAME_LENGTH"`
	MaxAccountNameLength                    uint64         `json:"GRAPHENE_MAX_ACCOUNT_NAME_LENGTH"`
	MinAssetSymbolLength                    uint64         `json:"GRAPHENE_MIN_ASSET_SYMBOL_LENGTH"`
	MaxAssetSymbolLength                    uint64         `json:"GRAPHENE_MAX_ASSET_SYMBOL_LENGTH"`
	MaxShareSupply                          string         `json:"GRAPHENE_MAX_SHARE_SUPPLY"`
	MaxSigCheckDepth                        uint64         `json:"GRAPHENE_MAX_SIG_CHECK_DEPTH"`
	MinTransactionSizeLimit                 uint64         `json:"GRAPHENE_MIN_TRANSACTION_SIZE_LIMIT"`
	MinBlockInterval                        uint8          `json:"GRAPHENE_MIN_BLOCK_INTERVAL"`
	MaxBlockInterval                        uint8          `json:"GRAPHENE_MAX_BLOCK_INTERVAL"`
	DefaultBlockInterval                    uint8          `json:"GRAPHENE_DEFAULT_BLOCK_INTERVAL"`
	DefaultMaxTransactionSize               uint64         `json:"GRAPHENE_DEFAULT_MAX_TRANSACTION_SIZE"`
	DefaultMaxBlockSize                     uint64         `json:"GRAPHENE_DEFAULT_MAX_BLOCK_SIZE"`
	DefaultMaxTimeUntilExpiration           uint64         `json:"GRAPHENE_DEFAULT_MAX_TIME_UNTIL_EXPIRATION"`
	DefaultMaintenanceInterval              uint64         `json:"GRAPHENE_DEFAULT_MAINTENANCE_INTERVAL"`
	DefaultMaintenanceSkipSlots             uint64         `json:"GRAPHENE_DEFAULT_MAINTENANCE_SKIP_SLOTS"`
	MinUndoHistory                          uint64         `json:"GRAPHENE_MIN_UNDO_HISTORY"`
	MaxUndoHistory                          uint64         `json:"GRAPHENE_MAX_UNDO_HISTORY"`
	MinBlockSizeLimit                       uint64         `json:"GRAPHENE_MIN_BLOCK_SIZE_LIMIT"`
	BlockchainPrecision                     uint64         `json:"GRAPHENE_BLOCKCHAIN_PRECISION"`
	BlockchainPrecisionDigits               uint64         `json:"GRAPHENE_BLOCKCHAIN_PRECISION_DIGITS"`
	Percent100                              uint16         `json:"GRAPHENE_100_PERCENT"`
	Percent1                                uint16         `json:"GRAPHENE_1_PERCENT"`
	MaxMarketFeePercent                     uint16         `json:"GRAPHENE_MAX_MARKET_FEE_PERCENT"`
	DefaultForceSettlementDelay             uint64         `json:"GRAPHENE_DEFAULT_FORCE_SETTLEMENT_DELAY"`
	DefaultForceSettlementOffset            uint64         `json:"GRAPHENE_DEFAULT_FORCE_SETTLEMENT_OFFSET"`
	DefaultForceSettlementMaxVolume         uint64         `json:"GRAPHENE_DEFAULT_FORCE_SETTLEMENT_MAX_VOLUME"`
	DefaultPriceFeedLifetime                uint64         `json:"GRAPHENE_DEFAULT_PRICE_FEED_LIFETIME"`
	DefaultMaxAuthorityMembership           uint64         `json:"GRAPHENE_DEFAULT_MAX_AUTHORITY_MEMBERSHIP"`
	DefaultMaxAssetWhitelistAuthorities     uint64         `json:"GRAPHENE_DEFAULT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	DefaultMaxAssetFeedPublishers           uint64         `json:"GRAPHENE_DEFAULT_MAX_ASSET_FEED_PUBLISHERS"`
	CollateralRatioDenom                    uint64         `json:"GRAPHENE_COLLATERAL_RATIO_DENOM"`
	MinCollateralRatio                      uint64         `json:"GRAPHENE_MIN_COLLATERAL_RATIO"`
	MaxCollateralRatio                      uint64         `json:"GRAPHENE_MAX_COLLATERAL_RATIO"`
	DefaultMaintenanceCollateralRatio       uint64         `json:"GRAPHENE_DEFAULT_MAINTENANCE_COLLATERAL_RATIO"`
	DefaultMaxShortSqueezeRatio             uint64         `json:"GRAPHENE_DEFAULT_MAX_SHORT_SQUEEZE_RATIO"`
	DefaultMaxWitnesses                     uint64         `json:"GRAPHENE_DEFAULT_MAX_WITNESSES"`
	DefaultMaxCommittee                     uint64         `json:"GRAPHENE_DEFAULT_MAX_COMMITTEE"`
	DefaultMaxProposalLifetimeSec           uint64         `json:"GRAPHENE_DEFAULT_MAX_PROPOSAL_LIFETIME_SEC"`
	DefaultCommitteeProposalReviewPeriodSec uint64         `json:"GRAPHENE_DEFAULT_COMMITTEE_PROPOSAL_REVIEW_PERIOD_SEC"`
	DefaultNetworkPercentOfFee              uint64         `json:"GRAPHENE_DEFAULT_NETWORK_PERCENT_OF_FEE"`
	DefaultLifetimeReferrerPercentOfFee     uint64         `json:"GRAPHENE_DEFAULT_LIFETIME_REFERRER_PERCENT_OF_FEE"`
	DefaultCashbackVestingPeriodSec         uint64         `json:"GRAPHENE_DEFAULT_CASHBACK_VESTING_PERIOD_SEC"`
	DefaultCashbackVestingThreshold         uint64         `json:"GRAPHENE_DEFAULT_CASHBACK_VESTING_THRESHOLD"`
	DefaultBurnPercentOfFee                 uint64         `json:"GRAPHENE_DEFAULT_BURN_PERCENT_OF_FEE"`
	DefaultMaxAssertOpcode                  uint64         `json:"GRAPHENE_DEFAULT_MAX_ASSERT_OPCODE"`
	DefaultFeeLiquidationThreshold          uint64         `json:"GRAPHENE_DEFAULT_FEE_LIQUIDATION_THRESHOLD"`
	DefaultAccountsPerFeeScale              uint64         `json:"GRAPHENE_DEFAULT_ACCOUNTS_PER_FEE_SCALE"`
	DefaultAccountFeeScaleBitshifts         uint64         `json:"GRAPHENE_DEFAULT_ACCOUNT_FEE_SCALE_BITSHIFTS"`
	MaxWorkerNameLength                     uint64         `json:"GRAPHENE_MAX_WORKER_NAME_LENGTH"`
	MaxUrlLength                            uint64         `json:"GRAPHENE_MAX_URL_LENGTH"`
	CoreAssetCycleRate                      uint64         `json:"GRAPHENE_CORE_ASSET_CYCLE_RATE"`
	CoreAssetCycleRateBits                  uint64         `json:"GRAPHENE_CORE_ASSET_CYCLE_RATE_BITS"`
	DefaultWitnessPayPerBlock               uint64         `json:"GRAPHENE_DEFAULT_WITNESS_PAY_PER_BLOCK"`
	DefaultWitnessPayVestingSeconds         uint64         `json:"GRAPHENE_DEFAULT_WITNESS_PAY_VESTING_SECONDS"`
	DefaultWorkerBudgetPerDay               string         `json:"GRAPHENE_DEFAULT_WORKER_BUDGET_PER_DAY"`
	CommitteeAccount                        types.ObjectID `json:"GRAPHENE_COMMITTEE_ACCOUNT"`
	WitnessAccount                          types.ObjectID `json:"GRAPHENE_WITNESS_ACCOUNT"`
	RelaxedCommitteeAccount                 types.ObjectID `json:"GRAPHENE_RELAXED_COMMITTEE_ACCOUNT"`
	NullAccount                             types.ObjectID `json:"GRAPHENE_NULL_ACCOUNT"`
	TempAccount                             types.ObjectID `json:"GRAPHENE_TEMP_ACCOUNT"`
}
