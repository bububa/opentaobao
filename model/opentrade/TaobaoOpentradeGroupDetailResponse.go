package opentrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组团购场景查询团详情 APIResponse
taobao.opentrade.group.detail

组团购场景下，查询团详情
*/
type TaobaoOpentradeGroupDetailAPIResponse struct {
    model.CommonResponse
    TaobaoOpentradeGroupDetailResponse
}

type TaobaoOpentradeGroupDetailResponse struct {
    XMLName xml.Name `xml:"opentrade_group_detail_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 团信息
    
    Result   *GroupDetailResponse `json:"result,omitempty" xml:"result,omitempty"`

    
}
