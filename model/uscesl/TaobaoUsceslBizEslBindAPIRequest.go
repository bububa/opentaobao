package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizeslbindAPIRequest 电子价签绑定接口 API请求
// taobao.uscesl.biz.esl.bind
//
// 电子价签商品绑定接口
type TaobaousceslbizeslbindAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 商品条码
	_itemBarCode string
	// 商家ID
	_bizBrandKey string
	// 额外扩展信息
	_extendInfo string
	// 门店storeId
	_storeId int64
}

// NewTaobaousceslbizeslbindRequest 初始化TaobaousceslbizeslbindAPIRequest对象
func NewTaobaousceslbizeslbindRequest() *TaobaousceslbizeslbindAPIRequest {
	return &TaobaousceslbizeslbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizeslbindAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.esl.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizeslbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizeslbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEslBarCode is EslBarCode Setter
// 价签条码
func (r *TaobaousceslbizeslbindAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// GetEslBarCode EslBarCode Getter
func (r TaobaousceslbizeslbindAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// SetItemBarCode is ItemBarCode Setter
// 商品条码
func (r *TaobaousceslbizeslbindAPIRequest) SetItemBarCode(_itemBarCode string) error {
	r._itemBarCode = _itemBarCode
	r.Set("item_bar_code", _itemBarCode)
	return nil
}

// GetItemBarCode ItemBarCode Getter
func (r TaobaousceslbizeslbindAPIRequest) GetItemBarCode() string {
	return r._itemBarCode
}

// SetBizBrandKey is BizBrandKey Setter
// 商家ID
func (r *TaobaousceslbizeslbindAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaousceslbizeslbindAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetExtendInfo is ExtendInfo Setter
// 额外扩展信息
func (r *TaobaousceslbizeslbindAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r TaobaousceslbizeslbindAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}

// SetStoreId is StoreId Setter
// 门店storeId
func (r *TaobaousceslbizeslbindAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaousceslbizeslbindAPIRequest) GetStoreId() int64 {
	return r._storeId
}
