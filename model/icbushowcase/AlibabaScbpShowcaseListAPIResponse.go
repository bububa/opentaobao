package icbushowcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseListAPIResponse 橱窗查询 API返回值
// alibaba.scbp.showcase.list
//
// 橱窗查询
type AlibabaScbpShowcaseListAPIResponse struct {
	model.CommonResponse
	AlibabaScbpShowcaseListAPIResponseModel
}

// AlibabaScbpShowcaseListAPIResponseModel is 橱窗查询 成功返回结果
type AlibabaScbpShowcaseListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Results []Showcase `json:"results,omitempty" xml:"results>showcase,omitempty"`
}
