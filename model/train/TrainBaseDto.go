package train

// TrainBaseDto 
type TrainBaseDto struct {

    // 车次号
    
    TrainNumber   string `json:"train_number,omitempty" xml:"train_number,omitempty"`
    

    // 车次类型
    
    TrainType   int64 `json:"train_type,omitempty" xml:"train_type,omitempty"`
    

    // 运行时长
    
    RunTime   string `json:"run_time,omitempty" xml:"run_time,omitempty"`
    

    // 到达时间
    
    ToStationTime   string `json:"to_station_time,omitempty" xml:"to_station_time,omitempty"`
    

    // 到达站
    
    ToStationName   string `json:"to_station_name,omitempty" xml:"to_station_name,omitempty"`
    

    // 出发时间
    
    FromStationTime   string `json:"from_station_time,omitempty" xml:"from_station_time,omitempty"`
    

    // 出发站
    
    FromStationName   string `json:"from_station_name,omitempty" xml:"from_station_name,omitempty"`
    

    // 坐席类型
    
    SeatType   string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
    

}
