package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticspartnersgetAPIRequest 查询支持起始地到目的地范围的物流公司 API请求
// taobao.logistics.partners.get
//
// 查询物流公司信息（可以查询目的地可不可达情况）
type TaobaologisticspartnersgetAPIRequest struct {
	model.Params
	// 物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
	_sourceId string
	// 物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
	_targetId string
	// 服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。
	_serviceType string
	// 货物价格.只有当选择货到付款此参数才会有效
	_goodsValue string
	// 是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。
	_isNeedCarriage bool
}

// NewTaobaologisticspartnersgetRequest 初始化TaobaologisticspartnersgetAPIRequest对象
func NewTaobaologisticspartnersgetRequest() *TaobaologisticspartnersgetAPIRequest {
	return &TaobaologisticspartnersgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticspartnersgetAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.partners.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticspartnersgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticspartnersgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSourceId is SourceId Setter
// 物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
func (r *TaobaologisticspartnersgetAPIRequest) SetSourceId(_sourceId string) error {
	r._sourceId = _sourceId
	r.Set("source_id", _sourceId)
	return nil
}

// GetSourceId SourceId Getter
func (r TaobaologisticspartnersgetAPIRequest) GetSourceId() string {
	return r._sourceId
}

// SetTargetId is TargetId Setter
// 物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
func (r *TaobaologisticspartnersgetAPIRequest) SetTargetId(_targetId string) error {
	r._targetId = _targetId
	r.Set("target_id", _targetId)
	return nil
}

// GetTargetId TargetId Getter
func (r TaobaologisticspartnersgetAPIRequest) GetTargetId() string {
	return r._targetId
}

// SetServiceType is ServiceType Setter
// 服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。
func (r *TaobaologisticspartnersgetAPIRequest) SetServiceType(_serviceType string) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r TaobaologisticspartnersgetAPIRequest) GetServiceType() string {
	return r._serviceType
}

// SetGoodsValue is GoodsValue Setter
// 货物价格.只有当选择货到付款此参数才会有效
func (r *TaobaologisticspartnersgetAPIRequest) SetGoodsValue(_goodsValue string) error {
	r._goodsValue = _goodsValue
	r.Set("goods_value", _goodsValue)
	return nil
}

// GetGoodsValue GoodsValue Getter
func (r TaobaologisticspartnersgetAPIRequest) GetGoodsValue() string {
	return r._goodsValue
}

// SetIsNeedCarriage is IsNeedCarriage Setter
// 是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。
func (r *TaobaologisticspartnersgetAPIRequest) SetIsNeedCarriage(_isNeedCarriage bool) error {
	r._isNeedCarriage = _isNeedCarriage
	r.Set("is_need_carriage", _isNeedCarriage)
	return nil
}

// GetIsNeedCarriage IsNeedCarriage Getter
func (r TaobaologisticspartnersgetAPIRequest) GetIsNeedCarriage() bool {
	return r._isNeedCarriage
}
