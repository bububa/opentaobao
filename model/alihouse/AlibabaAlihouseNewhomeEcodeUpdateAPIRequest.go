package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeEcodeUpdateAPIRequest 新房货变更E码 API请求
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
type AlibabaAlihouseNewhomeEcodeUpdateAPIRequest struct {
	model.Params
	// 房源请求体
	_house *UpdateNewHomeECodeInfoDto
}

// NewAlibabaAlihouseNewhomeEcodeUpdateRequest 初始化AlibabaAlihouseNewhomeEcodeUpdateAPIRequest对象
func NewAlibabaAlihouseNewhomeEcodeUpdateRequest() *AlibabaAlihouseNewhomeEcodeUpdateAPIRequest {
	return &AlibabaAlihouseNewhomeEcodeUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) Reset() {
	r._house = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.ecode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouse is House Setter
// 房源请求体
func (r *AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) SetHouse(_house *UpdateNewHomeECodeInfoDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) GetHouse() *UpdateNewHomeECodeInfoDto {
	return r._house
}

var poolAlibabaAlihouseNewhomeEcodeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeEcodeUpdateRequest()
	},
}

// GetAlibabaAlihouseNewhomeEcodeUpdateRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeEcodeUpdateAPIRequest
func GetAlibabaAlihouseNewhomeEcodeUpdateAPIRequest() *AlibabaAlihouseNewhomeEcodeUpdateAPIRequest {
	return poolAlibabaAlihouseNewhomeEcodeUpdateAPIRequest.Get().(*AlibabaAlihouseNewhomeEcodeUpdateAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeEcodeUpdateAPIRequest 将 AlibabaAlihouseNewhomeEcodeUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeEcodeUpdateAPIRequest(v *AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeEcodeUpdateAPIRequest.Put(v)
}
