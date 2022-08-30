package travel

// PontusTravelItemBaseInfo 结构体
type PontusTravelItemBaseInfo struct {
	// 商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。商品主图需要关联用户图片空间的绝对url。 这个url所对应的图片必须要属于当前用户。
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 商品亮点。1）目前最多支持4个亮点，超过4个的亮点描述不会被保存 2）每个亮点描述35个字符以内
	SubTitles []string `json:"sub_titles,omitempty" xml:"sub_titles>string,omitempty"`
	// 宝贝所在地省份
	Prov string `json:"prov,omitempty" xml:"prov,omitempty"`
	// pc端更多商品描述，不超过50000个字符
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 手机端描述。格式： 使用xml标签（shortDesc，txt和img）进行图文混排，shortDesc标签表示添加小标题，txt标签表示添加文本描述，img标签表示添加一张图片。  注意： 1）shortDesc，txt和img三个xml标签独立，可以乱序使用，但不能嵌套 2）图片路径只能使用商家图片空间内图片的绝对路径
	WapDesc string `json:"wap_desc,omitempty" xml:"wap_desc,omitempty"`
	// 所有商品必填，目的地（城市或国家），填写中文。城市名，以英文逗号分隔，最多可填12个， 城市的值可从基础数据API中获取（http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.hSqBrl&amp;apiId=25760），如果国家下还有城市，则必须精确到城市。【重要提示-此字段会作为搜索/筛选条件】
	ToLocations string `json:"to_locations,omitempty" xml:"to_locations,omitempty"`
	// 商家定义的编码
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 宝贝所在市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 跟团游商品必填。跟团游商品必填。出发地（城市），填写中文，必须对应类目下存在的出发地，否则会报出发地不存在，可用的城市名称可从基础API中获取：http://open.taobao.com/doc2/apiDetail?spm=0.0.0.0.hSqBrl&amp;apiId=25760  【重要提示-此属性会作为搜索/筛选条件】
	FromLocations string `json:"from_locations,omitempty" xml:"from_locations,omitempty"`
	// 商品标题。30个字符以内
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 预留，商家备注
	OuterTitle string `json:"outer_title,omitempty" xml:"outer_title,omitempty"`
	// 商品标签
	ItemTagContent string `json:"item_tag_content,omitempty" xml:"item_tag_content,omitempty"`
	// 最大行程天数
	TripMaxDays int64 `json:"trip_max_days,omitempty" xml:"trip_max_days,omitempty"`
	// 必填，商品类型。1-国内跟团游 2- 国内自由行 3-出境跟团游 4- 出境自由行 5-境外当地玩乐 6-国际邮轮 9-国内邮轮
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// 行程晚数，必须大于0，且小于等于行程天数 这里注意下，trip_days（行程天数）&gt;=accom_nights（行程晚数）&gt;=hotel_days（住宿晚数）
	AccomNights int64 `json:"accom_nights,omitempty" xml:"accom_nights,omitempty"`
	// 最小行程天数
	TripMinDays int64 `json:"trip_min_days,omitempty" xml:"trip_min_days,omitempty"`
}
