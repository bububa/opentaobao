package flight

// IssueList 
type IssueList struct {

    // 证件信息
    CertNo   string `json:"cert_no,omitempty"`

    // 乘机人姓名
    PassengerName   string `json:"passenger_name,omitempty"`

    // 证件类型
    CertType   int64 `json:"cert_type,omitempty"`

    // 票号
    Tickets   []String `json:"tickets,omitempty"`

    // 乘客类型
    PassengerType   int64 `json:"passenger_type,omitempty"`

    // 票面价
    TicketPrice   int64 `json:"ticket_price,omitempty"`

    // pnr
    Pnr   string `json:"pnr,omitempty"`

    // 税项对象
    Taxes   []Taxes `json:"taxes,omitempty"`

    // 优惠价格
    Promotion   int64 `json:"promotion,omitempty"`

    // 航段
    Segments   []Segments `json:"segments,omitempty"`

    // 联系电话
    Mobile   string `json:"mobile,omitempty"`

}
