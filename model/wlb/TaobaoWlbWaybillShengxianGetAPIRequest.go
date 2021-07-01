package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWaybillShengxianGetAPIRequest
商家获取生鲜电子面单号 API请求
taobao.wlb.waybill.shengxian.get

商家通过交易订单号获取电子面单接口 */
type TaobaoWlbWaybillShengxianGetAPIRequest struct {
	model.Params
	// 物流服务方代码，生鲜配送：YXSR
	_bizCode string
	// 物流服务类型。冷链：602，常温：502
	_deliveryType string
	// 商家淘宝地址信息ID
	_senderAddressId string
	// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)
	_serviceCode string
	// 订单渠道： 淘宝平台订单："TB"; 天猫平台订单："TM"; 京东："JD"; 拍拍："PP"; 易迅平台订单："YX"; 一号店平台订单："YHD"; 当当网平台订单："DD"; EBAY："EBAY"; QQ网购："QQ"; 苏宁："SN"; 国美："GM"; 唯品会："WPH"; 聚美："Jm"; 乐峰："LF"; 蘑菇街："MGJ"; 聚尚："JS"; 银泰："YT"; VANCL："VANCL"; 邮乐："YL"; 优购："YG"; 拍鞋："PX"; 1688平台："1688";
	_orderChannelsType string
	// 交易号，传入交易号不能存在拆单场景。
	_tradeId string
	// 预留扩展字段
	_feature string
}

// NewTaobaoWlbWaybillShengxianGetRequest 初始化TaobaoWlbWaybillShengxianGetAPIRequest对象
func NewTaobaoWlbWaybillShengxianGetRequest() *TaobaoWlbWaybillShengxianGetAPIRequest {
	return &TaobaoWlbWaybillShengxianGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.waybill.shengxian.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BizCode Setter
// 物流服务方代码，生鲜配送：YXSR
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// Get BizCode Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetBizCode() string {
	return r._bizCode
}

// Set is DeliveryType Setter
// 物流服务类型。冷链：602，常温：502
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetDeliveryType(_deliveryType string) error {
	r._deliveryType = _deliveryType
	r.Set("delivery_type", _deliveryType)
	return nil
}

// Get DeliveryType Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetDeliveryType() string {
	return r._deliveryType
}

// Set is SenderAddressId Setter
// 商家淘宝地址信息ID
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetSenderAddressId(_senderAddressId string) error {
	r._senderAddressId = _senderAddressId
	r.Set("sender_address_id", _senderAddressId)
	return nil
}

// Get SenderAddressId Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetSenderAddressId() string {
	return r._senderAddressId
}

// Set is ServiceCode Setter
// 仓库的服务代码标示，代码一个仓库的实体。(可以通过taobao.wlb.stores.baseinfo.get接口查询)
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetServiceCode(_serviceCode string) error {
	r._serviceCode = _serviceCode
	r.Set("service_code", _serviceCode)
	return nil
}

// Get ServiceCode Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetServiceCode() string {
	return r._serviceCode
}

// Set is OrderChannelsType Setter
// 订单渠道： 淘宝平台订单："TB"; 天猫平台订单："TM"; 京东："JD"; 拍拍："PP"; 易迅平台订单："YX"; 一号店平台订单："YHD"; 当当网平台订单："DD"; EBAY："EBAY"; QQ网购："QQ"; 苏宁："SN"; 国美："GM"; 唯品会："WPH"; 聚美："Jm"; 乐峰："LF"; 蘑菇街："MGJ"; 聚尚："JS"; 银泰："YT"; VANCL："VANCL"; 邮乐："YL"; 优购："YG"; 拍鞋："PX"; 1688平台："1688";
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetOrderChannelsType(_orderChannelsType string) error {
	r._orderChannelsType = _orderChannelsType
	r.Set("order_channels_type", _orderChannelsType)
	return nil
}

// Get OrderChannelsType Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetOrderChannelsType() string {
	return r._orderChannelsType
}

// Set is TradeId Setter
// 交易号，传入交易号不能存在拆单场景。
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetTradeId(_tradeId string) error {
	r._tradeId = _tradeId
	r.Set("trade_id", _tradeId)
	return nil
}

// Get TradeId Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetTradeId() string {
	return r._tradeId
}

// Set is Feature Setter
// 预留扩展字段
func (r *TaobaoWlbWaybillShengxianGetAPIRequest) SetFeature(_feature string) error {
	r._feature = _feature
	r.Set("feature", _feature)
	return nil
}

// Get Feature Getter
func (r TaobaoWlbWaybillShengxianGetAPIRequest) GetFeature() string {
	return r._feature
}
