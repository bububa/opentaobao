package fenxiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryQueryAPIRequest 查询商品库存信息 API请求
// taobao.inventory.query
//
// 建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
// 商家查询商品总体库存信息
type TaobaoInventoryQueryAPIRequest struct {
	model.Params
	// 后端商品ID 列表，控制到50个
	_scItemIds string
	// 后端商品的商家编码列表，控制到50个
	_scItemCodes string
	// 卖家昵称
	_sellerNick string
	// 仓库列表:GLY001^GLY002
	_storeCodes string
}

// NewTaobaoInventoryQueryRequest 初始化TaobaoInventoryQueryAPIRequest对象
func NewTaobaoInventoryQueryRequest() *TaobaoInventoryQueryAPIRequest {
	return &TaobaoInventoryQueryAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryQueryAPIRequest) Reset() {
	r._scItemIds = ""
	r._scItemCodes = ""
	r._sellerNick = ""
	r._storeCodes = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryQueryAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScItemIds is ScItemIds Setter
// 后端商品ID 列表，控制到50个
func (r *TaobaoInventoryQueryAPIRequest) SetScItemIds(_scItemIds string) error {
	r._scItemIds = _scItemIds
	r.Set("sc_item_ids", _scItemIds)
	return nil
}

// GetScItemIds ScItemIds Getter
func (r TaobaoInventoryQueryAPIRequest) GetScItemIds() string {
	return r._scItemIds
}

// SetScItemCodes is ScItemCodes Setter
// 后端商品的商家编码列表，控制到50个
func (r *TaobaoInventoryQueryAPIRequest) SetScItemCodes(_scItemCodes string) error {
	r._scItemCodes = _scItemCodes
	r.Set("sc_item_codes", _scItemCodes)
	return nil
}

// GetScItemCodes ScItemCodes Getter
func (r TaobaoInventoryQueryAPIRequest) GetScItemCodes() string {
	return r._scItemCodes
}

// SetSellerNick is SellerNick Setter
// 卖家昵称
func (r *TaobaoInventoryQueryAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoInventoryQueryAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetStoreCodes is StoreCodes Setter
// 仓库列表:GLY001^GLY002
func (r *TaobaoInventoryQueryAPIRequest) SetStoreCodes(_storeCodes string) error {
	r._storeCodes = _storeCodes
	r.Set("store_codes", _storeCodes)
	return nil
}

// GetStoreCodes StoreCodes Getter
func (r TaobaoInventoryQueryAPIRequest) GetStoreCodes() string {
	return r._storeCodes
}

var poolTaobaoInventoryQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryQueryRequest()
	},
}

// GetTaobaoInventoryQueryRequest 从 sync.Pool 获取 TaobaoInventoryQueryAPIRequest
func GetTaobaoInventoryQueryAPIRequest() *TaobaoInventoryQueryAPIRequest {
	return poolTaobaoInventoryQueryAPIRequest.Get().(*TaobaoInventoryQueryAPIRequest)
}

// ReleaseTaobaoInventoryQueryAPIRequest 将 TaobaoInventoryQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryQueryAPIRequest(v *TaobaoInventoryQueryAPIRequest) {
	v.Reset()
	poolTaobaoInventoryQueryAPIRequest.Put(v)
}
