package wlb

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家获取生鲜电子面单号 APIRequest
taobao.wlb.waybill.shengxian.get

商家通过交易订单号获取电子面单接口
*/
type TaobaoWlbWaybillShengxianGetRequest struct {
    model.Params

    // 物流服务方代码，生鲜配送：YXSR
    bizCode   string 

    // 物流服务类型。冷链：602，常温：502
    deliveryType   string 

    // 商家淘宝地址信息ID
    senderAddressId   string 

    // 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)
    serviceCode   string 

    // 订单渠道： 淘宝平台订单："TB"; 天猫平台订单："TM"; 京东："JD"; 拍拍："PP"; 易迅平台订单："YX"; 一号店平台订单："YHD"; 当当网平台订单："DD"; EBAY："EBAY"; QQ网购："QQ"; 苏宁："SN"; 国美："GM"; 唯品会："WPH"; 聚美："Jm"; 乐峰："LF"; 蘑菇街："MGJ"; 聚尚："JS"; 银泰："YT"; VANCL："VANCL"; 邮乐："YL"; 优购："YG"; 拍鞋："PX"; 1688平台："1688";
    orderChannelsType   string 

    // 交易号，传入交易号不能存在拆单场景。
    tradeId   string 

    // 预留扩展字段
    feature   string 

}

func NewTaobaoWlbWaybillShengxianGetRequest() *TaobaoWlbWaybillShengxianGetRequest{
    return &TaobaoWlbWaybillShengxianGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetApiMethodName() string {
    return "taobao.wlb.waybill.shengxian.get"
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoWlbWaybillShengxianGetRequest) SetBizCode(bizCode string) error {
    r.bizCode = bizCode
    r.Set("biz_code", bizCode)
    return nil
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetBizCode() string {
    return r.bizCode
}

func (r *TaobaoWlbWaybillShengxianGetRequest) SetDeliveryType(deliveryType string) error {
    r.deliveryType = deliveryType
    r.Set("delivery_type", deliveryType)
    return nil
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetDeliveryType() string {
    return r.deliveryType
}

func (r *TaobaoWlbWaybillShengxianGetRequest) SetSenderAddressId(senderAddressId string) error {
    r.senderAddressId = senderAddressId
    r.Set("sender_address_id", senderAddressId)
    return nil
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetSenderAddressId() string {
    return r.senderAddressId
}

func (r *TaobaoWlbWaybillShengxianGetRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetServiceCode() string {
    return r.serviceCode
}

func (r *TaobaoWlbWaybillShengxianGetRequest) SetOrderChannelsType(orderChannelsType string) error {
    r.orderChannelsType = orderChannelsType
    r.Set("order_channels_type", orderChannelsType)
    return nil
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetOrderChannelsType() string {
    return r.orderChannelsType
}

func (r *TaobaoWlbWaybillShengxianGetRequest) SetTradeId(tradeId string) error {
    r.tradeId = tradeId
    r.Set("trade_id", tradeId)
    return nil
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetTradeId() string {
    return r.tradeId
}

func (r *TaobaoWlbWaybillShengxianGetRequest) SetFeature(feature string) error {
    r.feature = feature
    r.Set("feature", feature)
    return nil
}

func (r TaobaoWlbWaybillShengxianGetRequest) GetFeature() string {
    return r.feature
}

