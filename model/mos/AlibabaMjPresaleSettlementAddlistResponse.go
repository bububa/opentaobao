package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预售结算数据回传 APIResponse
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。
*/
type AlibabaMjPresaleSettlementAddlistAPIResponse struct {
    model.CommonResponse
    AlibabaMjPresaleSettlementAddlistResponse
}

type AlibabaMjPresaleSettlementAddlistResponse struct {
    XMLName xml.Name `xml:"alibaba_mj_presale_settlement_addlist_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
