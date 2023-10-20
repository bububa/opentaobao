package uscesl

import (
	"net/url"
	"sync"

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
	// 商家标识key
	_bizBrandKey string
	// 价签系统注册的门店storeId
	_storeId int64
}

// NewTaobaoUsceslBizEslUnbindRequest 初始化TaobaoUsceslBizEslUnbindAPIRequest对象
func NewTaobaoUsceslBizEslUnbindRequest() *TaobaoUsceslBizEslUnbindAPIRequest {
	return &TaobaoUsceslBizEslUnbindAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUsceslBizEslUnbindAPIRequest) Reset() {
	r._eslBarCode = ""
	r._bizBrandKey = ""
	r._storeId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetApiMethodName() string {
	return "taobao.uscesl.biz.esl.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEslBarCode is EslBarCode Setter
// 价签条码
func (r *TaobaoUsceslBizEslUnbindAPIRequest) SetEslBarCode(_eslBarCode string) error {
	r._eslBarCode = _eslBarCode
	r.Set("esl_bar_code", _eslBarCode)
	return nil
}

// GetEslBarCode EslBarCode Getter
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetEslBarCode() string {
	return r._eslBarCode
}

// SetBizBrandKey is BizBrandKey Setter
// 商家标识key
func (r *TaobaoUsceslBizEslUnbindAPIRequest) SetBizBrandKey(_bizBrandKey string) error {
	r._bizBrandKey = _bizBrandKey
	r.Set("biz_brand_key", _bizBrandKey)
	return nil
}

// GetBizBrandKey BizBrandKey Getter
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetBizBrandKey() string {
	return r._bizBrandKey
}

// SetStoreId is StoreId Setter
// 价签系统注册的门店storeId
func (r *TaobaoUsceslBizEslUnbindAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoUsceslBizEslUnbindAPIRequest) GetStoreId() int64 {
	return r._storeId
}

var poolTaobaoUsceslBizEslUnbindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUsceslBizEslUnbindRequest()
	},
}

// GetTaobaoUsceslBizEslUnbindRequest 从 sync.Pool 获取 TaobaoUsceslBizEslUnbindAPIRequest
func GetTaobaoUsceslBizEslUnbindAPIRequest() *TaobaoUsceslBizEslUnbindAPIRequest {
	return poolTaobaoUsceslBizEslUnbindAPIRequest.Get().(*TaobaoUsceslBizEslUnbindAPIRequest)
}

// ReleaseTaobaoUsceslBizEslUnbindAPIRequest 将 TaobaoUsceslBizEslUnbindAPIRequest 放入 sync.Pool
func ReleaseTaobaoUsceslBizEslUnbindAPIRequest(v *TaobaoUsceslBizEslUnbindAPIRequest) {
	v.Reset()
	poolTaobaoUsceslBizEslUnbindAPIRequest.Put(v)
}
