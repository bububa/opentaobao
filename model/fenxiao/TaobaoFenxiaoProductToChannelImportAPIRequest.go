package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFenxiaoProductToChannelImportAPIRequest
产品导入到渠道 API请求
taobao.fenxiao.product.to.channel.import

支持供应商将已有产品导入到某个渠道销售 */
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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.product.to.channel.import"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Channel Setter
// 要导入的渠道[21 零售PLUS]目前仅支持此渠道
func (r *TaobaoFenxiaoProductToChannelImportAPIRequest) SetChannel(_channel int64) error {
	r._channel = _channel
	r.Set("channel", _channel)
	return nil
}

// Get Channel Getter
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetChannel() int64 {
	return r._channel
}

// Set is ProductId Setter
// 要导入的产品id
func (r *TaobaoFenxiaoProductToChannelImportAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// Get ProductId Getter
func (r TaobaoFenxiaoProductToChannelImportAPIRequest) GetProductId() int64 {
	return r._productId
}
