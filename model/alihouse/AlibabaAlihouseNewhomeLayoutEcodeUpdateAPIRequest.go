package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomelayoutecodeupdateAPIRequest 新房户型变更E码 API请求
// alibaba.alihouse.newhome.layout.ecode.update
//
// 新房户型变更E码
type AlibabaalihousenewhomelayoutecodeupdateAPIRequest struct {
	model.Params
	// 数据
	_layout *UpdateProjectLayoutEcodeDto
}

// NewAlibabaalihousenewhomelayoutecodeupdateRequest 初始化AlibabaalihousenewhomelayoutecodeupdateAPIRequest对象
func NewAlibabaalihousenewhomelayoutecodeupdateRequest() *AlibabaalihousenewhomelayoutecodeupdateAPIRequest {
	return &AlibabaalihousenewhomelayoutecodeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomelayoutecodeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.layout.ecode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomelayoutecodeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomelayoutecodeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLayout is Layout Setter
// 数据
func (r *AlibabaalihousenewhomelayoutecodeupdateAPIRequest) SetLayout(_layout *UpdateProjectLayoutEcodeDto) error {
	r._layout = _layout
	r.Set("layout", _layout)
	return nil
}

// GetLayout Layout Getter
func (r AlibabaalihousenewhomelayoutecodeupdateAPIRequest) GetLayout() *UpdateProjectLayoutEcodeDto {
	return r._layout
}
