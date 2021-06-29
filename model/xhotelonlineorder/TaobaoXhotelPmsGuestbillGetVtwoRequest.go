package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
客人PMS账单信息查询 API请求
taobao.xhotel.pms.guestbill.get.vtwo

从pms获取客人账单信息
*/
type TaobaoXhotelPmsGuestbillGetVtwoRequest struct {
    model.Params
    // 开票点税号
    _taxNum   string
    // 身份证后4位
    _shortIdNum   string
    // 房间号
    _roomNum   string
    // 请求id (32位唯一值)
    _requestId   string
    // 用户渠道(0:未知,1:淘宝)
    _userChannel   int64
}

// 初始化TaobaoXhotelPmsGuestbillGetVtwoRequest对象
func NewTaobaoXhotelPmsGuestbillGetVtwoRequest() *TaobaoXhotelPmsGuestbillGetVtwoRequest{
    return &TaobaoXhotelPmsGuestbillGetVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetApiMethodName() string {
    return "taobao.xhotel.pms.guestbill.get.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TaxNum Setter
// 开票点税号
func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetTaxNum(_taxNum string) error {
    r._taxNum = _taxNum
    r.Set("tax_num", _taxNum)
    return nil
}

// TaxNum Getter
func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetTaxNum() string {
    return r._taxNum
}
// ShortIdNum Setter
// 身份证后4位
func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetShortIdNum(_shortIdNum string) error {
    r._shortIdNum = _shortIdNum
    r.Set("short_id_num", _shortIdNum)
    return nil
}

// ShortIdNum Getter
func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetShortIdNum() string {
    return r._shortIdNum
}
// RoomNum Setter
// 房间号
func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetRoomNum(_roomNum string) error {
    r._roomNum = _roomNum
    r.Set("room_num", _roomNum)
    return nil
}

// RoomNum Getter
func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetRoomNum() string {
    return r._roomNum
}
// RequestId Setter
// 请求id (32位唯一值)
func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetRequestId(_requestId string) error {
    r._requestId = _requestId
    r.Set("request_id", _requestId)
    return nil
}

// RequestId Getter
func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetRequestId() string {
    return r._requestId
}
// UserChannel Setter
// 用户渠道(0:未知,1:淘宝)
func (r *TaobaoXhotelPmsGuestbillGetVtwoRequest) SetUserChannel(_userChannel int64) error {
    r._userChannel = _userChannel
    r.Set("user_channel", _userChannel)
    return nil
}

// UserChannel Getter
func (r TaobaoXhotelPmsGuestbillGetVtwoRequest) GetUserChannel() int64 {
    return r._userChannel
}
