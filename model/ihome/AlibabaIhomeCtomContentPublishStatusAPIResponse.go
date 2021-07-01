package ihome

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIhomeCtomContentPublishStatusAPIResponse
实拍图发布审核状态查询API API返回值
alibaba.ihome.ctom.content.publish.status

实拍图发布审核状态查询API */
type AlibabaIhomeCtomContentPublishStatusAPIResponse struct {
	model.CommonResponse
	AlibabaIhomeCtomContentPublishStatusAPIResponseModel
}

// AlibabaIhomeCtomContentPublishStatusAPIResponseModel is 实拍图发布审核状态查询API 成功返回结果
type AlibabaIhomeCtomContentPublishStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ihome_ctom_content_publish_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 根据站点名称查询产品
	Result *AlibabaIhomeCtomContentPublishStatusApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
