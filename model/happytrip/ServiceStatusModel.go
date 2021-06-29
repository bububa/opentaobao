package happytrip

// ServiceStatusModel 
type ServiceStatusModel struct {

    // 供应商服务在各地区的支持状态
    
    Cities   []CityServiceStatus `json:"cities,omitempty" xml:"cities,omitempty"`
    

}
