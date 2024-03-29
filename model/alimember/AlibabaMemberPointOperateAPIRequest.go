package alimember

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberPointOperateAPIRequest 积分操作 API请求
// alibaba.member.point.operate
//
// 消费会员积分
type AlibabaMemberPointOperateAPIRequest struct {
	model.Params
	// 请求参数
	_pointOperateRequest *PointOperateDto
}

// NewAlibabaMemberPointOperateRequest 初始化AlibabaMemberPointOperateAPIRequest对象
func NewAlibabaMemberPointOperateRequest() *AlibabaMemberPointOperateAPIRequest {
	return &AlibabaMemberPointOperateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberPointOperateAPIRequest) Reset() {
	r._pointOperateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberPointOperateAPIRequest) GetApiMethodName() string {
	return "alibaba.member.point.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberPointOperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberPointOperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPointOperateRequest is PointOperateRequest Setter
// 请求参数
func (r *AlibabaMemberPointOperateAPIRequest) SetPointOperateRequest(_pointOperateRequest *PointOperateDto) error {
	r._pointOperateRequest = _pointOperateRequest
	r.Set("point_operate_request", _pointOperateRequest)
	return nil
}

// GetPointOperateRequest PointOperateRequest Getter
func (r AlibabaMemberPointOperateAPIRequest) GetPointOperateRequest() *PointOperateDto {
	return r._pointOperateRequest
}

var poolAlibabaMemberPointOperateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberPointOperateRequest()
	},
}

// GetAlibabaMemberPointOperateRequest 从 sync.Pool 获取 AlibabaMemberPointOperateAPIRequest
func GetAlibabaMemberPointOperateAPIRequest() *AlibabaMemberPointOperateAPIRequest {
	return poolAlibabaMemberPointOperateAPIRequest.Get().(*AlibabaMemberPointOperateAPIRequest)
}

// ReleaseAlibabaMemberPointOperateAPIRequest 将 AlibabaMemberPointOperateAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberPointOperateAPIRequest(v *AlibabaMemberPointOperateAPIRequest) {
	v.Reset()
	poolAlibabaMemberPointOperateAPIRequest.Put(v)
}
