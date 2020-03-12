package api

//login_api

func (api *API) Login(username, password string) (bool, error) {
	var resp bool
	err := api.call(1,"login", []interface{}{username, password}, &resp)
	return resp, err
}

func (api *API) DataBaseAPIID() (*uint8, error) {
	var id uint8
	err := api.call(1,"database", EmptyParams, &id)
	return &id, err
}

func (api *API) HistoryAPIID() (*uint8, error) {
	var id uint8
	err := api.call(1,"history", EmptyParams, &id)
	return &id, err
}

func (api *API) NetworkBroadcastAPIID() (*uint8, error) {
	var id uint8
	err := api.call(1,"network_broadcast", EmptyParams, &id)
	return &id, err
}