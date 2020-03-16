package api

import (
	"encoding/json"
)

//database_api

//GetConfig api request get_config
func (api *API) GetConfig() (*Config, error) {
	var resp Config
	err := api.call(*api.DatabaseID, "get_config", EmptyParams, &resp)
	return &resp, err
}

//GetChainID api request get_chain_id
func (api *API) GetChainID() (*string, error) {
	var resp string
	err := api.call(*api.DatabaseID, "get_chain_id", EmptyParams, &resp)
	return &resp, err
}

//GetGlobalProperties api request get_global_properties
func (api *API) GetGlobalProperties() (*GlobalProperty, error) {
	var resp GlobalProperty
	err := api.call(*api.DatabaseID, "get_global_properties", EmptyParams, &resp)
	return &resp, err
}

//GetChainProperties api request get_chain_properties
func (api *API) GetChainProperties() (*ChainProperty, error) {
	var resp ChainProperty
	err := api.call(*api.DatabaseID, "get_chain_properties", EmptyParams, &resp)
	return &resp, err
}

//GetDynamicGlobalProperties api request get_dynamic_global_properties
func (api *API) GetDynamicGlobalProperties() (*DynamicGlobalProperties, error) {
	var resp DynamicGlobalProperties
	err := api.call(*api.DatabaseID, "get_dynamic_global_properties", EmptyParams, &resp)
	return &resp, err
}

//GetBlockHeader api request get_block_header
func (api *API) GetBlockHeader(blockNum uint32) (*BlockHeader, error) {
	var resp BlockHeader
	err := api.call(*api.DatabaseID, "get_block_header", []uint32{blockNum}, &resp)
	resp.Number = blockNum
	return &resp, err
}

//GetBlockHeaderBatch api request get_block_header_batch
func (api *API) GetBlockHeaderBatch(blockNum ...uint32) ([]BlockHeaderBatch, error) {
	var resp []BlockHeaderBatch
	err := api.call(*api.DatabaseID, "get_block_header_batch", []interface{}{blockNum}, &resp)
	return resp, err
}

//GetKeyReferences returns a list of accounts by a list of public keys
func (api *API) GetKeyReferences(pubkey ...string) ([]*string, error) {
	var resp []*string
	err := api.call(*api.DatabaseID, "get_key_references", []interface{}{pubkey}, &resp)
	return resp, err
}

//IsPublicKeyRegistered api request is_public_key_registered
func (api *API) IsPublicKeyRegistered(pubkey string) (*bool, error) {
	var resp bool
	err := api.call(*api.DatabaseID, "is_public_key_registered", []interface{}{pubkey}, &resp)
	return &resp, err
}

//GetAccountCount returns the number of registered users.
func (api *API) GetAccountCount() (*uint64, error) {
	var resp uint64
	err := api.call(*api.DatabaseID, "get_account_count", EmptyParams, &resp)
	return &resp, err
}

// Set callback to invoke as soon as a new block is applied
func (api *API) SetBlockAppliedCallback(notice func(header *CallbackBlockResponse, error error)) (err error) {
	err = api.setCallback(*api.DatabaseID, "set_block_applied_callback", func(raw json.RawMessage) {
		var header CallbackBlockResponse
		if err := json.Unmarshal(raw, &header); err != nil {
			notice(nil, err)
		}
		notice(&header, nil)
	})
	return
}
