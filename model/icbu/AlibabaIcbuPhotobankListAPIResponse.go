package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuphotobanklistAPIResponse 国际站图片银行查询接口 API返回值
// alibaba.icbu.photobank.list
//
// 国际站图片银行查询接口
type AlibabaicbuphotobanklistAPIResponse struct {
	model.CommonResponse
	AlibabaicbuphotobanklistAPIResponseModel
}

// AlibabaicbuphotobanklistAPIResponseModel is 国际站图片银行查询接口 成功返回结果
type AlibabaicbuphotobanklistAPIResponseModel struct {
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
