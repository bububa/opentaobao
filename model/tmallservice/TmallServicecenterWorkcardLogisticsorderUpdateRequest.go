package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
物流工单履约状态更新 API请求
tmall.servicecenter.workcard.logisticsorder.update

天猫寄送类服务对接外部物流服务商回传物流状态信息
*/
type TmallServicecenterWorkcardLogisticsorderUpdateRequest struct {
    model.Params
    // 体积 单位 立方毫米
    _volume   int64
    // 重量 单位 克
    _weight   int64
    // 备注说明
    _comment   string
    // 物流单号（展示给消费者）
    _expressCode   string
    // 物流公司名词（展示给消费者）
    _expressCompany   string
    // 小件员手机号码
    _courierMobile   string
    // 小件员姓名
    _courierName   string
    // 取件码
    _gotCode   string
    // 物流订单号
    _logisticsOrderId   int64
    // 金额 单位分
    _cost   int64
    // 1、以下状态时必填： 包裹已揽收完成 2、字符串格式为：json 串 例子： [{ "name": "上衣", "count": 1 }, { "name": "裤子", "count": 2 }]
    _goodsInfo   string
    // 物流创建 ：create 物流取消 ：cancel 分派小件员：assign 已经分派小件员: assigned 包裹上门揽收: pickup_door 包裹已揽收完成: pickup_finished 包裹派送中: dispatching 包裹已签收: signed
    _statusCode   string
    // 物流子单号
    _subExpressCodes   []string
    // 预计送达时间  dispatching节点时必填
    _deliveryTime   string
    // 签收时间 signed节点时必填
    _signTime   string
    // 揽收完成时间 pickup_finished节点时必填
    _pickupFinishTime   string
    // 上门揽收时间 pickup_door节点时必填
    _pickupDoorTime   string
}

// 初始化TmallServicecenterWorkcardLogisticsorderUpdateRequest对象
func NewTmallServicecenterWorkcardLogisticsorderUpdateRequest() *TmallServicecenterWorkcardLogisticsorderUpdateRequest{
    return &TmallServicecenterWorkcardLogisticsorderUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.workcard.logisticsorder.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Volume Setter
// 体积 单位 立方毫米
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetVolume(_volume int64) error {
    r._volume = _volume
    r.Set("volume", _volume)
    return nil
}

// Volume Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetVolume() int64 {
    return r._volume
}
// Weight Setter
// 重量 单位 克
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetWeight(_weight int64) error {
    r._weight = _weight
    r.Set("weight", _weight)
    return nil
}

// Weight Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetWeight() int64 {
    return r._weight
}
// Comment Setter
// 备注说明
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetComment(_comment string) error {
    r._comment = _comment
    r.Set("comment", _comment)
    return nil
}

// Comment Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetComment() string {
    return r._comment
}
// ExpressCode Setter
// 物流单号（展示给消费者）
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetExpressCode(_expressCode string) error {
    r._expressCode = _expressCode
    r.Set("express_code", _expressCode)
    return nil
}

// ExpressCode Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetExpressCode() string {
    return r._expressCode
}
// ExpressCompany Setter
// 物流公司名词（展示给消费者）
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetExpressCompany(_expressCompany string) error {
    r._expressCompany = _expressCompany
    r.Set("express_company", _expressCompany)
    return nil
}

// ExpressCompany Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetExpressCompany() string {
    return r._expressCompany
}
// CourierMobile Setter
// 小件员手机号码
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetCourierMobile(_courierMobile string) error {
    r._courierMobile = _courierMobile
    r.Set("courier_mobile", _courierMobile)
    return nil
}

// CourierMobile Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetCourierMobile() string {
    return r._courierMobile
}
// CourierName Setter
// 小件员姓名
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetCourierName(_courierName string) error {
    r._courierName = _courierName
    r.Set("courier_name", _courierName)
    return nil
}

// CourierName Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetCourierName() string {
    return r._courierName
}
// GotCode Setter
// 取件码
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetGotCode(_gotCode string) error {
    r._gotCode = _gotCode
    r.Set("got_code", _gotCode)
    return nil
}

// GotCode Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetGotCode() string {
    return r._gotCode
}
// LogisticsOrderId Setter
// 物流订单号
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
    r._logisticsOrderId = _logisticsOrderId
    r.Set("logistics_order_id", _logisticsOrderId)
    return nil
}

// LogisticsOrderId Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetLogisticsOrderId() int64 {
    return r._logisticsOrderId
}
// Cost Setter
// 金额 单位分
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetCost(_cost int64) error {
    r._cost = _cost
    r.Set("cost", _cost)
    return nil
}

// Cost Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetCost() int64 {
    return r._cost
}
// GoodsInfo Setter
// 1、以下状态时必填： 包裹已揽收完成 2、字符串格式为：json 串 例子： [{ "name": "上衣", "count": 1 }, { "name": "裤子", "count": 2 }]
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetGoodsInfo(_goodsInfo string) error {
    r._goodsInfo = _goodsInfo
    r.Set("goods_info", _goodsInfo)
    return nil
}

// GoodsInfo Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetGoodsInfo() string {
    return r._goodsInfo
}
// StatusCode Setter
// 物流创建 ：create 物流取消 ：cancel 分派小件员：assign 已经分派小件员: assigned 包裹上门揽收: pickup_door 包裹已揽收完成: pickup_finished 包裹派送中: dispatching 包裹已签收: signed
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetStatusCode(_statusCode string) error {
    r._statusCode = _statusCode
    r.Set("status_code", _statusCode)
    return nil
}

// StatusCode Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetStatusCode() string {
    return r._statusCode
}
// SubExpressCodes Setter
// 物流子单号
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetSubExpressCodes(_subExpressCodes []string) error {
    r._subExpressCodes = _subExpressCodes
    r.Set("sub_express_codes", _subExpressCodes)
    return nil
}

// SubExpressCodes Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetSubExpressCodes() []string {
    return r._subExpressCodes
}
// DeliveryTime Setter
// 预计送达时间  dispatching节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetDeliveryTime(_deliveryTime string) error {
    r._deliveryTime = _deliveryTime
    r.Set("delivery_time", _deliveryTime)
    return nil
}

// DeliveryTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetDeliveryTime() string {
    return r._deliveryTime
}
// SignTime Setter
// 签收时间 signed节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetSignTime(_signTime string) error {
    r._signTime = _signTime
    r.Set("sign_time", _signTime)
    return nil
}

// SignTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetSignTime() string {
    return r._signTime
}
// PickupFinishTime Setter
// 揽收完成时间 pickup_finished节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetPickupFinishTime(_pickupFinishTime string) error {
    r._pickupFinishTime = _pickupFinishTime
    r.Set("pickup_finish_time", _pickupFinishTime)
    return nil
}

// PickupFinishTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetPickupFinishTime() string {
    return r._pickupFinishTime
}
// PickupDoorTime Setter
// 上门揽收时间 pickup_door节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateRequest) SetPickupDoorTime(_pickupDoorTime string) error {
    r._pickupDoorTime = _pickupDoorTime
    r.Set("pickup_door_time", _pickupDoorTime)
    return nil
}

// PickupDoorTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateRequest) GetPickupDoorTime() string {
    return r._pickupDoorTime
}
