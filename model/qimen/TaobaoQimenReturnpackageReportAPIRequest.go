package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenReturnpackageReportAPIRequest 退货包裹状态通知接口 API请求
// taobao.qimen.returnpackage.report
//
// 退货包裹状态通知接口
type TaobaoQimenReturnpackageReportAPIRequest struct {
	model.Params
	//
	_request *TaobaoQimenReturnpackageReportRequest
}

// NewTaobaoQimenReturnpackageReportRequest 初始化TaobaoQimenReturnpackageReportAPIRequest对象
func NewTaobaoQimenReturnpackageReportRequest() *TaobaoQimenReturnpackageReportAPIRequest {
	return &TaobaoQimenReturnpackageReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnpackageReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.returnpackage.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnpackageReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenReturnpackageReportAPIRequest) SetRequest(_request *TaobaoQimenReturnpackageReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenReturnpackageReportAPIRequest) GetRequest() *TaobaoQimenReturnpackageReportRequest {
	return r._request
}
