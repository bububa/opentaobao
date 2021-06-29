package happytrip

// OrderInfo 
type OrderInfo struct {
    // 订单id
    Id   string `json:"id,omitempty" xml:"id,omitempty"`
    // 城市id
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 订单类型
    Type   int64 `json:"type,omitempty" xml:"type,omitempty"`
    // 叫车人手机号
    CallPhone   string `json:"call_phone,omitempty" xml:"call_phone,omitempty"`
    // 乘车人手机号
    PassengerPhone   string `json:"passenger_phone,omitempty" xml:"passenger_phone,omitempty"`
    // 订单状态
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 出发地纬度
    Flat   string `json:"flat,omitempty" xml:"flat,omitempty"`
    // 出发地经度
    Flng   string `json:"flng,omitempty" xml:"flng,omitempty"`
    // 目的地纬度
    Tlat   string `json:"tlat,omitempty" xml:"tlat,omitempty"`
    // 目的地经度
    Tlng   string `json:"tlng,omitempty" xml:"tlng,omitempty"`
    // 当前纬度
    Clat   string `json:"clat,omitempty" xml:"clat,omitempty"`
    // 当前经度
    Clng   string `json:"clng,omitempty" xml:"clng,omitempty"`
    // 出发地名称
    StartName   string `json:"start_name,omitempty" xml:"start_name,omitempty"`
    // 出发地地址
    StartAddress   string `json:"start_address,omitempty" xml:"start_address,omitempty"`
    // 目的地名称
    EndName   string `json:"end_name,omitempty" xml:"end_name,omitempty"`
    // 目的地地址
    EndAddress   string `json:"end_address,omitempty" xml:"end_address,omitempty"`
    // 备注
    ExtraInfo   string `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
    // 出发时间
    DepartureTime   string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
    // 下单时间
    OrderTime   string `json:"order_time,omitempty" xml:"order_time,omitempty"`
    // 所需车型代码
    RequireLevel   string `json:"require_level,omitempty" xml:"require_level,omitempty"`
    // 备注
    Remark   string `json:"remark,omitempty" xml:"remark,omitempty"`
    // 成本中心代码，用于区分不同的分账账号
    CostCenter   string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
    // 订单详细状态码
    SubStatus   int64 `json:"sub_status,omitempty" xml:"sub_status,omitempty"`
    // 司机称呼
    DriverName   string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
    // 号码保护中间号（如无号码保护，为司机真实手机号）
    DriverPhone   string `json:"driver_phone,omitempty" xml:"driver_phone,omitempty"`
    // 司机真实手机号
    DriverPhoneReal   string `json:"driver_phone_real,omitempty" xml:"driver_phone_real,omitempty"`
    // 已通知司机数量
    DriverNum   int64 `json:"driver_num,omitempty" xml:"driver_num,omitempty"`
    // 司机车型
    DriverCarType   string `json:"driver_car_type,omitempty" xml:"driver_car_type,omitempty"`
    // 司机车牌
    DriverCard   string `json:"driver_card,omitempty" xml:"driver_card,omitempty"`
    // 司机头像
    DriverAvatar   string `json:"driver_avatar,omitempty" xml:"driver_avatar,omitempty"`
    // 司机星级
    DriverLevel   string `json:"driver_level,omitempty" xml:"driver_level,omitempty"`
    // 司机抢单数
    DriverOrderCount   int64 `json:"driver_order_count,omitempty" xml:"driver_order_count,omitempty"`
    // 司机当前实时经度
    Dlng   string `json:"dlng,omitempty" xml:"dlng,omitempty"`
    // 司机当前实时纬度
    Dlat   string `json:"dlat,omitempty" xml:"dlat,omitempty"`
    // 司机接单时间
    StriveTime   string `json:"strive_time,omitempty" xml:"strive_time,omitempty"`
    // 司机到达时间
    ReachTime   string `json:"reach_time,omitempty" xml:"reach_time,omitempty"`
    // 开始计价时间
    BeginChargeTime   string `json:"begin_charge_time,omitempty" xml:"begin_charge_time,omitempty"`
    // 行程结束时间
    FinishTime   string `json:"finish_time,omitempty" xml:"finish_time,omitempty"`
    // 迟到计费时间
    DelayTimeStart   string `json:"delay_time_start,omitempty" xml:"delay_time_start,omitempty"`
    // 实际行驶公里数
    NormalDistance   string `json:"normal_distance,omitempty" xml:"normal_distance,omitempty"`
    // 实际行驶时长（分钟）
    NormalTime   int64 `json:"normal_time,omitempty" xml:"normal_time,omitempty"`
    // 实际车型代码
    StriveLevel   string `json:"strive_level,omitempty" xml:"strive_level,omitempty"`
    // 汽车的颜色
    DriverCarColor   string `json:"driver_car_color,omitempty" xml:"driver_car_color,omitempty"`
}
