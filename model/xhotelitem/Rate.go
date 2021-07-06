package xhotelitem

// Rate 结构体
type Rate struct {
	// 名称
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 创建时间
	CreatedTime string `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 修改时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 价格和库存信息。<br/>A:use_room_inventory:是否使用room级别共享库存，可选值 true false 1、true时：使用room级别共享库存（即使用gid对应的XRoom中的inventory），rate_quota_map 的json 数据中不需要录入库存信息,录入的库存信息会忽略 2、false时：使用rate级别私有库存，此时要求价格和库存必填。<br/>B:date  日期必须为 T---T+90 日内的日期（T为当天），且不能重复<br/>C:price 价格 int类型 取值范围1-99999999 单位为分<br/>D:quota 库存 int 类型 取值范围  0-999（数量库存）  60000(状态库存关) 61000(状态库存开)
	InventoryPrice string `json:"inventory_price,omitempty" xml:"inventory_price,omitempty"`
	// invPriceWithSwitch
	InvPriceWithSwitch string `json:"inv_price_with_switch,omitempty" xml:"inv_price_with_switch,omitempty"`
	// rate 维度下特殊标签含义 json: {"ebk-tail-room-Rate":1}, key:ebk-tail-room-Rate 表示rate维度ebk尾房标
	TagJson string `json:"tag_json,omitempty" xml:"tag_json,omitempty"`
	// 酒店商品id
	Gid int64 `json:"gid,omitempty" xml:"gid,omitempty"`
	// 酒店RPID
	Rpid int64 `json:"rpid,omitempty" xml:"rpid,omitempty"`
	// 额外服务-是否可以加床，1：不可以，2：可以
	AddBed int64 `json:"add_bed,omitempty" xml:"add_bed,omitempty"`
	// 额外服务-加床价格
	AddBedPrice int64 `json:"add_bed_price,omitempty" xml:"add_bed_price,omitempty"`
	// 币种（仅支持CNY）
	CurrencyCode int64 `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
	// 实价有房标签（RP支付类型为全额支付）
	ShijiaTag int64 `json:"shijia_tag,omitempty" xml:"shijia_tag,omitempty"`
	// 即时确认状态，表示此rate预订后是否可以直接发货。可取范围：0,1。可以为空
	JishiquerenTag int64 `json:"jishiqueren_tag,omitempty" xml:"jishiqueren_tag,omitempty"`
	// 是否使用RoomInventory库存   仅当Rate上使用时有意义
	UseRoomInventory bool `json:"use_room_inventory,omitempty" xml:"use_room_inventory,omitempty"`
}
