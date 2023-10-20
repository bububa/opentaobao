package wms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsInventoryProfitlossGetAPIRequest 通过订单列表批量获取库存损益单信息 API请求
// taobao.wlb.wms.inventory.profitloss.get
//
// 通过订单列表批量获取库存损益单信息
type TaobaoWlbWmsInventoryProfitlossGetAPIRequest struct {
	model.Params
	// 菜鸟订单编码
	_cnOrderCode string
}

// NewTaobaoWlbWmsInventoryProfitlossGetRequest 初始化TaobaoWlbWmsInventoryProfitlossGetAPIRequest对象
func NewTaobaoWlbWmsInventoryProfitlossGetRequest() *TaobaoWlbWmsInventoryProfitlossGetAPIRequest {
	return &TaobaoWlbWmsInventoryProfitlossGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWmsInventoryProfitlossGetAPIRequest) Reset() {
	r._cnOrderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsInventoryProfitlossGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.inventory.profitloss.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsInventoryProfitlossGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWmsInventoryProfitlossGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCnOrderCode is CnOrderCode Setter
// 菜鸟订单编码
func (r *TaobaoWlbWmsInventoryProfitlossGetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
	r._cnOrderCode = _cnOrderCode
	r.Set("cn_order_code", _cnOrderCode)
	return nil
}

// GetCnOrderCode CnOrderCode Getter
func (r TaobaoWlbWmsInventoryProfitlossGetAPIRequest) GetCnOrderCode() string {
	return r._cnOrderCode
}

var poolTaobaoWlbWmsInventoryProfitlossGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWmsInventoryProfitlossGetRequest()
	},
}

// GetTaobaoWlbWmsInventoryProfitlossGetRequest 从 sync.Pool 获取 TaobaoWlbWmsInventoryProfitlossGetAPIRequest
func GetTaobaoWlbWmsInventoryProfitlossGetAPIRequest() *TaobaoWlbWmsInventoryProfitlossGetAPIRequest {
	return poolTaobaoWlbWmsInventoryProfitlossGetAPIRequest.Get().(*TaobaoWlbWmsInventoryProfitlossGetAPIRequest)
}

// ReleaseTaobaoWlbWmsInventoryProfitlossGetAPIRequest 将 TaobaoWlbWmsInventoryProfitlossGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWmsInventoryProfitlossGetAPIRequest(v *TaobaoWlbWmsInventoryProfitlossGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbWmsInventoryProfitlossGetAPIRequest.Put(v)
}
