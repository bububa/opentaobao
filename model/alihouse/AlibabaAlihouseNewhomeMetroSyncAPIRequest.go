package alihouse

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.metro.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeMetroSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
