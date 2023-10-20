package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest 新房户型变更E码 API请求
// alibaba.alihouse.newhome.layout.ecode.update
//
// 新房户型变更E码
type AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest struct {
	model.Params
	// 数据
	_layout *UpdateProjectLayoutECodeDto
}

// NewAlibabaAlihouseNewhomeLayoutEcodeUpdateRequest 初始化AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest对象
func NewAlibabaAlihouseNewhomeLayoutEcodeUpdateRequest() *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest {
	return &AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) Reset() {
	r._layout = nil
	r.Params.ToZero()
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
func (r *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) SetLayout(_layout *UpdateProjectLayoutECodeDto) error {
	r._layout = _layout
	r.Set("layout", _layout)
	return nil
}

// GetLayout Layout Getter
func (r AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) GetLayout() *UpdateProjectLayoutECodeDto {
	return r._layout
}

var poolAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeLayoutEcodeUpdateRequest()
	},
}

// GetAlibabaAlihouseNewhomeLayoutEcodeUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest
func GetAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest() *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest {
	return poolAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest.Get().(*AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest 将 AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest(v *AlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeLayoutEcodeUpdateAPIRequest.Put(v)
}
