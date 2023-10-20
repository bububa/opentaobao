package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenWarehouseinfoQueryAPIRequest 货主仓库资源查询接口 API请求
// taobao.qimen.warehouseinfo.query
//
// 货主仓库资源查询
type TaobaoQimenWarehouseinfoQueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenWarehouseinfoQueryRequest
}

// NewTaobaoQimenWarehouseinfoQueryRequest 初始化TaobaoQimenWarehouseinfoQueryAPIRequest对象
func NewTaobaoQimenWarehouseinfoQueryRequest() *TaobaoQimenWarehouseinfoQueryAPIRequest {
	return &TaobaoQimenWarehouseinfoQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenWarehouseinfoQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.warehouseinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenWarehouseinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenWarehouseinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenWarehouseinfoQueryAPIRequest) SetRequest(_request *TaobaoQimenWarehouseinfoQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenWarehouseinfoQueryAPIRequest) GetRequest() *TaobaoQimenWarehouseinfoQueryRequest {
	return r._request
}
