package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenitemlackreportAPIRequest 发货单缺货通知接口 API请求
// taobao.qimen.itemlack.report
//
// WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoqimenitemlackreportAPIRequest struct {
	model.Params
	//
	_request *ItemLackReportRequest
}

// NewTaobaoqimenitemlackreportRequest 初始化TaobaoqimenitemlackreportAPIRequest对象
func NewTaobaoqimenitemlackreportRequest() *TaobaoqimenitemlackreportAPIRequest {
	return &TaobaoqimenitemlackreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenitemlackreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemlack.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenitemlackreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenitemlackreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenitemlackreportAPIRequest) SetRequest(_request *ItemLackReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenitemlackreportAPIRequest) GetRequest() *ItemLackReportRequest {
	return r._request
}
