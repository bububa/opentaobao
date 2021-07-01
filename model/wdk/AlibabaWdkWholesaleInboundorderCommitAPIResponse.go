package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkWholesaleInboundorderCommitAPIResponse
盒马帮退货信息回传接口 API返回值
alibaba.wdk.wholesale.inboundorder.commit

盒马帮退货信息回传接口 */
type AlibabaWdkWholesaleInboundorderCommitAPIResponse struct {
	model.CommonResponse
	AlibabaWdkWholesaleInboundorderCommitAPIResponseModel
}

// AlibabaWdkWholesaleInboundorderCommitAPIResponseModel is 盒马帮退货信息回传接口 成功返回结果
type AlibabaWdkWholesaleInboundorderCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_wholesale_inboundorder_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkWholesaleInboundorderCommitApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
