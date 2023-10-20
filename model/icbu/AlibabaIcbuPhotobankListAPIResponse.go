package icbu

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuPhotobankListAPIResponse 国际站图片银行查询接口 API返回值
// alibaba.icbu.photobank.list
//
// 国际站图片银行查询接口
type AlibabaIcbuPhotobankListAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuPhotobankListAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIcbuPhotobankListAPIResponseModel).Reset()
}

// AlibabaIcbuPhotobankListAPIResponseModel is 国际站图片银行查询接口 成功返回结果
type AlibabaIcbuPhotobankListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_photobank_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// error code
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// error message
	Errormsg string `json:"errormsg,omitempty" xml:"errormsg,omitempty"`
	// PaginationQueryList
	PaginationQueryList *PaginationQueryList `json:"pagination_query_list,omitempty" xml:"pagination_query_list,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIcbuPhotobankListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TraceId = ""
	m.Errorcode = ""
	m.Errormsg = ""
	m.PaginationQueryList = nil
}

var poolAlibabaIcbuPhotobankListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIcbuPhotobankListAPIResponse)
	},
}

// GetAlibabaIcbuPhotobankListAPIResponse 从 sync.Pool 获取 AlibabaIcbuPhotobankListAPIResponse
func GetAlibabaIcbuPhotobankListAPIResponse() *AlibabaIcbuPhotobankListAPIResponse {
	return poolAlibabaIcbuPhotobankListAPIResponse.Get().(*AlibabaIcbuPhotobankListAPIResponse)
}

// ReleaseAlibabaIcbuPhotobankListAPIResponse 将 AlibabaIcbuPhotobankListAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIcbuPhotobankListAPIResponse(v *AlibabaIcbuPhotobankListAPIResponse) {
	v.Reset()
	poolAlibabaIcbuPhotobankListAPIResponse.Put(v)
}
