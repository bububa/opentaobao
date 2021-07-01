package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillIQuerydetailAPIRequest
查面单号状态v1.0 API请求
taobao.wlb.waybill.i.querydetail

查看面单号的当前状态，如签收、发货、失效等。 */
type TaobaoWlbWaybillIQuerydetailAPIRequest struct {
	model.Params
	// 面单查询请求
	_waybillDetailQueryRequest *WaybillDetailQueryRequest
}

// New
