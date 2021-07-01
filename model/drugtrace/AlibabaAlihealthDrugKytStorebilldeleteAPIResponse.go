package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytStorebilldeleteAPIResponse
零售端单据删除 API返回值
alibaba.alihealth.drug.kyt.storebilldelete

零售端单据删除 */
type AlibabaAlihealthDrugKytStorebilldeleteAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel
}

// AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel is 零售端单据删除 成功返回结果
type AlibabaAlihealthDrugKytStorebilldeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_storebilldelete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果说明
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 调用信息编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 调用信息描述
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回结果
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
