package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripGrouptoursProductUploadAPIRequest 跟团游 产品维护接口 API请求
// alitrip.grouptours.product.upload
//
// 跟团游 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
type AlitripGrouptoursProductUploadAPIRequest struct {
	model.Params
	// 可选，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
	_itineraries []string
	// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_feeExclude []string
	// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_orderInfo []string
	// 特殊可选，当refund_type=1或7时，需要上传自定义退改内容。自定义退改规则，最多可含5组规则，每组规则间以英文逗号分隔。 1）当refund_type为1时格式为：a_b_num,b-1_c_num。含义：提前a天至提前b天发起退款，买家需支付num比例违约费。 2）当refund_type为7时格式为：a_b_num1_num2_0,b-1_c_num1_num2_0。含义：提前a天至提前b天发起退款，买家需支付num1比例违约费，卖家需支付num2比例违约费，最后一个数字代表是否节假日规则（0-不是，1-是）。特别注意，当refund_type为7时，自定义退改规则必须设置 n天以上违约规则 以及 行程当日违约规则，即第一组规则需要以-1_a_num1_num2_0或-1_a_num1_num2_1开头，且最后一组规则需要以0_0_num1_num2_0或0_0_num1_num2_1结尾。
	_refundRegulations []string
	// 可选，跟团时的集合地点，列表中每一个元素对应一个集合地点
	_gatherPlaces []GatherPlaceInfo
	// 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
	_feeInclude []string
	// 行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
	_structItineraries []StructItinerary
	// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
	_picUrls []string
	// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
	_subTitles []string
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
	// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
	_fromLocations string
	// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
	_descHtml string
	// 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船
	_goTrafficType int64
	// 新发布商品时必填。旅游天数
	_tripDay int64
	// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
	_subStock int64
	// 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船
	_backTrafficType int64
	// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
	_confirmTime int64
	// 新发布商品时必填。参团线路类型。0 -目的地参团，1-为出发地参团
	_routeType int64
	// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
	_confirmType int64
	// 可选，旅游晚数，不传默认旅游天数-1
	_tripNight int64
	// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
	_itemId int64
	// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
	_refundType int64
	// 可选，电子合同信息设置。
	_electronContract *ElectronContract
	// 可选，出行人模板id，预留，暂不支持
	_travellerTemplateId int64
	// 可选，是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
	_purePlay int64
	// 新发布商品时必填。是否出境游，0-不是，1-是。
	_isOverseasTour int64
}

// NewAlitripGrouptoursProductUploadRequest 初始化AlitripGrouptoursProductUploadAPIRequest对象
func NewAlitripGrouptoursProductUploadRequest() *AlitripGrouptoursProductUploadAPIRequest {
	return &AlitripGrouptoursProductUploadAPIRequest{
		Params: model.NewParams(32),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripGrouptoursProductUploadAPIRequest) Reset() {
	r._itineraries = r._itineraries[:0]
	r._feeExclude = r._feeExclude[:0]
	r._orderInfo = r._orderInfo[:0]
	r._refundRegulations = r._refundRegulations[:0]
	r._gatherPlaces = r._gatherPlaces[:0]
	r._feeInclude = r._feeInclude[:0]
	r._structItineraries = r._structItineraries[:0]
	r._picUrls = r._picUrls[:0]
	r._subTitles = r._subTitles[:0]
	r._wapDesc = ""
	r._descXml = ""
	r._title = ""
	r._itemCustomTag = ""
	r._outProductId = ""
	r._toLocations = ""
	r._reserveLimit = ""
	r._fromLocations = ""
	r._descHtml = ""
	r._goTrafficType = 0
	r._tripDay = 0
	r._subStock = 0
	r._backTrafficType = 0
	r._confirmTime = 0
	r._routeType = 0
	r._confirmType = 0
	r._tripNight = 0
	r._itemId = 0
	r._refundType = 0
	r._electronContract = nil
	r._travellerTemplateId = 0
	r._purePlay = 0
	r._isOverseasTour = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripGrouptoursProductUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.grouptours.product.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripGrouptoursProductUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripGrouptoursProductUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItineraries is Itineraries Setter
// 可选，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadAPIRequest) SetItineraries(_itineraries []string) error {
	r._itineraries = _itineraries
	r.Set("itineraries", _itineraries)
	return nil
}

// GetItineraries Itineraries Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetItineraries() []string {
	return r._itineraries
}

// SetFeeExclude is FeeExclude Setter
// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadAPIRequest) SetFeeExclude(_feeExclude []string) error {
	r._feeExclude = _feeExclude
	r.Set("fee_exclude", _feeExclude)
	return nil
}

// GetFeeExclude FeeExclude Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetFeeExclude() []string {
	return r._feeExclude
}

// SetOrderInfo is OrderInfo Setter
// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadAPIRequest) SetOrderInfo(_orderInfo []string) error {
	r._orderInfo = _orderInfo
	r.Set("order_info", _orderInfo)
	return nil
}

// GetOrderInfo OrderInfo Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetOrderInfo() []string {
	return r._orderInfo
}

// SetRefundRegulations is RefundRegulations Setter
// 特殊可选，当refund_type=1或7时，需要上传自定义退改内容。自定义退改规则，最多可含5组规则，每组规则间以英文逗号分隔。 1）当refund_type为1时格式为：a_b_num,b-1_c_num。含义：提前a天至提前b天发起退款，买家需支付num比例违约费。 2）当refund_type为7时格式为：a_b_num1_num2_0,b-1_c_num1_num2_0。含义：提前a天至提前b天发起退款，买家需支付num1比例违约费，卖家需支付num2比例违约费，最后一个数字代表是否节假日规则（0-不是，1-是）。特别注意，当refund_type为7时，自定义退改规则必须设置 n天以上违约规则 以及 行程当日违约规则，即第一组规则需要以-1_a_num1_num2_0或-1_a_num1_num2_1开头，且最后一组规则需要以0_0_num1_num2_0或0_0_num1_num2_1结尾。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetRefundRegulations(_refundRegulations []string) error {
	r._refundRegulations = _refundRegulations
	r.Set("refund_regulations", _refundRegulations)
	return nil
}

// GetRefundRegulations RefundRegulations Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetRefundRegulations() []string {
	return r._refundRegulations
}

// SetGatherPlaces is GatherPlaces Setter
// 可选，跟团时的集合地点，列表中每一个元素对应一个集合地点
func (r *AlitripGrouptoursProductUploadAPIRequest) SetGatherPlaces(_gatherPlaces []GatherPlaceInfo) error {
	r._gatherPlaces = _gatherPlaces
	r.Set("gather_places", _gatherPlaces)
	return nil
}

// GetGatherPlaces GatherPlaces Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetGatherPlaces() []GatherPlaceInfo {
	return r._gatherPlaces
}

// SetFeeInclude is FeeInclude Setter
// 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadAPIRequest) SetFeeInclude(_feeInclude []string) error {
	r._feeInclude = _feeInclude
	r.Set("fee_include", _feeInclude)
	return nil
}

// GetFeeInclude FeeInclude Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetFeeInclude() []string {
	return r._feeInclude
}

// SetStructItineraries is StructItineraries Setter
// 行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetStructItineraries(_structItineraries []StructItinerary) error {
	r._structItineraries = _structItineraries
	r.Set("struct_itineraries", _structItineraries)
	return nil
}

// GetStructItineraries StructItineraries Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetStructItineraries() []StructItinerary {
	return r._structItineraries
}

// SetPicUrls is PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadAPIRequest) SetPicUrls(_picUrls []string) error {
	r._picUrls = _picUrls
	r.Set("pic_urls", _picUrls)
	return nil
}

// GetPicUrls PicUrls Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetPicUrls() []string {
	return r._picUrls
}

// SetSubTitles is SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadAPIRequest) SetSubTitles(_subTitles []string) error {
	r._subTitles = _subTitles
	r.Set("sub_titles", _subTitles)
	return nil
}

// GetSubTitles SubTitles Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetSubTitles() []string {
	return r._subTitles
}

// SetWapDesc is WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetWapDesc(_wapDesc string) error {
	r._wapDesc = _wapDesc
	r.Set("wap_desc", _wapDesc)
	return nil
}

// GetWapDesc WapDesc Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetWapDesc() string {
	return r._wapDesc
}

// SetDescXml is DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetDescXml(_descXml string) error {
	r._descXml = _descXml
	r.Set("desc_xml", _descXml)
	return nil
}

// GetDescXml DescXml Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetDescXml() string {
	return r._descXml
}

// SetTitle is Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripGrouptoursProductUploadAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetTitle() string {
	return r._title
}

// SetItemCustomTag is ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripGrouptoursProductUploadAPIRequest) SetItemCustomTag(_itemCustomTag string) error {
	r._itemCustomTag = _itemCustomTag
	r.Set("item_custom_tag", _itemCustomTag)
	return nil
}

// GetItemCustomTag ItemCustomTag Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetItemCustomTag() string {
	return r._itemCustomTag
}

// SetOutProductId is OutProductId Setter
// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetToLocations is ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripGrouptoursProductUploadAPIRequest) SetToLocations(_toLocations string) error {
	r._toLocations = _toLocations
	r.Set("to_locations", _toLocations)
	return nil
}

// GetToLocations ToLocations Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetToLocations() string {
	return r._toLocations
}

// SetReserveLimit is ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripGrouptoursProductUploadAPIRequest) SetReserveLimit(_reserveLimit string) error {
	r._reserveLimit = _reserveLimit
	r.Set("reserve_limit", _reserveLimit)
	return nil
}

// GetReserveLimit ReserveLimit Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetReserveLimit() string {
	return r._reserveLimit
}

// SetFromLocations is FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripGrouptoursProductUploadAPIRequest) SetFromLocations(_fromLocations string) error {
	r._fromLocations = _fromLocations
	r.Set("from_locations", _fromLocations)
	return nil
}

// GetFromLocations FromLocations Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetFromLocations() string {
	return r._fromLocations
}

// SetDescHtml is DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetDescHtml(_descHtml string) error {
	r._descHtml = _descHtml
	r.Set("desc_html", _descHtml)
	return nil
}

// GetDescHtml DescHtml Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetDescHtml() string {
	return r._descHtml
}

// SetGoTrafficType is GoTrafficType Setter
// 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船
func (r *AlitripGrouptoursProductUploadAPIRequest) SetGoTrafficType(_goTrafficType int64) error {
	r._goTrafficType = _goTrafficType
	r.Set("go_traffic_type", _goTrafficType)
	return nil
}

// GetGoTrafficType GoTrafficType Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetGoTrafficType() int64 {
	return r._goTrafficType
}

// SetTripDay is TripDay Setter
// 新发布商品时必填。旅游天数
func (r *AlitripGrouptoursProductUploadAPIRequest) SetTripDay(_tripDay int64) error {
	r._tripDay = _tripDay
	r.Set("trip_day", _tripDay)
	return nil
}

// GetTripDay TripDay Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetTripDay() int64 {
	return r._tripDay
}

// SetSubStock is SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripGrouptoursProductUploadAPIRequest) SetSubStock(_subStock int64) error {
	r._subStock = _subStock
	r.Set("sub_stock", _subStock)
	return nil
}

// GetSubStock SubStock Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetSubStock() int64 {
	return r._subStock
}

// SetBackTrafficType is BackTrafficType Setter
// 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船
func (r *AlitripGrouptoursProductUploadAPIRequest) SetBackTrafficType(_backTrafficType int64) error {
	r._backTrafficType = _backTrafficType
	r.Set("back_traffic_type", _backTrafficType)
	return nil
}

// GetBackTrafficType BackTrafficType Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetBackTrafficType() int64 {
	return r._backTrafficType
}

// SetConfirmTime is ConfirmTime Setter
// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripGrouptoursProductUploadAPIRequest) SetConfirmTime(_confirmTime int64) error {
	r._confirmTime = _confirmTime
	r.Set("confirm_time", _confirmTime)
	return nil
}

// GetConfirmTime ConfirmTime Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetConfirmTime() int64 {
	return r._confirmTime
}

// SetRouteType is RouteType Setter
// 新发布商品时必填。参团线路类型。0 -目的地参团，1-为出发地参团
func (r *AlitripGrouptoursProductUploadAPIRequest) SetRouteType(_routeType int64) error {
	r._routeType = _routeType
	r.Set("route_type", _routeType)
	return nil
}

// GetRouteType RouteType Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetRouteType() int64 {
	return r._routeType
}

// SetConfirmType is ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripGrouptoursProductUploadAPIRequest) SetConfirmType(_confirmType int64) error {
	r._confirmType = _confirmType
	r.Set("confirm_type", _confirmType)
	return nil
}

// GetConfirmType ConfirmType Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetConfirmType() int64 {
	return r._confirmType
}

// SetTripNight is TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1
func (r *AlitripGrouptoursProductUploadAPIRequest) SetTripNight(_tripNight int64) error {
	r._tripNight = _tripNight
	r.Set("trip_night", _tripNight)
	return nil
}

// GetTripNight TripNight Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetTripNight() int64 {
	return r._tripNight
}

// SetItemId is ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetRefundType is RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
func (r *AlitripGrouptoursProductUploadAPIRequest) SetRefundType(_refundType int64) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetRefundType() int64 {
	return r._refundType
}

// SetElectronContract is ElectronContract Setter
// 可选，电子合同信息设置。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetElectronContract(_electronContract *ElectronContract) error {
	r._electronContract = _electronContract
	r.Set("electron_contract", _electronContract)
	return nil
}

// GetElectronContract ElectronContract Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetElectronContract() *ElectronContract {
	return r._electronContract
}

// SetTravellerTemplateId is TravellerTemplateId Setter
// 可选，出行人模板id，预留，暂不支持
func (r *AlitripGrouptoursProductUploadAPIRequest) SetTravellerTemplateId(_travellerTemplateId int64) error {
	r._travellerTemplateId = _travellerTemplateId
	r.Set("traveller_template_id", _travellerTemplateId)
	return nil
}

// GetTravellerTemplateId TravellerTemplateId Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetTravellerTemplateId() int64 {
	return r._travellerTemplateId
}

// SetPurePlay is PurePlay Setter
// 可选，是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
func (r *AlitripGrouptoursProductUploadAPIRequest) SetPurePlay(_purePlay int64) error {
	r._purePlay = _purePlay
	r.Set("pure_play", _purePlay)
	return nil
}

// GetPurePlay PurePlay Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetPurePlay() int64 {
	return r._purePlay
}

// SetIsOverseasTour is IsOverseasTour Setter
// 新发布商品时必填。是否出境游，0-不是，1-是。
func (r *AlitripGrouptoursProductUploadAPIRequest) SetIsOverseasTour(_isOverseasTour int64) error {
	r._isOverseasTour = _isOverseasTour
	r.Set("is_overseas_tour", _isOverseasTour)
	return nil
}

// GetIsOverseasTour IsOverseasTour Getter
func (r AlitripGrouptoursProductUploadAPIRequest) GetIsOverseasTour() int64 {
	return r._isOverseasTour
}

var poolAlitripGrouptoursProductUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripGrouptoursProductUploadRequest()
	},
}

// GetAlitripGrouptoursProductUploadRequest 从 sync.Pool 获取 AlitripGrouptoursProductUploadAPIRequest
func GetAlitripGrouptoursProductUploadAPIRequest() *AlitripGrouptoursProductUploadAPIRequest {
	return poolAlitripGrouptoursProductUploadAPIRequest.Get().(*AlitripGrouptoursProductUploadAPIRequest)
}

// ReleaseAlitripGrouptoursProductUploadAPIRequest 将 AlitripGrouptoursProductUploadAPIRequest 放入 sync.Pool
func ReleaseAlitripGrouptoursProductUploadAPIRequest(v *AlitripGrouptoursProductUploadAPIRequest) {
	v.Reset()
	poolAlitripGrouptoursProductUploadAPIRequest.Put(v)
}
