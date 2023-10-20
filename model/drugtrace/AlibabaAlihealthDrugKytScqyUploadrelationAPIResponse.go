package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse 关联关系上传 API返回值
// alibaba.alihealth.drug.kyt.scqy.uploadrelation
//
// 关联关系上传
type AlibabaAlihealthDrugKytScqyUploadrelationAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel
}

// AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel is 关联关系上传 成功返回结果
type AlibabaAlihealthDrugKytScqyUploadrelationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_scqy_uploadrelation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytScqyUploadrelationResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
