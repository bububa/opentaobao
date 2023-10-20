package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse 码流向查询 API返回值
// alibaba.alihealth.drug.code.kyt.querycodeflow
//
// 追溯码流向查询
type AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel
}

// AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel is 码流向查询 成功返回结果
type AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_code_kyt_querycodeflow_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugCodeKytQuerycodeflowResult `json:"result,omitempty" xml:"result,omitempty"`
}
