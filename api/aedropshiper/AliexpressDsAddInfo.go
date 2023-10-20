package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// Aliexpressdsaddinfo 上报DS信息
// aliexpress.ds.add.info
//
// ISV用户上报下游DS信息
func Aliexpressdsaddinfo(clt *core.SDKClient, req *aedropshiper.AliexpressdsaddinfoAPIRequest, session string) (*aedropshiper.AliexpressdsaddinfoAPIResponse, error) {
	var resp aedropshiper.AliexpressdsaddinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
