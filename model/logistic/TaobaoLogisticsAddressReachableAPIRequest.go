package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsAddressReachableAPIRequest 判定服务是否可达 API请求
// taobao.logistics.address.reachable
//
// 根据输入的目标地址，判断服务是否可达。
// 现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
type TaobaoLogisticsAddressReachableAPIRequest struct {
	model.Params
	// 标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105
	_areaCode string
	// 详细地址
	_address string
	// 物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”
	_partnerIds string
	// 服务对应的数字编码，如揽派范围对应88
	_serviceType int64
	// 发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011
	_sourceAreaCode string
}

// NewTaobaoLogisticsAddressReachableRequest 初始化TaobaoLogisticsAddressReachableAPIRequest对象
func NewTaobaoLogisticsAddressReachableRequest() *TaobaoLogisticsAddressReachableAPIRequest {
	return &TaobaoLogisticsAddressReachableAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressReachableAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.address.reachable"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressReachableAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAreaCode is AreaCode Setter
// 标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105
func (r *TaobaoLogisticsAddressReachableAPIRequest) SetAreaCode(_areaCode string) error {
	r._areaCode = _areaCode
	r.Set("area_code", _areaCode)
	return nil
}

// GetAreaCode AreaCode Getter
func (r TaobaoLogisticsAddressReachableAPIRequest) GetAreaCode() string {
	return r._areaCode
}

// SetAddress is Address Setter
// 详细地址
func (r *TaobaoLogisticsAddressReachableAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r TaobaoLogisticsAddressReachableAPIRequest) GetAddress() string {
	return r._address
}

// SetPartnerIds is PartnerIds Setter
// 物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”
func (r *TaobaoLogisticsAddressReachableAPIRequest) SetPartnerIds(_partnerIds string) error {
	r._partnerIds = _partnerIds
	r.Set("partner_ids", _partnerIds)
	return nil
}

// GetPartnerIds PartnerIds Getter
func (r TaobaoLogisticsAddressReachableAPIRequest) GetPartnerIds() string {
	return r._partnerIds
}

// SetServiceType is ServiceType Setter
// 服务对应的数字编码，如揽派范围对应88
func (r *TaobaoLogisticsAddressReachableAPIRequest) SetServiceType(_serviceType int64) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// GetServiceType ServiceType Getter
func (r TaobaoLogisticsAddressReachableAPIRequest) GetServiceType() int64 {
	return r._serviceType
}

// SetSourceAreaCode is SourceAreaCode Setter
// 发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011
func (r *TaobaoLogisticsAddressReachableAPIRequest) SetSourceAreaCode(_sourceAreaCode string) error {
	r._sourceAreaCode = _sourceAreaCode
	r.Set("source_area_code", _sourceAreaCode)
	return nil
}

// GetSourceAreaCode SourceAreaCode Getter
func (r TaobaoLogisticsAddressReachableAPIRequest) GetSourceAreaCode() string {
	return r._sourceAreaCode
}
