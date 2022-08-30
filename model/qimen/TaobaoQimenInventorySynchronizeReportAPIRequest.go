package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenInventorySynchronizeReportAPIRequest 库存状态同步确认接口 API请求
// taobao.qimen.inventory.synchronize.report
//
// 库存状态同步确认接口
type TaobaoQimenInventorySynchronizeReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenInventorySynchronizeReportRequest
}

// NewTaobaoQimenInventorySynchronizeReportRequest 初始化TaobaoQimenInventorySynchronizeReportAPIRequest对象
func NewTaobaoQimenInventorySynchronizeReportRequest() *TaobaoQimenInventorySynchronizeReportAPIRequest {
	return &TaobaoQimenInventorySynchronizeReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventorySynchronizeReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.inventory.synchronize.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventorySynchronizeReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequest is Request Setter
//
func (r *TaobaoQimenInventorySynchronizeReportAPIRequest) SetRequest(_request *TaobaoQimenInventorySynchronizeReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenInventorySynchronizeReportAPIRequest) GetRequest() *TaobaoQimenInventorySynchronizeReportRequest {
	return r._request
}
