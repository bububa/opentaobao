package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkreverserefundAPIRequest 退款打款 API请求
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
type AlibabawdkreverserefundAPIRequest struct {
	model.Params
	// 退款打款请求
	_openRefundReqDTO *OpenRefundReqDto
}

// NewAlibabawdkreverserefundRequest 初始化AlibabawdkreverserefundAPIRequest对象
func NewAlibabawdkreverserefundRequest() *AlibabawdkreverserefundAPIRequest {
	return &AlibabawdkreverserefundAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkreverserefundAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.reverse.refund"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkreverserefundAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkreverserefundAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenRefundReqDTO is OpenRefundReqDTO Setter
// 退款打款请求
func (r *AlibabawdkreverserefundAPIRequest) SetOpenRefundReqDTO(_openRefundReqDTO *OpenRefundReqDto) error {
	r._openRefundReqDTO = _openRefundReqDTO
	r.Set("open_refund_req_d_t_o", _openRefundReqDTO)
	return nil
}

// GetOpenRefundReqDTO OpenRefundReqDTO Getter
func (r AlibabawdkreverserefundAPIRequest) GetOpenRefundReqDTO() *OpenRefundReqDto {
	return r._openRefundReqDTO
}
