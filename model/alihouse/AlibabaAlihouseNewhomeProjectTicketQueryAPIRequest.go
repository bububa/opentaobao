package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest 根据商品id查询核销卷信息 API请求
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
type AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest struct {
	model.Params
	// 入参对象
	_queryDto *ProjectVerifyTicketQueryDto
}

// NewAlibabaAlihouseNewhomeProjectTicketQueryRequest 初始化AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectTicketQueryRequest() *AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest {
	return &AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) Reset() {
	r._queryDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.ticket.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQueryDto is QueryDto Setter
// 入参对象
func (r *AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) SetQueryDto(_queryDto *ProjectVerifyTicketQueryDto) error {
	r._queryDto = _queryDto
	r.Set("query_dto", _queryDto)
	return nil
}

// GetQueryDto QueryDto Getter
func (r AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) GetQueryDto() *ProjectVerifyTicketQueryDto {
	return r._queryDto
}

var poolAlibabaAlihouseNewhomeProjectTicketQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectTicketQueryRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectTicketQueryRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest
func GetAlibabaAlihouseNewhomeProjectTicketQueryAPIRequest() *AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectTicketQueryAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectTicketQueryAPIRequest 将 AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectTicketQueryAPIRequest(v *AlibabaAlihouseNewhomeProjectTicketQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectTicketQueryAPIRequest.Put(v)
}
