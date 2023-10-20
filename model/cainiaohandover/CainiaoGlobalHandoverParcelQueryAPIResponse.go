package cainiaohandover

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalhandoverparcelqueryAPIResponse 获取交接单小包信息 API返回值
// cainiao.global.handover.parcel.query
//
// 提供给ISV通过该接口查询小包信息
type CainiaoglobalhandoverparcelqueryAPIResponse struct {
	model.CommonResponse
	CainiaoglobalhandoverparcelqueryAPIResponseModel
}

// CainiaoglobalhandoverparcelqueryAPIResponseModel is 获取交接单小包信息 成功返回结果
type CainiaoglobalhandoverparcelqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_handover_parcel_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}
