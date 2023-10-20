package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponse 码拼箱建立父子关系接口 API返回值
// alibaba.alihealth.drug.kyt.scqy.putpackagebind
//
// 码拼箱建立父子关系接口
type AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponseModel
}

// AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponseModel is 码拼箱建立父子关系接口 成功返回结果
type AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_putpackagebind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyPutpackagebindResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
