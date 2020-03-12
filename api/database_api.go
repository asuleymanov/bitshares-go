package api

//login_api

func (api *API) GetConfig() (*Config, error) {
	var resp Config
	err := api.call(*api.DatabaseID, "get_config", EmptyParams, &resp)
	return &resp, err
}

func (api *API) GetChainID() (*string, error) {
	var resp string
	err := api.call(*api.DatabaseID, "get_chain_id", EmptyParams, &resp)
	return &resp, err
}