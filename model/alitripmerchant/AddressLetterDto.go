package alitripmerchant

// AddressLetterDto 
type AddressLetterDto struct {
    // 字母分类列表
    AddressList   []AddressSearchDto `json:"address_list,omitempty" xml:"address_list>address_search_dto,omitempty"`
    // 字母
    Letter   string `json:"letter,omitempty" xml:"letter,omitempty"`
}
