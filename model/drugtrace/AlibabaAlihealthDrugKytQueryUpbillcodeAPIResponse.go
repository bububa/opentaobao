package drugtrace

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse 通过一个码查询上游出库单 API返回值
// alibaba.alihealth.drug.kyt.query.upbillcode
//
// 一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
type AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel
}

// AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel is 通过一个码查询上游出库单 成功返回结果
type AlibabaAlihealthDrugKytQueryUpbillcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_query_upbillcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihealthDrugKytQueryUpbillcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
