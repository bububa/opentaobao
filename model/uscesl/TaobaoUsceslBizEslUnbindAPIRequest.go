package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousceslbizeslunbindAPIRequest 电子价签解绑接口 API请求
// taobao.uscesl.biz.esl.unbind
//
// 电子价签解绑接口
type TaobaousceslbizeslunbindAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 商家标识key
	_bizBrandKey string
	// 价签系统注册的门店storeId
	_storeId int64
}

// NewTaobaousceslbizeslunbindRequest 初始化TaobaousceslbizeslunbindAPIRequest对象
func NewTaobaousceslbizeslunbindRequest() *TaobaousceslbizeslunbindAPIRequest {
	return &TaobaousceslbizeslunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousceslbizeslunbindAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.esl.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousceslbizeslunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousceslbizeslunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEslBarCode is EslBarCode Setter
// 价签条码
func (r *TaobaousceslbizeslunbindAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// GetEslBarCode EslBarCode Getter
func (r TaobaousceslbizeslunbindAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// SetBizBrandKey is BizBrandKey Setter
// 商家标识key
func (r *TaobaousceslbizeslunbindAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaousceslbizeslunbindAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 价签系统注册的门店storeId
func (r *TaobaousceslbizeslunbindAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaousceslbizeslunbindAPIRequest) GetStoreId() int64 {
	return r._storeId
}
