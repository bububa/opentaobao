package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseMerchantTradeConfigBindAPIRequest 交易场景绑定 API请求
// alibaba.alihouse.merchant.trade.config.bind
//
// 交易场景绑定
type AlibabaAlihouseMerchantTradeConfigBindAPIRequest struct {
	model.Params
	// 进件资料对象
	_cis *TradeSceneAddInfoDto
}

// NewAlibabaAlihouseMerchantTradeConfigBindRequest 初始化AlibabaAlihouseMerchantTradeConfigBindAPIRequest对象
func NewAlibabaAlihouseMerchantTradeConfigBindRequest() *AlibabaAlihouseMerchantTradeConfigBindAPIRequest {
	return &AlibabaAlihouseMerchantTradeConfigBindAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseMerchantTradeConfigBindAPIRequest) Reset() {
	r._cis = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseMerchantTradeConfigBindAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.merchant.trade.config.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseMerchantTradeConfigBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseMerchantTradeConfigBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCis is Cis Setter
// 进件资料对象
func (r *AlibabaAlihouseMerchantTradeConfigBindAPIRequest) SetCis(_cis *TradeSceneAddInfoDto) error {
	r._cis = _cis
	r.Set("cis", _cis)
	return nil
}

// GetCis Cis Getter
func (r AlibabaAlihouseMerchantTradeConfigBindAPIRequest) GetCis() *TradeSceneAddInfoDto {
	return r._cis
}

var poolAlibabaAlihouseMerchantTradeConfigBindAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseMerchantTradeConfigBindRequest()
	},
}

// GetAlibabaAlihouseMerchantTradeConfigBindRequest 从 sync.Pool 获取 AlibabaAlihouseMerchantTradeConfigBindAPIRequest
func GetAlibabaAlihouseMerchantTradeConfigBindAPIRequest() *AlibabaAlihouseMerchantTradeConfigBindAPIRequest {
	return poolAlibabaAlihouseMerchantTradeConfigBindAPIRequest.Get().(*AlibabaAlihouseMerchantTradeConfigBindAPIRequest)
}

// ReleaseAlibabaAlihouseMerchantTradeConfigBindAPIRequest 将 AlibabaAlihouseMerchantTradeConfigBindAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseMerchantTradeConfigBindAPIRequest(v *AlibabaAlihouseMerchantTradeConfigBindAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseMerchantTradeConfigBindAPIRequest.Put(v)
}
