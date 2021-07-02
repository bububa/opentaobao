package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymRequirementsReceive 交易猫需求接单接口
// alibaba.jym.requirements.receive
//
// 交易猫需求接单接口
func AlibabaJymRequirementsReceive(clt *core.SDKClient, req *jym.AlibabaJymRequirementsReceiveAPIRequest, session string) (*jym.AlibabaJymRequirementsReceiveAPIResponse, error) {
	var resp jym.AlibabaJymRequirementsReceiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
