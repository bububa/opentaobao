package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabatianjidistributorordersubmit 分销商提交受理订单
// alibaba.tianji.distributor.order.submit
//
// 分销商提交受理订单，如合约订购、充值受理等
func Alibabatianjidistributorordersubmit(clt *core.SDKClient, req *alicom.AlibabatianjidistributorordersubmitAPIRequest, session string) (*alicom.AlibabatianjidistributorordersubmitAPIResponse, error) {
	var resp alicom.AlibabatianjidistributorordersubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
