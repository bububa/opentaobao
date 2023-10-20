package alihealth2

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihealth2"
)

// Alibabaalihealthtmscutconfirm 配拦截失败CP确认结果并回告
// alibaba.alihealth.tms.cut.confirm
//
// 配拦截失败CP确认结果并回告
func Alibabaalihealthtmscutconfirm(clt *core.SDKClient, req *alihealth2.AlibabaalihealthtmscutconfirmAPIRequest, session string) (*alihealth2.AlibabaalihealthtmscutconfirmAPIResponse, error) {
	var resp alihealth2.AlibabaalihealthtmscutconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
