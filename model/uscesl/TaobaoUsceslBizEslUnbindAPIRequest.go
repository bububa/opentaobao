package uscesl

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizEslUnbindAPIRequest 电子价签解绑接口 API请求
// taobao.uscesl.biz.esl.unbind
//
// 电子价签解绑接口
type TaobaoUsceslBizEslUnbindAPIRequest struct {
	model.Params
	// 价签条码
	_eslBarCode string
	// 价签系统注册的门店storeId
	_storeId int64
	// 商家标识key
	_bizBrandKey string
}

// NewTaobaoUsceslBizEslUnbindRequest 初始化TaobaoUsceslBizEslUnbindAPIRequest对象
func NewTaobaoUsceslBizEslUnbindRequest() *TaobaoUsceslBizEslUnbindAPIRequest {
	return &TaobaoUsceslBizEslUnbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.esl.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizEslUnbindAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// Get EslBarCode Getter
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// Set is StoreId Setter
// 价签系统注册的门店storeId
func (r *TaobaoUsceslBizEslUnbindAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// Get StoreId Getter
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// Set is BizBrandKey Setter
// 商家标识key
func (r *TaobaoUsceslBizEslUnbindAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// Get BizBrandKey Getter
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}
