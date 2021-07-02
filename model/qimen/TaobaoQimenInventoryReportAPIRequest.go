package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventoryReportAPIRequest 库存盘点通知接口 API请求
// taobao.qimen.inventory.report
//
// WMS调用奇门的接口,将库存盘点情况回传ERP
type TaobaoQimenInventoryReportAPIRequest struct {
	model.Params
	//
	_request *InventoryReportRequest
}

// NewTaobaoQimenInventoryReportRequest 初始化TaobaoQimenInventoryReportAPIRequest对象
func NewTaobaoQimenInventoryReportRequest() *TaobaoQimenInventoryReportAPIRequest {
	return &TaobaoQimenInventoryReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequest is Request Setter
//
func (r *TaobaoQimenInventoryReportAPIRequest) SetRequest(_request *InventoryReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventoryReportAPIRequest) GetRequest() *InventoryReportRequest {
	return r._request
}
