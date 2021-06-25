package train

// ToAgentTicketInfo 
type ToAgentTicketInfo struct {

    // 学生信息
    StudentInfo   *StudentInfo `json:"student_info,omitempty"`

    // 淘宝火车票子订单id.
    SubOrderId   string `json:"sub_order_id,omitempty"`

    // 出发站
    FromStation   string `json:"from_station,omitempty"`

    // 出发时间
    FromTime   string `json:"from_time,omitempty"`

    // 到达站
    ToStation   string `json:"to_station,omitempty"`

    // 坐席
    Seat   int64 `json:"seat,omitempty"`

    // 车次
    TrainNum   string `json:"train_num,omitempty"`

    // 乘客姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 证件编号
    CertificateNum   string `json:"certificate_num,omitempty"`

    // 证件类型，0:身份证 1:护照 4:港澳通行证 5:台湾通行证
    CertificateType   string `json:"certificate_type,omitempty"`

    // 保险价格，精确到分，例如10元，输入1000。
    InsurancePrice   int64 `json:"insurance_price,omitempty"`

    // 单张票价(不包含保险价格),例如100元,输出为10000,精确到分.
    TicketPrice   int64 `json:"ticket_price,omitempty"`

    // 乘客生日
    Birthday   string `json:"birthday,omitempty"`

    // 到站时间
    ToTime   string `json:"to_time,omitempty"`

    // 1:单程票
    Tag   int64 `json:"tag,omitempty"`

    // 保险的单一价格
    InsuranceUnitPrice   int64 `json:"insurance_unit_price,omitempty"`

    // 0:成人 1:儿童 2：学生
    PassengerType   int64 `json:"passenger_type,omitempty"`

    // segmentIndex
    SegmentIndex   int64 `json:"segment_index,omitempty"`

    // 国籍
    Nationality   string `json:"nationality,omitempty"`

    // 国籍代码
    NationalityCode   string `json:"nationality_code,omitempty"`

    // 证件有效期
    ValidUntil   string `json:"valid_until,omitempty"`

    // 性别  1:男  0:女
    Gender   string `json:"gender,omitempty"`

    // 联系人电话
    Telephone   string `json:"telephone,omitempty"`

    // 
    RealTicketPrice   int64 `json:"real_ticket_price,omitempty"`

}
