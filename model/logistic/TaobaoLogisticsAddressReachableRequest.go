package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判定服务是否可达 APIRequest
taobao.logistics.address.reachable

根据输入的目标地址，判断服务是否可达。
现已支持筛单的快递公司共15家：中国邮政、EMS、国通、汇通、快捷、全峰、优速、圆通、宅急送、中通、顺丰、天天、韵达、德邦快递、申通
*/
type TaobaoLogisticsAddressReachableRequest struct {
    model.Params

    // 标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105
    areaCode   string 

    // 详细地址
    address   string 

    // 物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”
    partnerIds   string 

    // 服务对应的数字编码，如揽派范围对应88
    serviceType   int64 

    // 发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011
    sourceAreaCode   string 

}

func NewTaobaoLogisticsAddressReachableRequest() *TaobaoLogisticsAddressReachableRequest{
    return &TaobaoLogisticsAddressReachableRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoLogisticsAddressReachableRequest) GetApiMethodName() string {
    return "taobao.logistics.address.reachable"
}

func (r TaobaoLogisticsAddressReachableRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoLogisticsAddressReachableRequest) SetAreaCode(areaCode string) error {
    r.areaCode = areaCode
    r.Set("area_code", areaCode)
    return nil
}

func (r TaobaoLogisticsAddressReachableRequest) GetAreaCode() string {
    return r.areaCode
}

func (r *TaobaoLogisticsAddressReachableRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r TaobaoLogisticsAddressReachableRequest) GetAddress() string {
    return r.address
}

func (r *TaobaoLogisticsAddressReachableRequest) SetPartnerIds(partnerIds string) error {
    r.partnerIds = partnerIds
    r.Set("partner_ids", partnerIds)
    return nil
}

func (r TaobaoLogisticsAddressReachableRequest) GetPartnerIds() string {
    return r.partnerIds
}

func (r *TaobaoLogisticsAddressReachableRequest) SetServiceType(serviceType int64) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

func (r TaobaoLogisticsAddressReachableRequest) GetServiceType() int64 {
    return r.serviceType
}

func (r *TaobaoLogisticsAddressReachableRequest) SetSourceAreaCode(sourceAreaCode string) error {
    r.sourceAreaCode = sourceAreaCode
    r.Set("source_area_code", sourceAreaCode)
    return nil
}

func (r TaobaoLogisticsAddressReachableRequest) GetSourceAreaCode() string {
    return r.sourceAreaCode
}

