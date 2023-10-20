package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomelinesyncAPIRequest 环线数据同步 API请求
// alibaba.alihouse.newhome.line.sync
//
// 环线数据同步
type AlibabaalihousenewhomelinesyncAPIRequest struct {
	model.Params
	// 环线入参
	_baseLoopLineDto *BaseLoopLineDto
}

// NewAlibabaalihousenewhomelinesyncRequest 初始化AlibabaalihousenewhomelinesyncAPIRequest对象
func NewAlibabaalihousenewhomelinesyncRequest() *AlibabaalihousenewhomelinesyncAPIRequest {
	return &AlibabaalihousenewhomelinesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomelinesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.line.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomelinesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomelinesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBaseLoopLineDto is BaseLoopLineDto Setter
// 环线入参
func (r *AlibabaalihousenewhomelinesyncAPIRequest) SetBaseLoopLineDto(_baseLoopLineDto *BaseLoopLineDto) error {
	r._baseLoopLineDto = _baseLoopLineDto
	r.Set("base_loop_line_dto", _baseLoopLineDto)
	return nil
}

// GetBaseLoopLineDto BaseLoopLineDto Getter
func (r AlibabaalihousenewhomelinesyncAPIRequest) GetBaseLoopLineDto() *BaseLoopLineDto {
	return r._baseLoopLineDto
}
