package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeEcodeUpdateAPIRequest 新房货变更E码 API请求
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
type AlibabaAlihouseNewhomeEcodeUpdateAPIRequest struct {
	model.Params
	// 房源请求体
	_house *UpdateNewHomeEcodeInfoDto
}

// NewAlibabaAlihouseNewhomeEcodeUpdateRequest 初始化AlibabaAlihouseNewhomeEcodeUpdateAPIRequest对象
func NewAlibabaAlihouseNewhomeEcodeUpdateRequest() *AlibabaAlihouseNewhomeEcodeUpdateAPIRequest {
	return &AlibabaAlihouseNewhomeEcodeUpdateAPIRequest{
		Params: model.NewParams(),
	}
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
func (r *AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) SetHouse(_house *UpdateNewHomeEcodeInfoDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaAlihouseNewhomeEcodeUpdateAPIRequest) GetHouse() *UpdateNewHomeEcodeInfoDto {
	return r._house
}
