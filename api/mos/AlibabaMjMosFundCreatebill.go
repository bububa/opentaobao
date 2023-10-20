package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMosFundCreatebill 创建一个付款单
// alibaba.mj.mos.fund.createbill
//
// 创建一个付款单
func AlibabaMjMosFundCreatebill(clt *core.SDKClient, req *mos.AlibabaMjMosFundCreatebillAPIRequest, resp *mos.AlibabaMjMosFundCreatebillAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
