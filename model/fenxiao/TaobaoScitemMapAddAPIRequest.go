package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoscitemmapaddAPIRequest 创建IC商品与后端商品的映射关系 API请求
// taobao.scitem.map.add
//
// 创建IC商品或分销商品与后端商品的映射关系
type TaobaoscitemmapaddAPIRequest struct {
	model.Params
	// sc_item_id和outer_code 其中一个不能为空
	_outerCode string
	// 前台ic商品id
	_itemId int64
	// 前台ic商品skuid
	_skuId int64
	// sc_item_id和outer_code 其中一个不能为空
	_scItemId int64
	// 默认值为false<br/>true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联<br/>false:不进行高级校验
	_needCheck bool
}

// NewTaobaoscitemmapaddRequest 初始化TaobaoscitemmapaddAPIRequest对象
func NewTaobaoscitemmapaddRequest() *TaobaoscitemmapaddAPIRequest {
	return &TaobaoscitemmapaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoscitemmapaddAPIRequest) GetApiMethodName() string {
	return "taobao.scitem.map.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoscitemmapaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoscitemmapaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCode is OuterCode Setter
// sc_item_id和outer_code 其中一个不能为空
func (r *TaobaoscitemmapaddAPIRequest) SetOuterCode(_outerCode string) error {
	r._outerCode = _outerCode
	r.Set("outer_code", _outerCode)
	return nil
}

// GetOuterCode OuterCode Getter
func (r TaobaoscitemmapaddAPIRequest) GetOuterCode() string {
	return r._outerCode
}

// SetItemId is ItemId Setter
// 前台ic商品id
func (r *TaobaoscitemmapaddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoscitemmapaddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetSkuId is SkuId Setter
// 前台ic商品skuid
func (r *TaobaoscitemmapaddAPIRequest) SetSkuId(_skuId int64) error {
	r._skuId = _skuId
	r.Set("sku_id", _skuId)
	return nil
}

// GetSkuId SkuId Getter
func (r TaobaoscitemmapaddAPIRequest) GetSkuId() int64 {
	return r._skuId
}

// SetScItemId is ScItemId Setter
// sc_item_id和outer_code 其中一个不能为空
func (r *TaobaoscitemmapaddAPIRequest) SetScItemId(_scItemId int64) error {
	r._scItemId = _scItemId
	r.Set("sc_item_id", _scItemId)
	return nil
}

// GetScItemId ScItemId Getter
func (r TaobaoscitemmapaddAPIRequest) GetScItemId() int64 {
	return r._scItemId
}

// SetNeedCheck is NeedCheck Setter
// 默认值为false&lt;br/&gt;true:进行高级校验,前端商品或SKU的商家编码必须与后端商品的商家编码一致，否则会拒绝关联&lt;br/&gt;false:不进行高级校验
func (r *TaobaoscitemmapaddAPIRequest) SetNeedCheck(_needCheck bool) error {
	r._needCheck = _needCheck
	r.Set("need_check", _needCheck)
	return nil
}

// GetNeedCheck NeedCheck Getter
func (r TaobaoscitemmapaddAPIRequest) GetNeedCheck() bool {
	return r._needCheck
}
