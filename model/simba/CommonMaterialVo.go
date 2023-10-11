package simba

// CommonMaterialVo 结构体
type CommonMaterialVo struct {
	// 生命周期指标-越迁阶段列表,列表中最多有2个元素。在页面上展示的时候，按照元素的索引，从左到右依次展示即
	LifeCycleList []ItemLifeCycleViewVo `json:"life_cycle_list,omitempty" xml:"life_cycle_list>item_life_cycle_view_vo,omitempty"`
	// 物料名称
	MaterialName string `json:"material_name,omitempty" xml:"material_name,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 图片地址
	ImgUrl string `json:"img_url,omitempty" xml:"img_url,omitempty"`
	// 跳转链接
	LinkUrl string `json:"link_url,omitempty" xml:"link_url,omitempty"`
	// 价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 首次上架时间
	FirstStartsTime string `json:"first_starts_time,omitempty" xml:"first_starts_time,omitempty"`
	// 最近上架时间
	Starts string `json:"starts,omitempty" xml:"starts,omitempty"`
	// 主站类目路径，空格分隔多级
	CategoryId string `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 推广主体类型,item:商品,item_private_mini:独享橱窗,shop:店铺,content:内容,short_video:短视频,user_define:自定义;
	PromotionType string `json:"promotion_type,omitempty" xml:"promotion_type,omitempty"`
	// 推广子主体类型,item:商品，item_private_mini:独享橱窗，shop:店铺，you_hao_huo:有好货，picture:图文，short_video:短视频，live_room:直播间，live_spot:看点，tao_blocks:淘积木，user_define_url:自定义url
	SubPromotionType string `json:"sub_promotion_type,omitempty" xml:"sub_promotion_type,omitempty"`
	// 生命周期指标-同叶子类目均值对比
	LifeCycleDiffWithAvg string `json:"life_cycle_diff_with_avg,omitempty" xml:"life_cycle_diff_with_avg,omitempty"`
	// 生命周期指标-同叶子类目均值对比提示文案
	LifeCycleDiffWithAvgTips string `json:"life_cycle_diff_with_avg_tips,omitempty" xml:"life_cycle_diff_with_avg_tips,omitempty"`
	// 商品成长指标-成长文案
	GrowDesc string `json:"grow_desc,omitempty" xml:"grow_desc,omitempty"`
	// 物料id
	MaterialId int64 `json:"material_id,omitempty" xml:"material_id,omitempty"`
	// 物料类型,1:淘宝宝贝,2:淘宝店铺,3:cms LandingPage页面,4:cms 分流总页面,5:自定义,201:PC店铺LandingPage页面,202:无线店铺LandingPage页面,203:内容LandingPage页面,204:淘宝直播,210:淘宝直播用户ID,205:有好货,209:无线店铺Tab3页面,212:超级品牌日LandingPage页面,211:超级发布会LandingPage页面,301:ICBU商品,302:ICBU PC店铺LandingPage页面,303:ICBU PC店铺,8:用户ID,9:信息流淘积木用户ID,10:口碑用户ID,206:店铺优惠券,207:商品优惠券,208:直播单品,309:直播间商品半屏卡,304:1688商品,305:内容平台-短视频,306:附近门店-用户ID,307:饿了么-菜品,401:LAZADA商品,501:飞猪酒店,502:飞猪商品,503:飞猪短信,601:淘积木页面,1001:AE宝贝,1101:淘宝宝贝SKU,310:内容平台-订阅内容
	MaterialType int64 `json:"material_type,omitempty" xml:"material_type,omitempty"`
	// 销量
	BidCount int64 `json:"bid_count,omitempty" xml:"bid_count,omitempty"`
	// 库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 生命周期指标-天数
	LifeCycleDays int64 `json:"life_cycle_days,omitempty" xml:"life_cycle_days,omitempty"`
	// 商品成长指标-成交排名
	GrowRank int64 `json:"grow_rank,omitempty" xml:"grow_rank,omitempty"`
}
