package idle

// RecycleSpuTemplate 结构体
type RecycleSpuTemplate struct {
	// 可选， 回收V1版本：3C数码&#34;3C&#34;，奢侈品&#34;LUXURIES&#34;，大件&#34;BULKS&#34;，车&#34;CAR&#34;,母婴&#34;BABY&#34;，办公设备&#34;OFFICE&#34;，美妆&#34;MAKEUP&#34;，服装&#34;CLOTHING&#34;，低残值&#34;LOWVALUE&#34;，虚拟卡券&#34;VIRTUAL&#34;
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 业务支持的产品类型：BM：帮卖，HS：回收，YJHX:天猫以旧换新，SYHS：淘宝私域回收
	PdCode string `json:"pd_code,omitempty" xml:"pd_code,omitempty"`
	// 服务商支持的SPU挂载：0和2：第一次请求预挂载新增，二次请求上线，上线后状态不会再变化；1：下线；
	ActionType int64 `json:"action_type,omitempty" xml:"action_type,omitempty"`
	// 淘宝商家ID，recycle_type = 2时才有效
	RecycleSupplierId int64 `json:"recycle_supplier_id,omitempty" xml:"recycle_supplier_id,omitempty"`
	// SPUID
	SpuId int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	// 回收商类型：1：专业服务商  2：淘宝商家
	RecycleType int64 `json:"recycle_type,omitempty" xml:"recycle_type,omitempty"`
}
