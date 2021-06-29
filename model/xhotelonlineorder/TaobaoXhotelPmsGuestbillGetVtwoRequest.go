package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客人PMS账单信息查询 APIRequest
taobao.xhotel.pms.guestbill.get.vtwo

从pms获取客人账单信息
*/
type TaobaoXhotelPmsGuestbillGetVtwoRequest struct {
    model.Params

    // 开票点税号
    taxNum   string 

    // 身份证后4位
    shortIdNum   string 

    // 房间号
    roomNum   string 

    // 请求id (32位唯一值)
    requestId   string 

    // 用户渠道(0:未知,1:淘宝)
    userChannel   int64 

}

func NewTaobaoXhotelPmsGuestbillGetVtwoRequest() *TaobaoXhotelPmsGuestbillGetVtwoRequest{
    return &TaobaoXhotelPmsGuestbillGetVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetApiMethodName() string {
    return "taobao.xhotel.pms.guestbill.get.vtwo"
}

func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetTaxNum(taxNum string) error {
    r.taxNum = taxNum
    r.Set("tax_num", taxNum)
    return nil
}

func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetTaxNum() string {
    return r.taxNum
}

func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetShortIdNum(shortIdNum string) error {
    r.shortIdNum = shortIdNum
    r.Set("short_id_num", shortIdNum)
    return nil
}

func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetShortIdNum() string {
    return r.shortIdNum
}

func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetRoomNum(roomNum string) error {
    r.roomNum = roomNum
    r.Set("room_num", roomNum)
    return nil
}

func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetRoomNum() string {
    return r.roomNum
}

func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetRequestId() string {
    return r.requestId
}

func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetUserChannel(userChannel int64) error {
    r.userChannel = userChannel
    r.Set("user_channel", userChannel)
    return nil
}

func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetUserChannel() int64 {
    return r.userChannel
}

