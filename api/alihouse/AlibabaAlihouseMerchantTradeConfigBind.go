package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousemerchanttradeconfigbind 交易场景绑定
// alibaba.alihouse.merchant.trade.config.bind
//
// 交易场景绑定
func Alibabaalihousemerchanttradeconfigbind(clt *core.SDKClient, req *alihouse.AlibabaalihousemerchanttradeconfigbindAPIRequest, session string) (*alihouse.AlibabaalihousemerchanttradeconfigbindAPIResponse, error) {
	var resp alihouse.AlibabaalihousemerchanttradeconfigbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
