package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest 赠品绑赠回滚 API请求
// alibaba.dchain.aoxiang.consignorder.gift.reverse
//
// 赠品绑赠回滚
type AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest struct {
	model.Params
	// 赠品绑赠回滚入参
	_reverseConsignorderGiftRequest *ReverseConsignOrderGiftRequest
}

// NewAlibabaDchainAoxiangConsignorderGiftReverseRequest 初始化AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest对象
func NewAlibabaDchainAoxiangConsignorderGiftReverseRequest() *AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest {
	return &AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.gift.reverse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReverseConsignorderGiftRequest is ReverseConsignorderGiftRequest Setter
// 赠品绑赠回滚入参
func (r *AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) SetReverseConsignorderGiftRequest(_reverseConsignorderGiftRequest *ReverseConsignOrderGiftRequest) error {
	r._reverseConsignorderGiftRequest = _reverseConsignorderGiftRequest
	r.Set("reverse_consignorder_gift_request", _reverseConsignorderGiftRequest)
	return nil
}

// GetReverseConsignorderGiftRequest ReverseConsignorderGiftRequest Getter
func (r AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) GetReverseConsignorderGiftRequest() *ReverseConsignOrderGiftRequest {
	return r._reverseConsignorderGiftRequest
}
