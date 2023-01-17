package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeBusinessSyncAPIRequest 商圈数据同步 API请求
// alibaba.alihouse.newhome.business.sync
//
// 商圈数据同步
type AlibabaAlihouseNewhomeBusinessSyncAPIRequest struct {
	model.Params
	// 入参数据
	_baseBusinessDto *BaseBusinessDto
}

// NewAlibabaAlihouseNewhomeBusinessSyncRequest 初始化AlibabaAlihouseNewhomeBusinessSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeBusinessSyncRequest() *AlibabaAlihouseNewhomeBusinessSyncAPIRequest {
	return &AlibabaAlihouseNewhomeBusinessSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeBusinessSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.business.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeBusinessSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeBusinessSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseBusinessDto is BaseBusinessDto Setter
// 入参数据
func (r *AlibabaAlihouseNewhomeBusinessSyncAPIRequest) SetBaseBusinessDto(_baseBusinessDto *BaseBusinessDto) error {
	r._baseBusinessDto = _baseBusinessDto
	r.Set("base_business_dto", _baseBusinessDto)
	return nil
}

// GetBaseBusinessDto BaseBusinessDto Getter
func (r AlibabaAlihouseNewhomeBusinessSyncAPIRequest) GetBaseBusinessDto() *BaseBusinessDto {
	return r._baseBusinessDto
}
