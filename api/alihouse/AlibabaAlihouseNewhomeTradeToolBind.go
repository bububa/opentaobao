package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeToolBind 批量绑定交易工具
// alibaba.alihouse.newhome.trade.tool.bind
//
// 批量绑定交易工具
func AlibabaAlihouseNewhomeTradeToolBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeToolBindAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeTradeToolBindAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeTradeToolBindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
