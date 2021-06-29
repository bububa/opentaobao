package drugtrace

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系处理详情 APIResponse
alibaba.alihealth.drug.kyt.relationdetail

关联关系处理详情
*/
type AlibabaAlihealthDrugKytRelationdetailAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthDrugKytRelationdetailResponse
}

type AlibabaAlihealthDrugKytRelationdetailResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_drug_kyt_relationdetail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaAlihealthDrugKytRelationdetailResultModel `json:"result,omitempty" xml:"result,omitempty"`

    
}
