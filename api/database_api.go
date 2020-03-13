package api

//login_api

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
