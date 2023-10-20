package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoinventoryqueryAPIRequest 查询商品库存信息 API请求
// taobao.inventory.query
//
// 建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
// 商家查询商品总体库存信息
type TaobaoinventoryqueryAPIRequest struct {
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

// NewTaobaoinventoryqueryRequest 初始化TaobaoinventoryqueryAPIRequest对象
func NewTaobaoinventoryqueryRequest() *TaobaoinventoryqueryAPIRequest {
	return &TaobaoinventoryqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoinventoryqueryAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoinventoryqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoinventoryqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetScItemIds is ScItemIds Setter
// 后端商品ID 列表，控制到50个
func (r *TaobaoinventoryqueryAPIRequest) SetScItemIds(_scItemIds string) error {
	r._scItemIds = _scItemIds
	r.Set("sc_item_ids", _scItemIds)
	return nil
}

// GetScItemIds ScItemIds Getter
func (r TaobaoinventoryqueryAPIRequest) GetScItemIds() string {
	return r._scItemIds
}

// SetScItemCodes is ScItemCodes Setter
// 后端商品的商家编码列表，控制到50个
func (r *TaobaoinventoryqueryAPIRequest) SetScItemCodes(_scItemCodes string) error {
	r._scItemCodes = _scItemCodes
	r.Set("sc_item_codes", _scItemCodes)
	return nil
}

// GetScItemCodes ScItemCodes Getter
func (r TaobaoinventoryqueryAPIRequest) GetScItemCodes() string {
	return r._scItemCodes
}

// SetSellerNick is SellerNick Setter
// 卖家昵称
func (r *TaobaoinventoryqueryAPIRequest) SetSellerNick(_sellerNick string) error {
	r._sellerNick = _sellerNick
	r.Set("seller_nick", _sellerNick)
	return nil
}

// GetSellerNick SellerNick Getter
func (r TaobaoinventoryqueryAPIRequest) GetSellerNick() string {
	return r._sellerNick
}

// SetStoreCodes is StoreCodes Setter
// 仓库列表:GLY001^GLY002
func (r *TaobaoinventoryqueryAPIRequest) SetStoreCodes(_storeCodes string) error {
	r._storeCodes = _storeCodes
	r.Set("store_codes", _storeCodes)
	return nil
}

// GetStoreCodes StoreCodes Getter
func (r TaobaoinventoryqueryAPIRequest) GetStoreCodes() string {
	return r._storeCodes
}
