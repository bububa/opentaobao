package ship

// ShipAgentConfirmBookTicketInfo 
type ShipAgentConfirmBookTicketInfo struct {
    // 电子票号
    ETicketNo   string `json:"e_ticket_no,omitempty" xml:"e_ticket_no,omitempty"`
    // 票描述
    TicketDesc   string `json:"ticket_desc,omitempty" xml:"ticket_desc,omitempty"`
    // 票Id
    TicketId   string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
    // 票号
    TicketNo   string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
    // 票价格(分)
    TicketPrice   int64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
    // 取票密码
    TicketPwd   string `json:"ticket_pwd,omitempty" xml:"ticket_pwd,omitempty"`
    // 座位号
    TicketSeatNo   string `json:"ticket_seat_no,omitempty" xml:"ticket_seat_no,omitempty"`
    // 票状态：1-出票中2-已出票3-无票4-退票中5-已退票6-退票失败
    TicketStatus   string `json:"ticket_status,omitempty" xml:"ticket_status,omitempty"`
    // 子票种类型11-去程;12-往返;21普通门票；31酒店
    TicketSubType   string `json:"ticket_sub_type,omitempty" xml:"ticket_sub_type,omitempty"`
    // 票标题
    TicketTitle   string `json:"ticket_title,omitempty" xml:"ticket_title,omitempty"`
    // 票种类别1-船票;2-门票;3-酒店;4-其他
    TicketType   string `json:"ticket_type,omitempty" xml:"ticket_type,omitempty"`
    // 扩展属性数据
    ExtAttr   string `json:"ext_attr,omitempty" xml:"ext_attr,omitempty"`
    // 过期时间
    ExpireTime   string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
}
