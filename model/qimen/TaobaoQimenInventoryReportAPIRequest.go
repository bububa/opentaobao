package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeninventoryreportAPIRequest 库存盘点通知接口 API请求
// taobao.qimen.inventory.report
//
// WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoqimeninventoryreportAPIRequest struct {
	model.Params
	//
	_request *InventoryReportRequest
}

// NewTaobaoqimeninventoryreportRequest 初始化TaobaoqimeninventoryreportAPIRequest对象
func NewTaobaoqimeninventoryreportRequest() *TaobaoqimeninventoryreportAPIRequest {
	return &TaobaoqimeninventoryreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeninventoryreportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeninventoryreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeninventoryreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimeninventoryreportAPIRequest) SetRequest(_request *InventoryReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimeninventoryreportAPIRequest) GetRequest() *InventoryReportRequest {
	return r._request
}
