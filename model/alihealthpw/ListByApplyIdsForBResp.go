package alihealthpw

// ListByApplyIdsForBResp 结构体
type ListByApplyIdsForBResp struct {
	// 列表
	List []PendingListDto `json:"list,omitempty" xml:"list>pending_list_dto,omitempty"`
}
