package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoProductToChannelImportAPIRequest 产品导入到渠道 API请求
// taobao.fenxiao.product.to.channel.import
//
// 支持供应商将已有产品导入到某个渠道销售
type TaobaoFenxiaoProductToChannelImportAPIRequest struct {
	model.Params
	// 要导入的渠道[21 零售PLUS]目前仅支持此渠道
	_channel int64
	// 要导入的产品id
	_productId int64
}

// NewTaobaoFenxiaoProductToChannelImportRequest 初始化TaobaoFenxiaoProductToChannelImportAPIRequest对象
func NewTaobaoFenxiaoProductToChannelImportRequest() *TaobaoFenxiaoProductToChannelImportAPIRequest {
	return &TaobaoFenxiaoProductToChannelImportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFenxiaoProductToChannelImportAPIRequest) Reset() {
	r._channel = 0
	r._productId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.to.channel.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChannel is Channel Setter
// 要导入的渠道[21 零售PLUS]目前仅支持此渠道
func (r *TaobaoFenxiaoProductToChannelImportAPIRequest) SetChannel(_channel int64) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// GetChannel Channel Getter
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetChannel() int64 {
	return r._channel
}

// SetProductId is ProductId Setter
// 要导入的产品id
func (r *TaobaoFenxiaoProductToChannelImportAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetProductId() int64 {
	return r._productId
}

var poolTaobaoFenxiaoProductToChannelImportAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFenxiaoProductToChannelImportRequest()
	},
}

// GetTaobaoFenxiaoProductToChannelImportRequest 从 sync.Pool 获取 TaobaoFenxiaoProductToChannelImportAPIRequest
func GetTaobaoFenxiaoProductToChannelImportAPIRequest() *TaobaoFenxiaoProductToChannelImportAPIRequest {
	return poolTaobaoFenxiaoProductToChannelImportAPIRequest.Get().(*TaobaoFenxiaoProductToChannelImportAPIRequest)
}

// ReleaseTaobaoFenxiaoProductToChannelImportAPIRequest 将 TaobaoFenxiaoProductToChannelImportAPIRequest 放入 sync.Pool
func ReleaseTaobaoFenxiaoProductToChannelImportAPIRequest(v *TaobaoFenxiaoProductToChannelImportAPIRequest) {
	v.Reset()
	poolTaobaoFenxiaoProductToChannelImportAPIRequest.Put(v)
}
