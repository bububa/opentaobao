package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkSkuChannelskuQueryAPIResponse
查询渠道商品 API返回值
alibaba.wdk.sku.channelsku.query

查询渠道商品 */
type AlibabaWdkSkuChannelskuQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkSkuChannelskuQueryAPIResponseModel
}

// AlibabaWdkSkuChannelskuQueryAPIResponseModel is 查询渠道商品 成功返回结果
type AlibabaWdkSkuChannelskuQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_channelsku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabaWdkSkuChannelskuQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
