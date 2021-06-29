package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自己联系物流（线下物流）发货 API请求
taobao.logistics.offline.send

用户调用该接口可实现自己联系发货（线下物流），使用该接口发货，交易订单状态会直接变成卖家已发货。不支持货到付款、在线下单类型的订单。
*/
type TaobaoLogisticsOfflineSendRequest struct {
    model.Params
    // 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
    subTid   []int64
    // 淘宝交易ID
    tid   int64
    // 表明是否是拆单，默认值0，1表示拆单
    isSplit   int64
    // 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    outSid   string
    // 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
    companyCode   string
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    senderId   int64
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
    cancelId   int64
    // feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
    feature   string
    // 商家的IP地址
    sellerIp   string
}

// 初始化TaobaoLogisticsOfflineSendRequest对象
func NewTaobaoLogisticsOfflineSendRequest() *TaobaoLogisticsOfflineSendRequest{
    return &TaobaoLogisticsOfflineSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOfflineSendRequest) GetApiMethodName() string {
    return "taobao.logistics.offline.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOfflineSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubTid Setter
// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
func (r *TaobaoLogisticsOfflineSendRequest) SetSubTid(subTid []int64) error {
    r.subTid = subTid
    r.Set("sub_tid", subTid)
    return nil
}

// SubTid Getter
func (r TaobaoLogisticsOfflineSendRequest) GetSubTid() []int64 {
    return r.subTid
}
// Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOfflineSendRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoLogisticsOfflineSendRequest) GetTid() int64 {
    return r.tid
}
// IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaoLogisticsOfflineSendRequest) SetIsSplit(isSplit int64) error {
    r.isSplit = isSplit
    r.Set("is_split", isSplit)
    return nil
}

// IsSplit Getter
func (r TaobaoLogisticsOfflineSendRequest) GetIsSplit() int64 {
    return r.isSplit
}
// OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaoLogisticsOfflineSendRequest) SetOutSid(outSid string) error {
    r.outSid = outSid
    r.Set("out_sid", outSid)
    return nil
}

// OutSid Getter
func (r TaobaoLogisticsOfflineSendRequest) GetOutSid() string {
    return r.outSid
}
// CompanyCode Setter
// 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
func (r *TaobaoLogisticsOfflineSendRequest) SetCompanyCode(companyCode string) error {
    r.companyCode = companyCode
    r.Set("company_code", companyCode)
    return nil
}

// CompanyCode Getter
func (r TaobaoLogisticsOfflineSendRequest) GetCompanyCode() string {
    return r.companyCode
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoLogisticsOfflineSendRequest) SetSenderId(senderId int64) error {
    r.senderId = senderId
    r.Set("sender_id", senderId)
    return nil
}

// SenderId Getter
func (r TaobaoLogisticsOfflineSendRequest) GetSenderId() int64 {
    return r.senderId
}
// CancelId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
func (r *TaobaoLogisticsOfflineSendRequest) SetCancelId(cancelId int64) error {
    r.cancelId = cancelId
    r.Set("cancel_id", cancelId)
    return nil
}

// CancelId Getter
func (r TaobaoLogisticsOfflineSendRequest) GetCancelId() int64 {
    return r.cancelId
}
// Feature Setter
// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
func (r *TaobaoLogisticsOfflineSendRequest) SetFeature(feature string) error {
    r.feature = feature
    r.Set("feature", feature)
    return nil
}

// Feature Getter
func (r TaobaoLogisticsOfflineSendRequest) GetFeature() string {
    return r.feature
}
// SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsOfflineSendRequest) SetSellerIp(sellerIp string) error {
    r.sellerIp = sellerIp
    r.Set("seller_ip", sellerIp)
    return nil
}

// SellerIp Getter
func (r TaobaoLogisticsOfflineSendRequest) GetSellerIp() string {
    return r.sellerIp
}
