package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjmosfundcreatebill 创建一个付款单
// alibaba.mj.mos.fund.createbill
//
// 创建一个付款单
func Alibabamjmosfundcreatebill(clt *core.SDKClient, req *mos.AlibabamjmosfundcreatebillAPIRequest, session string) (*mos.AlibabamjmosfundcreatebillAPIResponse, error) {
	var resp mos.AlibabamjmosfundcreatebillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
