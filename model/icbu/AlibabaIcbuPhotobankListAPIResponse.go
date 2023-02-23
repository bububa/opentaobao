package icbu

import (
	"encoding/xml"

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

// AlibabaIcbuPhotobankListAPIResponseModel is 国际站图片银行查询接口 成功返回结果
type AlibabaIcbuPhotobankListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_photobank_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
