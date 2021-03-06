package aliqin

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotQryPersoninfoAPIResponse 查询物联卡个人实人认证信息 API返回值
// alibaba.aliqin.fc.iot.qry.personinfo
//
// 查询物联卡个人实人认证信息
type AlibabaAliqinFcIotQryPersoninfoAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinFcIotQryPersoninfoAPIResponseModel
}

// AlibabaAliqinFcIotQryPersoninfoAPIResponseModel is 查询物联卡个人实人认证信息 成功返回结果
type AlibabaAliqinFcIotQryPersoninfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_fc_iot_qry_personinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinFcIotQryPersoninfoResult `json:"result,omitempty" xml:"result,omitempty"`
}
