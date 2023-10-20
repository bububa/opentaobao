package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkorderrefundget 五道口订单退款按ID查询
// alibaba.wdk.order.refund.get
//
// 按照退款ID或者五道口中台订单ID查询退款信息详情
func Alibabawdkorderrefundget(clt *core.SDKClient, req *wdk.AlibabawdkorderrefundgetAPIRequest, session string) (*wdk.AlibabawdkorderrefundgetAPIResponse, error) {
	var resp wdk.AlibabawdkorderrefundgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
