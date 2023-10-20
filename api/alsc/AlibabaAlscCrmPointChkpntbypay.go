package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmpointchkpntbypay 校验支付链路中的积分抵扣是否合法
// alibaba.alsc.crm.point.chkpntbypay
//
// 校验支付链路中的积分抵扣是否合法
func Alibabaalsccrmpointchkpntbypay(clt *core.SDKClient, req *alsc.AlibabaalsccrmpointchkpntbypayAPIRequest, session string) (*alsc.AlibabaalsccrmpointchkpntbypayAPIResponse, error) {
	var resp alsc.AlibabaalsccrmpointchkpntbypayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
