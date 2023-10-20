package util

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimeneventproduceAPIRequest 发出奇门事件 API请求
// taobao.qimen.event.produce
//
// 当订单被处理时，用于通知奇门系统。
type TaobaoqimeneventproduceAPIRequest struct {
	model.Params
	// 事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK
	_status string
	// JSON格式扩展字段
	_ext string
	// 淘宝订单号
	_tid string
	// 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
	_platform string
	// 外部商家名称。必须同时填写platform
	_nick string
	// 主单号对应的erp单号，转单、审单、通知配货、出库 需要填。拆单、合单场景下不用填
	_erpOrderId string
	// 淘宝子订单id（拆单、合单场景下不用填，其他场景需要回传,用英文逗号隔开）
	_taobaoSubOrderIds string
	// 触发事件的时间
	_eventTime string
	// 订单创建时间,数字
	_create int64
}

// NewTaobaoqimeneventproduceRequest 初始化TaobaoqimeneventproduceAPIRequest对象
func NewTaobaoqimeneventproduceRequest() *TaobaoqimeneventproduceAPIRequest {
	return &TaobaoqimeneventproduceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimeneventproduceAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.event.produce"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimeneventproduceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimeneventproduceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatus is Status Setter
// 事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK
func (r *TaobaoqimeneventproduceAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoqimeneventproduceAPIRequest) GetStatus() string {
	return r._status
}

// SetExt is Ext Setter
// JSON格式扩展字段
func (r *TaobaoqimeneventproduceAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoqimeneventproduceAPIRequest) GetExt() string {
	return r._ext
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoqimeneventproduceAPIRequest) SetTid(_tid string) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoqimeneventproduceAPIRequest) GetTid() string {
	return r._tid
}

// SetPlatform is Platform Setter
// 商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他
func (r *TaobaoqimeneventproduceAPIRequest) SetPlatform(_platform string) error {
	r._platform = _platform
	r.Set("platform", _platform)
	return nil
}

// GetPlatform Platform Getter
func (r TaobaoqimeneventproduceAPIRequest) GetPlatform() string {
	return r._platform
}

// SetNick is Nick Setter
// 外部商家名称。必须同时填写platform
func (r *TaobaoqimeneventproduceAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoqimeneventproduceAPIRequest) GetNick() string {
	return r._nick
}

// SetErpOrderId is ErpOrderId Setter
// 主单号对应的erp单号，转单、审单、通知配货、出库 需要填。拆单、合单场景下不用填
func (r *TaobaoqimeneventproduceAPIRequest) SetErpOrderId(_erpOrderId string) error {
	r._erpOrderId = _erpOrderId
	r.Set("erp_order_id", _erpOrderId)
	return nil
}

// GetErpOrderId ErpOrderId Getter
func (r TaobaoqimeneventproduceAPIRequest) GetErpOrderId() string {
	return r._erpOrderId
}

// SetTaobaoSubOrderIds is TaobaoSubOrderIds Setter
// 淘宝子订单id（拆单、合单场景下不用填，其他场景需要回传,用英文逗号隔开）
func (r *TaobaoqimeneventproduceAPIRequest) SetTaobaoSubOrderIds(_taobaoSubOrderIds string) error {
	r._taobaoSubOrderIds = _taobaoSubOrderIds
	r.Set("taobao_sub_order_ids", _taobaoSubOrderIds)
	return nil
}

// GetTaobaoSubOrderIds TaobaoSubOrderIds Getter
func (r TaobaoqimeneventproduceAPIRequest) GetTaobaoSubOrderIds() string {
	return r._taobaoSubOrderIds
}

// SetEventTime is EventTime Setter
// 触发事件的时间
func (r *TaobaoqimeneventproduceAPIRequest) SetEventTime(_eventTime string) error {
	r._eventTime = _eventTime
	r.Set("event_time", _eventTime)
	return nil
}

// GetEventTime EventTime Getter
func (r TaobaoqimeneventproduceAPIRequest) GetEventTime() string {
	return r._eventTime
}

// SetCreate is Create Setter
// 订单创建时间,数字
func (r *TaobaoqimeneventproduceAPIRequest) SetCreate(_create int64) error {
	r._create = _create
	r.Set("create", _create)
	return nil
}

// GetCreate Create Getter
func (r TaobaoqimeneventproduceAPIRequest) GetCreate() int64 {
	return r._create
}
