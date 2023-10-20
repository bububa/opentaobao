package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeecodeupdateAPIRequest 新房货变更E码 API请求
// alibaba.alihouse.newhome.ecode.update
//
// 新房货变更E码
type AlibabaalihousenewhomeecodeupdateAPIRequest struct {
	model.Params
	// 房源请求体
	_house *UpdateNewHomeEcodeInfoDto
}

// NewAlibabaalihousenewhomeecodeupdateRequest 初始化AlibabaalihousenewhomeecodeupdateAPIRequest对象
func NewAlibabaalihousenewhomeecodeupdateRequest() *AlibabaalihousenewhomeecodeupdateAPIRequest {
	return &AlibabaalihousenewhomeecodeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeecodeupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.ecode.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeecodeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeecodeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHouse is House Setter
// 房源请求体
func (r *AlibabaalihousenewhomeecodeupdateAPIRequest) SetHouse(_house *UpdateNewHomeEcodeInfoDto) error {
	r._house = _house
	r.Set("house", _house)
	return nil
}

// GetHouse House Getter
func (r AlibabaalihousenewhomeecodeupdateAPIRequest) GetHouse() *UpdateNewHomeEcodeInfoDto {
	return r._house
}
