package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseMerchantTradeConfigBind 交易场景绑定
// alibaba.alihouse.merchant.trade.config.bind
//
// 交易场景绑定
func AlibabaAlihouseMerchantTradeConfigBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseMerchantTradeConfigBindAPIRequest, session string) (*alihouse.AlibabaAlihouseMerchantTradeConfigBindAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseMerchantTradeConfigBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
