package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeLineSyncAPIRequest 环线数据同步 API请求
// alibaba.alihouse.newhome.line.sync
//
// 环线数据同步
type AlibabaAlihouseNewhomeLineSyncAPIRequest struct {
	model.Params
	// 环线入参
	_baseLoopLineDto *BaseLoopLineDto
}

// NewAlibabaAlihouseNewhomeLineSyncRequest 初始化AlibabaAlihouseNewhomeLineSyncAPIRequest对象
func NewAlibabaAlihouseNewhomeLineSyncRequest() *AlibabaAlihouseNewhomeLineSyncAPIRequest {
	return &AlibabaAlihouseNewhomeLineSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeLineSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.line.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeLineSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeLineSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseLoopLineDto is BaseLoopLineDto Setter
// 环线入参
func (r *AlibabaAlihouseNewhomeLineSyncAPIRequest) SetBaseLoopLineDto(_baseLoopLineDto *BaseLoopLineDto) error {
	r._baseLoopLineDto = _baseLoopLineDto
	r.Set("base_loop_line_dto", _baseLoopLineDto)
	return nil
}

// GetBaseLoopLineDto BaseLoopLineDto Getter
func (r AlibabaAlihouseNewhomeLineSyncAPIRequest) GetBaseLoopLineDto() *BaseLoopLineDto {
	return r._baseLoopLineDto
}
