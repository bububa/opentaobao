package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeCommunitySubmitAPIRequest 提交小区信息 API请求
// alibaba.alihouse.newhome.community.submit
//
// 提交小区信息
type AlibabaAlihouseNewhomeCommunitySubmitAPIRequest struct {
	model.Params
	// 小区对象
	_ebbasCommunityDto *EbbasCommunityDto
}

// NewAlibabaAlihouseNewhomeCommunitySubmitRequest 初始化AlibabaAlihouseNewhomeCommunitySubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeCommunitySubmitRequest() *AlibabaAlihouseNewhomeCommunitySubmitAPIRequest {
	return &AlibabaAlihouseNewhomeCommunitySubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.community.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetEbbasCommunityDto is EbbasCommunityDto Setter
// 小区对象
func (r *AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) SetEbbasCommunityDto(_ebbasCommunityDto *EbbasCommunityDto) error {
	r._ebbasCommunityDto = _ebbasCommunityDto
	r.Set("ebbas_community_dto", _ebbasCommunityDto)
	return nil
}

// GetEbbasCommunityDto EbbasCommunityDto Getter
func (r AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) GetEbbasCommunityDto() *EbbasCommunityDto {
	return r._ebbasCommunityDto
}
