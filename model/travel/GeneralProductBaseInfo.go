package travel

// GeneralProductBaseInfo 结构体
type GeneralProductBaseInfo struct {
	// 品图片路径。最多支持5张，第一张为主图 必填，其余四张可选填（多张图片间使用英文逗号分隔）。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为png、jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。
	PicUrls []string `json:"pic_urls,omitempty" xml:"pic_urls>string,omitempty"`
	// 商品亮点。1）目前最多支持4个亮点，超过4个的亮点描述不会被保存 2）每个亮点描述35个字符以内 3）每个亮点间用英文逗号分隔
	SubTitles []string `json:"sub_titles,omitempty" xml:"sub_titles>string,omitempty"`
	// 商品属性列表；由类目的属性PID和VID组成，属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。
	Props []CatPropInfo `json:"props,omitempty" xml:"props>cat_prop_info,omitempty"`
	// 商品描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 手机描述 格式:标题 描述图片路径
	WapDesc string `json:"wap_desc,omitempty" xml:"wap_desc,omitempty"`
	// 商家编码
	OutId string `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 目的地，填写中文，以英文逗号分隔，最多可填12个，如果国家底下还有城市，则必须精确到城市。如果接口报错提示目的地不存在，则可登录商家基础信息映射管理后台（https://sell.alitrip.com/sell/basicdata/BasicDataMapping.htm）修正并生效映射关系，映射关系生效后，无法识别的目的地将自动替换为映射值 【重要提示-此字段会作为搜索/筛选条件】
	ToLocations string `json:"to_locations,omitempty" xml:"to_locations,omitempty"`
	// 必填，商品标题。30个字符以内 【重要-此字段会作为搜索条件】
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 最晚收客时间:分钟。仅个别类目支持
	ReserveDeadlineMinutes int64 `json:"reserve_deadline_minutes,omitempty" xml:"reserve_deadline_minutes,omitempty"`
	// 商品类目id，发布商品必填，编辑选填；支持的线上类目ID，船票：50018298
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 淘宝平台商品ID 产品更新时使用
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
	ConfirmTime int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
	// 至少提前天数，最晚成团提前天数，最小0天，最大为300天；不传则为0
	Duration int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// 最晚收客时间:小时。仅个别类目支持
	ReserveDeadlineHours int64 `json:"reserve_deadline_hours,omitempty" xml:"reserve_deadline_hours,omitempty"`
	// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
	ConfirmType int64 `json:"confirm_type,omitempty" xml:"confirm_type,omitempty"`
}
