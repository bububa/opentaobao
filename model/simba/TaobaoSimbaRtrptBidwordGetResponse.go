package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广词实时报表数据 API返回值 
taobao.simba.rtrpt.bidword.get

获取推广词报表数据
*/
type TaobaoSimbaRtrptBidwordGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRtrptBidwordGetResponse
}

// 获取推广词实时报表数据 成功返回结果
type TaobaoSimbaRtrptBidwordGetResponse struct {
    XMLName xml.Name `xml:"simba_rtrpt_bidword_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // bidword result
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
}
