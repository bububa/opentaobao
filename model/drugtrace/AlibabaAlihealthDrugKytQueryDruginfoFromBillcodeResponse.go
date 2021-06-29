package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单据编号查询单据明细 APIResponse
alibaba.alihealth.drug.kyt.query.druginfo.from.billcode

根据单据编号查询单据明细
*/
type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResponse
}

type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_query_druginfo_from_billcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 监控宝推送网站监控信息，返回结果
    
    Result   *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
