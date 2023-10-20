package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterServiceTypeQueryallAPIRequest 服务供应链服务类型 API请求
// tmall.servicecenter.service.type.queryall
//
// 查询天猫服务类型列表
type TmallServicecenterServiceTypeQueryallAPIRequest struct {
	model.Params
}

// NewTmallServicecenterServiceTypeQueryallRequest 初始化TmallServicecenterServiceTypeQueryallAPIRequest对象
func NewTmallServicecenterServiceTypeQueryallRequest() *TmallServicecenterServiceTypeQueryallAPIRequest {
	return &TmallServicecenterServiceTypeQueryallAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterServiceTypeQueryallAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterServiceTypeQueryallAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.service.type.queryall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterServiceTypeQueryallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterServiceTypeQueryallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTmallServicecenterServiceTypeQueryallAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterServiceTypeQueryallRequest()
	},
}

// GetTmallServicecenterServiceTypeQueryallRequest 从 sync.Pool 获取 TmallServicecenterServiceTypeQueryallAPIRequest
func GetTmallServicecenterServiceTypeQueryallAPIRequest() *TmallServicecenterServiceTypeQueryallAPIRequest {
	return poolTmallServicecenterServiceTypeQueryallAPIRequest.Get().(*TmallServicecenterServiceTypeQueryallAPIRequest)
}

// ReleaseTmallServicecenterServiceTypeQueryallAPIRequest 将 TmallServicecenterServiceTypeQueryallAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterServiceTypeQueryallAPIRequest(v *TmallServicecenterServiceTypeQueryallAPIRequest) {
	v.Reset()
	poolTmallServicecenterServiceTypeQueryallAPIRequest.Put(v)
}
