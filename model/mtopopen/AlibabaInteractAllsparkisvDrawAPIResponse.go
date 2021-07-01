package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractAllsparkisvDrawAPIResponse
allspark提供抽奖tida接口对应鉴权接口 API返回值
alibaba.interact.allsparkisv.draw

该接口没有实际对外使用。只是内部鉴权使用，不会有三方应用调用 */
type AlibabaInteractAllsparkisvDrawAPIResponse struct {
	model.CommonResponse
	AlibabaInteractAllsparkisvDrawAPIResponseModel
}

// AlibabaInteractAllsparkisvDrawAPIResponseModel is allspark提供抽奖tida接口对应鉴权接口 成功返回结果
type AlibabaInteractAllsparkisvDrawAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_allsparkisv_draw_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ddd
	Ddd string `json:"ddd,omitempty" xml:"ddd,omitempty"`
}
