package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterSpserviceorderCreateAPIRequest 服务单创建 API请求
// tmall.servicecenter.spserviceorder.create
//
// 服务单创建
type TmallServicecenterSpserviceorderCreateAPIRequest struct {
	model.Params
	// 服务对象类型
	_serviceObjectType string
	// 服务对象类型名称
	_serviceObjectName string
	// 服务内容
	_serviceContent string
	// 服务计价json
	_serviceProperties string
	// 外部单号
	_outOrderId string
	// 外部渠道
	_source string
	// 额外业务信息
	_extJson string
	// 备注
	_memo string
	// 服务产品id
	_serviceProductId int64
	// 品牌id
	_brandId int64
	// 消费者信息
	_buyer *BuyerDto
	// 预约信息
	_reservation *ReservationDto
	// 数量
	_serviceCount int64
}

// NewTmallServicecenterSpserviceorderCreateRequest 初始化TmallServicecenterSpserviceorderCreateAPIRequest对象
func NewTmallServicecenterSpserviceorderCreateRequest() *TmallServicecenterSpserviceorderCreateAPIRequest {
	return &TmallServicecenterSpserviceorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceObjectType is ServiceObjectType Setter
// 服务对象类型
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetServiceObjectType(_serviceObjectType string) error {
	r._serviceObjectType = _serviceObjectType
	r.Set("service_object_type", _serviceObjectType)
	return nil
}

// GetServiceObjectType ServiceObjectType Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetServiceObjectType() string {
	return r._serviceObjectType
}

// SetServiceObjectName is ServiceObjectName Setter
// 服务对象类型名称
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetServiceObjectName(_serviceObjectName string) error {
	r._serviceObjectName = _serviceObjectName
	r.Set("service_object_name", _serviceObjectName)
	return nil
}

// GetServiceObjectName ServiceObjectName Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetServiceObjectName() string {
	return r._serviceObjectName
}

// SetServiceContent is ServiceContent Setter
// 服务内容
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetServiceContent(_serviceContent string) error {
	r._serviceContent = _serviceContent
	r.Set("service_content", _serviceContent)
	return nil
}

// GetServiceContent ServiceContent Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetServiceContent() string {
	return r._serviceContent
}

// SetServiceProperties is ServiceProperties Setter
// 服务计价json
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetServiceProperties(_serviceProperties string) error {
	r._serviceProperties = _serviceProperties
	r.Set("service_properties", _serviceProperties)
	return nil
}

// GetServiceProperties ServiceProperties Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetServiceProperties() string {
	return r._serviceProperties
}

// SetOutOrderId is OutOrderId Setter
// 外部单号
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetOutOrderId(_outOrderId string) error {
	r._outOrderId = _outOrderId
	r.Set("out_order_id", _outOrderId)
	return nil
}

// GetOutOrderId OutOrderId Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetOutOrderId() string {
	return r._outOrderId
}

// SetSource is Source Setter
// 外部渠道
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetSource() string {
	return r._source
}

// SetExtJson is ExtJson Setter
// 额外业务信息
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetMemo is Memo Setter
// 备注
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetMemo(_memo string) error {
	r._memo = _memo
	r.Set("memo", _memo)
	return nil
}

// GetMemo Memo Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetMemo() string {
	return r._memo
}

// SetServiceProductId is ServiceProductId Setter
// 服务产品id
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetServiceProductId(_serviceProductId int64) error {
	r._serviceProductId = _serviceProductId
	r.Set("service_product_id", _serviceProductId)
	return nil
}

// GetServiceProductId ServiceProductId Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetServiceProductId() int64 {
	return r._serviceProductId
}

// SetBrandId is BrandId Setter
// 品牌id
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetBrandId() int64 {
	return r._brandId
}

// SetBuyer is Buyer Setter
// 消费者信息
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetBuyer(_buyer *BuyerDto) error {
	r._buyer = _buyer
	r.Set("buyer", _buyer)
	return nil
}

// GetBuyer Buyer Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetBuyer() *BuyerDto {
	return r._buyer
}

// SetReservation is Reservation Setter
// 预约信息
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetReservation(_reservation *ReservationDto) error {
	r._reservation = _reservation
	r.Set("reservation", _reservation)
	return nil
}

// GetReservation Reservation Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetReservation() *ReservationDto {
	return r._reservation
}

// SetServiceCount is ServiceCount Setter
// 数量
func (r *TmallServicecenterSpserviceorderCreateAPIRequest) SetServiceCount(_serviceCount int64) error {
	r._serviceCount = _serviceCount
	r.Set("service_count", _serviceCount)
	return nil
}

// GetServiceCount ServiceCount Getter
func (r TmallServicecenterSpserviceorderCreateAPIRequest) GetServiceCount() int64 {
	return r._serviceCount
}
