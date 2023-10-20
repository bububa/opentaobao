package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenwarehouseinfoqueryAPIRequest 货主仓库资源查询接口 API请求
// taobao.qimen.warehouseinfo.query
//
// 货主仓库资源查询
type TaobaoqimenwarehouseinfoqueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenwarehouseinfoqueryRequest
}

// NewTaobaoqimenwarehouseinfoqueryRequest 初始化TaobaoqimenwarehouseinfoqueryAPIRequest对象
func NewTaobaoqimenwarehouseinfoqueryRequest() *TaobaoqimenwarehouseinfoqueryAPIRequest {
	return &TaobaoqimenwarehouseinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenwarehouseinfoqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.warehouseinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenwarehouseinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenwarehouseinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenwarehouseinfoqueryAPIRequest) SetRequest(_request *TaobaoqimenwarehouseinfoqueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenwarehouseinfoqueryAPIRequest) GetRequest() *TaobaoqimenwarehouseinfoqueryRequest {
	return r._request
}
