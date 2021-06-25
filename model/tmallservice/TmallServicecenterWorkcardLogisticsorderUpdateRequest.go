package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流工单履约状态更新 APIRequest
tmall.servicecenter.workcard.logisticsorder.update

天猫寄送类服务对接外部物流服务商回传物流状态信息
*/
type TmallServicecenterWorkcardLogisticsorderUpdateRequest struct {
    model.Params

    // 体积 单位 立方毫米
    volume   int64 

    // 重量 单位 克
    weight   int64 

    // 备注说明
    comment   string 

    // 物流单号（展示给消费者）
    expressCode   string 

    // 物流公司名词（展示给消费者）
    expressCompany   string 

    // 小件员手机号码
    courierMobile   string 

    // 小件员姓名
    courierName   string 

    // 取件码
    gotCode   string 

    // 物流订单号
    logisticsOrderId   int64 

    // 金额 单位分
    cost   int64 

    // 1、以下状态时必填： 包裹已揽收完成 2、字符串格式为：json 串 例子： [{ "name": "上衣", "count": 1 }, { "name": "裤子", "count": 2 }]
    goodsInfo   string 

    // 物流创建 ：create 物流取消 ：cancel 分派小件员：assign 已经分派小件员: assigned 包裹上门揽收: pickup_door 包裹已揽收完成: pickup_finished 包裹派送中: dispatching 包裹已签收: signed
    statusCode   string 

    // 物流子单号
    subExpressCodes   []String 

    // 预计送达时间  dispatching节点时必填
    deliveryTime   string 

    // 签收时间 signed节点时必填
    signTime   string 

    // 揽收完成时间 pickup_finished节点时必填
    pickupFinishTime   string 

    // 上门揽收时间 pickup_door节点时必填
    pickupDoorTime   string 

}

func NewTmallServicecenterWorkcardLogisticsorderUpdateRequest() *TmallServicecenterWorkcardLogisticsorderUpdateRequest{
    return &TmallServicecenterWorkcardLogisticsorderUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.logisticsorder.update"
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetVolume(volume int64) error {
    r.volume = volume
    r.Set("volume", volume)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetVolume() int64 {
    return r.volume
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetWeight(weight int64) error {
    r.weight = weight
    r.Set("weight", weight)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetWeight() int64 {
    return r.weight
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetComment(comment string) error {
    r.comment = comment
    r.Set("comment", comment)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetComment() string {
    return r.comment
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetExpressCode(expressCode string) error {
    r.expressCode = expressCode
    r.Set("express_code", expressCode)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetExpressCode() string {
    return r.expressCode
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetExpressCompany(expressCompany string) error {
    r.expressCompany = expressCompany
    r.Set("express_company", expressCompany)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetExpressCompany() string {
    return r.expressCompany
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetCourierMobile(courierMobile string) error {
    r.courierMobile = courierMobile
    r.Set("courier_mobile", courierMobile)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetCourierMobile() string {
    return r.courierMobile
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetCourierName(courierName string) error {
    r.courierName = courierName
    r.Set("courier_name", courierName)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetCourierName() string {
    return r.courierName
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetGotCode(gotCode string) error {
    r.gotCode = gotCode
    r.Set("got_code", gotCode)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetGotCode() string {
    return r.gotCode
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetLogisticsOrderId(logisticsOrderId int64) error {
    r.logisticsOrderId = logisticsOrderId
    r.Set("logistics_order_id", logisticsOrderId)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetLogisticsOrderId() int64 {
    return r.logisticsOrderId
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetCost(cost int64) error {
    r.cost = cost
    r.Set("cost", cost)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetCost() int64 {
    return r.cost
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetGoodsInfo(goodsInfo string) error {
    r.goodsInfo = goodsInfo
    r.Set("goods_info", goodsInfo)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetGoodsInfo() string {
    return r.goodsInfo
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetStatusCode(statusCode string) error {
    r.statusCode = statusCode
    r.Set("status_code", statusCode)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetStatusCode() string {
    return r.statusCode
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetSubExpressCodes(subExpressCodes []String) error {
    r.subExpressCodes = subExpressCodes
    r.Set("sub_express_codes", subExpressCodes)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetSubExpressCodes() []String {
    return r.subExpressCodes
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetDeliveryTime(deliveryTime string) error {
    r.deliveryTime = deliveryTime
    r.Set("delivery_time", deliveryTime)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetDeliveryTime() string {
    return r.deliveryTime
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetSignTime(signTime string) error {
    r.signTime = signTime
    r.Set("sign_time", signTime)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetSignTime() string {
    return r.signTime
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetPickupFinishTime(pickupFinishTime string) error {
    r.pickupFinishTime = pickupFinishTime
    r.Set("pickup_finish_time", pickupFinishTime)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetPickupFinishTime() string {
    return r.pickupFinishTime
}

func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetPickupDoorTime(pickupDoorTime string) error {
    r.pickupDoorTime = pickupDoorTime
    r.Set("pickup_door_time", pickupDoorTime)
    return nil
}

func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetPickupDoorTime() string {
    return r.pickupDoorTime
}

