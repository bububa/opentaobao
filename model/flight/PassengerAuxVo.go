package flight

// PassengerAuxVo 
type PassengerAuxVo struct {
    // 辅营产品规格信息
    AuxProductVo   *AuxProductVo `json:"aux_product_vo,omitempty" xml:"aux_product_vo,omitempty"`
    // 购买数量
    BookNum   int64 `json:"book_num,omitempty" xml:"book_num,omitempty"`
    // 航段信息
    FlightVo   *BookFlightVo `json:"flight_vo,omitempty" xml:"flight_vo,omitempty"`
    // 乘机人姓名
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
}
