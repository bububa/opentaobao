package cainiaohandover

// PickupDTO 
type PickupDTO struct {
    // 揽收地址
    Address   *AddressDTO `json:"address,omitempty" xml:"address,omitempty"`
    // 邮箱
    Email   string `json:"email,omitempty" xml:"email,omitempty"`
    // 移动电话, 校验格式：^1(3|4|5|6|7|8|9)\d{9}$
    Mobile   string `json:"mobile,omitempty" xml:"mobile,omitempty"`
    // 固定电话，可空，校验格式：(^0[\d]{2,3}-[\d]{7,8}$)|(^400[\d]{3,4}[\d]{3,4}$)|(400-[\d]{3,4}-[\d]{3,4}$)
    Phone   string `json:"phone,omitempty" xml:"phone,omitempty"`
    // 揽收联系人名称，必须包含中文字符
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // AE后台维护的发件地址ID
    AddressId   int64 `json:"address_id,omitempty" xml:"address_id,omitempty"`
}
