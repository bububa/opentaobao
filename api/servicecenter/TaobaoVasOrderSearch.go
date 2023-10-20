package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Taobaovasordersearch 订单记录导出
// taobao.vas.order.search
//
// 用于ISV查询自己名下的应用及收费项目的订单记录（已付款订单）。&lt;br/&gt;建议用于查询前一日的历史记录，不适合用作实时数据查询。&lt;br/&gt;现在只能查询90天以内的数据&lt;br/&gt;该接口限制每分钟所有appkey调用总和只能有800次。
func Taobaovasordersearch(clt *core.SDKClient, req *servicecenter.TaobaovasordersearchAPIRequest, session string) (*servicecenter.TaobaovasordersearchAPIResponse, error) {
	var resp servicecenter.TaobaovasordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
