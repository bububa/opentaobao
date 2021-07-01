package idle

// RecycleSpuTemplate 结构体
type RecycleSpuTemplate struct {
	// spuId
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 闲鱼端的spu变化，0新增，1删除，2修改名字
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 可选， 回收V1版本：3C数码"3C"，奢侈品"LUXURIES"，大件"BULKS"，车"CAR",母婴"BABY"，办公设备"OFFICE"，美妆"MAKEUP"，服装"CLOTHING"，低残值"LOWVALUE"，虚拟卡券"VIRTUAL" 回收V2版本：3C数码"3C_NEW"，奢侈品"LUXURIES_NEW"，大件"BULKS_NEW"，黄金"GOLD_NEW"
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
}
