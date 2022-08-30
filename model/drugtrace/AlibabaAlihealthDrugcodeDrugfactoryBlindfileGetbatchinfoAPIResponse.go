package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIResponse 获取盲底文件中的批次信息 API返回值
// alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo
//
// 获取盲底文件中的批次信息
type AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIResponseModel
}

// AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIResponseModel is 获取盲底文件中的批次信息 成功返回结果
type AlibabaAlihealthDrugcodeDrugfactoryBlindfileGetbatchinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drugcode_drugfactory_blindfile_getbatchinfo_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 和三方交互最外层model对象
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
