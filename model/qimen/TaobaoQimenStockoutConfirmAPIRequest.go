package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenStockoutConfirmAPIRequest 出库单确认接口 API请求
// taobao.qimen.stockout.confirm
//
// 货品出库后，WMS将状态回传给ERP
type TaobaoQimenStockoutConfirmAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenStockoutConfirmStruct
}

// NewTaobaoQimenStockoutConfirmRequest 初始化TaobaoQimenStockoutConfirmAPIRequest对象
func NewTaobaoQimenStockoutConfirmRequest() *TaobaoQimenStockoutConfirmAPIRequest {
	return &TaobaoQimenStockoutConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenStockoutConfirmAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockoutConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stockout.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockoutConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenStockoutConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenStockoutConfirmAPIRequest) SetRequest(_request *TaobaoQimenStockoutConfirmStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenStockoutConfirmAPIRequest) GetRequest() *TaobaoQimenStockoutConfirmStruct {
	return r._request
}

var poolTaobaoQimenStockoutConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenStockoutConfirmRequest()
	},
}

// GetTaobaoQimenStockoutConfirmRequest 从 sync.Pool 获取 TaobaoQimenStockoutConfirmAPIRequest
func GetTaobaoQimenStockoutConfirmAPIRequest() *TaobaoQimenStockoutConfirmAPIRequest {
	return poolTaobaoQimenStockoutConfirmAPIRequest.Get().(*TaobaoQimenStockoutConfirmAPIRequest)
}

// ReleaseTaobaoQimenStockoutConfirmAPIRequest 将 TaobaoQimenStockoutConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenStockoutConfirmAPIRequest(v *TaobaoQimenStockoutConfirmAPIRequest) {
	v.Reset()
	poolTaobaoQimenStockoutConfirmAPIRequest.Put(v)
}
