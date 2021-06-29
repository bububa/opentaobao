package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家数据-选品整体概况 APIResponse
taobao.xhotel.item.selection.seller.stat.summary

商家数据-选品整体概况
*/
type TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelItemSelectionSellerStatSummaryResponse
}

type TaobaoXhotelItemSelectionSellerStatSummaryResponse struct {
    XMLName xml.Name `xml:"xhotel_item_selection_seller_stat_summary_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回参数
    
    Result   *HsfResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
