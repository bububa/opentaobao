package tuanhotel

// TuanItemSellParamVO 结构体
type TuanItemSellParamVO struct {
	// 免费赠品
	Gift string `json:"gift,omitempty" xml:"gift,omitempty"`
	// 预约说明，字数限制200
	AppointExplain string `json:"appoint_explain,omitempty" xml:"appoint_explain,omitempty"`
	// PC端图片地址
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 秒杀商品平台：web-电脑，wap-手机
	SecKills []string `json:"sec_kills,omitempty" xml:"sec_kills>string,omitempty"`
	// 当前开始到哪天结束
	DuringEndDate string `json:"during_end_date,omitempty" xml:"during_end_date,omitempty"`
	// 不可用日期（范围）
	UnavailableDates string `json:"unavailable_dates,omitempty" xml:"unavailable_dates,omitempty"`
	// 购买成功后有效天数
	EffectiveDays int64 `json:"effective_days,omitempty" xml:"effective_days,omitempty"`
	// 若为酒店客房优惠券类目则显示价格，其他类目则显示第一条sku的价格
	Price string `json:"price,omitempty" xml:"price,omitempty"`
	// 售卖时间小时，售卖时间类型1时返回
	BeginHour string `json:"begin_hour,omitempty" xml:"begin_hour,omitempty"`
	// 确认时间，可选字段：2,6,9,18.确认类型为2时返回
	ConfirmTime string `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 橱窗推荐：0-不推荐，1-推荐
	PromotedStatus int64 `json:"promoted_status,omitempty" xml:"promoted_status,omitempty"`
	// 售卖时间分钟，售卖时间类型1返回
	BeginMin string `json:"begin_min,omitempty" xml:"begin_min,omitempty"`
	// 类目ID：国内酒店套餐-201189402；国际酒店套餐-201188002；酒店餐饮美食-201214101；酒店服务-201214201；酒店客房优惠券-201214301
	CatId int64 `json:"cat_id,omitempty" xml:"cat_id,omitempty"`
	// 宝贝ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 自动下架时间小时
	DownShelfHour string `json:"down_shelf_hour,omitempty" xml:"down_shelf_hour,omitempty"`
	// 提前预约天数，数值范围为（0，1，2，3，5，7，15，30），0代表无需设置
	AdvanceDays int64 `json:"advance_days,omitempty" xml:"advance_days,omitempty"`
	// 0-无早餐，1-单份早餐，2-两份早餐，3-三份早餐，4-多份早餐。若为酒店服务或酒店餐饮美食与早餐类型无关类目，此处为-1
	Breakfast int64 `json:"breakfast,omitempty" xml:"breakfast,omitempty"`
	// 主图视频封面图和缩略图
	MainVideoPicUrl string `json:"main_video_pic_url,omitempty" xml:"main_video_pic_url,omitempty"`
	// 支持过期退款比率
	Refund int64 `json:"refund,omitempty" xml:"refund,omitempty"`
	// 第一条sku的间夜
	NightCount int64 `json:"night_count,omitempty" xml:"night_count,omitempty"`
	// 城市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 若为酒店客房优惠券类目则显示价格，其他类目则显示第一条sku的原价
	OrigPrice string `json:"orig_price,omitempty" xml:"orig_price,omitempty"`
	// 商家店铺类目
	ShopCategoriesIdList []string `json:"shop_categories_id_list,omitempty" xml:"shop_categories_id_list>string,omitempty"`
	// 拍下减库存，0-否，1-是
	SubStockAtBuy int64 `json:"sub_stock_at_buy,omitempty" xml:"sub_stock_at_buy,omitempty"`
	// 预约时段结束日期，预约时段类型为1返回
	EndAvailableDate string `json:"end_available_date,omitempty" xml:"end_available_date,omitempty"`
	// 预约时段类型：1-开始-结束日期，2-购买成功到哪天，3-购买成功多少天内
	EffectiveDateType int64 `json:"effective_date_type,omitempty" xml:"effective_date_type,omitempty"`
	// 费用包含
	FeeInclude string `json:"fee_include,omitempty" xml:"fee_include,omitempty"`
	// 省份
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// 卖家昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 旺铺详情
	WangpuDetail string `json:"wangpu_detail,omitempty" xml:"wangpu_detail,omitempty"`
	// 无线端图片地址
	WirelessPicUrls []string `json:"wireless_pic_urls,omitempty" xml:"wireless_pic_urls>string,omitempty"`
	// 亮点2
	Sub2Title string `json:"sub2_title,omitempty" xml:"sub2_title,omitempty"`
	// 预约时段结束日期，预约时段类型为1返回
	EndEffectiveDate string `json:"end_effective_date,omitempty" xml:"end_effective_date,omitempty"`
	// C卖家有效。酒店代订服务-671582076,酒店住宿服务-2773361951
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 星期几不可入住
	UnavailableWeeks string `json:"unavailable_weeks,omitempty" xml:"unavailable_weeks,omitempty"`
	// 亮点3
	Sub3Title string `json:"sub3_title,omitempty" xml:"sub3_title,omitempty"`
	// 店铺卖家ID
	StoreSellerId string `json:"store_seller_id,omitempty" xml:"store_seller_id,omitempty"`
	// 售卖时间类型：0-立刻,1-设定时间,2-放入仓库
	BeginType int64 `json:"begin_type,omitempty" xml:"begin_type,omitempty"`
	// 亮点4
	Sub4Title string `json:"sub4_title,omitempty" xml:"sub4_title,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 是否电子凭证
	Etc int64 `json:"etc,omitempty" xml:"etc,omitempty"`
	// 预约时段开始日期，预约时段类型为1时返回
	StartEffectiveDate string `json:"start_effective_date,omitempty" xml:"start_effective_date,omitempty"`
	// 自动下架时间分钟
	DownShelfMin string `json:"down_shelf_min,omitempty" xml:"down_shelf_min,omitempty"`
	// 确认类型，1-即时确认，2-二次确认，日历库存下返回
	ConfirmType string `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
	// 无线端宝贝描述
	WlDescription string `json:"wl_description,omitempty" xml:"wl_description,omitempty"`
	// 库存类型，1-普通库存，5-日历库存
	InventoryType int64 `json:"inventory_type,omitempty" xml:"inventory_type,omitempty"`
	// 支持过期退款比率
	AutoRefund int64 `json:"auto_refund,omitempty" xml:"auto_refund,omitempty"`
	// 门店卖家昵称
	StoreSellerNick string `json:"store_seller_nick,omitempty" xml:"store_seller_nick,omitempty"`
	// 早餐描述
	BreakfastDesc string `json:"breakfast_desc,omitempty" xml:"breakfast_desc,omitempty"`
	// 拍卖点，对C卖家有效
	AuctionPoint int64 `json:"auction_point,omitempty" xml:"auction_point,omitempty"`
	// 包含元素
	ContainElements []string `json:"contain_elements,omitempty" xml:"contain_elements>string,omitempty"`
	// 参与会员打折 0-不参与，1-参与
	OptionPromoted int64 `json:"option_promoted,omitempty" xml:"option_promoted,omitempty"`
	// PC端宝贝描述
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 亮点1
	Sub1Title string `json:"sub1_title,omitempty" xml:"sub1_title,omitempty"`
	// 可入住日期开始
	StartAvailableDate string `json:"start_available_date,omitempty" xml:"start_available_date,omitempty"`
	// 酒店客房优惠券类目下的代金券数量
	Amount int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// 若为酒店客房优惠券类目则显示库存，其他类目则显示第一条sku的库存
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 费用包含
	FeeExclude string `json:"fee_exclude,omitempty" xml:"fee_exclude,omitempty"`
	// 发票说明
	InvoiceExplain string `json:"invoice_explain,omitempty" xml:"invoice_explain,omitempty"`
	// 分润比例（推广费比例）
	ProfitRate string `json:"profit_rate,omitempty" xml:"profit_rate,omitempty"`
	// 自动下架时间日期
	DownShelfDate string `json:"down_shelf_date,omitempty" xml:"down_shelf_date,omitempty"`
	// 附加费用
	AdditionalPay string `json:"additional_pay,omitempty" xml:"additional_pay,omitempty"`
	// 售卖时间日期，售卖时间类型1返回
	BeginDate string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// 商家编码
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 是否有发票 0-没有，1-店铺提供，2-卖家提供
	Invoice int64 `json:"invoice,omitempty" xml:"invoice,omitempty"`
}
