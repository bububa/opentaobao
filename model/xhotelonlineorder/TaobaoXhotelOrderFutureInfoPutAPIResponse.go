package xhotelonlineorder

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
订单信息上传更新 API返回值 
taobao.xhotel.order.future.info.put

商家调用推送信息给飞猪平台。 支持如下操作类型：21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
*/
type TaobaoXhotelOrderFutureInfoPutAPIResponse struct {
    model.CommonResponse
    TaobaoXhotelOrderFutureInfoPutAPIResponseModel
}

// 订单信息上传更新 成功返回结果
type TaobaoXhotelOrderFutureInfoPutAPIResponseModel struct {
    XMLName xml.Name `xml:"xhotel_order_future_info_put_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功标记
    Issuccess   bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
    // 错误码
    Errcode   string `json:"errcode,omitempty" xml:"errcode,omitempty"`
    // 错误描述
    Eerrmsg   string `json:"eerrmsg,omitempty" xml:"eerrmsg,omitempty"`
    // 是否更新失败。返回false表示更新成功。否则请读取错误码与错误描述
    Error   bool `json:"error,omitempty" xml:"error,omitempty"`
}
