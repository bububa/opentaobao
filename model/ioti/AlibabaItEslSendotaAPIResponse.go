package ioti

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaiteslsendotaAPIResponse 电子价签ota接口 API返回值
// alibaba.it.esl.sendota
//
// 厂测接口，电子价签ota接口
type AlibabaiteslsendotaAPIResponse struct {
	model.CommonResponse
	AlibabaiteslsendotaAPIResponseModel
}

// AlibabaiteslsendotaAPIResponseModel is 电子价签ota接口 成功返回结果
type AlibabaiteslsendotaAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_it_esl_sendota_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Can not find Corresponding AP MAC with ESL
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
