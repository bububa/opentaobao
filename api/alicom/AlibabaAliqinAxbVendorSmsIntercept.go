package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinaxbvendorsmsintercept AXB短信托收推送接口
// alibaba.aliqin.axb.vendor.sms.intercept
//
// 用于给供应商推送需要托收的短信
func Alibabaaliqinaxbvendorsmsintercept(clt *core.SDKClient, req *alicom.AlibabaaliqinaxbvendorsmsinterceptAPIRequest, session string) (*alicom.AlibabaaliqinaxbvendorsmsinterceptAPIResponse, error) {
	var resp alicom.AlibabaaliqinaxbvendorsmsinterceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
