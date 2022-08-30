package tbk

// DataMap 结构体
type DataMap struct {
	// 当入参不传pid的时候返回，表示账号关联的pid
	Pid string `json:"pid,omitempty" xml:"pid,omitempty"`
	// 当入参传入pid的时候返回，表示pid关联的relationId
	Rid string `json:"rid,omitempty" xml:"rid,omitempty"`
	// 0表示预警，1表示拦截，如果名单中同一个淘客同时有拦截和预警信息，以拦截为准
	Status string `json:"status,omitempty" xml:"status,omitempty"`
}
