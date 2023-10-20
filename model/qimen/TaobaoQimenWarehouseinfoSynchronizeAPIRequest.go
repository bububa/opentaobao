package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenwarehouseinfosynchronizeAPIRequest 仓库同步接口 API请求
// taobao.qimen.warehouseinfo.synchronize
//
// 仓库同步接口
type TaobaoqimenwarehouseinfosynchronizeAPIRequest struct {
	model.Params
	// 请求报文
	_request *TaobaoqimenwarehouseinfosynchronizeRequest
}

// NewTaobaoqimenwarehouseinfosynchronizeRequest 初始化TaobaoqimenwarehouseinfosynchronizeAPIRequest对象
func NewTaobaoqimenwarehouseinfosynchronizeRequest() *TaobaoqimenwarehouseinfosynchronizeAPIRequest {
	return &TaobaoqimenwarehouseinfosynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenwarehouseinfosynchronizeAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.warehouseinfo.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenwarehouseinfosynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenwarehouseinfosynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 请求报文
func (r *TaobaoqimenwarehouseinfosynchronizeAPIRequest) SetRequest(_request *TaobaoqimenwarehouseinfosynchronizeRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenwarehouseinfosynchronizeAPIRequest) GetRequest() *TaobaoqimenwarehouseinfosynchronizeRequest {
	return r._request
}
