package ott

// ChannelDto 结构体
type ChannelDto struct {
	// itemList
	ItemList []ItemDto `json:"item_list,omitempty" xml:"item_list>item_dto,omitempty"`
	// Description of Content Provider
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// Homepage URL of Content Provider
	Link string `json:"link,omitempty" xml:"link,omitempty"`
	// Last time the entry was modified
	PubDate string `json:"pub_date,omitempty" xml:"pub_date,omitempty"`
	// skip days for the channel
	SkipDays string `json:"skip_days,omitempty" xml:"skip_days,omitempty"`
	// skip hours for the channel
	SkipHours string `json:"skip_hours,omitempty" xml:"skip_hours,omitempty"`
	// Name of Content Provider
	Title string `json:"title,omitempty" xml:"title,omitempty"`
}
