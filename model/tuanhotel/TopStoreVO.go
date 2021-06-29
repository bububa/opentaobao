package tuanhotel

// TopStoreVO 
type TopStoreVO struct {
    // 物理酒店ID
    Hid   int64 `json:"hid,omitempty" xml:"hid,omitempty"`
    // 城市
    City   int64 `json:"city,omitempty" xml:"city,omitempty"`
    // 标准酒店ID
    Shid   int64 `json:"shid,omitempty" xml:"shid,omitempty"`
    // 门店ID
    StoreId   int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 门店分账下的分账信息，数据格式（淘宝账户名称:比例或者金额:分账方式（1-按照比例的方式进行分账，2-按照固定金额的方式进行分账））
    BillInfos   string `json:"bill_infos,omitempty" xml:"bill_infos,omitempty"`
    // 分账描述
    BillDescs   string `json:"bill_descs,omitempty" xml:"bill_descs,omitempty"`
    // 核销账号
    WriteOffAccounts   string `json:"write_off_accounts,omitempty" xml:"write_off_accounts,omitempty"`
    // 房型列表
    RoomTypes   []TopRoomTypeVO `json:"room_types,omitempty" xml:"room_types>top_room_type_vo,omitempty"`
    // 酒店名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 联系电话列表
    AppointPhones   []TopAppointPhoneVO `json:"appoint_phones,omitempty" xml:"appoint_phones>top_appoint_phone_vo,omitempty"`
}
