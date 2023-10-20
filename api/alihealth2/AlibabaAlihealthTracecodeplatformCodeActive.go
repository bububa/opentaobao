package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodeplatformcodeactive 正大鸡蛋激活追溯码
// alibaba.alihealth.tracecodeplatform.code.active
//
// 用于正大鸡蛋激活追溯码
func Alibabaalihealthtracecodeplatformcodeactive(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodeplatformcodeactiveAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodeplatformcodeactiveAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodeplatformcodeactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
