package alihealthpw

// ListByApplyIdsForBresp 结构体
type ListByApplyIdsForBresp struct {
	// 列表
	List []PendingListDto `json:"list,omitempty" xml:"list>pending_list_dto,omitempty"`
}
