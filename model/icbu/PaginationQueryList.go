package icbu

// PaginationQueryList 结构体
type PaginationQueryList struct {
	// list
	List []PhotobankImageDo `json:"list,omitempty" xml:"list>photobank_image_do,omitempty"`
}
