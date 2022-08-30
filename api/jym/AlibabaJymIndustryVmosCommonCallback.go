package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymIndustryVmosCommonCallback vmos游戏信息采集结果回调通知
// alibaba.jym.industry.vmos.common.callback
//
// vmos游戏信息采集结果回调通知
func AlibabaJymIndustryVmosCommonCallback(clt *core.SDKClient, req *jym.AlibabaJymIndustryVmosCommonCallbackAPIRequest, session string) (*jym.AlibabaJymIndustryVmosCommonCallbackAPIResponse, error) {
	var resp jym.AlibabaJymIndustryVmosCommonCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
