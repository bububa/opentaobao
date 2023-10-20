package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderStoreSdtquerystationAPIRequest 速店通查询站点信息 API请求
// taobao.omniorder.store.sdtquerystation
//
// 速店通查询站点信息
type TaobaoOmniorderStoreSdtquerystationAPIRequest struct {
	model.Params
	// 取号时返回的packageId
	_paramLong2 int64
}

// NewTaobaoOmniorderStoreSdtquerystationRequest 初始化TaobaoOmniorderStoreSdtquerystationAPIRequest对象
func NewTaobaoOmniorderStoreSdtquerystationRequest() *TaobaoOmniorderStoreSdtquerystationAPIRequest {
	return &TaobaoOmniorderStoreSdtquerystationAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniorderStoreSdtquerystationAPIRequest) Reset() {
	r._paramLong2 = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtquerystationAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.sdtquerystation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtquerystationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniorderStoreSdtquerystationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamLong2 is ParamLong2 Setter
// 取号时返回的packageId
func (r *TaobaoOmniorderStoreSdtquerystationAPIRequest) SetParamLong2(_paramLong2 int64) error {
	r._paramLong2 = _paramLong2
	r.Set("param_long2", _paramLong2)
	return nil
}

// GetParamLong2 ParamLong2 Getter
func (r TaobaoOmniorderStoreSdtquerystationAPIRequest) GetParamLong2() int64 {
	return r._paramLong2
}

var poolTaobaoOmniorderStoreSdtquerystationAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniorderStoreSdtquerystationRequest()
	},
}

// GetTaobaoOmniorderStoreSdtquerystationRequest 从 sync.Pool 获取 TaobaoOmniorderStoreSdtquerystationAPIRequest
func GetTaobaoOmniorderStoreSdtquerystationAPIRequest() *TaobaoOmniorderStoreSdtquerystationAPIRequest {
	return poolTaobaoOmniorderStoreSdtquerystationAPIRequest.Get().(*TaobaoOmniorderStoreSdtquerystationAPIRequest)
}

// ReleaseTaobaoOmniorderStoreSdtquerystationAPIRequest 将 TaobaoOmniorderStoreSdtquerystationAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniorderStoreSdtquerystationAPIRequest(v *TaobaoOmniorderStoreSdtquerystationAPIRequest) {
	v.Reset()
	poolTaobaoOmniorderStoreSdtquerystationAPIRequest.Put(v)
}
