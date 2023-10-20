package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductReleaseAPIRequest 供应商铺货 API请求
// tmall.supplychain.channel.product.release
//
// 供应商渠道铺货接口
type TmallSupplychainChannelProductReleaseAPIRequest struct {
	model.Params
	// 产品数字ID
	_productId int64
	// 渠道ID
	_channelCode int64
}

// NewTmallSupplychainChannelProductReleaseRequest 初始化TmallSupplychainChannelProductReleaseAPIRequest对象
func NewTmallSupplychainChannelProductReleaseRequest() *TmallSupplychainChannelProductReleaseAPIRequest {
	return &TmallSupplychainChannelProductReleaseAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallSupplychainChannelProductReleaseAPIRequest) Reset() {
	r._productId = 0
	r._channelCode = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallSupplychainChannelProductReleaseAPIRequest) GetApiMethodName() string {
	return "tmall.supplychain.channel.product.release"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallSupplychainChannelProductReleaseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallSupplychainChannelProductReleaseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProductId is ProductId Setter
// 产品数字ID
func (r *TmallSupplychainChannelProductReleaseAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TmallSupplychainChannelProductReleaseAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetChannelCode is ChannelCode Setter
// 渠道ID
func (r *TmallSupplychainChannelProductReleaseAPIRequest) SetChannelCode(_channelCode int64) error {
	r._channelCode = _channelCode
	r.Set("channel_code", _channelCode)
	return nil
}

// GetChannelCode ChannelCode Getter
func (r TmallSupplychainChannelProductReleaseAPIRequest) GetChannelCode() int64 {
	return r._channelCode
}

var poolTmallSupplychainChannelProductReleaseAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallSupplychainChannelProductReleaseRequest()
	},
}

// GetTmallSupplychainChannelProductReleaseRequest 从 sync.Pool 获取 TmallSupplychainChannelProductReleaseAPIRequest
func GetTmallSupplychainChannelProductReleaseAPIRequest() *TmallSupplychainChannelProductReleaseAPIRequest {
	return poolTmallSupplychainChannelProductReleaseAPIRequest.Get().(*TmallSupplychainChannelProductReleaseAPIRequest)
}

// ReleaseTmallSupplychainChannelProductReleaseAPIRequest 将 TmallSupplychainChannelProductReleaseAPIRequest 放入 sync.Pool
func ReleaseTmallSupplychainChannelProductReleaseAPIRequest(v *TmallSupplychainChannelProductReleaseAPIRequest) {
	v.Reset()
	poolTmallSupplychainChannelProductReleaseAPIRequest.Put(v)
}
