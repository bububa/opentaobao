package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIResponse
接收中央随机化系统和临床研究机构的绑定确认状态 API返回值
alibaba.alihealth.drugcode.center.receive.bound.status

临床用药试验-接收中央随机化系统和临床研究机构的绑定确认状态 */
type AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIResponseModel
}

// AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIResponseModel is 接收中央随机化系统和临床研究机构的绑定确认状态 成功返回结果
type AlibabaAlihealthDrugcodeCenterReceiveBoundStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_center_receive_bound_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
	// 调用结果码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 提示信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
