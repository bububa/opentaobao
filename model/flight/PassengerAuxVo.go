package flight

// PassengerAuxVo 
/* model for simplify = false
type PassengerAuxVo struct {

    // 辅营产品规格信息
    
    AuxProductVo  *struct {
        AuxProductVo  *AuxProductVo `json:"aux_product_vo,omitempty"`
    } `json:"aux_product_vo,omitempty"`
    

    // 购买数量
    
    BookNum   int64 `json:"book_num,omitempty"`
    

    // 航段信息
    
    FlightVo  *struct {
        BookFlightVo  *BookFlightVo `json:"book_flight_vo,omitempty"`
    } `json:"flight_vo,omitempty"`
    

    // 乘机人姓名
    
    Name   string `json:"name,omitempty"`
    

}
*/

// PassengerAuxVo 
type PassengerAuxVo struct {

    // 辅营产品规格信息
    AuxProductVo   *AuxProductVo `json:"aux_product_vo,omitempty"`

    // 购买数量
    BookNum   int64 `json:"book_num,omitempty"`

    // 航段信息
    FlightVo   *BookFlightVo `json:"flight_vo,omitempty"`

    // 乘机人姓名
    Name   string `json:"name,omitempty"`

}
