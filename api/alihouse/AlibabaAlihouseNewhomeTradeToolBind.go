package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeTradeToolBind 批量绑定交易工具
// alibaba.alihouse.newhome.trade.tool.bind
//
// 批量绑定交易工具
func AlibabaAlihouseNewhomeTradeToolBind(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeTradeToolBindAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeTradeToolBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
