package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest 新房户型变更E码 API请求
// alibaba.alihouse.newhome.layout.ecode.update
//
// 新房户型变更E码
type AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest struct {
	model.Params
	// 数据
	_layout *UpdateProjectLayoutEcodeDto
}

// NewAlibabaAlihouseNewhomeLayoutEcodeUpdateRequest 初始化AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest对象
func NewAlibabaAlihouseNewhomeLayoutEcodeUpdateRequest() *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest {
	return &AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.layout.ecode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLayout is Layout Setter
// 数据
func (r *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) SetLayout(_layout *UpdateProjectLayoutEcodeDto) error {
	r._layout = _layout
	r.Set("layout", _layout)
	return nil
}

// GetLayout Layout Getter
func (r AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) GetLayout() *UpdateProjectLayoutEcodeDto {
	return r._layout
}
