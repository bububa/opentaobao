package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单信息上传更新 API请求
taobao.xhotel.order.future.info.put

商家调用推送信息给飞猪平台。 支持如下操作类型：21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
*/
type TaobaoXhotelOrderFutureInfoPutRequest struct {
    model.Params
    // 商家请求流水号
    outUuid   string
    // 操作类型 21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
    operateType   int64
    // 酒店编码
    hotelCode   string
    // 字段详细介绍参见 https://open.alitrip.com/docs/doc.htm?docType=1&articleId=106153
    context   string
    // 商家vendor信息。具体值咨询淘宝技术
    vendor   string
    // 请求流水号
    requestId   string
}

// 初始化TaobaoXhotelOrderFutureInfoPutRequest对象
func NewTaobaoXhotelOrderFutureInfoPutRequest() *TaobaoXhotelOrderFutureInfoPutRequest{
    return &TaobaoXhotelOrderFutureInfoPutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.future.info.put"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutUuid Setter
// 商家请求流水号
func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

// OutUuid Getter
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetOutUuid() string {
    return r.outUuid
}
// OperateType Setter
// 操作类型 21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetOperateType(operateType int64) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

// OperateType Getter
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetOperateType() int64 {
    return r.operateType
}
// HotelCode Setter
// 酒店编码
func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetHotelCode() string {
    return r.hotelCode
}
// Context Setter
// 字段详细介绍参见 https://open.alitrip.com/docs/doc.htm?docType=1&articleId=106153
func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetContext(context string) error {
    r.context = context
    r.Set("context", context)
    return nil
}

// Context Getter
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetContext() string {
    return r.context
}
// Vendor Setter
// 商家vendor信息。具体值咨询淘宝技术
func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetVendor() string {
    return r.vendor
}
// RequestId Setter
// 请求流水号
func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

// RequestId Getter
func (r TaobaoXhotelOrderFutureInfoPutRequest) GetRequestId() string {
    return r.requestId
}
