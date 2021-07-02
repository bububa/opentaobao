package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingVersionCommitAPIResponse 提交版本号 API返回值
// alibaba.wdk.marketing.version.commit
//
// 提交版本号，标识结束此版本操作
type AlibabaWdkMarketingVersionCommitAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingVersionCommitAPIResponseModel
}

// AlibabaWdkMarketingVersionCommitAPIResponseModel is 提交版本号 成功返回结果
type AlibabaWdkMarketingVersionCommitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_version_commit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MarketResult `json:"result,omitempty" xml:"result,omitempty"`
}
