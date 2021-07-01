package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcSmsNumQueryAPIResponse
短信发送记录查询 API返回值
alibaba.aliqin.fc.sms.num.query

短信发送记录查询。 */
type AlibabaAliqinFcSmsNumQueryAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcSmsNumQueryAPIResponseModel
}

// AlibabaAliqinFcSmsNumQueryAPIResponseModel is 短信发送记录查询 成功返回结果
type AlibabaAliqinFcSmsNumQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_sms_num_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 1
	Values []FcPartnerSmsDetailDto `json:"values,omitempty" xml:"values>fc_partner_sms_detail_dto,omitempty"`
}
