package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkWholesaleOutboundorderCommitAPIResponse 盒马帮发货信息回传接口 API返回值
// alibaba.wdk.wholesale.outboundorder.commit
//
// 盒马帮发货信息回传接口
type AlibabaWdkWholesaleOutboundorderCommitAPIResponse struct {
	model.CommonResponse
	AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel
}

// AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel is 盒马帮发货信息回传接口 成功返回结果
type AlibabaWdkWholesaleOutboundorderCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_wholesale_outboundorder_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkWholesaleOutboundorderCommitApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
