package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemstoreskuupdateAPIRequest 五道口商品中心门店商品修改 API请求
// alibaba.wdk.item.storesku.update
//
// 五道口商品中心门店商品修改
type AlibabawdkitemstoreskuupdateAPIRequest struct {
	model.Params
	// 盒马门店id
	_storeId string
	// 商品编码
	_skuCode string
	// 1-可售   0-不可售
	_saleFlag int64
}

// NewAlibabawdkitemstoreskuupdateRequest 初始化AlibabawdkitemstoreskuupdateAPIRequest对象
func NewAlibabawdkitemstoreskuupdateRequest() *AlibabawdkitemstoreskuupdateAPIRequest {
	return &AlibabawdkitemstoreskuupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemstoreskuupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.storesku.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemstoreskuupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemstoreskuupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreId is StoreId Setter
// 盒马门店id
func (r *AlibabawdkitemstoreskuupdateAPIRequest) SetStoreId(_storeId string) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r AlibabawdkitemstoreskuupdateAPIRequest) GetStoreId() string {
	return r._storeId
}

// SetSkuCode is SkuCode Setter
// 商品编码
func (r *AlibabawdkitemstoreskuupdateAPIRequest) SetSkuCode(_skuCode string) error {
	r._skuCode = _skuCode
	r.Set("sku_code", _skuCode)
	return nil
}

// GetSkuCode SkuCode Getter
func (r AlibabawdkitemstoreskuupdateAPIRequest) GetSkuCode() string {
	return r._skuCode
}

// SetSaleFlag is SaleFlag Setter
// 1-可售   0-不可售
func (r *AlibabawdkitemstoreskuupdateAPIRequest) SetSaleFlag(_saleFlag int64) error {
	r._saleFlag = _saleFlag
	r.Set("sale_flag", _saleFlag)
	return nil
}

// GetSaleFlag SaleFlag Getter
func (r AlibabawdkitemstoreskuupdateAPIRequest) GetSaleFlag() int64 {
	return r._saleFlag
}
