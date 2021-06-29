package xhotelitem

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
选品曝光趋势 APIResponse
taobao.xhotel.item.selection.seller.stat.exposure

用于提供给商家获取选品曝光趋势
*/
type TaobaoXhotelItemSelectionSellerStatExposureAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelItemSelectionSellerStatExposureResponse
}

type TaobaoXhotelItemSelectionSellerStatExposureResponse struct {
    XMLName xml.Name `xml:"xhotel_item_selection_seller_stat_exposure_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *TaobaoXhotelItemSelectionSellerStatExposureResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
