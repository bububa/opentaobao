package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenexpressinfoqueryAPIRequest 配送公司信息查询接口 API请求
// taobao.qimen.expressinfo.query
//
// 配送公司信息查询
type TaobaoqimenexpressinfoqueryAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenexpressinfoqueryRequest
}

// NewTaobaoqimenexpressinfoqueryRequest 初始化TaobaoqimenexpressinfoqueryAPIRequest对象
func NewTaobaoqimenexpressinfoqueryRequest() *TaobaoqimenexpressinfoqueryAPIRequest {
	return &TaobaoqimenexpressinfoqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenexpressinfoqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.expressinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenexpressinfoqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenexpressinfoqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenexpressinfoqueryAPIRequest) SetRequest(_request *TaobaoqimenexpressinfoqueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenexpressinfoqueryAPIRequest) GetRequest() *TaobaoqimenexpressinfoqueryRequest {
	return r._request
}
