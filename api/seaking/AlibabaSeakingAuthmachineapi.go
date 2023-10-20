package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// Alibabaseakingauthmachineapi 机翻Api授权
// alibaba.seaking.authmachineapi
//
// 机翻Api授权
func Alibabaseakingauthmachineapi(clt *core.SDKClient, req *seaking.AlibabaseakingauthmachineapiAPIRequest, session string) (*seaking.AlibabaseakingauthmachineapiAPIResponse, error) {
	var resp seaking.AlibabaseakingauthmachineapiAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
