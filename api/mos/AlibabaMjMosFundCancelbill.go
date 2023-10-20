package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMosFundCancelbill 取消付款单
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
func AlibabaMjMosFundCancelbill(clt *core.SDKClient, req *mos.AlibabaMjMosFundCancelbillAPIRequest, resp *mos.AlibabaMjMosFundCancelbillAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
