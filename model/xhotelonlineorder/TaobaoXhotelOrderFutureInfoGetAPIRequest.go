package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取(查询)订单变更信息 API请求
taobao.xhotel.order.future.info.get

支持操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
*/
type TaobaoXhotelOrderFutureInfoGetAPIRequest struct {
    model.Params
    // 请求流水号
    _outUuid   string
    // 指定淘宝订单ID。以英文分号隔开的字符串“123455666;123455666;123455666”
    _tids   string
    // 酒店编码
    _hotelCode   string
    // 系统商分配的身份识别
    _vendor   string
    // 操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
    _operateType   int64
    // 开始时间
    _createdStart   string
    // 结束时间
    _createdEnd   string
}

// 初始化TaobaoXhotelOrderFutureInfoGetAPIRequest对象
func NewTaobaoXhotelOrderFutureInfoGetRequest() *TaobaoXhotelOrderFutureInfoGetAPIRequest{
    return &TaobaoXhotelOrderFutureInfoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.future.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutUuid Setter
// 请求流水号
func (r *TaobaoXhotelOrderFutureInfoGetAPIRequest) SetOutUuid(_outUuid string) error {
    r._outUuid = _outUuid
    r.Set("out_uuid", _outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetOutUuid() string {
    return r._outUuid
}
// Tids Setter
// 指定淘宝订单ID。以英文分号隔开的字符串“123455666;123455666;123455666”
func (r *TaobaoXhotelOrderFutureInfoGetAPIRequest) SetTids(_tids string) error {
    r._tids = _tids
    r.Set("tids", _tids)
    return nil
}

// Tids Getter
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetTids() string {
    return r._tids
}
// HotelCode Setter
// 酒店编码
func (r *TaobaoXhotelOrderFutureInfoGetAPIRequest) SetHotelCode(_hotelCode string) error {
    r._hotelCode = _hotelCode
    r.Set("hotel_code", _hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetHotelCode() string {
    return r._hotelCode
}
// Vendor Setter
// 系统商分配的身份识别
func (r *TaobaoXhotelOrderFutureInfoGetAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetVendor() string {
    return r._vendor
}
// OperateType Setter
// 操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
func (r *TaobaoXhotelOrderFutureInfoGetAPIRequest) SetOperateType(_operateType int64) error {
    r._operateType = _operateType
    r.Set("operate_type", _operateType)
    return nil
}

// OperateType Getter
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetOperateType() int64 {
    return r._operateType
}
// CreatedStart Setter
// 开始时间
func (r *TaobaoXhotelOrderFutureInfoGetAPIRequest) SetCreatedStart(_createdStart string) error {
    r._createdStart = _createdStart
    r.Set("created_start", _createdStart)
    return nil
}

// CreatedStart Getter
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetCreatedStart() string {
    return r._createdStart
}
// CreatedEnd Setter
// 结束时间
func (r *TaobaoXhotelOrderFutureInfoGetAPIRequest) SetCreatedEnd(_createdEnd string) error {
    r._createdEnd = _createdEnd
    r.Set("created_end", _createdEnd)
    return nil
}

// CreatedEnd Getter
func (r TaobaoXhotelOrderFutureInfoGetAPIRequest) GetCreatedEnd() string {
    return r._createdEnd
}
