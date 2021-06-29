package alsc

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
搜索卡实例列表(支持号段查询) APIResponse
alibaba.alsc.crm.card.searchcard

搜索卡实例列表(支持号段查询)
*/
type AlibabaAlscCrmCardSearchcardAPIResponse struct {
    model.CommonResponse
    AlibabaAlscCrmCardSearchcardResponse
}

type AlibabaAlscCrmCardSearchcardResponse struct {
    XMLName xml.Name `xml:"alibaba_alsc_crm_card_searchcard_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 分页返回模型
    
    Result   *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
