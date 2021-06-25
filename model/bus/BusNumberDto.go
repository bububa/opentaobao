package bus

// BusNumberDto 
type BusNumberDto struct {

    // 始发城市ID
    AgentFromCityId   string `json:"agent_from_city_id,omitempty"`

    // 始发车站ID
    AgentFromStationId   string `json:"agent_from_station_id,omitempty"`

    // 到达时间（格式：yyyy-MM-dd HH:mm）
    ArriveTime   string `json:"arrive_time,omitempty"`

    // 其他非格式化属性合并成一个json格式字符串，对接时明确约定
    Attribute   string `json:"attribute,omitempty"`

    // 业务类型（0:普通汽车票；1:机场巴士；2:景区巴士，默认为0）
    BizType   int64 `json:"biz_type,omitempty"`

    // 发车前多少分钟停止售票(单位：分钟)
    BookLimitTime   int64 `json:"book_limit_time,omitempty"`

    // 客车车次
    BusNumber   string `json:"bus_number,omitempty"`

    // 客车车型
    BusType   string `json:"bus_type,omitempty"`

    // 出发时间（格式：yyyy-MM-dd HH:mm）
    DepartTime   string `json:"depart_time,omitempty"`

    // 目的地名称
    DestinationName   string `json:"destination_name,omitempty"`

    // 里程(单位：km)
    Distance   int64 `json:"distance,omitempty"`

    // 是否加班车（1是，0否）
    ExtraSchedule   int64 `json:"extra_schedule,omitempty"`

    // 始发城市名
    FromCityName   string `json:"from_city_name,omitempty"`

    // 始发车站名
    FromStationName   string `json:"from_station_name,omitempty"`

    // 全票价,单位分
    FullPrice   int64 `json:"full_price,omitempty"`

    // 半票价，单位分
    HalfPrice   int64 `json:"half_price,omitempty"`

    // 是否可退票（1是，0否）
    IsRefund   int64 `json:"is_refund,omitempty"`

    // 最晚发车时间（格式：HH:mm）
    LastSchedule   string `json:"last_schedule,omitempty"`

    // 余票数
    RemainSeats   int64 `json:"remain_seats,omitempty"`

    // 运行时长(单位：分钟)
    RunTime   int64 `json:"run_time,omitempty"`

    // 车次ID，需保持唯一
    ScheduleId   string `json:"schedule_id,omitempty"`

    // 服务费，单位分
    ServicePrice   int64 `json:"service_price,omitempty"`

    // 班次类型（0固定班，1流水班）
    ShiftType   int64 `json:"shift_type,omitempty"`

    // 车次状态（1可售；0停售）
    Status   int64 `json:"status,omitempty"`

    // 检票口
    TicketWicket   string `json:"ticket_wicket,omitempty"`

    // 目的地城市
    ToCityName   string `json:"to_city_name,omitempty"`

    // 目的地省份
    ToProvinceName   string `json:"to_province_name,omitempty"`

    // 目的地所属车站
    ToStationName   string `json:"to_station_name,omitempty"`

    // 总座位数
    TotalSeats   int64 `json:"total_seats,omitempty"`

    // 途经站
    ViaStation   string `json:"via_station,omitempty"`

}
