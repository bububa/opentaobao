package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripDaytoursProductUploadAPIRequest 境外一日游/多日游 产品维护接口 API请求
// alitrip.daytours.product.upload
//
// 境外一日游/多日游 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
type AlitripDaytoursProductUploadAPIRequest struct {
	model.Params
	// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
	_outProductId string
	// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
	_itemId int64
	// 新发布商品时必填。商品标题，30个中文字符以内
	_title string
	// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
	_subTitles []string
	// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
	_picUrls []string
	// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
	_fromLocations string
	// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
	_toLocations string
	// 新发布商品时必填。旅游天数
	_tripDay int64
	// 可选，旅游晚数，不传默认旅游天数-1
	_tripNight int64
	// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
	_descXml string
	// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
	_descHtml string
	// 可选，手机端详情描述，xml格式，格式详见示例。
	_wapDesc string
	// 特殊必填，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
	_itineraries []string
	// 特殊必填，行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
	_structItineraries []StructItinerary
	// 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_feeInclude []string
	// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_feeExclude []string
	// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_orderInfo []string
	// 真实的旅游服务提供商
	_touristServiceProvider string
	// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
	_refundType int64
	// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
	_refundRegulationsJson string
	// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
	_reserveLimit string
	// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
	_confirmType int64
	// 特殊可选（confirm_type为2时必填），资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
	_confirmTime int64
	// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
	_subStock int64
	// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
	_travellerTemplateId int64
	// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
	_itemCustomTag string
	// 一日游 产品亮点
	_highLights []ProductHighLights
	// 必填，营业执照图片路径。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在3M以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间。
	_businessLicense string
	// 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
	_sellerCids []string
	// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
	_secondKill string
	// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
	_hasDiscount bool
}

// NewAlitripDaytoursProductUploadRequest 初始化AlitripDaytoursProductUploadAPIRequest对象
func NewAlitripDaytoursProductUploadRequest() *AlitripDaytoursProductUploadAPIRequest {
	return &AlitripDaytoursProductUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripDaytoursProductUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.daytours.product.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripDaytoursProductUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOutProductId is OutProductId Setter
// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripDaytoursProductUploadAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetItemId is ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripDaytoursProductUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetTitle is Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripDaytoursProductUploadAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetTitle() string {
	return r._title
}

// SetSubTitles is SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadAPIRequest) SetSubTitles(_subTitles []string) error {
	r._subTitles = _subTitles
	r.Set("sub_titles", _subTitles)
	return nil
}

// GetSubTitles SubTitles Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetSubTitles() []string {
	return r._subTitles
}

// SetPicUrls is PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadAPIRequest) SetPicUrls(_picUrls []string) error {
	r._picUrls = _picUrls
	r.Set("pic_urls", _picUrls)
	return nil
}

// GetPicUrls PicUrls Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetPicUrls() []string {
	return r._picUrls
}

// SetFromLocations is FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripDaytoursProductUploadAPIRequest) SetFromLocations(_fromLocations string) error {
	r._fromLocations = _fromLocations
	r.Set("from_locations", _fromLocations)
	return nil
}

// GetFromLocations FromLocations Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetFromLocations() string {
	return r._fromLocations
}

// SetToLocations is ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripDaytoursProductUploadAPIRequest) SetToLocations(_toLocations string) error {
	r._toLocations = _toLocations
	r.Set("to_locations", _toLocations)
	return nil
}

// GetToLocations ToLocations Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetToLocations() string {
	return r._toLocations
}

// SetTripDay is TripDay Setter
// 新发布商品时必填。旅游天数
func (r *AlitripDaytoursProductUploadAPIRequest) SetTripDay(_tripDay int64) error {
	r._tripDay = _tripDay
	r.Set("trip_day", _tripDay)
	return nil
}

// GetTripDay TripDay Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetTripDay() int64 {
	return r._tripDay
}

// SetTripNight is TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1
func (r *AlitripDaytoursProductUploadAPIRequest) SetTripNight(_tripNight int64) error {
	r._tripNight = _tripNight
	r.Set("trip_night", _tripNight)
	return nil
}

// GetTripNight TripNight Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetTripNight() int64 {
	return r._tripNight
}

// SetDescXml is DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripDaytoursProductUploadAPIRequest) SetDescXml(_descXml string) error {
	r._descXml = _descXml
	r.Set("desc_xml", _descXml)
	return nil
}

// GetDescXml DescXml Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetDescXml() string {
	return r._descXml
}

// SetDescHtml is DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripDaytoursProductUploadAPIRequest) SetDescHtml(_descHtml string) error {
	r._descHtml = _descHtml
	r.Set("desc_html", _descHtml)
	return nil
}

// GetDescHtml DescHtml Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetDescHtml() string {
	return r._descHtml
}

// SetWapDesc is WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripDaytoursProductUploadAPIRequest) SetWapDesc(_wapDesc string) error {
	r._wapDesc = _wapDesc
	r.Set("wap_desc", _wapDesc)
	return nil
}

// GetWapDesc WapDesc Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetWapDesc() string {
	return r._wapDesc
}

// SetItineraries is Itineraries Setter
// 特殊必填，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadAPIRequest) SetItineraries(_itineraries []string) error {
	r._itineraries = _itineraries
	r.Set("itineraries", _itineraries)
	return nil
}

// GetItineraries Itineraries Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetItineraries() []string {
	return r._itineraries
}

// SetStructItineraries is StructItineraries Setter
// 特殊必填，行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
func (r *AlitripDaytoursProductUploadAPIRequest) SetStructItineraries(_structItineraries []StructItinerary) error {
	r._structItineraries = _structItineraries
	r.Set("struct_itineraries", _structItineraries)
	return nil
}

// GetStructItineraries StructItineraries Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetStructItineraries() []StructItinerary {
	return r._structItineraries
}

// SetFeeInclude is FeeInclude Setter
// 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadAPIRequest) SetFeeInclude(_feeInclude []string) error {
	r._feeInclude = _feeInclude
	r.Set("fee_include", _feeInclude)
	return nil
}

// GetFeeInclude FeeInclude Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetFeeInclude() []string {
	return r._feeInclude
}

// SetFeeExclude is FeeExclude Setter
// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadAPIRequest) SetFeeExclude(_feeExclude []string) error {
	r._feeExclude = _feeExclude
	r.Set("fee_exclude", _feeExclude)
	return nil
}

// GetFeeExclude FeeExclude Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetFeeExclude() []string {
	return r._feeExclude
}

// SetOrderInfo is OrderInfo Setter
// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadAPIRequest) SetOrderInfo(_orderInfo []string) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetOrderInfo() []string {
	return r._orderInfo
}

// SetTouristServiceProvider is TouristServiceProvider Setter
// 真实的旅游服务提供商
func (r *AlitripDaytoursProductUploadAPIRequest) SetTouristServiceProvider(_touristServiceProvider string) error {
	r._touristServiceProvider = _touristServiceProvider
	r.Set("tourist_service_provider", _touristServiceProvider)
	return nil
}

// GetTouristServiceProvider TouristServiceProvider Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetTouristServiceProvider() string {
	return r._touristServiceProvider
}

// SetRefundType is RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
func (r *AlitripDaytoursProductUploadAPIRequest) SetRefundType(_refundType int64) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetRefundType() int64 {
	return r._refundType
}

// SetRefundRegulationsJson is RefundRegulationsJson Setter
// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
func (r *AlitripDaytoursProductUploadAPIRequest) SetRefundRegulationsJson(_refundRegulationsJson string) error {
	r._refundRegulationsJson = _refundRegulationsJson
	r.Set("refund_regulations_json", _refundRegulationsJson)
	return nil
}

// GetRefundRegulationsJson RefundRegulationsJson Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetRefundRegulationsJson() string {
	return r._refundRegulationsJson
}

// SetReserveLimit is ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripDaytoursProductUploadAPIRequest) SetReserveLimit(_reserveLimit string) error {
	r._reserveLimit = _reserveLimit
	r.Set("reserve_limit", _reserveLimit)
	return nil
}

// GetReserveLimit ReserveLimit Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetReserveLimit() string {
	return r._reserveLimit
}

// SetConfirmType is ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripDaytoursProductUploadAPIRequest) SetConfirmType(_confirmType int64) error {
	r._confirmType = _confirmType
	r.Set("confirm_type", _confirmType)
	return nil
}

// GetConfirmType ConfirmType Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetConfirmType() int64 {
	return r._confirmType
}

// SetConfirmTime is ConfirmTime Setter
// 特殊可选（confirm_type为2时必填），资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripDaytoursProductUploadAPIRequest) SetConfirmTime(_confirmTime int64) error {
	r._confirmTime = _confirmTime
	r.Set("confirm_time", _confirmTime)
	return nil
}

// GetConfirmTime ConfirmTime Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetConfirmTime() int64 {
	return r._confirmTime
}

// SetSubStock is SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripDaytoursProductUploadAPIRequest) SetSubStock(_subStock int64) error {
	r._subStock = _subStock
	r.Set("sub_stock", _subStock)
	return nil
}

// GetSubStock SubStock Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetSubStock() int64 {
	return r._subStock
}

// SetTravellerTemplateId is TravellerTemplateId Setter
// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
func (r *AlitripDaytoursProductUploadAPIRequest) SetTravellerTemplateId(_travellerTemplateId int64) error {
	r._travellerTemplateId = _travellerTemplateId
	r.Set("traveller_template_id", _travellerTemplateId)
	return nil
}

// GetTravellerTemplateId TravellerTemplateId Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetTravellerTemplateId() int64 {
	return r._travellerTemplateId
}

// SetItemCustomTag is ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripDaytoursProductUploadAPIRequest) SetItemCustomTag(_itemCustomTag string) error {
	r._itemCustomTag = _itemCustomTag
	r.Set("item_custom_tag", _itemCustomTag)
	return nil
}

// GetItemCustomTag ItemCustomTag Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetItemCustomTag() string {
	return r._itemCustomTag
}

// SetHighLights is HighLights Setter
// 一日游 产品亮点
func (r *AlitripDaytoursProductUploadAPIRequest) SetHighLights(_highLights []ProductHighLights) error {
	r._highLights = _highLights
	r.Set("high_lights", _highLights)
	return nil
}

// GetHighLights HighLights Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetHighLights() []ProductHighLights {
	return r._highLights
}

// SetBusinessLicense is BusinessLicense Setter
// 必填，营业执照图片路径。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在3M以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间。
func (r *AlitripDaytoursProductUploadAPIRequest) SetBusinessLicense(_businessLicense string) error {
	r._businessLicense = _businessLicense
	r.Set("business_license", _businessLicense)
	return nil
}

// GetBusinessLicense BusinessLicense Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetBusinessLicense() string {
	return r._businessLicense
}

// SetSellerCids is SellerCids Setter
// 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
func (r *AlitripDaytoursProductUploadAPIRequest) SetSellerCids(_sellerCids []string) error {
	r._sellerCids = _sellerCids
	r.Set("seller_cids", _sellerCids)
	return nil
}

// GetSellerCids SellerCids Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetSellerCids() []string {
	return r._sellerCids
}

// SetSecondKill is SecondKill Setter
// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
func (r *AlitripDaytoursProductUploadAPIRequest) SetSecondKill(_secondKill string) error {
	r._secondKill = _secondKill
	r.Set("second_kill", _secondKill)
	return nil
}

// GetSecondKill SecondKill Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetSecondKill() string {
	return r._secondKill
}

// SetHasDiscount is HasDiscount Setter
// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
func (r *AlitripDaytoursProductUploadAPIRequest) SetHasDiscount(_hasDiscount bool) error {
	r._hasDiscount = _hasDiscount
	r.Set("has_discount", _hasDiscount)
	return nil
}

// GetHasDiscount HasDiscount Getter
func (r AlitripDaytoursProductUploadAPIRequest) GetHasDiscount() bool {
	return r._hasDiscount
}
