package train

// TrainAgentStopInfo 结构体
type TrainAgentStopInfo struct {
	// 车次号
	TrainNo string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	// 发车时间
	DepTime string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// 出发车站
	FromStation string `json:"from_station,omitempty" xml:"from_station,omitempty"`
	// 到达车站
	ToStation string `json:"to_station,omitempty" xml:"to_station,omitempty"`
	// 车次状态：STOP 停运、NORMAL 正常开行、NO_TRAIN 无车次、UN_KNOWN未知状态
	TrainStatus string `json:"train_status,omitempty" xml:"train_status,omitempty"`
	// uuid唯一标识
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
}
