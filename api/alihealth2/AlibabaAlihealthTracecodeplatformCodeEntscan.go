package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// AlibabaAlihealthTracecodeplatformCodeEntscan 药品商家扫码
// alibaba.alihealth.tracecodeplatform.code.entscan
//
// 药品商家扫描药品监管码，只有该商家的药才返回
func AlibabaAlihealthTracecodeplatformCodeEntscan(clt *core.SDKClient, req *alihealth2.AlibabaAlihealthTracecodeplatformCodeEntscanAPIRequest, session string) (*alihealth2.AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse, error) {
	var resp alihealth2.AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
