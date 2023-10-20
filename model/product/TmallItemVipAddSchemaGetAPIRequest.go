package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemVipAddSchemaGetAPIRequest vip商家发布商品的获取规则接口 API请求
// tmall.item.vip.add.schema.get
//
// 获取vip商家发布商品的规则
type TmallItemVipAddSchemaGetAPIRequest struct {
	model.Params
}

// NewTmallItemVipAddSchemaGetRequest 初始化TmallItemVipAddSchemaGetAPIRequest对象
func NewTmallItemVipAddSchemaGetRequest() *TmallItemVipAddSchemaGetAPIRequest {
	return &TmallItemVipAddSchemaGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallItemVipAddSchemaGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallItemVipAddSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.item.vip.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallItemVipAddSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallItemVipAddSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallItemVipAddSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallItemVipAddSchemaGetRequest()
	},
}

// GetTmallItemVipAddSchemaGetRequest 从 sync.Pool 获取 TmallItemVipAddSchemaGetAPIRequest
func GetTmallItemVipAddSchemaGetAPIRequest() *TmallItemVipAddSchemaGetAPIRequest {
	return poolTmallItemVipAddSchemaGetAPIRequest.Get().(*TmallItemVipAddSchemaGetAPIRequest)
}

// ReleaseTmallItemVipAddSchemaGetAPIRequest 将 TmallItemVipAddSchemaGetAPIRequest 放入 sync.Pool
func ReleaseTmallItemVipAddSchemaGetAPIRequest(v *TmallItemVipAddSchemaGetAPIRequest) {
	v.Reset()
	poolTmallItemVipAddSchemaGetAPIRequest.Put(v)
}
