package ascpchannel

// Aicinventoryqueryrequest 结构体
type Aicinventoryqueryrequest struct {
	// 参数列表
	AicinventoryQueryList []Aicinventoryquerylist `json:"aicinventory_query_list,omitempty" xml:"aicinventory_query_list>aicinventoryquerylist,omitempty"`
}
