package seaking

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/seaking"
)

// AlibabaSeakingAuthmachineapi 机翻Api授权
// alibaba.seaking.authmachineapi
//
// 机翻Api授权
func AlibabaSeakingAuthmachineapi(clt *core.SDKClient, req *seaking.AlibabaSeakingAuthmachineapiAPIRequest, resp *seaking.AlibabaSeakingAuthmachineapiAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
