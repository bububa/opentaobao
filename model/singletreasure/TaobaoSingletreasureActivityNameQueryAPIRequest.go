package singletreasure

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSingletreasureActivityNameQueryAPIRequest 查询官方的活动名称接口 API请求
// taobao.singletreasure.activity.name.query
//
// 查询官方的活动名称列表接口
type TaobaoSingletreasureActivityNameQueryAPIRequest struct {
	model.Params
}

// NewTaobaoSingletreasureActivityNameQueryRequest 初始化TaobaoSingletreasureActivityNameQueryAPIRequest对象
func NewTaobaoSingletreasureActivityNameQueryRequest() *TaobaoSingletreasureActivityNameQueryAPIRequest {
	return &TaobaoSingletreasureActivityNameQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSingletreasureActivityNameQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityNameQueryAPIRequest) GetApiMethodName() string {
	return "taobao.singletreasure.activity.name.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityNameQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSingletreasureActivityNameQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoSingletreasureActivityNameQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSingletreasureActivityNameQueryRequest()
	},
}

// GetTaobaoSingletreasureActivityNameQueryRequest 从 sync.Pool 获取 TaobaoSingletreasureActivityNameQueryAPIRequest
func GetTaobaoSingletreasureActivityNameQueryAPIRequest() *TaobaoSingletreasureActivityNameQueryAPIRequest {
	return poolTaobaoSingletreasureActivityNameQueryAPIRequest.Get().(*TaobaoSingletreasureActivityNameQueryAPIRequest)
}

// ReleaseTaobaoSingletreasureActivityNameQueryAPIRequest 将 TaobaoSingletreasureActivityNameQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoSingletreasureActivityNameQueryAPIRequest(v *TaobaoSingletreasureActivityNameQueryAPIRequest) {
	v.Reset()
	poolTaobaoSingletreasureActivityNameQueryAPIRequest.Put(v)
}
