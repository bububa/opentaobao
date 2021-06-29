package alitripmerchant

// AddressLetterDTO 
type AddressLetterDTO struct {
    // 字母分类列表
    AddressList   []AddressSearchDTO `json:"address_list,omitempty" xml:"address_list>address_search_dto,omitempty"`
    // 字母
    Letter   string `json:"letter,omitempty" xml:"letter,omitempty"`
}
