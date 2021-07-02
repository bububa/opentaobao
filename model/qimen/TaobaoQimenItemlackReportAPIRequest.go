package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenItemlackReportAPIRequest 发货单缺货通知接口 API请求
// taobao.qimen.itemlack.report
//
// WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
type TaobaoQimenItemlackReportAPIRequest struct {
	model.Params
	//
	_request *ItemLackReportRequest
}

// NewTaobaoQimenItemlackReportRequest 初始化TaobaoQimenItemlackReportAPIRequest对象
func NewTaobaoQimenItemlackReportRequest() *TaobaoQimenItemlackReportAPIRequest {
	return &TaobaoQimenItemlackReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemlackReportAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.itemlack.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemlackReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRequest is Request Setter
//
func (r *TaobaoQimenItemlackReportAPIRequest) SetRequest(_request *ItemLackReportRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenItemlackReportAPIRequest) GetRequest() *ItemLackReportRequest {
	return r._request
}
