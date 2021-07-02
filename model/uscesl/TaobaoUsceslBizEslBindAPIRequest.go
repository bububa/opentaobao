package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizEslBindAPIRequest 电子价签绑定接口 API请求
// taobao.uscesl.biz.esl.bind
//
// 电子价签商品绑定接口
type TaobaoUsceslBizEslBindAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 商品条码
	_itemBarCode string
	// 门店storeId
	_storeId int64
	// 商家ID
	_bizBrandKey string
	// 额外扩展信息
	_extendInfo string
}

// NewTaobaoUsceslBizEslBindRequest 初始化TaobaoUsceslBizEslBindAPIRequest对象
func NewTaobaoUsceslBizEslBindRequest() *TaobaoUsceslBizEslBindAPIRequest {
	return &TaobaoUsceslBizEslBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizEslBindAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.esl.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizEslBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizEslBindAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// Get EslBarCode Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// Set is ItemBarCode Setter
// 商品条码
func (r *TaobaoUsceslBizEslBindAPIRequest) SetItemBarCode(_itemBarCode string) error {
	r._itemBarCode = _itemBarCode
	r.Set("item_bar_code", _itemBarCode)
	return nil
}

// Get ItemBarCode Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetItemBarCode() string {
	return r._itemBarCode
}

// Set is StoreId Setter
// 门店storeId
func (r *TaobaoUsceslBizEslBindAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is BizBrandKey Setter
// 商家ID
func (r *TaobaoUsceslBizEslBindAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// Get BizBrandKey Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// Set is ExtendInfo Setter
// 额外扩展信息
func (r *TaobaoUsceslBizEslBindAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// Get ExtendInfo Getter
func (r TaobaoUsceslBizEslBindAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}
