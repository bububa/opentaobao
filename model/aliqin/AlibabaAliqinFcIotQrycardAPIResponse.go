package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinfciotqrycardAPIResponse 查询终端信息 API返回值
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
type AlibabaaliqinfciotqrycardAPIResponse struct {
	model.CommonResponse
	AlibabaaliqinfciotqrycardAPIResponseModel
}

// AlibabaaliqinfciotqrycardAPIResponseModel is 查询终端信息 成功返回结果
type AlibabaaliqinfciotqrycardAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_qrycard_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *AlibabaaliqinfciotqrycardResult `json:"result,omitempty" xml:"result,omitempty"`
}
