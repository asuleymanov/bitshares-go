package types

import (
//"github.com/asuleymanov/bitshares-go/encoding/transaction"
)

type ChainParameters struct {
	CurrentFees                      map[string]interface{} `json:"current_fees"`
	BlockInterval                    uint8                  `json:"block_interval"`
	MaintenanceInterval              uint32                 `json:"maintenance_interval"`
	MaintenanceSkipSlots             uint8                  `json:"maintenance_skip_slots"`
	CommitteeProposalReviewPeriod    uint32                 `json:"committee_proposal_review_period"`
	MaximumTransactionSize           uint32                 `json:"maximum_transaction_size"`
	MaximumBlockSize                 uint32                 `json:"maximum_block_size"`
	MaximumTimeUntilExpiration       uint32                 `json:"maximum_time_until_expiration"`
	MaximumProposalLifetime          uint32                 `json:"maximum_proposal_lifetime"`
	MaximumAssetWhitelistAuthorities uint8                  `json:"maximum_asset_whitelist_authorities"`
	MaximumAssetFeedPublishers       uint8                  `json:"maximum_asset_feed_publishers"`
	MaximumWitnessCount              uint16                 `json:"maximum_witness_count"`
	MaximumCommitteeCount            uint16                 `json:"maximum_committee_count"`
	MaximumAuthorityMembership       uint16                 `json:"maximum_authority_membership"`
	ReservePercentOfFee              uint16                 `json:"reserve_percent_of_fee"`
	NetworkPercentOfFee              uint16                 `json:"network_percent_of_fee"`
	LifetimeReferrerPercentOfFee     uint16                 `json:"lifetime_referrer_percent_of_fee"`
	CashbackVestingPeriodSeconds     uint32                 `json:"cashback_vesting_period_seconds"`
	CashbackVestingThreshold         int64                  `json:"cashback_vesting_threshold"`
	CountNonMemberVotes              bool                   `json:"count_non_member_votes"`
	AllowNonMemberWhitelists         bool                   `json:"allow_non_member_whitelists"`
	WitnessPayPerBlock               int64                  `json:"witness_pay_per_block"`
	WitnessPayVestingSeconds         uint32                 `json:"witness_pay_vesting_seconds"`
	WorkerBudgetPerDay               string                 `json:"worker_budget_per_day"`
	MaxPredicateOpcode               uint16                 `json:"max_predicate_opcode"`
	FeeLiquidationThreshold          int64                  `json:"fee_liquidation_threshold"`
	AccountsPerFeeScale              uint16                 `json:"accounts_per_fee_scale"`
	AccountFeeScaleBitshifts         uint8                  `json:"account_fee_scale_bitshifts"`
	MaxAuthorityDepth                uint8                  `json:"max_authority_depth"`
	Extensions                       map[string]interface{} `json:"extensions"`
}
