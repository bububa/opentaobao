package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无需物流（虚拟）发货处理 APIRequest
taobao.logistics.dummy.send

用户调用该接口可实现无需物流（虚拟）发货,使用该接口发货，交易订单状态会直接变成卖家已发货
*/
type TaobaoLogisticsDummySendRequest struct {
    model.Params

    // 淘宝交易ID
    tid   int64 

    // feature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。
    feature   string 

    // 商家的IP地址
    sellerIp   string 

}

func NewTaobaoLogisticsDummySendRequest() *TaobaoLogisticsDummySendRequest{
    return &TaobaoLogisticsDummySendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsDummySendRequest) GetApiMethodName() string {
    return "taobao.logistics.dummy.send"
}

func (r TaobaoLogisticsDummySendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsDummySendRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoLogisticsDummySendRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoLogisticsDummySendRequest) SetFeature(feature string) error {
    r.feature = feature
    r.Set("feature", feature)
    return nil
}

func (r TaobaoLogisticsDummySendRequest) GetFeature() string {
    return r.feature
}

func (r *TaobaoLogisticsDummySendRequest) SetSellerIp(sellerIp string) error {
    r.sellerIp = sellerIp
    r.Set("seller_ip", sellerIp)
    return nil
}

func (r TaobaoLogisticsDummySendRequest) GetSellerIp() string {
    return r.sellerIp
}

