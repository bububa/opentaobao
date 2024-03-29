package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemElementManageAPIRequest 【API3.0】资源元素管理接口 API请求
// taobao.alitrip.travel.item.element.manage
//
// 资源元素管理接口：提供商家管理（增删改）基本资源元素信息。基本资源元素可供多个商品共享
type TaobaoAlitripTravelItemElementManageAPIRequest struct {
	model.Params
	// 必填，元素的外部商家编码，必须唯一。编辑、删除时将根据该编码找到对应元素。
	_outerId string
	// 元素名称，新增时必填； 注意：Wifi库的使用地和签证库所在国家均适用这个字段
	_name string
	// 元素所在城市，景点、酒店在新增时必填
	_city string
	// 元素的子类型，新增时必填。景点指门票类型，酒店指房型信息，交通接驳（接送机、接驳车、租车、船票、其他）选其一，餐饮（早餐、晚餐、午餐、下午茶及其他）选其一；签证（旅游签证、商务签证、工作签证、留学签证、探亲访友签证、入台证、其他）
	_type string
	// 当新增“交通接驳、餐饮、特色活动、其他”资源类型时 必填
	_desc string
	// 必填，操作类型：1-新增，2-修改，3-删除。。特别注意：删除 为逻辑删除，即该outer_id所对应的元素还存在但是会置为无效状态，重新编辑修改即可恢复为有效状态。因此该id一旦使用将不可重复
	_operation int64
	// 资源元素类型，新增时必填：1-景点，2-酒店，5-交通接驳，6-WIFI库，7-电话卡，8-餐饮，9-签证库，11-特色活动，999-其他
	_elementType int64
}

// NewTaobaoAlitripTravelItemElementManageRequest 初始化TaobaoAlitripTravelItemElementManageAPIRequest对象
func NewTaobaoAlitripTravelItemElementManageRequest() *TaobaoAlitripTravelItemElementManageAPIRequest {
	return &TaobaoAlitripTravelItemElementManageAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) Reset() {
	r._outerId = ""
	r._name = ""
	r._city = ""
	r._type = ""
	r._desc = ""
	r._operation = 0
	r._elementType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.item.element.manage"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 必填，元素的外部商家编码，必须唯一。编辑、删除时将根据该编码找到对应元素。
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetName is Name Setter
// 元素名称，新增时必填； 注意：Wifi库的使用地和签证库所在国家均适用这个字段
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetName() string {
	return r._name
}

// SetCity is City Setter
// 元素所在城市，景点、酒店在新增时必填
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) SetCity(_city string) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetCity() string {
	return r._city
}

// SetType is Type Setter
// 元素的子类型，新增时必填。景点指门票类型，酒店指房型信息，交通接驳（接送机、接驳车、租车、船票、其他）选其一，餐饮（早餐、晚餐、午餐、下午茶及其他）选其一；签证（旅游签证、商务签证、工作签证、留学签证、探亲访友签证、入台证、其他）
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetType() string {
	return r._type
}

// SetDesc is Desc Setter
// 当新增“交通接驳、餐饮、特色活动、其他”资源类型时 必填
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetDesc() string {
	return r._desc
}

// SetOperation is Operation Setter
// 必填，操作类型：1-新增，2-修改，3-删除。。特别注意：删除 为逻辑删除，即该outer_id所对应的元素还存在但是会置为无效状态，重新编辑修改即可恢复为有效状态。因此该id一旦使用将不可重复
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) SetOperation(_operation int64) error {
	r._operation = _operation
	r.Set("operation", _operation)
	return nil
}

// GetOperation Operation Getter
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetOperation() int64 {
	return r._operation
}

// SetElementType is ElementType Setter
// 资源元素类型，新增时必填：1-景点，2-酒店，5-交通接驳，6-WIFI库，7-电话卡，8-餐饮，9-签证库，11-特色活动，999-其他
func (r *TaobaoAlitripTravelItemElementManageAPIRequest) SetElementType(_elementType int64) error {
	r._elementType = _elementType
	r.Set("element_type", _elementType)
	return nil
}

// GetElementType ElementType Getter
func (r TaobaoAlitripTravelItemElementManageAPIRequest) GetElementType() int64 {
	return r._elementType
}

var poolTaobaoAlitripTravelItemElementManageAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelItemElementManageRequest()
	},
}

// GetTaobaoAlitripTravelItemElementManageRequest 从 sync.Pool 获取 TaobaoAlitripTravelItemElementManageAPIRequest
func GetTaobaoAlitripTravelItemElementManageAPIRequest() *TaobaoAlitripTravelItemElementManageAPIRequest {
	return poolTaobaoAlitripTravelItemElementManageAPIRequest.Get().(*TaobaoAlitripTravelItemElementManageAPIRequest)
}

// ReleaseTaobaoAlitripTravelItemElementManageAPIRequest 将 TaobaoAlitripTravelItemElementManageAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelItemElementManageAPIRequest(v *TaobaoAlitripTravelItemElementManageAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelItemElementManageAPIRequest.Put(v)
}
