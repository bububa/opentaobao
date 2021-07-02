package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponse 查询一个码的所有子码 API返回值
// alibaba.alihealth.drug.kyt.yy.querysubcodes
//
// 单码的了码查询
type AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponseModel
}

// AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponseModel is 查询一个码的所有子码 成功返回结果
type AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yy_querysubcodes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytYyQuerysubcodesResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
