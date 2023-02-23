package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeSupportSyncAPIRequest 周边配套数据同步 API请求
// alibaba.alihouse.newhome.support.sync
//
// 周边配套数据同步
type AlibabaAlihouseNewhomeSupportSyncAPIRequest struct {
	model.Params
	// 周边设施入参
	_baseSupportingDto *BaseSupportingDto
}

// NewAlibabaAlihouseNewhomeSupportSyncRequest 初始化AlibabaAlihouseNewhomeSupportSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeSupportSyncRequest() *AlibabaAlihouseNewhomeSupportSyncAPIRequest {
	return &AlibabaAlihouseNewhomeSupportSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeSupportSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.support.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeSupportSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeSupportSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseSupportingDto is BaseSupportingDto Setter
// 周边设施入参
func (r *AlibabaAlihouseNewhomeSupportSyncAPIRequest) SetBaseSupportingDto(_baseSupportingDto *BaseSupportingDto) error {
	r._baseSupportingDto = _baseSupportingDto
	r.Set("base_supporting_dto", _baseSupportingDto)
	return nil
}

// GetBaseSupportingDto BaseSupportingDto Getter
func (r AlibabaAlihouseNewhomeSupportSyncAPIRequest) GetBaseSupportingDto() *BaseSupportingDto {
	return r._baseSupportingDto
}
