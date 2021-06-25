package flight

// ModifyBeforeSegmentDTO 
type ModifyBeforeSegmentDTO struct {

    // 到达城市
    ArrCity   string `json:"arr_city,omitempty"`

    // 起飞城市
    DepCity   string `json:"dep_city,omitempty"`

}
