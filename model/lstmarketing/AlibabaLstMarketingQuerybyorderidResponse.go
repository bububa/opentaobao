package lstmarketing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据订单查询营销信息 APIResponse
alibaba.lst.marketing.querybyorderid

根据订单查询营销信息
*/
type AlibabaLstMarketingQuerybyorderidAPIResponse struct {
    model.CommonResponse
    AlibabaLstMarketingQuerybyorderidResponse
}

type AlibabaLstMarketingQuerybyorderidResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_marketing_querybyorderid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 异步获取历史数据接口返回结果
    
    Result   *AlibabaLstMarketingQuerybyorderidResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
