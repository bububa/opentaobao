package maitix

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest 分销项目列表查询（已过时，不推荐使用） API请求
// alibaba.damai.maitix.project.distribution.querylist
//
// 分销项目列表查询接口（已过时，不推荐使用）
type AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest struct {
	model.Params
}

// NewAlibabaDamaiMaitixProjectDistributionQuerylistRequest 初始化AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest对象
func NewAlibabaDamaiMaitixProjectDistributionQuerylistRequest() *AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest {
	return &AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest) GetApiMethodName() string {
	return "alibaba.damai.maitix.project.distribution.querylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDamaiMaitixProjectDistributionQuerylistRequest()
	},
}

// GetAlibabaDamaiMaitixProjectDistributionQuerylistRequest 从 sync.Pool 获取 AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest
func GetAlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest() *AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest {
	return poolAlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest.Get().(*AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest)
}

// ReleaseAlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest 将 AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest 放入 sync.Pool
func ReleaseAlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest(v *AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest) {
	v.Reset()
	poolAlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest.Put(v)
}
