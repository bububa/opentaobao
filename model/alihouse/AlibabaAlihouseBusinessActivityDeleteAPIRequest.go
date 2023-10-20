package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousebusinessactivitydeleteAPIRequest 电商券活动删除 API请求
// alibaba.alihouse.business.activity.delete
//
// 电商券活动删除
type AlibabaalihousebusinessactivitydeleteAPIRequest struct {
	model.Params
	// 数据
	_baiYiActivityDataDto *BusinessActivityDataDto
}

// NewAlibabaalihousebusinessactivitydeleteRequest 初始化AlibabaalihousebusinessactivitydeleteAPIRequest对象
func NewAlibabaalihousebusinessactivitydeleteRequest() *AlibabaalihousebusinessactivitydeleteAPIRequest {
	return &AlibabaalihousebusinessactivitydeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousebusinessactivitydeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.business.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousebusinessactivitydeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousebusinessactivitydeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaiYiActivityDataDto is BaiYiActivityDataDto Setter
// 数据
func (r *AlibabaalihousebusinessactivitydeleteAPIRequest) SetBaiYiActivityDataDto(_baiYiActivityDataDto *BusinessActivityDataDto) error {
	r._baiYiActivityDataDto = _baiYiActivityDataDto
	r.Set("bai_yi_activity_data_dto", _baiYiActivityDataDto)
	return nil
}

// GetBaiYiActivityDataDto BaiYiActivityDataDto Getter
func (r AlibabaalihousebusinessactivitydeleteAPIRequest) GetBaiYiActivityDataDto() *BusinessActivityDataDto {
	return r._baiYiActivityDataDto
}
