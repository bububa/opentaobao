package aedropshiper

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aedropshiper"
)

// AliexpressDsAddInfo 上报DS信息
// aliexpress.ds.add.info
//
// ISV用户上报下游DS信息
func AliexpressDsAddInfo(clt *core.SDKClient, req *aedropshiper.AliexpressDsAddInfoAPIRequest, session string) (*aedropshiper.AliexpressDsAddInfoAPIResponse, error) {
	var resp aedropshiper.AliexpressDsAddInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
