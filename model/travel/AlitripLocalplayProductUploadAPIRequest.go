package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripLocalplayProductUploadAPIRequest 当地玩乐 产品维护接口 API请求
// alitrip.localplay.product.upload
//
// 当地玩乐（境内当地玩乐/境外玩乐套餐） 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
type AlitripLocalplayProductUploadAPIRequest struct {
	model.Params
	// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_feeExclude []string
	// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_orderInfo []string
	// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
	_picUrls []string
	// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
	_subTitles []string
	// 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
	_sellerCids []string
	// 可选，手机端详情描述，xml格式，格式详见示例。
	_wapDesc string
	// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
	_descXml string
	// 新发布商品时必填。商品标题，30个中文字符以内
	_title string
	// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
	_itemCustomTag string
	// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
	_outProductId string
	// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
	_toLocations string
	// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
	_reserveLimit string
	// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
	_descHtml string
	// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
	_fromLocations string
	// 代订服务说明（请填写真实的旅游服务提供商）
	_touristServiceProvider string
	// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
	_refundRegulationsJson string
	// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
	_secondKill string
	// 新发布商品时必填。旅游天数
	_tripDay int64
	// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
	_subStock int64
	// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
	_confirmTime int64
	// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
	_confirmType int64
	// 可选，旅游晚数，不传默认旅游天数-1
	_tripNight int64
	// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
	_itemId int64
	// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-线路商品（跟团、自由行等）新版自定义退改规则。不传默认为0
	_refundType int64
	// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
	_travellerTemplateId int64
	// 新发布商品时必填。是否出境游，0-不是，1-是。
	_isOverseasTour int64
	// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
	_hasDiscount bool
}

// NewAlitripLocalplayProductUploadRequest 初始化AlitripLocalplayProductUploadAPIRequest对象
func NewAlitripLocalplayProductUploadRequest() *AlitripLocalplayProductUploadAPIRequest {
	return &AlitripLocalplayProductUploadAPIRequest{
		Params: model.NewParams(27),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripLocalplayProductUploadAPIRequest) Reset() {
	r._feeExclude = r._feeExclude[:0]
	r._orderInfo = r._orderInfo[:0]
	r._picUrls = r._picUrls[:0]
	r._subTitles = r._subTitles[:0]
	r._sellerCids = r._sellerCids[:0]
	r._wapDesc = ""
	r._descXml = ""
	r._title = ""
	r._itemCustomTag = ""
	r._outProductId = ""
	r._toLocations = ""
	r._reserveLimit = ""
	r._descHtml = ""
	r._fromLocations = ""
	r._touristServiceProvider = ""
	r._refundRegulationsJson = ""
	r._secondKill = ""
	r._tripDay = 0
	r._subStock = 0
	r._confirmTime = 0
	r._confirmType = 0
	r._tripNight = 0
	r._itemId = 0
	r._refundType = 0
	r._travellerTemplateId = 0
	r._isOverseasTour = 0
	r._hasDiscount = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripLocalplayProductUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.localplay.product.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripLocalplayProductUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripLocalplayProductUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFeeExclude is FeeExclude Setter
// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadAPIRequest) SetFeeExclude(_feeExclude []string) error {
	r._feeExclude = _feeExclude
	r.Set("fee_exclude", _feeExclude)
	return nil
}

// GetFeeExclude FeeExclude Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetFeeExclude() []string {
	return r._feeExclude
}

// SetOrderInfo is OrderInfo Setter
// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadAPIRequest) SetOrderInfo(_orderInfo []string) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetOrderInfo() []string {
	return r._orderInfo
}

// SetPicUrls is PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadAPIRequest) SetPicUrls(_picUrls []string) error {
	r._picUrls = _picUrls
	r.Set("pic_urls", _picUrls)
	return nil
}

// GetPicUrls PicUrls Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetPicUrls() []string {
	return r._picUrls
}

// SetSubTitles is SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadAPIRequest) SetSubTitles(_subTitles []string) error {
	r._subTitles = _subTitles
	r.Set("sub_titles", _subTitles)
	return nil
}

// GetSubTitles SubTitles Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetSubTitles() []string {
	return r._subTitles
}

// SetSellerCids is SellerCids Setter
// 关联商品与店铺类目 结构:&#34;cid1,cid2,...,&#34;。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
func (r *AlitripLocalplayProductUploadAPIRequest) SetSellerCids(_sellerCids []string) error {
	r._sellerCids = _sellerCids
	r.Set("seller_cids", _sellerCids)
	return nil
}

// GetSellerCids SellerCids Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetSellerCids() []string {
	return r._sellerCids
}

// SetWapDesc is WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripLocalplayProductUploadAPIRequest) SetWapDesc(_wapDesc string) error {
	r._wapDesc = _wapDesc
	r.Set("wap_desc", _wapDesc)
	return nil
}

// GetWapDesc WapDesc Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetWapDesc() string {
	return r._wapDesc
}

// SetDescXml is DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripLocalplayProductUploadAPIRequest) SetDescXml(_descXml string) error {
	r._descXml = _descXml
	r.Set("desc_xml", _descXml)
	return nil
}

// GetDescXml DescXml Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetDescXml() string {
	return r._descXml
}

// SetTitle is Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripLocalplayProductUploadAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetTitle() string {
	return r._title
}

// SetItemCustomTag is ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripLocalplayProductUploadAPIRequest) SetItemCustomTag(_itemCustomTag string) error {
	r._itemCustomTag = _itemCustomTag
	r.Set("item_custom_tag", _itemCustomTag)
	return nil
}

// GetItemCustomTag ItemCustomTag Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetItemCustomTag() string {
	return r._itemCustomTag
}

// SetOutProductId is OutProductId Setter
// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripLocalplayProductUploadAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetToLocations is ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripLocalplayProductUploadAPIRequest) SetToLocations(_toLocations string) error {
	r._toLocations = _toLocations
	r.Set("to_locations", _toLocations)
	return nil
}

// GetToLocations ToLocations Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetToLocations() string {
	return r._toLocations
}

// SetReserveLimit is ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripLocalplayProductUploadAPIRequest) SetReserveLimit(_reserveLimit string) error {
	r._reserveLimit = _reserveLimit
	r.Set("reserve_limit", _reserveLimit)
	return nil
}

// GetReserveLimit ReserveLimit Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetReserveLimit() string {
	return r._reserveLimit
}

// SetDescHtml is DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripLocalplayProductUploadAPIRequest) SetDescHtml(_descHtml string) error {
	r._descHtml = _descHtml
	r.Set("desc_html", _descHtml)
	return nil
}

// GetDescHtml DescHtml Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetDescHtml() string {
	return r._descHtml
}

// SetFromLocations is FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripLocalplayProductUploadAPIRequest) SetFromLocations(_fromLocations string) error {
	r._fromLocations = _fromLocations
	r.Set("from_locations", _fromLocations)
	return nil
}

// GetFromLocations FromLocations Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetFromLocations() string {
	return r._fromLocations
}

// SetTouristServiceProvider is TouristServiceProvider Setter
// 代订服务说明（请填写真实的旅游服务提供商）
func (r *AlitripLocalplayProductUploadAPIRequest) SetTouristServiceProvider(_touristServiceProvider string) error {
	r._touristServiceProvider = _touristServiceProvider
	r.Set("tourist_service_provider", _touristServiceProvider)
	return nil
}

// GetTouristServiceProvider TouristServiceProvider Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetTouristServiceProvider() string {
	return r._touristServiceProvider
}

// SetRefundRegulationsJson is RefundRegulationsJson Setter
// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
func (r *AlitripLocalplayProductUploadAPIRequest) SetRefundRegulationsJson(_refundRegulationsJson string) error {
	r._refundRegulationsJson = _refundRegulationsJson
	r.Set("refund_regulations_json", _refundRegulationsJson)
	return nil
}

// GetRefundRegulationsJson RefundRegulationsJson Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetRefundRegulationsJson() string {
	return r._refundRegulationsJson
}

// SetSecondKill is SecondKill Setter
// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
func (r *AlitripLocalplayProductUploadAPIRequest) SetSecondKill(_secondKill string) error {
	r._secondKill = _secondKill
	r.Set("second_kill", _secondKill)
	return nil
}

// GetSecondKill SecondKill Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetSecondKill() string {
	return r._secondKill
}

// SetTripDay is TripDay Setter
// 新发布商品时必填。旅游天数
func (r *AlitripLocalplayProductUploadAPIRequest) SetTripDay(_tripDay int64) error {
	r._tripDay = _tripDay
	r.Set("trip_day", _tripDay)
	return nil
}

// GetTripDay TripDay Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetTripDay() int64 {
	return r._tripDay
}

// SetSubStock is SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripLocalplayProductUploadAPIRequest) SetSubStock(_subStock int64) error {
	r._subStock = _subStock
	r.Set("sub_stock", _subStock)
	return nil
}

// GetSubStock SubStock Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetSubStock() int64 {
	return r._subStock
}

// SetConfirmTime is ConfirmTime Setter
// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripLocalplayProductUploadAPIRequest) SetConfirmTime(_confirmTime int64) error {
	r._confirmTime = _confirmTime
	r.Set("confirm_time", _confirmTime)
	return nil
}

// GetConfirmTime ConfirmTime Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetConfirmTime() int64 {
	return r._confirmTime
}

// SetConfirmType is ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripLocalplayProductUploadAPIRequest) SetConfirmType(_confirmType int64) error {
	r._confirmType = _confirmType
	r.Set("confirm_type", _confirmType)
	return nil
}

// GetConfirmType ConfirmType Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetConfirmType() int64 {
	return r._confirmType
}

// SetTripNight is TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1
func (r *AlitripLocalplayProductUploadAPIRequest) SetTripNight(_tripNight int64) error {
	r._tripNight = _tripNight
	r.Set("trip_night", _tripNight)
	return nil
}

// GetTripNight TripNight Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetTripNight() int64 {
	return r._tripNight
}

// SetItemId is ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripLocalplayProductUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetRefundType is RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-线路商品（跟团、自由行等）新版自定义退改规则。不传默认为0
func (r *AlitripLocalplayProductUploadAPIRequest) SetRefundType(_refundType int64) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetRefundType() int64 {
	return r._refundType
}

// SetTravellerTemplateId is TravellerTemplateId Setter
// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具-&gt;出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
func (r *AlitripLocalplayProductUploadAPIRequest) SetTravellerTemplateId(_travellerTemplateId int64) error {
	r._travellerTemplateId = _travellerTemplateId
	r.Set("traveller_template_id", _travellerTemplateId)
	return nil
}

// GetTravellerTemplateId TravellerTemplateId Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetTravellerTemplateId() int64 {
	return r._travellerTemplateId
}

// SetIsOverseasTour is IsOverseasTour Setter
// 新发布商品时必填。是否出境游，0-不是，1-是。
func (r *AlitripLocalplayProductUploadAPIRequest) SetIsOverseasTour(_isOverseasTour int64) error {
	r._isOverseasTour = _isOverseasTour
	r.Set("is_overseas_tour", _isOverseasTour)
	return nil
}

// GetIsOverseasTour IsOverseasTour Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetIsOverseasTour() int64 {
	return r._isOverseasTour
}

// SetHasDiscount is HasDiscount Setter
// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
func (r *AlitripLocalplayProductUploadAPIRequest) SetHasDiscount(_hasDiscount bool) error {
	r._hasDiscount = _hasDiscount
	r.Set("has_discount", _hasDiscount)
	return nil
}

// GetHasDiscount HasDiscount Getter
func (r AlitripLocalplayProductUploadAPIRequest) GetHasDiscount() bool {
	return r._hasDiscount
}

var poolAlitripLocalplayProductUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripLocalplayProductUploadRequest()
	},
}

// GetAlitripLocalplayProductUploadRequest 从 sync.Pool 获取 AlitripLocalplayProductUploadAPIRequest
func GetAlitripLocalplayProductUploadAPIRequest() *AlitripLocalplayProductUploadAPIRequest {
	return poolAlitripLocalplayProductUploadAPIRequest.Get().(*AlitripLocalplayProductUploadAPIRequest)
}

// ReleaseAlitripLocalplayProductUploadAPIRequest 将 AlitripLocalplayProductUploadAPIRequest 放入 sync.Pool
func ReleaseAlitripLocalplayProductUploadAPIRequest(v *AlitripLocalplayProductUploadAPIRequest) {
	v.Reset()
	poolAlitripLocalplayProductUploadAPIRequest.Put(v)
}
