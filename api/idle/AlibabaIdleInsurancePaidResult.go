package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleInsurancePaidResult 闲鱼行业保险理赔结果同步
// alibaba.idle.insurance.paid.result
//
// 闲鱼行业保险理赔结果同步
func AlibabaIdleInsurancePaidResult(clt *core.SDKClient, req *idle.AlibabaIdleInsurancePaidResultAPIRequest, session string) (*idle.AlibabaIdleInsurancePaidResultAPIResponse, error) {
	var resp idle.AlibabaIdleInsurancePaidResultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
