package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberpointoperateAPIRequest 积分操作 API请求
// alibaba.member.point.operate
//
// 消费会员积分
type AlibabamemberpointoperateAPIRequest struct {
	model.Params
	// 请求参数
	_pointOperateRequest *PointOperateDto
}

// NewAlibabamemberpointoperateRequest 初始化AlibabamemberpointoperateAPIRequest对象
func NewAlibabamemberpointoperateRequest() *AlibabamemberpointoperateAPIRequest {
	return &AlibabamemberpointoperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamemberpointoperateAPIRequest) GetApiMethodName() string {
	return "alibaba.member.point.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamemberpointoperateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamemberpointoperateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPointOperateRequest is PointOperateRequest Setter
// 请求参数
func (r *AlibabamemberpointoperateAPIRequest) SetPointOperateRequest(_pointOperateRequest *PointOperateDto) error {
	r._pointOperateRequest = _pointOperateRequest
	r.Set("point_operate_request", _pointOperateRequest)
	return nil
}

// GetPointOperateRequest PointOperateRequest Getter
func (r AlibabamemberpointoperateAPIRequest) GetPointOperateRequest() *PointOperateDto {
	return r._pointOperateRequest
}
