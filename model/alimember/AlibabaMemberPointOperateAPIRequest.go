package alimember

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberPointOperateAPIRequest) GetApiMethodName() string {
	return "alibaba.member.point.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberPointOperateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
