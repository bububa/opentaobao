package alihouse

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) Reset() {
	r._ebbasCommunityDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.community.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihouseNewhomeCommunitySubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeCommunitySubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeCommunitySubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeCommunitySubmitAPIRequest
func GetAlibabaAlihouseNewhomeCommunitySubmitAPIRequest() *AlibabaAlihouseNewhomeCommunitySubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeCommunitySubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeCommunitySubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeCommunitySubmitAPIRequest 将 AlibabaAlihouseNewhomeCommunitySubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeCommunitySubmitAPIRequest(v *AlibabaAlihouseNewhomeCommunitySubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeCommunitySubmitAPIRequest.Put(v)
}
