package tbk

// NTbkItem 结构体
type NTbkItem struct {
	// 商品小图列表
	Smallimages []string `json:"small_images,omitempty" xml:"small_images>string,omitempty"`
	// 一级类目名称
	Catname string `json:"cat_name,omitempty" xml:"cat_name,omitempty"`
	// 商品ID
	Numiid string `json:"num_iid,omitempty" xml:"num_iid,omitempty"`
	// 商品标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 商品主图
	Picturl string `json:"pict_url,omitempty" xml:"pict_url,omitempty"`
	// 商品一口价格
	Reserveprice string `json:"reserve_price,omitempty" xml:"reserve_price,omitempty"`
	// 折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
	Zkfinalprice string `json:"zk_final_price,omitempty" xml:"zk_final_price,omitempty"`
	// 商品所在地
	Provcity string `json:"provcity,omitempty" xml:"provcity,omitempty"`
	// 商品链接
	Itemurl string `json:"item_url,omitempty" xml:"item_url,omitempty"`
	// 店铺名称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 叶子类目名称
	Catleafname string `json:"cat_leaf_name,omitempty" xml:"cat_leaf_name,omitempty"`
	// 商品库类型，支持多库类型输出，以英文逗号分隔“,”分隔，1:营销商品主推库，如果值为空则不属于1这种商品类型
	Materiallibtype string `json:"material_lib_type,omitempty" xml:"material_lib_type,omitempty"`
	// 预售商品-商品优惠信息
	Presalediscountfeetext string `json:"presale_discount_fee_text,omitempty" xml:"presale_discount_fee_text,omitempty"`
	// 预售商品-定金（元）
	Presaledeposit string `json:"presale_deposit,omitempty" xml:"presale_deposit,omitempty"`
	// 1聚划算满减：满N件减X元，满N件X折，满N件X元）  2天猫限时抢：前N分钟每件X元，前N分钟满N件每件X元，前N件每件X元）
	Playinfo string `json:"play_info,omitempty" xml:"play_info,omitempty"`
	// 聚划算信息-聚淘开始时间（毫秒）
	Juonlinestarttime string `json:"ju_online_start_time,omitempty" xml:"ju_online_start_time,omitempty"`
	// 聚划算信息-聚淘结束时间（毫秒）
	Juonlineendtime string `json:"ju_online_end_time,omitempty" xml:"ju_online_end_time,omitempty"`
	// 聚划算信息-商品预热开始时间（毫秒）
	Jupreshowstarttime string `json:"ju_pre_show_start_time,omitempty" xml:"ju_pre_show_start_time,omitempty"`
	// 聚划算信息-商品预热结束时间（毫秒）
	Jupreshowendtime string `json:"ju_pre_show_end_time,omitempty" xml:"ju_pre_show_end_time,omitempty"`
	// 活动价
	Saleprice string `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 跨店满减信息
	Kuadianpromotioninfo string `json:"kuadian_promotion_info,omitempty" xml:"kuadian_promotion_info,omitempty"`
	// 是否品牌精选，0不是，1是
	Superiorbrand string `json:"superior_brand,omitempty" xml:"superior_brand,omitempty"`
	// 是否是热门商品，0不是，1是
	Hotflag string `json:"hot_flag,omitempty" xml:"hot_flag,omitempty"`
	// 入参的（新）商品ID
	Inputnumiid string `json:"input_num_iid,omitempty" xml:"input_num_iid,omitempty"`
	// 卖家类型，0表示集市，1表示商城，3表示特价版
	Usertype int64 `json:"user_type,omitempty" xml:"user_type,omitempty"`
	// 卖家id
	Sellerid int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 30天销量
	Volume int64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// 店铺dsr 评分
	Shopdsr int64 `json:"shop_dsr,omitempty" xml:"shop_dsr,omitempty"`
	// 卖家等级
	Ratesum int64 `json:"ratesum,omitempty" xml:"ratesum,omitempty"`
	// 预售商品-付定金结束时间（毫秒）
	Presaletailendtime int64 `json:"presale_tail_end_time,omitempty" xml:"presale_tail_end_time,omitempty"`
	// 预售商品-付尾款开始时间（毫秒）
	Presaletailstarttime int64 `json:"presale_tail_start_time,omitempty" xml:"presale_tail_start_time,omitempty"`
	// 预售商品-付定金结束时间（毫秒）
	Presaleendtime int64 `json:"presale_end_time,omitempty" xml:"presale_end_time,omitempty"`
	// 预售商品-付定金开始时间（毫秒）
	Presalestarttime int64 `json:"presale_start_time,omitempty" xml:"presale_start_time,omitempty"`
	// 聚划算满减  -结束时间（毫秒）
	Juplayendtime int64 `json:"ju_play_end_time,omitempty" xml:"ju_play_end_time,omitempty"`
	// 聚划算满减  -开始时间（毫秒）
	Juplaystarttime int64 `json:"ju_play_start_time,omitempty" xml:"ju_play_start_time,omitempty"`
	// 天猫限时抢可售  -结束时间（毫秒）
	Tmallplayactivityendtime int64 `json:"tmall_play_activity_end_time,omitempty" xml:"tmall_play_activity_end_time,omitempty"`
	// 天猫限时抢可售  -开始时间（毫秒）
	Tmallplayactivitystarttime int64 `json:"tmall_play_activity_start_time,omitempty" xml:"tmall_play_activity_start_time,omitempty"`
	// 是否加入消费者保障
	Isprepay bool `json:"is_prepay,omitempty" xml:"is_prepay,omitempty"`
	// 退款率是否低于行业均值
	Irfdrate bool `json:"i_rfd_rate,omitempty" xml:"i_rfd_rate,omitempty"`
	// 好评率是否高于行业均值
	Hgoodrate bool `json:"h_good_rate,omitempty" xml:"h_good_rate,omitempty"`
	// 成交转化是否高于行业均值
	Hpayrate30 bool `json:"h_pay_rate30,omitempty" xml:"h_pay_rate30,omitempty"`
	// 是否包邮
	Freeshipment bool `json:"free_shipment,omitempty" xml:"free_shipment,omitempty"`
}
