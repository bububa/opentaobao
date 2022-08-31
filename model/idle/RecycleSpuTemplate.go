package idle

// RecycleSpuTemplate 结构体
type RecycleSpuTemplate struct {
	// 可选， 回收V1版本：3C数码&#34;3C&#34;，奢侈品&#34;LUXURIES&#34;，大件&#34;BULKS&#34;，车&#34;CAR&#34;,母婴&#34;BABY&#34;，办公设备&#34;OFFICE&#34;，美妆&#34;MAKEUP&#34;，服装&#34;CLOTHING&#34;，低残值&#34;LOWVALUE&#34;，虚拟卡券&#34;VIRTUAL&#34; 回收V2版本：3C数码&#34;3C_NEW&#34;，奢侈品&#34;LUXURIES_NEW&#34;，大件&#34;BULKS_NEW&#34;，黄金&#34;GOLD_NEW&#34;
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// spuId
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 闲鱼端的spu变化，0新增，1删除，2修改名字
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 回收商类型 1专业回收商（省心卖） 2淘宝商家兼职回收
	RecycleType int64 `json:"recycle_type,omitempty" xml:"recycle_type,omitempty"`
	// 回收卖家的id,当 recyclerType 为2时才会有值
	RecycleSupplierId int64 `json:"recycle_supplier_id,omitempty" xml:"recycle_supplier_id,omitempty"`
}
