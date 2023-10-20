package campus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacampusguardantgatesyncAPIResponse 网点数据同步 API返回值
// alibaba.campus.guardant.gate.sync
//
// 门禁供应商创建网点同步
type AlibabacampusguardantgatesyncAPIResponse struct {
	model.CommonResponse
	AlibabacampusguardantgatesyncAPIResponseModel
}

// AlibabacampusguardantgatesyncAPIResponseModel is 网点数据同步 成功返回结果
type AlibabacampusguardantgatesyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_guardant_gate_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}
