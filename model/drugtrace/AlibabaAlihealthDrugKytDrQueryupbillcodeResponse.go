package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询上游企业出库单据号 APIResponse
alibaba.alihealth.drug.kyt.dr.queryupbillcode

疫苗温度合规补充需求-增加一个查询上游出库单号的接口。疾控在扫码入库时，接口通过扫到的码判定这个码对应所属的出库单据号
*/
type AlibabaAlihealthDrugKytDrQueryupbillcodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytDrQueryupbillcodeResponse
}

type AlibabaAlihealthDrugKytDrQueryupbillcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_dr_queryupbillcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytDrQueryupbillcodeResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
