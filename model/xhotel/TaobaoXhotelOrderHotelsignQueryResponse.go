package xhotel

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取直连酒店（客栈）签约上线进度信息 APIResponse
taobao.xhotel.order.hotelsign.query

获取直连酒店（客栈）签约上线进度信息
*/
type TaobaoXhotelOrderHotelsignQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoXhotelOrderHotelsignQueryResponse `json:"taobao_xhotel_order_hotelsign_query_response,omitempty"`
}

type TaobaoXhotelOrderHotelsignQueryResponse struct {

    // outUuid
    OutUuid   string `json:"out_uuid,omitempty"`

    // hotelSignInfo，当入参中包含hotelcode和vendor的时候，返回该对象
    HotelSignInfo   *DchotelSignDo `json:"hotel_sign_info,omitempty"`

    // dsNhotelInfoDOList，当入参不包含hotelcode的时候，只有vendor的时候返回该对象
    DsNhotelInfoDOList   []DsNhotelInfoDo `json:"ds_nhotel_info_d_o_list,omitempty"`

}
