package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse 码拼箱解除父子关系接口 API返回值
// alibaba.alihealth.drug.kyt.scqy.putpackageunbind
//
// 码拼箱解除父子关系接口
type AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel
}

// AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel is 码拼箱解除父子关系接口 成功返回结果
type AlibabaAlihealthDrugKytScqyPutpackageunbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_putpackageunbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 最外层结果
	Result *AlibabaAlihealthDrugKytScqyPutpackageunbindResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
