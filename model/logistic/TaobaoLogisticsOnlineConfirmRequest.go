package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
确认发货通知接口 API请求
taobao.logistics.online.confirm

<br><font color='red'>仅在使用taobao.logistics.online.send 发货时未输入运单号的情况下，需要使用该接口补充填写运单号，来确认发货。<br>
确认发货的目的是让交易流程继续走下去，确认发货后交易状态会由【买家已付款】变为【卖家已发货】。</font>
*/
type TaobaoLogisticsOnlineConfirmRequest struct {
    model.Params
    // 淘宝交易ID
    tid   int64
    // 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
    subTid   []int64
    // 表明是否是拆单，默认值0，1表示拆单
    isSplit   int64
    // 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    outSid   string
    // 商家的IP地址
    sellerIp   string
}

// 初始化TaobaoLogisticsOnlineConfirmRequest对象
func NewTaobaoLogisticsOnlineConfirmRequest() *TaobaoLogisticsOnlineConfirmRequest{
    return &TaobaoLogisticsOnlineConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsOnlineConfirmRequest) GetApiMethodName() string {
    return "taobao.logistics.online.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsOnlineConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Tid Setter
// 淘宝交易ID
func (r *TaobaoLogisticsOnlineConfirmRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoLogisticsOnlineConfirmRequest) GetTid() int64 {
    return r.tid
}
// SubTid Setter
// 拆单子订单列表，对应的数据是：子订单号的列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集
func (r *TaobaoLogisticsOnlineConfirmRequest) SetSubTid(subTid []int64) error {
    r.subTid = subTid
    r.Set("sub_tid", subTid)
    return nil
}

// SubTid Getter
func (r TaobaoLogisticsOnlineConfirmRequest) GetSubTid() []int64 {
    return r.subTid
}
// IsSplit Setter
// 表明是否是拆单，默认值0，1表示拆单
func (r *TaobaoLogisticsOnlineConfirmRequest) SetIsSplit(isSplit int64) error {
    r.isSplit = isSplit
    r.Set("is_split", isSplit)
    return nil
}

// IsSplit Getter
func (r TaobaoLogisticsOnlineConfirmRequest) GetIsSplit() int64 {
    return r.isSplit
}
// OutSid Setter
// 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
func (r *TaobaoLogisticsOnlineConfirmRequest) SetOutSid(outSid string) error {
    r.outSid = outSid
    r.Set("out_sid", outSid)
    return nil
}

// OutSid Getter
func (r TaobaoLogisticsOnlineConfirmRequest) GetOutSid() string {
    return r.outSid
}
// SellerIp Setter
// 商家的IP地址
func (r *TaobaoLogisticsOnlineConfirmRequest) SetSellerIp(sellerIp string) error {
    r.sellerIp = sellerIp
    r.Set("seller_ip", sellerIp)
    return nil
}

// SellerIp Getter
func (r TaobaoLogisticsOnlineConfirmRequest) GetSellerIp() string {
    return r.sellerIp
}
