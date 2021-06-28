package logistic

// AddressReachable 
type AddressReachable struct {

    // 标准区域编码(三级行政区编码或是四级行政区)，可以通过taobao.areas.get获取，如北京市朝阳区为110105
    
    AreaCode   string `json:"area_code,omitempty" xml:"area_code,omitempty"`
    

    // 详细地址
    
    Address   string `json:"address,omitempty" xml:"address,omitempty"`
    

    // 发货地，标准区域编码(四级行政)，可以通过taobao.areas.get获取，如浙江省杭州市余杭区闲林街道为 330110011
    
    SourceAreaCode   string `json:"source_area_code,omitempty" xml:"source_area_code,omitempty"`
    

    // 物流公司编码ID，可以从这个接口获取所有物流公司的标准编码taobao.logistics.companies.get，可以传入多个值，以英文逗号分隔，如“1000000952,1000000953”
    
    PartnerId   string `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
    

    // 服务对应的数字编码，如揽派范围对应88
    
    ServiceType   int64 `json:"service_type,omitempty" xml:"service_type,omitempty"`
    

}
