package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenTransferorderReportAPIRequest
调拨单通知 API请求
taobao.qimen.transferorder.report

调拨单通知 */
type TaobaoQimenTransferorderReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenTransferorderReportStruct
}

// NewTaobaoQimenTransferorderReportRequest 初始化TaobaoQimenTransferorderReportAPIRequest对象
func NewTaobaoQimenTransferorderReportRequest() *TaobaoQimenTransferorderReportAPIRequest {
	return &TaobaoQimenTransferorderReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenTransferorderReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.transferorder.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenTransferorderReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenTransferorderReportAPIRequest) SetRequest(_request *TaobaoQimenTransferorderReportStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenTransferorderReportAPIRequest) GetRequest() *TaobaoQimenTransferorderReportStruct {
	return r._request
}
