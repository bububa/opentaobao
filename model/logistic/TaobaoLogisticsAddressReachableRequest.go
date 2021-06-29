package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判定服务是否可达 API请求
taobao.logistics.address.reachable

根据输入的目标地址，判断服务是否可达。
现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
*/
type TaobaoLogisticsAddressReachableRequest struct {
    model.Params
    // 标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105
    _areaCode   string
    // 详细地址
    _address   string
    // 物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”
    _partnerIds   string
    // 服务对应的数字编码，如揽派范围对应88
    _serviceType   int64
    // 发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011
    _sourceAreaCode   string
}

// 初始化TaobaoLogisticsAddressReachableRequest对象
func NewTaobaoLogisticsAddressReachableRequest() *TaobaoLogisticsAddressReachableRequest{
    return &TaobaoLogisticsAddressReachableRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsAddressReachableRequest) GetApiMethodName() string {
    return "taobao.logistics.address.reachable"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsAddressReachableRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AreaCode Setter
// 标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105
func (r *TaobaoLogisticsAddressReachableRequest) SetAreaCode(_areaCode string) error {
    r._areaCode = _areaCode
    r.Set("area_code", _areaCode)
    return nil
}

// AreaCode Getter
func (r TaobaoLogisticsAddressReachableRequest) GetAreaCode() string {
    return r._areaCode
}
// Address Setter
// 详细地址
func (r *TaobaoLogisticsAddressReachableRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoLogisticsAddressReachableRequest) GetAddress() string {
    return r._address
}
// PartnerIds Setter
// 物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”
func (r *TaobaoLogisticsAddressReachableRequest) SetPartnerIds(_partnerIds string) error {
    r._partnerIds = _partnerIds
    r.Set("partner_ids", _partnerIds)
    return nil
}

// PartnerIds Getter
func (r TaobaoLogisticsAddressReachableRequest) GetPartnerIds() string {
    return r._partnerIds
}
// ServiceType Setter
// 服务对应的数字编码，如揽派范围对应88
func (r *TaobaoLogisticsAddressReachableRequest) SetServiceType(_serviceType int64) error {
    r._serviceType = _serviceType
    r.Set("service_type", _serviceType)
    return nil
}

// ServiceType Getter
func (r TaobaoLogisticsAddressReachableRequest) GetServiceType() int64 {
    return r._serviceType
}
// SourceAreaCode Setter
// 发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011
func (r *TaobaoLogisticsAddressReachableRequest) SetSourceAreaCode(_sourceAreaCode string) error {
    r._sourceAreaCode = _sourceAreaCode
    r.Set("source_area_code", _sourceAreaCode)
    return nil
}

// SourceAreaCode Getter
func (r TaobaoLogisticsAddressReachableRequest) GetSourceAreaCode() string {
    return r._sourceAreaCode
}
