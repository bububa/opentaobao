package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询一个码的所有子码 APIResponse
alibaba.alihealth.drug.kyt.yy.querysubcodes

单码的了码查询
*/
type AlibabaAlihealthDrugKytYyQuerysubcodesAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytYyQuerysubcodesResponse
}

type AlibabaAlihealthDrugKytYyQuerysubcodesResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_yy_querysubcodes_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytYyQuerysubcodesResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
