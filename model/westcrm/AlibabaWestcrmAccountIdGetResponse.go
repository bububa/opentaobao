package westcrm

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据支付宝id查询IB的id API返回值 
alibaba.westcrm.account.id.get

根据支付宝id查询IB的id
*/
type AlibabaWestcrmAccountIdGetAPIResponse struct {
    model.CommonResponse
    AlibabaWestcrmAccountIdGetResponse
}

// 根据支付宝id查询IB的id 成功返回结果
type AlibabaWestcrmAccountIdGetResponse struct {
    XMLName xml.Name `xml:"alibaba_westcrm_account_id_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回对象封装
    Result   *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
