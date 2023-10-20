package nrt

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtBrandinfoQueryAPIRequest 品牌数据查询 API请求
// tmall.nrt.brandinfo.query
//
// 商家获取自己旗舰店授权的品牌id列表
type TmallNrtBrandinfoQueryAPIRequest struct {
	model.Params
}

// NewTmallNrtBrandinfoQueryRequest 初始化TmallNrtBrandinfoQueryAPIRequest对象
func NewTmallNrtBrandinfoQueryRequest() *TmallNrtBrandinfoQueryAPIRequest {
	return &TmallNrtBrandinfoQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtBrandinfoQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtBrandinfoQueryAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.brandinfo.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtBrandinfoQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtBrandinfoQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallNrtBrandinfoQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtBrandinfoQueryRequest()
	},
}

// GetTmallNrtBrandinfoQueryRequest 从 sync.Pool 获取 TmallNrtBrandinfoQueryAPIRequest
func GetTmallNrtBrandinfoQueryAPIRequest() *TmallNrtBrandinfoQueryAPIRequest {
	return poolTmallNrtBrandinfoQueryAPIRequest.Get().(*TmallNrtBrandinfoQueryAPIRequest)
}

// ReleaseTmallNrtBrandinfoQueryAPIRequest 将 TmallNrtBrandinfoQueryAPIRequest 放入 sync.Pool
func ReleaseTmallNrtBrandinfoQueryAPIRequest(v *TmallNrtBrandinfoQueryAPIRequest) {
	v.Reset()
	poolTmallNrtBrandinfoQueryAPIRequest.Put(v)
}
