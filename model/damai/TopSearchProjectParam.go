package damai

// TopSearchProjectParam 结构体
type TopSearchProjectParam struct {
	// 搜索关键字
	Keyword string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// 一级分类名，支持多个（之间是OR关系），以"|"隔开。演唱会,音乐会,话剧歌剧,舞蹈芭蕾,曲苑杂坛,体育比赛,度假休闲,儿童亲子,旅游演艺,韩流地带,动漫,旅游展览
	CategoryName string `json:"category_name,omitempty" xml:"category_name,omitempty"`
	// 页码
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 每页大小,上限50条，不传默认10条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 关联艺人姓名（之间是OR关系），以"|"隔开
	ArtistName string `json:"artist_name,omitempty" xml:"artist_name,omitempty"`
	// 子分类，支持多个（之间是OR关系），以"|"隔开           一级分类对应子分类如下，演唱会:(摇滚,民族,流行,音乐节,其他) 音乐会:(声乐及合唱,室内乐及古乐,独奏,管弦乐,其他) 话剧歌剧:(儿童剧,歌剧,歌舞剧,话剧,音乐剧) 舞蹈芭蕾:(舞剧,舞蹈,芭蕾) 曲苑杂坛:(戏曲,杂技,相声,马戏,魔术,其他) 体育比赛:(冰雪,搏击运动,格斗,球类运动,田径,电竞,篮球,网球,赛车,足球,其它竞技,其他) 度假休闲:(主题公园,代金券,展会,度假村,温泉,游览线路,滑雪,特色体验,风景区)
	SubCategoryName string `json:"sub_category_name,omitempty" xml:"sub_category_name,omitempty"`
	// 过滤城市名，支持多个（之间是OR关系），以"|"隔开
	FilterCityName string `json:"filter_city_name,omitempty" xml:"filter_city_name,omitempty"`
	// 排序类型 1: 上架时间倒排序 2: 最近场次演出时间排序 0:按相关度降序 3:热门排序,  默认 0
	SortType int64 `json:"sort_type,omitempty" xml:"sort_type,omitempty"`
	// 0：无强指定 1:今天 2：明天 3：周末 4:30天内 5:自定义 6: 本周 7： 本月 8：本月周末场
	DateType int64 `json:"date_type,omitempty" xml:"date_type,omitempty"`
	// 开始日期，格式yyyy-MM-dd
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 结束日期，格式yyyy-MM-dd
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 渠道（之间是OR关系），以"|"隔开。10001, "大麦app"      *      10002, "大麦pc"      *      10003, "大麦h5"      *      20000, "天猫"      *      30000, "支付宝"      *      40000, "淘票票"      *      50000, "其他"
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
}
