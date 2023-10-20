package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenreturnpackagereportAPIRequest 退货包裹状态通知接口 API请求
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
type TaobaoqimenreturnpackagereportAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenreturnpackagereportRequest
}

// NewTaobaoqimenreturnpackagereportRequest 初始化TaobaoqimenreturnpackagereportAPIRequest对象
func NewTaobaoqimenreturnpackagereportRequest() *TaobaoqimenreturnpackagereportAPIRequest {
	return &TaobaoqimenreturnpackagereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenreturnpackagereportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.returnpackage.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenreturnpackagereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenreturnpackagereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenreturnpackagereportAPIRequest) SetRequest(_request *TaobaoqimenreturnpackagereportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenreturnpackagereportAPIRequest) GetRequest() *TaobaoqimenreturnpackagereportRequest {
	return r._request
}
