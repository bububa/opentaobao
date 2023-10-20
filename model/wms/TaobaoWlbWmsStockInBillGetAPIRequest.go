package wms

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbWmsStockInBillGetAPIRequest 获取入库单信息 API请求
// taobao.wlb.wms.stock.in.bill.get
//
// 获取入库单信息
type TaobaoWlbWmsStockInBillGetAPIRequest struct {
	model.Params
	// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_orderCode string
	// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
	_cnOrderCode string
}

// NewTaobaoWlbWmsStockInBillGetRequest 初始化TaobaoWlbWmsStockInBillGetAPIRequest对象
func NewTaobaoWlbWmsStockInBillGetRequest() *TaobaoWlbWmsStockInBillGetAPIRequest {
	return &TaobaoWlbWmsStockInBillGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbWmsStockInBillGetAPIRequest) Reset() {
	r._orderCode = ""
	r._cnOrderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWmsStockInBillGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.wms.stock.in.bill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWmsStockInBillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbWmsStockInBillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// ERP订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockInBillGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbWmsStockInBillGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

// SetCnOrderCode is CnOrderCode Setter
// 菜鸟订单编码，查询单个订单时orderCode与cnOrderCode必须有一个参数值不为空，两个参数都赋值时，以cnOrderCode值检索数据
func (r *TaobaoWlbWmsStockInBillGetAPIRequest) SetCnOrderCode(_cnOrderCode string) error {
	r._cnOrderCode = _cnOrderCode
	r.Set("cn_order_code", _cnOrderCode)
	return nil
}

// GetCnOrderCode CnOrderCode Getter
func (r TaobaoWlbWmsStockInBillGetAPIRequest) GetCnOrderCode() string {
	return r._cnOrderCode
}

var poolTaobaoWlbWmsStockInBillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbWmsStockInBillGetRequest()
	},
}

// GetTaobaoWlbWmsStockInBillGetRequest 从 sync.Pool 获取 TaobaoWlbWmsStockInBillGetAPIRequest
func GetTaobaoWlbWmsStockInBillGetAPIRequest() *TaobaoWlbWmsStockInBillGetAPIRequest {
	return poolTaobaoWlbWmsStockInBillGetAPIRequest.Get().(*TaobaoWlbWmsStockInBillGetAPIRequest)
}

// ReleaseTaobaoWlbWmsStockInBillGetAPIRequest 将 TaobaoWlbWmsStockInBillGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbWmsStockInBillGetAPIRequest(v *TaobaoWlbWmsStockInBillGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbWmsStockInBillGetAPIRequest.Put(v)
}
