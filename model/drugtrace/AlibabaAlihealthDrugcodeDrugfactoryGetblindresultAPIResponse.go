package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse 获取盲底文件处理结果 API返回值
// alibaba.alihealth.drugcode.drugfactory.getblindresult
//
// 获取盲底文件处理结果
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel
}

// AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel is 获取盲底文件处理结果 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_getblindresult_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
