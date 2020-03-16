package api

import (
	"encoding/json"
	"fmt"

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
	MaxSigCheckDepth                        uint8          `json:"GRAPHENE_MAX_SIG_CHECK_DEPTH"`
	MinTransactionSizeLimit                 uint64         `json:"GRAPHENE_MIN_TRANSACTION_SIZE_LIMIT"`
	MinBlockInterval                        uint8          `json:"GRAPHENE_MIN_BLOCK_INTERVAL"`
	MaxBlockInterval                        uint8          `json:"GRAPHENE_MAX_BLOCK_INTERVAL"`
	DefaultBlockInterval                    uint8          `json:"GRAPHENE_DEFAULT_BLOCK_INTERVAL"`
	DefaultMaxTransactionSize               uint32         `json:"GRAPHENE_DEFAULT_MAX_TRANSACTION_SIZE"`
	DefaultMaxBlockSize                     uint32         `json:"GRAPHENE_DEFAULT_MAX_BLOCK_SIZE"`
	DefaultMaxTimeUntilExpiration           uint32         `json:"GRAPHENE_DEFAULT_MAX_TIME_UNTIL_EXPIRATION"`
	DefaultMaintenanceInterval              uint32         `json:"GRAPHENE_DEFAULT_MAINTENANCE_INTERVAL"`
	DefaultMaintenanceSkipSlots             uint8          `json:"GRAPHENE_DEFAULT_MAINTENANCE_SKIP_SLOTS"`
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
	DefaultMaxAuthorityMembership           uint16         `json:"GRAPHENE_DEFAULT_MAX_AUTHORITY_MEMBERSHIP"`
	DefaultMaxAssetWhitelistAuthorities     uint8          `json:"GRAPHENE_DEFAULT_MAX_ASSET_WHITELIST_AUTHORITIES"`
	DefaultMaxAssetFeedPublishers           uint8          `json:"GRAPHENE_DEFAULT_MAX_ASSET_FEED_PUBLISHERS"`
	CollateralRatioDenom                    uint64         `json:"GRAPHENE_COLLATERAL_RATIO_DENOM"`
	MinCollateralRatio                      uint64         `json:"GRAPHENE_MIN_COLLATERAL_RATIO"`
	MaxCollateralRatio                      uint64         `json:"GRAPHENE_MAX_COLLATERAL_RATIO"`
	DefaultMaintenanceCollateralRatio       uint64         `json:"GRAPHENE_DEFAULT_MAINTENANCE_COLLATERAL_RATIO"`
	DefaultMaxShortSqueezeRatio             uint64         `json:"GRAPHENE_DEFAULT_MAX_SHORT_SQUEEZE_RATIO"`
	DefaultMaxWitnesses                     uint16         `json:"GRAPHENE_DEFAULT_MAX_WITNESSES"`
	DefaultMaxCommittee                     uint16         `json:"GRAPHENE_DEFAULT_MAX_COMMITTEE"`
	DefaultMaxProposalLifetimeSec           uint32         `json:"GRAPHENE_DEFAULT_MAX_PROPOSAL_LIFETIME_SEC"`
	DefaultCommitteeProposalReviewPeriodSec uint32         `json:"GRAPHENE_DEFAULT_COMMITTEE_PROPOSAL_REVIEW_PERIOD_SEC"`
	DefaultNetworkPercentOfFee              uint16         `json:"GRAPHENE_DEFAULT_NETWORK_PERCENT_OF_FEE"`
	DefaultLifetimeReferrerPercentOfFee     uint16         `json:"GRAPHENE_DEFAULT_LIFETIME_REFERRER_PERCENT_OF_FEE"`
	DefaultCashbackVestingPeriodSec         uint32         `json:"GRAPHENE_DEFAULT_CASHBACK_VESTING_PERIOD_SEC"`
	DefaultCashbackVestingThreshold         int64          `json:"GRAPHENE_DEFAULT_CASHBACK_VESTING_THRESHOLD"`
	DefaultBurnPercentOfFee                 uint16         `json:"GRAPHENE_DEFAULT_BURN_PERCENT_OF_FEE"`
	DefaultMaxAssertOpcode                  uint16         `json:"GRAPHENE_DEFAULT_MAX_ASSERT_OPCODE"`
	DefaultFeeLiquidationThreshold          int64          `json:"GRAPHENE_DEFAULT_FEE_LIQUIDATION_THRESHOLD"`
	DefaultAccountsPerFeeScale              uint16         `json:"GRAPHENE_DEFAULT_ACCOUNTS_PER_FEE_SCALE"`
	DefaultAccountFeeScaleBitshifts         uint8          `json:"GRAPHENE_DEFAULT_ACCOUNT_FEE_SCALE_BITSHIFTS"`
	MaxWorkerNameLength                     uint64         `json:"GRAPHENE_MAX_WORKER_NAME_LENGTH"`
	MaxUrlLength                            uint64         `json:"GRAPHENE_MAX_URL_LENGTH"`
	CoreAssetCycleRate                      uint64         `json:"GRAPHENE_CORE_ASSET_CYCLE_RATE"`
	CoreAssetCycleRateBits                  uint64         `json:"GRAPHENE_CORE_ASSET_CYCLE_RATE_BITS"`
	DefaultWitnessPayPerBlock               int64          `json:"GRAPHENE_DEFAULT_WITNESS_PAY_PER_BLOCK"`
	DefaultWitnessPayVestingSeconds         uint32         `json:"GRAPHENE_DEFAULT_WITNESS_PAY_VESTING_SECONDS"`
	DefaultWorkerBudgetPerDay               string         `json:"GRAPHENE_DEFAULT_WORKER_BUDGET_PER_DAY"`
	CommitteeAccount                        types.ObjectID `json:"GRAPHENE_COMMITTEE_ACCOUNT"`
	WitnessAccount                          types.ObjectID `json:"GRAPHENE_WITNESS_ACCOUNT"`
	RelaxedCommitteeAccount                 types.ObjectID `json:"GRAPHENE_RELAXED_COMMITTEE_ACCOUNT"`
	NullAccount                             types.ObjectID `json:"GRAPHENE_NULL_ACCOUNT"`
	TempAccount                             types.ObjectID `json:"GRAPHENE_TEMP_ACCOUNT"`
}

//GlobalProperty structure for the GetGlobalProperties function
type GlobalProperty struct {
	Parameters             types.ChainParameters    `json:"parameters"`
	PendingParameters      []*types.ChainParameters `json:"pending_parameters,omitempty"`
	NextAvailableVoteID    uint32                   `json:"next_available_vote_id"`
	ActiveCommitteeMembers []types.ObjectID         `json:"active_committee_members"`
	ActiveWitnesses        []types.ObjectID         `json:"active_witnesses"`
}

//ChainProperty structure for the GetChainProperties function
type ChainProperty struct {
	ChainID                  string                  `json:"chain_id"`
	ImmutableChainParameters ImmutableChainParameter `json:"immutable_chain_parameters"`
}

//ImmutableChainParameter is an additional structure used by ChainProperty structure.
type ImmutableChainParameter struct {
	MinCommitteeMemberCount uint16 `json:"min_committee_member_count"`
	MinWitnessCount         uint16 `json:"min_witness_count"`
	NumSpecialAccounts      uint32 `json:"num_special_accounts"`
	NumSpecialAssets        uint32 `json:"num_special_assets"`
}

//DynamicGlobalProperties structure for the GetDynamicGlobalProperties function
type DynamicGlobalProperties struct {
	HeadBlockNumber                uint32     `json:"head_block_number"`
	HeadBlockID                    string     `json:"head_block_id"`
	Time                           types.Time `json:"time"`
	CurrentWitness                 string     `json:"current_witness"`
	NextMaintenanceTime            types.Time `json:"next_maintenance_time"`
	LastBudgetTime                 types.Time `json:"last_budget_time"`
	WitnessBudget                  int64      `json:"witness_budget"`
	AccountsRegisteredThisInterval uint32     `json:"accounts_registered_this_interval"`
	RecentlyMissedCount            uint32     `json:"recently_missed_count"`
	CurrentAslot                   uint64     `json:"current_aslot"`
	RecentSlotsFilled              string     `json:"recent_slots_filled"`
	DynamicFlags                   uint32     `json:"dynamic_flags"`
	LastIrreversibleBlockNum       uint32     `json:"last_irreversible_block_num"`
	DynamicFlagBits                uint8      `json:"dynamic_flag_bits"`
}

//BlockHeader structure for the GetBlockHeader function
type BlockHeader struct {
	Number                uint32         `json:"-"`
	Previous              string         `json:"previous"`
	Timestamp             string         `json:"timestamp"`
	Witness               types.ObjectID `json:"witness"`
	TransactionMerkleRoot string         `json:"transaction_merkle_root"`
	Extensions            []interface{}  `json:"extensions"`
}

//BlockHeader structure for the GetBlockHeaderBatch function
type BlockHeaderBatch struct {
	BlockNum  uint32
	BlockHead BlockHeader
}

func (n *BlockHeaderBatch) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&n.BlockNum, &n.BlockHead}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in BlockHeaderBatch: %d != %d", g, e)
	}
	return nil
}

//BroadcastResponse structure for the BroadcastTransactionSynchronous function
type BroadcastResponse struct {
	ID       string `json:"id"`
	BlockNum int32  `json:"block_num"`
	TrxNum   int32  `json:"trx_num"`
	Expired  bool   `json:"expired"`
}

//CallbackBlockResponse structure for the SetBlockAppliedCallback function
type CallbackBlockResponse struct {
	BlockNum   int           `json:"block_num"`
	Operations []interface{} `json:"operations"`
}
