package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) Reset() {
	r._reverseConsignorderGiftRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.consignorder.gift.reverse"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaDchainAoxiangConsignorderGiftReverseAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangConsignorderGiftReverseRequest()
	},
}

// GetAlibabaDchainAoxiangConsignorderGiftReverseRequest 从 sync.Pool 获取 AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest
func GetAlibabaDchainAoxiangConsignorderGiftReverseAPIRequest() *AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest {
	return poolAlibabaDchainAoxiangConsignorderGiftReverseAPIRequest.Get().(*AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest)
}

// ReleaseAlibabaDchainAoxiangConsignorderGiftReverseAPIRequest 将 AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangConsignorderGiftReverseAPIRequest(v *AlibabaDchainAoxiangConsignorderGiftReverseAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangConsignorderGiftReverseAPIRequest.Put(v)
}
