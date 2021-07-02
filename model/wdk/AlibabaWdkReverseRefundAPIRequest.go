package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseRefundAPIRequest 退款打款 API请求
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
type AlibabaWdkReverseRefundAPIRequest struct {
	model.Params
	// 退款打款请求
	_openRefundReqDTO *OpenRefundReqDto
}

// NewAlibabaWdkReverseRefundRequest 初始化AlibabaWdkReverseRefundAPIRequest对象
func NewAlibabaWdkReverseRefundRequest() *AlibabaWdkReverseRefundAPIRequest {
	return &AlibabaWdkReverseRefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkReverseRefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkReverseRefundAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OpenRefundReqDTO Setter
// 退款打款请求
func (r *AlibabaWdkReverseRefundAPIRequest) SetOpenRefundReqDTO(_openRefundReqDTO *OpenRefundReqDto) error {
	r._openRefundReqDTO = _openRefundReqDTO
	r.Set("open_refund_req_d_t_o", _openRefundReqDTO)
	return nil
}

// Get OpenRefundReqDTO Getter
func (r AlibabaWdkReverseRefundAPIRequest) GetOpenRefundReqDTO() *OpenRefundReqDto {
	return r._openRefundReqDTO
}
