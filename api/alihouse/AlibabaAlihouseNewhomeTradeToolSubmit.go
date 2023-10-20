package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhometradetoolsubmit 交易工具信息上翻
// alibaba.alihouse.newhome.trade.tool.submit
//
// 交易工具信息上翻
func Alibabaalihousenewhometradetoolsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhometradetoolsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhometradetoolsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhometradetoolsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
