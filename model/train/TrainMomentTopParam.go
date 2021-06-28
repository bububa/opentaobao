package train

// TrainMomentTopParam 
/* model for simplify = false
type TrainMomentTopParam struct {

    // 到达站
    
    ArrStation   string `json:"arr_station,omitempty"`
    

    // 出发日期
    
    DepDate   string `json:"dep_date,omitempty"`
    

    // 出发站
    
    DepStation   string `json:"dep_station,omitempty"`
    

    // 车次号
    
    TrainNo   string `json:"train_no,omitempty"`
    

    // 长车次号
    
    TrainNoLong   string `json:"train_no_long,omitempty"`
    

    // 渠道标识
    
    Ttid   string `json:"ttid,omitempty"`
    

}
*/

// TrainMomentTopParam 
type TrainMomentTopParam struct {

    // 到达站
    ArrStation   string `json:"arr_station,omitempty"`

    // 出发日期
    DepDate   string `json:"dep_date,omitempty"`

    // 出发站
    DepStation   string `json:"dep_station,omitempty"`

    // 车次号
    TrainNo   string `json:"train_no,omitempty"`

    // 长车次号
    TrainNoLong   string `json:"train_no_long,omitempty"`

    // 渠道标识
    Ttid   string `json:"ttid,omitempty"`

}
