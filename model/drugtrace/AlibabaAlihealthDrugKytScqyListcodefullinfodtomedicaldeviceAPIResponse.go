package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse 医疗器械的码查询信息接口 API返回值
// alibaba.alihealth.drug.kyt.scqy.listcodefullinfodtomedicaldevice
//
// 医疗器械的码查询信息接口
type AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel
}

// AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel is 医疗器械的码查询信息接口 成功返回结果
type AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_listcodefullinfodtomedicaldevice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyListcodefullinfodtomedicaldeviceResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
