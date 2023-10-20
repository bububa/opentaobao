package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeMetroSyncAPIRequest 地铁数据同步 API请求
// alibaba.alihouse.newhome.metro.sync
//
// 地铁数据同步
type AlibabaAlihouseNewhomeMetroSyncAPIRequest struct {
	model.Params
	// 地铁入参数据
	_baseMetroLineDto *BaseMetroLineDto
}

// NewAlibabaAlihouseNewhomeMetroSyncRequest 初始化AlibabaAlihouseNewhomeMetroSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeMetroSyncRequest() *AlibabaAlihouseNewhomeMetroSyncAPIRequest {
	return &AlibabaAlihouseNewhomeMetroSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeMetroSyncAPIRequest) Reset() {
	r._baseMetroLineDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.metro.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseMetroLineDto is BaseMetroLineDto Setter
// 地铁入参数据
func (r *AlibabaAlihouseNewhomeMetroSyncAPIRequest) SetBaseMetroLineDto(_baseMetroLineDto *BaseMetroLineDto) error {
	r._baseMetroLineDto = _baseMetroLineDto
	r.Set("base_metro_line_dto", _baseMetroLineDto)
	return nil
}

// GetBaseMetroLineDto BaseMetroLineDto Getter
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetBaseMetroLineDto() *BaseMetroLineDto {
	return r._baseMetroLineDto
}

var poolAlibabaAlihouseNewhomeMetroSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeMetroSyncRequest()
	},
}

// GetAlibabaAlihouseNewhomeMetroSyncRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeMetroSyncAPIRequest
func GetAlibabaAlihouseNewhomeMetroSyncAPIRequest() *AlibabaAlihouseNewhomeMetroSyncAPIRequest {
	return poolAlibabaAlihouseNewhomeMetroSyncAPIRequest.Get().(*AlibabaAlihouseNewhomeMetroSyncAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeMetroSyncAPIRequest 将 AlibabaAlihouseNewhomeMetroSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeMetroSyncAPIRequest(v *AlibabaAlihouseNewhomeMetroSyncAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeMetroSyncAPIRequest.Put(v)
}
