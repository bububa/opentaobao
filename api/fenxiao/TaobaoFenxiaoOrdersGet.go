package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoOrdersGet 查询采购单信息
// taobao.fenxiao.orders.get
//
// 查询代销采购单单据。
//
// 1. 支持商家按照供应商、分销商两种角色来查询数据。如果没有指定角色角色，系统会自动判断，此时如果商家存在供应商、分销商两种角色时，按照供应商角色查询。
// 2. 同时此接口还可以查询除供销经销外的其他经营模式的数据。如果需要查询供销经销单据请参考接口：taobao.fenxiao.dealer.requisitionorder.query
//
// 3. 发货请调用物流API中的发货接口taobao.logistics.offline.send 进行发货，需要注意的是这里是供应商发货，因此调发货接口时需要传人供应商账号对应的sessionkey，tid 需传入供销平台的采购单（即fenxiao_id  分销流水号）)。
func TaobaoFenxiaoOrdersGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoOrdersGetAPIRequest, resp *fenxiao.TaobaoFenxiaoOrdersGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
