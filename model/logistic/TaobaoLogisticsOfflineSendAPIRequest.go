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
type TaobaoLogisticsOfflineSendAPIRequest struct {
    model.Params
    // 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
    _subTid   []int64
    // 淘宝交易ID
    _tid   int64
    // 表明是否是拆单，默认值0，1表示拆单
    _isSplit   int64
    // 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    _outSid   string
    // 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
    _companyCode   string
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
    _senderId   int64
    // 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
    _cancelId   int64
    // feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
    _feature   string
    // 商家的IP地址
    _sellerIp   string
}

// 初始化TaobaoLogisticsOfflineSendAPIRequest对象
func NewTaobaoLogisticsOfflineSendRequest() *TaobaoLogisticsOfflineSendAPIRequest{
    return &TaobaoLogisticsOfflineSendAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOfflineSendAPIRequest) GetApiMethodName() string {
    return "taobao.logistics.offline.send"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOfflineSendAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SubTid Setter
// 需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetSubTid(_subTid []int64) error {
    r._subTid = _subTid
    r.Set("sub_tid", _subTid)
    return nil
}

// SubTid Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetSubTid() []int64 {
    return r._subTid
}
// Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetTid() int64 {
    return r._tid
}
// IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetIsSplit(_isSplit int64) error {
    r._isSplit = _isSplit
    r.Set("is_split", _isSplit)
    return nil
}

// IsSplit Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetIsSplit() int64 {
    return r._isSplit
}
// OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetOutSid(_outSid string) error {
    r._outSid = _outSid
    r.Set("out_sid", _outSid)
    return nil
}

// OutSid Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetOutSid() string {
    return r._outSid
}
// CompanyCode Setter
// 物流公司代码.如"POST"就代表中国邮政,"ZJS"就代表宅急送.调用 taobao.logistics.companies.get 获取。
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetCompanyCode(_companyCode string) error {
    r._companyCode = _companyCode
    r.Set("company_code", _companyCode)
    return nil
}

// CompanyCode Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetCompanyCode() string {
    return r._companyCode
}
// SenderId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetSenderId(_senderId int64) error {
    r._senderId = _senderId
    r.Set("sender_id", _senderId)
    return nil
}

// SenderId Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetSenderId() int64 {
    return r._senderId
}
// CancelId Setter
// 卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetCancelId(_cancelId int64) error {
    r._cancelId = _cancelId
    r.Set("cancel_id", _cancelId)
    return nil
}

// CancelId Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetCancelId() int64 {
    return r._cancelId
}
// Feature Setter
// feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段。
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetFeature(_feature string) error {
    r._feature = _feature
    r.Set("feature", _feature)
    return nil
}

// Feature Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetFeature() string {
    return r._feature
}
// SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsOfflineSendAPIRequest) SetSellerIp(_sellerIp string) error {
    r._sellerIp = _sellerIp
    r.Set("seller_ip", _sellerIp)
    return nil
}

// SellerIp Getter
func (r TaobaoLogisticsOfflineSendAPIRequest) GetSellerIp() string {
    return r._sellerIp
}
