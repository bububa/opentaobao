package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemtraceurlgetAPIRequest 根据shopId和skuCode返回商品静态溯源url API请求
// alibaba.wdk.item.trace.url.get
//
// 根据shopId和skuCode返回商品静态溯源url
type AlibabawdkitemtraceurlgetAPIRequest struct {
	model.Params
	// 商品skuCode
	_skuCode string
	// 所属门店code
	_shopId string
	// 来源编码
	_sourceCode string
}

// NewAlibabawdkitemtraceurlgetRequest 初始化AlibabawdkitemtraceurlgetAPIRequest对象
func NewAlibabawdkitemtraceurlgetRequest() *AlibabawdkitemtraceurlgetAPIRequest {
	return &AlibabawdkitemtraceurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemtraceurlgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.trace.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemtraceurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemtraceurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkuCode is SkuCode Setter
// 商品skuCode
func (r *AlibabawdkitemtraceurlgetAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemtraceurlgetAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetShopId is ShopId Setter
// 所属门店code
func (r *AlibabawdkitemtraceurlgetAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabawdkitemtraceurlgetAPIRequest) GetShopId() string {
	return r._shopId
}

// SetSourceCode is SourceCode Setter
// 来源编码
func (r *AlibabawdkitemtraceurlgetAPIRequest) SetSourceCode(_sourceCode string) error {
	r._sourceCode = _sourceCode
	r.Set("source_code", _sourceCode)
	return nil
}

// GetSourceCode SourceCode Getter
func (r AlibabawdkitemtraceurlgetAPIRequest) GetSourceCode() string {
	return r._sourceCode
}
