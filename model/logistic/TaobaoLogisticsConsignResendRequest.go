package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
修改物流公司和运单号 APIRequest
taobao.logistics.consign.resend

支持卖家发货后修改运单号;支持在线下单和自己联系两种发货方式;使用条件：<br><br/>1、必须是已发货订单，自己联系发货的必须24小时内才可修改；在线下单的，必须下单后物流公司未揽收成功前才可修改；<br/>2、自己联系只能切换为自己联系的公司，在线下单也只能切换为在线下单的物流公司。
*/
type TaobaoLogisticsConsignResendRequest struct {
    model.Params

    // 淘宝交易ID
    tid   int64 

    // 拆单子订单列表，对应的数据是：子订单号列表。可以不传，但是如果传了则必须符合传递的规则。子订单必须是操作的物流订单的子订单的真子集！
    subTid   []Number 

    // 表明是否是拆单，默认值0，1表示拆单
    isSplit   int64 

    // 运单号.具体一个物流公司的真实运单号码。淘宝官方物流会校验，请谨慎传入；
    outSid   string 

    // 物流公司代码.如"POST"代表中国邮政,"ZJS"代表宅急送。调用 taobao.logistics.companies.get 获取。<br><font color='red'>如果是货到付款订单，选择的物流公司必须支持货到付款发货方式</font>
    companyCode   string 

    // feature参数格式<br>范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B<br>identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔<br>“tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。"|"不同商品间的分隔符。<br>例1商品和2商品，之间就用"|"分开。<br>TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上)<br>":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。<br>"," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。<br>具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。<br>参数格式：identCode=TIDA:12345,67890。TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。<br>当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。
    feature   string 

    // 商家的IP地址
    sellerIp   string 

}

func NewTaobaoLogisticsConsignResendRequest() *TaobaoLogisticsConsignResendRequest{
    return &TaobaoLogisticsConsignResendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsConsignResendRequest) GetApiMethodName() string {
    return "taobao.logistics.consign.resend"
}

func (r TaobaoLogisticsConsignResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsConsignResendRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

func (r TaobaoLogisticsConsignResendRequest) GetTid() int64 {
    return r.tid
}

func (r *TaobaoLogisticsConsignResendRequest) SetSubTid(subTid []Number) error {
    r.subTid = subTid
    r.Set("sub_tid", subTid)
    return nil
}

func (r TaobaoLogisticsConsignResendRequest) GetSubTid() []Number {
    return r.subTid
}

func (r *TaobaoLogisticsConsignResendRequest) SetIsSplit(isSplit int64) error {
    r.isSplit = isSplit
    r.Set("is_split", isSplit)
    return nil
}

func (r TaobaoLogisticsConsignResendRequest) GetIsSplit() int64 {
    return r.isSplit
}

func (r *TaobaoLogisticsConsignResendRequest) SetOutSid(outSid string) error {
    r.outSid = outSid
    r.Set("out_sid", outSid)
    return nil
}

func (r TaobaoLogisticsConsignResendRequest) GetOutSid() string {
    return r.outSid
}

func (r *TaobaoLogisticsConsignResendRequest) SetCompanyCode(companyCode string) error {
    r.companyCode = companyCode
    r.Set("company_code", companyCode)
    return nil
}

func (r TaobaoLogisticsConsignResendRequest) GetCompanyCode() string {
    return r.companyCode
}

func (r *TaobaoLogisticsConsignResendRequest) SetFeature(feature string) error {
    r.feature = feature
    r.Set("feature", feature)
    return nil
}

func (r TaobaoLogisticsConsignResendRequest) GetFeature() string {
    return r.feature
}

func (r *TaobaoLogisticsConsignResendRequest) SetSellerIp(sellerIp string) error {
    r.sellerIp = sellerIp
    r.Set("seller_ip", sellerIp)
    return nil
}

func (r TaobaoLogisticsConsignResendRequest) GetSellerIp() string {
    return r.sellerIp
}

