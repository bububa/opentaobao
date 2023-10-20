package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinflowalipaypublish 流量钱包流量发放-面向支付宝用户
// alibaba.aliqin.flow.alipay.publish
//
// 用户淘宝流量钱包商家给支付宝用户发放流量
func Alibabaaliqinflowalipaypublish(clt *core.SDKClient, req *alicom.AlibabaaliqinflowalipaypublishAPIRequest, session string) (*alicom.AlibabaaliqinflowalipaypublishAPIResponse, error) {
	var resp alicom.AlibabaaliqinflowalipaypublishAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
