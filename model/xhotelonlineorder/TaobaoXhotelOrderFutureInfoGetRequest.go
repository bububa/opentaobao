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
type TaobaoXhotelOrderFutureInfoGetRequest struct {
    model.Params
    // 请求流水号
    outUuid   string
    // 指定淘宝订单ID。以英文分号隔开的字符串“123455666;123455666;123455666”
    tids   string
    // 酒店编码
    hotelCode   string
    // 系统商分配的身份识别
    vendor   string
    // 操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
    operateType   int64
    // 开始时间
    createdStart   string
    // 结束时间
    createdEnd   string
}

// 初始化TaobaoXhotelOrderFutureInfoGetRequest对象
func NewTaobaoXhotelOrderFutureInfoGetRequest() *TaobaoXhotelOrderFutureInfoGetRequest{
    return &TaobaoXhotelOrderFutureInfoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.future.info.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutUuid Setter
// 请求流水号
func (r *TaobaoXhotelOrderFutureInfoGetRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetOutUuid() string {
    return r.outUuid
}
// Tids Setter
// 指定淘宝订单ID。以英文分号隔开的字符串“123455666;123455666;123455666”
func (r *TaobaoXhotelOrderFutureInfoGetRequest) SetTids(tids string) error {
    r.tids = tids
    r.Set("tids", tids)
    return nil
}

// Tids Getter
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetTids() string {
    return r.tids
}
// HotelCode Setter
// 酒店编码
func (r *TaobaoXhotelOrderFutureInfoGetRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetHotelCode() string {
    return r.hotelCode
}
// Vendor Setter
// 系统商分配的身份识别
func (r *TaobaoXhotelOrderFutureInfoGetRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetVendor() string {
    return r.vendor
}
// OperateType Setter
// 操作类型 1.在线开发票请求 3.在线选房请求 4.自助checkIn请求 13.扫脸入住身份信息请求 10.房态信息查询请求 103.通用任务取消指令
func (r *TaobaoXhotelOrderFutureInfoGetRequest) SetOperateType(operateType int64) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

// OperateType Getter
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetOperateType() int64 {
    return r.operateType
}
// CreatedStart Setter
// 开始时间
func (r *TaobaoXhotelOrderFutureInfoGetRequest) SetCreatedStart(createdStart string) error {
    r.createdStart = createdStart
    r.Set("created_start", createdStart)
    return nil
}

// CreatedStart Getter
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetCreatedStart() string {
    return r.createdStart
}
// CreatedEnd Setter
// 结束时间
func (r *TaobaoXhotelOrderFutureInfoGetRequest) SetCreatedEnd(createdEnd string) error {
    r.createdEnd = createdEnd
    r.Set("created_end", createdEnd)
    return nil
}

// CreatedEnd Getter
func (r TaobaoXhotelOrderFutureInfoGetRequest) GetCreatedEnd() string {
    return r.createdEnd
}
