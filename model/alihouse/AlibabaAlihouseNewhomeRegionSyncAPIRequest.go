package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeRegionSyncAPIRequest 城区数据同步 API请求
// alibaba.alihouse.newhome.region.sync
//
// 城区数据同步
type AlibabaAlihouseNewhomeRegionSyncAPIRequest struct {
	model.Params
	// 城区数据
	_baseRegionDto *BaseRegionDto
}

// NewAlibabaAlihouseNewhomeRegionSyncRequest 初始化AlibabaAlihouseNewhomeRegionSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeRegionSyncRequest() *AlibabaAlihouseNewhomeRegionSyncAPIRequest {
	return &AlibabaAlihouseNewhomeRegionSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeRegionSyncAPIRequest) Reset() {
	r._baseRegionDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeRegionSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.region.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeRegionSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeRegionSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseRegionDto is BaseRegionDto Setter
// 城区数据
func (r *AlibabaAlihouseNewhomeRegionSyncAPIRequest) SetBaseRegionDto(_baseRegionDto *BaseRegionDto) error {
	r._baseRegionDto = _baseRegionDto
	r.Set("base_region_dto", _baseRegionDto)
	return nil
}

// GetBaseRegionDto BaseRegionDto Getter
func (r AlibabaAlihouseNewhomeRegionSyncAPIRequest) GetBaseRegionDto() *BaseRegionDto {
	return r._baseRegionDto
}

var poolAlibabaAlihouseNewhomeRegionSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeRegionSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeRegionSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeRegionSyncAPIRequest
func GetAlibabaAlihouseNewhomeRegionSyncAPIRequest() *AlibabaAlihouseNewhomeRegionSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeRegionSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeRegionSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeRegionSyncAPIRequest 将 AlibabaAlihouseNewhomeRegionSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeRegionSyncAPIRequest(v *AlibabaAlihouseNewhomeRegionSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeRegionSyncAPIRequest.Put(v)
}
