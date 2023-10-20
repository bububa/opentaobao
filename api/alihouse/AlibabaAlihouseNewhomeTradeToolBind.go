package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhometradetoolbind 批量绑定交易工具
// alibaba.alihouse.newhome.trade.tool.bind
//
// 批量绑定交易工具
func Alibabaalihousenewhometradetoolbind(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhometradetoolbindAPIRequest, session string) (*alihouse.AlibabaalihousenewhometradetoolbindAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhometradetoolbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
