package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应链选品热销标准酒店覆盖情况 APIResponse
taobao.xhotel.item.selection.seller.stat.hotshid

供应链选品热销标准酒店覆盖情况
*/
type TaobaoXhotelItemSelectionSellerStatHotshidAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelItemSelectionSellerStatHotshidResponse
}

type TaobaoXhotelItemSelectionSellerStatHotshidResponse struct {
    XMLName xml.Name `xml:"xhotel_item_selection_seller_stat_hotshid_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoXhotelItemSelectionSellerStatHotshidResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
