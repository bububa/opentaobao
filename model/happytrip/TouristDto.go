package happytrip

// TouristDto 
type TouristDto struct {
    // 证件签发国
    CertNation   string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
    // 姓
    FirstName   string `json:"first_name,omitempty" xml:"first_name,omitempty"`
    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
    // 主键
    Id   int64 `json:"id,omitempty" xml:"id,omitempty"`
    // 名
    LastName   string `json:"last_name,omitempty" xml:"last_name,omitempty"`
    // 出行人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 订单id
    OrderId   int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
    // 1男，2女
    Sex   int64 `json:"sex,omitempty" xml:"sex,omitempty"`
    // 成人0，儿童1，婴儿2
    TouristType   int64 `json:"tourist_type,omitempty" xml:"tourist_type,omitempty"`
    // 出行人0，同行人1，外部人员2
    TravelBusinessType   int64 `json:"travel_business_type,omitempty" xml:"travel_business_type,omitempty"`
    // 用户id
    UserId   string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
