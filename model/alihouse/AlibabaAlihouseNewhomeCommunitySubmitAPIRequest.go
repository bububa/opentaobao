package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomecommunitysubmitAPIRequest 提交小区信息 API请求
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
type AlibabaalihousenewhomecommunitysubmitAPIRequest struct {
	model.Params
	// 小区对象
	_ebbasCommunityDto *EbbasCommunityDto
}

// NewAlibabaalihousenewhomecommunitysubmitRequest 初始化AlibabaalihousenewhomecommunitysubmitAPIRequest对象
func NewAlibabaalihousenewhomecommunitysubmitRequest() *AlibabaalihousenewhomecommunitysubmitAPIRequest {
	return &AlibabaalihousenewhomecommunitysubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomecommunitysubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.community.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomecommunitysubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomecommunitysubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEbbasCommunityDto is EbbasCommunityDto Setter
// 小区对象
func (r *AlibabaalihousenewhomecommunitysubmitAPIRequest) SetEbbasCommunityDto(_ebbasCommunityDto *EbbasCommunityDto) error {
	r._ebbasCommunityDto = _ebbasCommunityDto
	r.Set("ebbas_community_dto", _ebbasCommunityDto)
	return nil
}

// GetEbbasCommunityDto EbbasCommunityDto Getter
func (r AlibabaalihousenewhomecommunitysubmitAPIRequest) GetEbbasCommunityDto() *EbbasCommunityDto {
	return r._ebbasCommunityDto
}
