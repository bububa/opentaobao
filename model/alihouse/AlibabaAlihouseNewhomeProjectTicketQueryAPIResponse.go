package alihouse

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse 根据商品id查询核销卷信息 API返回值
// alibaba.alihouse.newhome.project.ticket.query
//
// 根据商品id查询核销卷信息
type AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel).Reset()
}

// AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel is 根据商品id查询核销卷信息 成功返回结果
type AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_project_ticket_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeProjectTicketQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihouseNewhomeProjectTicketQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihouseNewhomeProjectTicketQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse)
	},
}

// GetAlibabaAlihouseNewhomeProjectTicketQueryAPIResponse 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse
func GetAlibabaAlihouseNewhomeProjectTicketQueryAPIResponse() *AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse {
	return poolAlibabaAlihouseNewhomeProjectTicketQueryAPIResponse.Get().(*AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse)
}

// ReleaseAlibabaAlihouseNewhomeProjectTicketQueryAPIResponse 将 AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectTicketQueryAPIResponse(v *AlibabaAlihouseNewhomeProjectTicketQueryAPIResponse) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectTicketQueryAPIResponse.Put(v)
}
