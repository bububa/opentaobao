package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenstockoutconfirmAPIRequest 出库单确认接口 API请求
// taobao.qimen.stockout.confirm
//
// 货品出库后，WMS将状态回传给ERP
type TaobaoqimenstockoutconfirmAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenstockoutconfirmStruct
}

// NewTaobaoqimenstockoutconfirmRequest 初始化TaobaoqimenstockoutconfirmAPIRequest对象
func NewTaobaoqimenstockoutconfirmRequest() *TaobaoqimenstockoutconfirmAPIRequest {
	return &TaobaoqimenstockoutconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenstockoutconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.stockout.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenstockoutconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenstockoutconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenstockoutconfirmAPIRequest) SetRequest(_request *TaobaoqimenstockoutconfirmStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenstockoutconfirmAPIRequest) GetRequest() *TaobaoqimenstockoutconfirmStruct {
	return r._request
}
