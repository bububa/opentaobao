package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSupplierSynchronizeAPIRequest 供应商同步接口 API请求
// taobao.qimen.supplier.synchronize
//
// 这个接口用来同步供应商信息
type TaobaoQimenSupplierSynchronizeAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenSupplierSynchronizeRequest
}

// NewTaobaoQimenSupplierSynchronizeRequest 初始化TaobaoQimenSupplierSynchronizeAPIRequest对象
func NewTaobaoQimenSupplierSynchronizeRequest() *TaobaoQimenSupplierSynchronizeAPIRequest {
	return &TaobaoQimenSupplierSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenSupplierSynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.supplier.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenSupplierSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenSupplierSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenSupplierSynchronizeAPIRequest) SetRequest(_request *TaobaoQimenSupplierSynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenSupplierSynchronizeAPIRequest) GetRequest() *TaobaoQimenSupplierSynchronizeRequest {
	return r._request
}
