package tmallcampus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCampusAuthstatusQueryAPIRequest 学生认证状态查询 API请求
// tmall.campus.authstatus.query
//
// 学生认证状态查询
type TmallCampusAuthstatusQueryAPIRequest struct {
	model.Params
}

// NewTmallCampusAuthstatusQueryRequest 初始化TmallCampusAuthstatusQueryAPIRequest对象
func NewTmallCampusAuthstatusQueryRequest() *TmallCampusAuthstatusQueryAPIRequest {
	return &TmallCampusAuthstatusQueryAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCampusAuthstatusQueryAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCampusAuthstatusQueryAPIRequest) GetApiMethodName() string {
	return "tmall.campus.authstatus.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCampusAuthstatusQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCampusAuthstatusQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallCampusAuthstatusQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCampusAuthstatusQueryRequest()
	},
}

// GetTmallCampusAuthstatusQueryRequest 从 sync.Pool 获取 TmallCampusAuthstatusQueryAPIRequest
func GetTmallCampusAuthstatusQueryAPIRequest() *TmallCampusAuthstatusQueryAPIRequest {
	return poolTmallCampusAuthstatusQueryAPIRequest.Get().(*TmallCampusAuthstatusQueryAPIRequest)
}

// ReleaseTmallCampusAuthstatusQueryAPIRequest 将 TmallCampusAuthstatusQueryAPIRequest 放入 sync.Pool
func ReleaseTmallCampusAuthstatusQueryAPIRequest(v *TmallCampusAuthstatusQueryAPIRequest) {
	v.Reset()
	poolTmallCampusAuthstatusQueryAPIRequest.Put(v)
}
