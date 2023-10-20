package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsPartnersGetAPIRequest 查询支持起始地到目的地范围的物流公司 API请求
// taobao.logistics.partners.get
//
// 查询物流公司信息（可以查询目的地可不可达情况）
type TaobaoLogisticsPartnersGetAPIRequest struct {
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

// NewTaobaoLogisticsPartnersGetRequest 初始化TaobaoLogisticsPartnersGetAPIRequest对象
func NewTaobaoLogisticsPartnersGetRequest() *TaobaoLogisticsPartnersGetAPIRequest {
	return &TaobaoLogisticsPartnersGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsPartnersGetAPIRequest) Reset() {
	r._sourceId = ""
	r._targetId = ""
	r._serviceType = ""
	r._goodsValue = ""
	r._isNeedCarriage = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsPartnersGetAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.partners.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsPartnersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsPartnersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSourceId is SourceId Setter
// 物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
func (r *TaobaoLogisticsPartnersGetAPIRequest) SetSourceId(_sourceId string) error {
	r._sourceId = _sourceId
	r.Set("source_id", _sourceId)
	return nil
}

// GetSourceId SourceId Getter
func (r TaobaoLogisticsPartnersGetAPIRequest) GetSourceId() string {
	return r._sourceId
}

// SetTargetId is TargetId Setter
// 物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取
func (r *TaobaoLogisticsPartnersGetAPIRequest) SetTargetId(_targetId string) error {
	r._targetId = _targetId
	r.Set("target_id", _targetId)
	return nil
}

// GetTargetId TargetId Getter
func (r TaobaoLogisticsPartnersGetAPIRequest) GetTargetId() string {
	return r._targetId
}

// SetServiceType is ServiceType Setter
// 服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。
func (r *TaobaoLogisticsPartnersGetAPIRequest) SetServiceType(_serviceType string) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r TaobaoLogisticsPartnersGetAPIRequest) GetServiceType() string {
	return r._serviceType
}

// SetGoodsValue is GoodsValue Setter
// 货物价格.只有当选择货到付款此参数才会有效
func (r *TaobaoLogisticsPartnersGetAPIRequest) SetGoodsValue(_goodsValue string) error {
	r._goodsValue = _goodsValue
	r.Set("goods_value", _goodsValue)
	return nil
}

// GetGoodsValue GoodsValue Getter
func (r TaobaoLogisticsPartnersGetAPIRequest) GetGoodsValue() string {
	return r._goodsValue
}

// SetIsNeedCarriage is IsNeedCarriage Setter
// 是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。
func (r *TaobaoLogisticsPartnersGetAPIRequest) SetIsNeedCarriage(_isNeedCarriage bool) error {
	r._isNeedCarriage = _isNeedCarriage
	r.Set("is_need_carriage", _isNeedCarriage)
	return nil
}

// GetIsNeedCarriage IsNeedCarriage Getter
func (r TaobaoLogisticsPartnersGetAPIRequest) GetIsNeedCarriage() bool {
	return r._isNeedCarriage
}

var poolTaobaoLogisticsPartnersGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsPartnersGetRequest()
	},
}

// GetTaobaoLogisticsPartnersGetRequest 从 sync.Pool 获取 TaobaoLogisticsPartnersGetAPIRequest
func GetTaobaoLogisticsPartnersGetAPIRequest() *TaobaoLogisticsPartnersGetAPIRequest {
	return poolTaobaoLogisticsPartnersGetAPIRequest.Get().(*TaobaoLogisticsPartnersGetAPIRequest)
}

// ReleaseTaobaoLogisticsPartnersGetAPIRequest 将 TaobaoLogisticsPartnersGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsPartnersGetAPIRequest(v *TaobaoLogisticsPartnersGetAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsPartnersGetAPIRequest.Put(v)
}
