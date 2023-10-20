package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktraceurlgetAPIRequest 溯源url透出 API请求
// alibaba.wdk.trace.url.get
//
// 根据shopId和skuCode返回商品溯源url
type AlibabawdktraceurlgetAPIRequest struct {
	model.Params
	// barCode 或者skuCode
	_scanCode string
	// 所属门店code
	_shopId string
	// 来源编码
	_sourceCode string
}

// NewAlibabawdktraceurlgetRequest 初始化AlibabawdktraceurlgetAPIRequest对象
func NewAlibabawdktraceurlgetRequest() *AlibabawdktraceurlgetAPIRequest {
	return &AlibabawdktraceurlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdktraceurlgetAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.trace.url.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdktraceurlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdktraceurlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScanCode is ScanCode Setter
// barCode 或者skuCode
func (r *AlibabawdktraceurlgetAPIRequest) SetScanCode(_scanCode string) error {
	r._scanCode = _scanCode
	r.Set("scan_code", _scanCode)
	return nil
}

// GetScanCode ScanCode Getter
func (r AlibabawdktraceurlgetAPIRequest) GetScanCode() string {
	return r._scanCode
}

// SetShopId is ShopId Setter
// 所属门店code
func (r *AlibabawdktraceurlgetAPIRequest) SetShopId(_shopId string) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r AlibabawdktraceurlgetAPIRequest) GetShopId() string {
	return r._shopId
}

// SetSourceCode is SourceCode Setter
// 来源编码
func (r *AlibabawdktraceurlgetAPIRequest) SetSourceCode(_sourceCode string) error {
	r._sourceCode = _sourceCode
	r.Set("source_code", _sourceCode)
	return nil
}

// GetSourceCode SourceCode Getter
func (r AlibabawdktraceurlgetAPIRequest) GetSourceCode() string {
	return r._sourceCode
}
