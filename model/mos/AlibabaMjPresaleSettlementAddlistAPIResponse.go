package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
预售结算数据回传 API返回值 
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。
*/
type AlibabaMjPresaleSettlementAddlistAPIResponse struct {
    model.CommonResponse
    AlibabaMjPresaleSettlementAddlistAPIResponseModel
}

// 预售结算数据回传 成功返回结果
type AlibabaMjPresaleSettlementAddlistAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mj_presale_settlement_addlist_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // success
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
