package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtracecodeplatformcodeentscan 药品商家扫码
// alibaba.alihealth.tracecodeplatform.code.entscan
//
// 药品商家扫描药品监管码，只有该商家的药才返回
func Alibabaalihealthtracecodeplatformcodeentscan(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtracecodeplatformcodeentscanAPIRequest, session string) (*alihealth2.AlibabaalihealthtracecodeplatformcodeentscanAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtracecodeplatformcodeentscanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
