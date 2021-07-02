package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// AlibabaServicecenterFulfiltaskBuyeraddressChange 修改消费者服务地址
// alibaba.servicecenter.fulfiltask.buyeraddress.change
//
// 当消费者反馈自己的服务地址错误时，可以电话联系服务商修改为正确地址，服务商只能修改派给自己的单子
func AlibabaServicecenterFulfiltaskBuyeraddressChange(clt *core.SDKClient, req *tmallservice.AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIRequest, session string) (*tmallservice.AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponse, error) {
	var resp tmallservice.AlibabaServicecenterFulfiltaskBuyeraddressChangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
