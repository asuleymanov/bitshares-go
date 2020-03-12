package api

import (
	"encoding/json"
	"fmt"

	"github.com/asuleymanov/bitshares-go/transports"
)

//API plug-in structure
type API struct {
	caller                                    transports.Caller
	DatabaseID, HistoryID, NetworkBroadcastID *uint8
}

var (
	EmptyParams = []struct{}{}
)

//NewAPI plug-in initialization
func NewAPI(caller transports.Caller) *API {
	return &API{caller: caller}
}

func (api *API) SetAPIID() error {
	var err error
	api.DatabaseID, err = api.DataBaseAPIID()
	if err != nil {
		return fmt.Errorf("Set DataBaseAPIID :", err)
	}
	api.NetworkBroadcastID, err = api.NetworkBroadcastAPIID()
	if err != nil {
		return fmt.Errorf("Set NetworkBroadcastAPIID :", err)
	}
	api.HistoryID, err = api.HistoryAPIID()
	if err != nil {
		return fmt.Errorf("Set HistoryAPIID :", err)
	}
	return nil
}

func (api *API) call(apiID uint8, method string, params, resp interface{}) error {
	return api.caller.Call("call", []interface{}{apiID, method, params}, resp)
}

func (api *API) setCallback(apiID string, method string, callback func(raw json.RawMessage)) error {
	return api.caller.SetCallback(apiID, method, callback)
}

func queryJson(query []interface{}) (string, error) {
	b, err := json.Marshal(query)
	return string(b[1 : len(b)-1]), err
}
