package btrip

// BusLineInfoVo 
type BusLineInfoVo struct {

    // 到达站
    
    ArrStation   string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
    

    // 到达时间（yyyy-MM-dd HH:mm:ss）
    
    ArrivalTime   string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
    

    // 距离
    
    BusDistance   string `json:"bus_distance,omitempty" xml:"bus_distance,omitempty"`
    

    // 车次id
    
    BusNoId   int64 `json:"bus_no_id,omitempty" xml:"bus_no_id,omitempty"`
    

    // 客车车次
    
    BusNumber   string `json:"bus_number,omitempty" xml:"bus_number,omitempty"`
    

    // 车次全局唯一id
    
    BusNumberUuid   int64 `json:"bus_number_uuid,omitempty" xml:"bus_number_uuid,omitempty"`
    

    // 车型id
    
    BusType   string `json:"bus_type,omitempty" xml:"bus_type,omitempty"`
    

    // 出发城市
    
    DepCity   string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
    

    // 出发站
    
    DepStation   string `json:"dep_station,omitempty" xml:"dep_station,omitempty"`
    

    // 出发时间
    
    DepartTime   string `json:"depart_time,omitempty" xml:"depart_time,omitempty"`
    

    // 是否加班车
    
    ExtraSchedule   bool `json:"extra_schedule,omitempty" xml:"extra_schedule,omitempty"`
    

    // 全价票价（单位：分）
    
    FullPrice   int64 `json:"full_price,omitempty" xml:"full_price,omitempty"`
    

    // 半票票价（单位：分）
    
    HalfPrice   int64 `json:"half_price,omitempty" xml:"half_price,omitempty"`
    

    // 是否支持电子票，0：不支持；1：支持
    
    IsSupportETicket   int64 `json:"is_support_e_ticket,omitempty" xml:"is_support_e_ticket,omitempty"`
    

    // 最晚出发时间
    
    LastDepartTime   string `json:"last_depart_time,omitempty" xml:"last_depart_time,omitempty"`
    

    // 到达城市
    
    LastPlaceName   string `json:"last_place_name,omitempty" xml:"last_place_name,omitempty"`
    

    // 是否是预约购票订单
    
    PreOrder   bool `json:"pre_order,omitempty" xml:"pre_order,omitempty"`
    

    // 实名制校验规则,'|'分隔,例如2|3|4；1 不需要进行实名制校验；2 订单内乘客身份证不能重复；3 当天同一班次乘客身份证限购1张票；4 同一取票人身份证当天限购3张票；5 同一乘车人身份证1天在该网站只能买1张票；6 取票人的姓名、证件信息必须在乘车人中
    
    RealNameGrade   string `json:"real_name_grade,omitempty" xml:"real_name_grade,omitempty"`
    

    // 服务费
    
    ServicePrice   int64 `json:"service_price,omitempty" xml:"service_price,omitempty"`
    

    // 流水班次 0 固定班次 1：流水班次 null 固定班次
    
    ShiftType   int64 `json:"shift_type,omitempty" xml:"shift_type,omitempty"`
    

    // 耗时，单位：分钟，30分钟为单位向上取整
    
    SpendTime   int64 `json:"spend_time,omitempty" xml:"spend_time,omitempty"`
    

    // 标准城市名称
    
    StandardCityName   string `json:"standard_city_name,omitempty" xml:"standard_city_name,omitempty"`
    

    // 状态 1：可售 2：不可售，非车站营业时间 3：不可售，超出预售期 4：暂停售卖
    
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    

    // 余票
    
    Stock   int64 `json:"stock,omitempty" xml:"stock,omitempty"`
    

    // 路线信息（途径站："廊坊,天津"）
    
    ViaStation   string `json:"via_station,omitempty" xml:"via_station,omitempty"`
    

    // 途经站类型 0 未知 1 途经站 2 终点站
    
    ViaStationType   int64 `json:"via_station_type,omitempty" xml:"via_station_type,omitempty"`
    

}
