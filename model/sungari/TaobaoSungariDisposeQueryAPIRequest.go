package sungari

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSungariDisposeQueryAPIRequest 商品商家处置结果查询 API请求
// taobao.sungari.dispose.query
//
// 红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
type TaobaoSungariDisposeQueryAPIRequest struct {
	model.Params
	// 查询的key列表
	_paramList []string
}

// NewTaobaoSungariDisposeQueryRequest 初始化TaobaoSungariDisposeQueryAPIRequest对象
func NewTaobaoSungariDisposeQueryRequest() *TaobaoSungariDisposeQueryAPIRequest {
	return &TaobaoSungariDisposeQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSungariDisposeQueryAPIRequest) Reset() {
	r._paramList = r._paramList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSungariDisposeQueryAPIRequest) GetApiMethodName() string {
	return "taobao.sungari.dispose.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSungariDisposeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSungariDisposeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamList is ParamList Setter
// 查询的key列表
func (r *TaobaoSungariDisposeQueryAPIRequest) SetParamList(_paramList []string) error {
	r._paramList = _paramList
	r.Set("param_list", _paramList)
	return nil
}

// GetParamList ParamList Getter
func (r TaobaoSungariDisposeQueryAPIRequest) GetParamList() []string {
	return r._paramList
}

var poolTaobaoSungariDisposeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSungariDisposeQueryRequest()
	},
}

// GetTaobaoSungariDisposeQueryRequest 从 sync.Pool 获取 TaobaoSungariDisposeQueryAPIRequest
func GetTaobaoSungariDisposeQueryAPIRequest() *TaobaoSungariDisposeQueryAPIRequest {
	return poolTaobaoSungariDisposeQueryAPIRequest.Get().(*TaobaoSungariDisposeQueryAPIRequest)
}

// ReleaseTaobaoSungariDisposeQueryAPIRequest 将 TaobaoSungariDisposeQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoSungariDisposeQueryAPIRequest(v *TaobaoSungariDisposeQueryAPIRequest) {
	v.Reset()
	poolTaobaoSungariDisposeQueryAPIRequest.Put(v)
}
