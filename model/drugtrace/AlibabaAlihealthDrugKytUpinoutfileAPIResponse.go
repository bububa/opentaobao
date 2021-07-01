package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUpinoutfileAPIResponse
上传出入库单据(传文件) API返回值
alibaba.alihealth.drug.kyt.upinoutfile

上传出入库单据(传文件) */
type AlibabaAlihealthDrugKytUpinoutfileAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytUpinoutfileAPIResponseModel
}

// AlibabaAlihealthDrugKytUpinoutfileAPIResponseModel is 上传出入库单据(传文件) 成功返回结果
type AlibabaAlihealthDrugKytUpinoutfileAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_upinoutfile_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 上传的ID
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 返回编码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 描述信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回值
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
