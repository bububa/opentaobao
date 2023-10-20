package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemAddSimpleschemaGetAPIRequest 天猫发布商品规则获取 API请求
// tmall.item.add.simpleschema.get
//
// 通过商家信息获取商品发布字段和规则。
type TmallItemAddSimpleschemaGetAPIRequest struct {
	model.Params
}

// NewTmallItemAddSimpleschemaGetRequest 初始化TmallItemAddSimpleschemaGetAPIRequest对象
func NewTmallItemAddSimpleschemaGetRequest() *TmallItemAddSimpleschemaGetAPIRequest {
	return &TmallItemAddSimpleschemaGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemAddSimpleschemaGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemAddSimpleschemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.add.simpleschema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemAddSimpleschemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemAddSimpleschemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallItemAddSimpleschemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemAddSimpleschemaGetRequest()
	},
}

// GetTmallItemAddSimpleschemaGetRequest 从 sync.Pool 获取 TmallItemAddSimpleschemaGetAPIRequest
func GetTmallItemAddSimpleschemaGetAPIRequest() *TmallItemAddSimpleschemaGetAPIRequest {
	return poolTmallItemAddSimpleschemaGetAPIRequest.Get().(*TmallItemAddSimpleschemaGetAPIRequest)
}

// ReleaseTmallItemAddSimpleschemaGetAPIRequest 将 TmallItemAddSimpleschemaGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemAddSimpleschemaGetAPIRequest(v *TmallItemAddSimpleschemaGetAPIRequest) {
	v.Reset()
	poolTmallItemAddSimpleschemaGetAPIRequest.Put(v)
}
