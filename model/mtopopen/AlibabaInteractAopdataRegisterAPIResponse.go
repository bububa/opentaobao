package mtopopen

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractAopdataRegisterAPIResponse 资源位数据推送接口 API返回值
// alibaba.interact.aopdata.register
//
// 提供给isv，查询以及推送浮层资源位的三方活动数据
type AlibabaInteractAopdataRegisterAPIResponse struct {
	model.CommonResponse
	AlibabaInteractAopdataRegisterAPIResponseModel
}

// AlibabaInteractAopdataRegisterAPIResponseModel is 资源位数据推送接口 成功返回结果
type AlibabaInteractAopdataRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_interact_aopdata_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaInteractAopdataRegisterResult `json:"result,omitempty" xml:"result,omitempty"`
}
