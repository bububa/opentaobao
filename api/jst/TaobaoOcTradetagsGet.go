package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

// TaobaoOcTradetagsGet 根据订单查询订单标签
// taobao.oc.tradetags.get
//
// 根据订单查询订单标签。&lt;br/&gt;
// 返回的tag说明:1为官方标，2为自定义标，3为主站只读标签。&lt;br/&gt;
// 官方标签和自定义标签请看taobao.oc.tradetag.attach 接口说明&lt;br/&gt;
// 主站只读标签请看:http://open.taobao.com/doc/detail.htm?id=102865&lt;br/&gt;
func TaobaoOcTradetagsGet(clt *core.SDKClient, req *jst.TaobaoOcTradetagsGetAPIRequest, resp *jst.TaobaoOcTradetagsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
