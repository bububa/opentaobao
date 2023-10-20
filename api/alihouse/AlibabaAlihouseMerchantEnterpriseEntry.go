package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousemerchantenterpriseentry 机构入驻
// alibaba.alihouse.merchant.enterprise.entry
//
// 机构入驻
func Alibabaalihousemerchantenterpriseentry(clt *core.SDKClient, req *alihouse.AlibabaalihousemerchantenterpriseentryAPIRequest, session string) (*alihouse.AlibabaalihousemerchantenterpriseentryAPIResponse, error) {
	var resp alihouse.AlibabaalihousemerchantenterpriseentryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
