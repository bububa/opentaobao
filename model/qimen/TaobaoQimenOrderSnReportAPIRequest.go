package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenordersnreportAPIRequest 订单SN通知接口 API请求
// taobao.qimen.order.sn.report
//
// WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
type TaobaoqimenordersnreportAPIRequest struct {
	model.Params
	//
	_request *TaobaoqimenordersnreportRequest
}

// NewTaobaoqimenordersnreportRequest 初始化TaobaoqimenordersnreportAPIRequest对象
func NewTaobaoqimenordersnreportRequest() *TaobaoqimenordersnreportAPIRequest {
	return &TaobaoqimenordersnreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenordersnreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.sn.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenordersnreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenordersnreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenordersnreportAPIRequest) SetRequest(_request *TaobaoqimenordersnreportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenordersnreportAPIRequest) GetRequest() *TaobaoqimenordersnreportRequest {
	return r._request
}
