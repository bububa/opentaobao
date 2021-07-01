package westcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIbizapiBrandSubscribeAPIResponse
关注品牌号 API返回值
alibaba.ibizapi.brand.subscribe

关注品牌号服务 */
type AlibabaIbizapiBrandSubscribeAPIResponse struct {
	model.CommonResponse
	AlibabaIbizapiBrandSubscribeAPIResponseModel
}

// AlibabaIbizapiBrandSubscribeAPIResponseModel is 关注品牌号 成功返回结果
type AlibabaIbizapiBrandSubscribeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ibizapi_brand_subscribe_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否关注
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
