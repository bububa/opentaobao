package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMosFundCancelbill 取消付款单
// alibaba.mj.mos.fund.cancelbill
//
// 取消付款单
func AlibabaMjMosFundCancelbill(clt *core.SDKClient, req *mos.AlibabaMjMosFundCancelbillAPIRequest, session string) (*mos.AlibabaMjMosFundCancelbillAPIResponse, error) {
	var resp mos.AlibabaMjMosFundCancelbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
