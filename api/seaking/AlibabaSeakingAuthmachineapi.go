package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingAuthmachineapi 机翻Api授权
// alibaba.seaking.authmachineapi
//
// 机翻Api授权
func AlibabaSeakingAuthmachineapi(clt *core.SDKClient, req *seaking.AlibabaSeakingAuthmachineapiAPIRequest, session string) (*seaking.AlibabaSeakingAuthmachineapiAPIResponse, error) {
	var resp seaking.AlibabaSeakingAuthmachineapiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
