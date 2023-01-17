package qimen

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
