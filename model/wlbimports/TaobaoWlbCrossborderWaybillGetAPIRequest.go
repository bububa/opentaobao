package wlbimports

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbCrossborderWaybillGetAPIRequest 集货商家pdf和云打印面单获取，pdf需要配置白名单 API请求
// taobao.wlb.crossborder.waybill.get
//
// 【TOF】欧洲供应商PDF格式电子面单渲染下发
//
//	需求链接：https://aone.alibaba-inc.com/req/21210808
type TaobaoWlbCrossborderWaybillGetAPIRequest struct {
	model.Params
	// 菜鸟物流单号
	_orderCode string
}

// NewTaobaoWlbCrossborderWaybillGetRequest 初始化TaobaoWlbCrossborderWaybillGetAPIRequest对象
func NewTaobaoWlbCrossborderWaybillGetRequest() *TaobaoWlbCrossborderWaybillGetAPIRequest {
	return &TaobaoWlbCrossborderWaybillGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbCrossborderWaybillGetAPIRequest) Reset() {
	r._orderCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbCrossborderWaybillGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.crossborder.waybill.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbCrossborderWaybillGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbCrossborderWaybillGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCode is OrderCode Setter
// 菜鸟物流单号
func (r *TaobaoWlbCrossborderWaybillGetAPIRequest) SetOrderCode(_orderCode string) error {
	r._orderCode = _orderCode
	r.Set("order_code", _orderCode)
	return nil
}

// GetOrderCode OrderCode Getter
func (r TaobaoWlbCrossborderWaybillGetAPIRequest) GetOrderCode() string {
	return r._orderCode
}

var poolTaobaoWlbCrossborderWaybillGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbCrossborderWaybillGetRequest()
	},
}

// GetTaobaoWlbCrossborderWaybillGetRequest 从 sync.Pool 获取 TaobaoWlbCrossborderWaybillGetAPIRequest
func GetTaobaoWlbCrossborderWaybillGetAPIRequest() *TaobaoWlbCrossborderWaybillGetAPIRequest {
	return poolTaobaoWlbCrossborderWaybillGetAPIRequest.Get().(*TaobaoWlbCrossborderWaybillGetAPIRequest)
}

// ReleaseTaobaoWlbCrossborderWaybillGetAPIRequest 将 TaobaoWlbCrossborderWaybillGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbCrossborderWaybillGetAPIRequest(v *TaobaoWlbCrossborderWaybillGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbCrossborderWaybillGetAPIRequest.Put(v)
}
