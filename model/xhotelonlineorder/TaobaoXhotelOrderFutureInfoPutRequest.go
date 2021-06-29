package xhotelonlineorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单信息上传更新 APIRequest
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

func NewTaobaoXhotelOrderFutureInfoPutRequest() *TaobaoXhotelOrderFutureInfoPutRequest{
    return &TaobaoXhotelOrderFutureInfoPutRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetApiMethodName() string {
    return "taobao.xhotel.order.future.info.put"
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetOutUuid(outUuid string) error {
    r.outUuid = outUuid
    r.Set("out_uuid", outUuid)
    return nil
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetOutUuid() string {
    return r.outUuid
}

func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetOperateType(operateType int64) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetOperateType() int64 {
    return r.operateType
}

func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetHotelCode() string {
    return r.hotelCode
}

func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetContext(context string) error {
    r.context = context
    r.Set("context", context)
    return nil
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetContext() string {
    return r.context
}

func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelOrderFutureInfoPutRequest) SetRequestId(requestId string) error {
    r.requestId = requestId
    r.Set("request_id", requestId)
    return nil
}

func (r TaobaoXhotelOrderFutureInfoPutRequest) GetRequestId() string {
    return r.requestId
}

