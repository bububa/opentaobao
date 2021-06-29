package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据支付宝id查询IB的id APIResponse
alibaba.westcrm.account.id.get

根据支付宝id查询IB的id
*/
type AlibabaWestcrmAccountIdGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmAccountIdGetResponse
}

type AlibabaWestcrmAccountIdGetResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_account_id_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回对象封装
    
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
