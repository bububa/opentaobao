package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuchannelskuaddAPIResponse 新增渠道商品 API返回值
// alibaba.wdk.sku.channelsku.add
//
// 盒马帮1期新增渠道商品
type AlibabawdkskuchannelskuaddAPIResponse struct {
	model.CommonResponse
	AlibabawdkskuchannelskuaddAPIResponseModel
}

// AlibabawdkskuchannelskuaddAPIResponseModel is 新增渠道商品 成功返回结果
type AlibabawdkskuchannelskuaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_channelsku_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlibabawdkskuchannelskuaddApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
