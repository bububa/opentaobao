package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest 物流工单履约状态更新 API请求
// tmall.servicecenter.workcard.logisticsorder.update
//
// 天猫寄送类服务对接外部物流服务商回传物流状态信息
type TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest struct {
	model.Params
	// 物流子单号
	_subExpressCodes []string
	// 备注说明
	_comment string
	// 物流单号（展示给消费者）
	_expressCode string
	// 物流公司名词（展示给消费者）
	_expressCompany string
	// 小件员手机号码
	_courierMobile string
	// 小件员姓名
	_courierName string
	// 取件码
	_gotCode string
	// 1、以下状态时必填： 包裹已揽收完成 2、字符串格式为：json 串 例子： [{ "name": "上衣", "count": 1 }, { "name": "裤子", "count": 2 }]
	_goodsInfo string
	// 物流创建 ：create 物流取消 ：cancel 分派小件员：assign 已经分派小件员: assigned 包裹上门揽收: pickup_door 包裹已揽收完成: pickup_finished 包裹派送中: dispatching 包裹已签收: signed
	_statusCode string
	// 预计送达时间  dispatching节点时必填
	_deliveryTime string
	// 签收时间 signed节点时必填
	_signTime string
	// 揽收完成时间 pickup_finished节点时必填
	_pickupFinishTime string
	// 上门揽收时间 pickup_door节点时必填
	_pickupDoorTime string
	// 扩展属性
	_extendInfo string
	// 体积 单位 立方毫米
	_volume int64
	// 重量 单位 克
	_weight int64
	// 物流订单号
	_logisticsOrderId int64
	// 金额 单位分
	_cost int64
	// 是否更新新物流数据
	_newLogistics bool
}

// NewTmallServicecenterWorkcardLogisticsorderUpdateRequest 初始化TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest对象
func NewTmallServicecenterWorkcardLogisticsorderUpdateRequest() *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest {
	return &TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.logisticsorder.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubExpressCodes is SubExpressCodes Setter
// 物流子单号
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetSubExpressCodes(_subExpressCodes []string) error {
	r._subExpressCodes = _subExpressCodes
	r.Set("sub_express_codes", _subExpressCodes)
	return nil
}

// GetSubExpressCodes SubExpressCodes Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetSubExpressCodes() []string {
	return r._subExpressCodes
}

// SetComment is Comment Setter
// 备注说明
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetComment(_comment string) error {
	r._comment = _comment
	r.Set("comment", _comment)
	return nil
}

// GetComment Comment Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetComment() string {
	return r._comment
}

// SetExpressCode is ExpressCode Setter
// 物流单号（展示给消费者）
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetExpressCode(_expressCode string) error {
	r._expressCode = _expressCode
	r.Set("express_code", _expressCode)
	return nil
}

// GetExpressCode ExpressCode Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetExpressCode() string {
	return r._expressCode
}

// SetExpressCompany is ExpressCompany Setter
// 物流公司名词（展示给消费者）
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetExpressCompany(_expressCompany string) error {
	r._expressCompany = _expressCompany
	r.Set("express_company", _expressCompany)
	return nil
}

// GetExpressCompany ExpressCompany Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetExpressCompany() string {
	return r._expressCompany
}

// SetCourierMobile is CourierMobile Setter
// 小件员手机号码
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetCourierMobile(_courierMobile string) error {
	r._courierMobile = _courierMobile
	r.Set("courier_mobile", _courierMobile)
	return nil
}

// GetCourierMobile CourierMobile Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetCourierMobile() string {
	return r._courierMobile
}

// SetCourierName is CourierName Setter
// 小件员姓名
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetCourierName(_courierName string) error {
	r._courierName = _courierName
	r.Set("courier_name", _courierName)
	return nil
}

// GetCourierName CourierName Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetCourierName() string {
	return r._courierName
}

// SetGotCode is GotCode Setter
// 取件码
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetGotCode(_gotCode string) error {
	r._gotCode = _gotCode
	r.Set("got_code", _gotCode)
	return nil
}

// GetGotCode GotCode Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetGotCode() string {
	return r._gotCode
}

// SetGoodsInfo is GoodsInfo Setter
// 1、以下状态时必填： 包裹已揽收完成 2、字符串格式为：json 串 例子： [{ &#34;name&#34;: &#34;上衣&#34;, &#34;count&#34;: 1 }, { &#34;name&#34;: &#34;裤子&#34;, &#34;count&#34;: 2 }]
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetGoodsInfo(_goodsInfo string) error {
	r._goodsInfo = _goodsInfo
	r.Set("goods_info", _goodsInfo)
	return nil
}

// GetGoodsInfo GoodsInfo Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetGoodsInfo() string {
	return r._goodsInfo
}

// SetStatusCode is StatusCode Setter
// 物流创建 ：create 物流取消 ：cancel 分派小件员：assign 已经分派小件员: assigned 包裹上门揽收: pickup_door 包裹已揽收完成: pickup_finished 包裹派送中: dispatching 包裹已签收: signed
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetStatusCode(_statusCode string) error {
	r._statusCode = _statusCode
	r.Set("status_code", _statusCode)
	return nil
}

// GetStatusCode StatusCode Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetStatusCode() string {
	return r._statusCode
}

// SetDeliveryTime is DeliveryTime Setter
// 预计送达时间  dispatching节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetDeliveryTime(_deliveryTime string) error {
	r._deliveryTime = _deliveryTime
	r.Set("delivery_time", _deliveryTime)
	return nil
}

// GetDeliveryTime DeliveryTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetDeliveryTime() string {
	return r._deliveryTime
}

// SetSignTime is SignTime Setter
// 签收时间 signed节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetSignTime(_signTime string) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetSignTime() string {
	return r._signTime
}

// SetPickupFinishTime is PickupFinishTime Setter
// 揽收完成时间 pickup_finished节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetPickupFinishTime(_pickupFinishTime string) error {
	r._pickupFinishTime = _pickupFinishTime
	r.Set("pickup_finish_time", _pickupFinishTime)
	return nil
}

// GetPickupFinishTime PickupFinishTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetPickupFinishTime() string {
	return r._pickupFinishTime
}

// SetPickupDoorTime is PickupDoorTime Setter
// 上门揽收时间 pickup_door节点时必填
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetPickupDoorTime(_pickupDoorTime string) error {
	r._pickupDoorTime = _pickupDoorTime
	r.Set("pickup_door_time", _pickupDoorTime)
	return nil
}

// GetPickupDoorTime PickupDoorTime Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetPickupDoorTime() string {
	return r._pickupDoorTime
}

// SetExtendInfo is ExtendInfo Setter
// 扩展属性
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}

// SetVolume is Volume Setter
// 体积 单位 立方毫米
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetVolume(_volume int64) error {
	r._volume = _volume
	r.Set("volume", _volume)
	return nil
}

// GetVolume Volume Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetVolume() int64 {
	return r._volume
}

// SetWeight is Weight Setter
// 重量 单位 克
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetWeight(_weight int64) error {
	r._weight = _weight
	r.Set("weight", _weight)
	return nil
}

// GetWeight Weight Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetWeight() int64 {
	return r._weight
}

// SetLogisticsOrderId is LogisticsOrderId Setter
// 物流订单号
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetLogisticsOrderId(_logisticsOrderId int64) error {
	r._logisticsOrderId = _logisticsOrderId
	r.Set("logistics_order_id", _logisticsOrderId)
	return nil
}

// GetLogisticsOrderId LogisticsOrderId Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetLogisticsOrderId() int64 {
	return r._logisticsOrderId
}

// SetCost is Cost Setter
// 金额 单位分
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetCost(_cost int64) error {
	r._cost = _cost
	r.Set("cost", _cost)
	return nil
}

// GetCost Cost Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetCost() int64 {
	return r._cost
}

// SetNewLogistics is NewLogistics Setter
// 是否更新新物流数据
func (r *TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) SetNewLogistics(_newLogistics bool) error {
	r._newLogistics = _newLogistics
	r.Set("new_logistics", _newLogistics)
	return nil
}

// GetNewLogistics NewLogistics Getter
func (r TmallServicecenterWorkcardLogisticsorderUpdateAPIRequest) GetNewLogistics() bool {
	return r._newLogistics
}
