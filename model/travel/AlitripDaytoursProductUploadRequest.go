package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
境外一日游/多日游 产品维护接口 API请求
alitrip.daytours.product.upload

境外一日游/多日游 产品维护接口。
接口同时支持新商品发布 和 现有商品编辑：
1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
*/
type AlitripDaytoursProductUploadRequest struct {
    model.Params
    // 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
    _outProductId   string
    // 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
    _itemId   int64
    // 新发布商品时必填。商品标题，30个中文字符以内
    _title   string
    // 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
    _subTitles   []string
    // 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
    _picUrls   []string
    // 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
    _fromLocations   string
    // 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
    _toLocations   string
    // 新发布商品时必填。旅游天数
    _tripDay   int64
    // 可选，旅游晚数，不传默认旅游天数-1
    _tripNight   int64
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
    _descXml   string
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
    _descHtml   string
    // 可选，手机端详情描述，xml格式，格式详见示例。
    _wapDesc   string
    // 特殊必填，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
    _itineraries   []string
    // 特殊必填，行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
    _structItineraries   []StructItinerary
    // 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    _feeInclude   []string
    // 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    _feeExclude   []string
    // 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    _orderInfo   []string
    // 真实的旅游服务提供商
    _touristServiceProvider   string
    // 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
    _refundType   int64
    // 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
    _refundRegulationsJson   string
    // 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
    _reserveLimit   string
    // 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
    _confirmType   int64
    // 特殊可选（confirm_type为2时必填），资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
    _confirmTime   int64
    // 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
    _subStock   int64
    // 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
    _travellerTemplateId   int64
    // 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
    _itemCustomTag   string
    // 一日游 产品亮点
    _highLights   []ProductHighLights
    // 必填，营业执照图片路径。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在3M以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间。
    _businessLicense   string
    // 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
    _sellerCids   []string
    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    _secondKill   string
    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    _hasDiscount   bool
}

// 初始化AlitripDaytoursProductUploadRequest对象
func NewAlitripDaytoursProductUploadRequest() *AlitripDaytoursProductUploadRequest{
    return &AlitripDaytoursProductUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripDaytoursProductUploadRequest) GetApiMethodName() string {
    return "alitrip.daytours.product.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlitripDaytoursProductUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutProductId Setter
// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripDaytoursProductUploadRequest) SetOutProductId(_outProductId string) error {
    r._outProductId = _outProductId
    r.Set("out_product_id", _outProductId)
    return nil
}

// OutProductId Getter
func (r AlitripDaytoursProductUploadRequest) GetOutProductId() string {
    return r._outProductId
}
// ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripDaytoursProductUploadRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlitripDaytoursProductUploadRequest) GetItemId() int64 {
    return r._itemId
}
// Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripDaytoursProductUploadRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlitripDaytoursProductUploadRequest) GetTitle() string {
    return r._title
}
// SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadRequest) SetSubTitles(_subTitles []string) error {
    r._subTitles = _subTitles
    r.Set("sub_titles", _subTitles)
    return nil
}

// SubTitles Getter
func (r AlitripDaytoursProductUploadRequest) GetSubTitles() []string {
    return r._subTitles
}
// PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadRequest) SetPicUrls(_picUrls []string) error {
    r._picUrls = _picUrls
    r.Set("pic_urls", _picUrls)
    return nil
}

// PicUrls Getter
func (r AlitripDaytoursProductUploadRequest) GetPicUrls() []string {
    return r._picUrls
}
// FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripDaytoursProductUploadRequest) SetFromLocations(_fromLocations string) error {
    r._fromLocations = _fromLocations
    r.Set("from_locations", _fromLocations)
    return nil
}

// FromLocations Getter
func (r AlitripDaytoursProductUploadRequest) GetFromLocations() string {
    return r._fromLocations
}
// ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripDaytoursProductUploadRequest) SetToLocations(_toLocations string) error {
    r._toLocations = _toLocations
    r.Set("to_locations", _toLocations)
    return nil
}

// ToLocations Getter
func (r AlitripDaytoursProductUploadRequest) GetToLocations() string {
    return r._toLocations
}
// TripDay Setter
// 新发布商品时必填。旅游天数
func (r *AlitripDaytoursProductUploadRequest) SetTripDay(_tripDay int64) error {
    r._tripDay = _tripDay
    r.Set("trip_day", _tripDay)
    return nil
}

// TripDay Getter
func (r AlitripDaytoursProductUploadRequest) GetTripDay() int64 {
    return r._tripDay
}
// TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1
func (r *AlitripDaytoursProductUploadRequest) SetTripNight(_tripNight int64) error {
    r._tripNight = _tripNight
    r.Set("trip_night", _tripNight)
    return nil
}

// TripNight Getter
func (r AlitripDaytoursProductUploadRequest) GetTripNight() int64 {
    return r._tripNight
}
// DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripDaytoursProductUploadRequest) SetDescXml(_descXml string) error {
    r._descXml = _descXml
    r.Set("desc_xml", _descXml)
    return nil
}

// DescXml Getter
func (r AlitripDaytoursProductUploadRequest) GetDescXml() string {
    return r._descXml
}
// DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripDaytoursProductUploadRequest) SetDescHtml(_descHtml string) error {
    r._descHtml = _descHtml
    r.Set("desc_html", _descHtml)
    return nil
}

// DescHtml Getter
func (r AlitripDaytoursProductUploadRequest) GetDescHtml() string {
    return r._descHtml
}
// WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripDaytoursProductUploadRequest) SetWapDesc(_wapDesc string) error {
    r._wapDesc = _wapDesc
    r.Set("wap_desc", _wapDesc)
    return nil
}

// WapDesc Getter
func (r AlitripDaytoursProductUploadRequest) GetWapDesc() string {
    return r._wapDesc
}
// Itineraries Setter
// 特殊必填，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadRequest) SetItineraries(_itineraries []string) error {
    r._itineraries = _itineraries
    r.Set("itineraries", _itineraries)
    return nil
}

// Itineraries Getter
func (r AlitripDaytoursProductUploadRequest) GetItineraries() []string {
    return r._itineraries
}
// StructItineraries Setter
// 特殊必填，行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
func (r *AlitripDaytoursProductUploadRequest) SetStructItineraries(_structItineraries []StructItinerary) error {
    r._structItineraries = _structItineraries
    r.Set("struct_itineraries", _structItineraries)
    return nil
}

// StructItineraries Getter
func (r AlitripDaytoursProductUploadRequest) GetStructItineraries() []StructItinerary {
    return r._structItineraries
}
// FeeInclude Setter
// 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadRequest) SetFeeInclude(_feeInclude []string) error {
    r._feeInclude = _feeInclude
    r.Set("fee_include", _feeInclude)
    return nil
}

// FeeInclude Getter
func (r AlitripDaytoursProductUploadRequest) GetFeeInclude() []string {
    return r._feeInclude
}
// FeeExclude Setter
// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadRequest) SetFeeExclude(_feeExclude []string) error {
    r._feeExclude = _feeExclude
    r.Set("fee_exclude", _feeExclude)
    return nil
}

// FeeExclude Getter
func (r AlitripDaytoursProductUploadRequest) GetFeeExclude() []string {
    return r._feeExclude
}
// OrderInfo Setter
// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，每点描述小于100个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripDaytoursProductUploadRequest) SetOrderInfo(_orderInfo []string) error {
    r._orderInfo = _orderInfo
    r.Set("order_info", _orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlitripDaytoursProductUploadRequest) GetOrderInfo() []string {
    return r._orderInfo
}
// TouristServiceProvider Setter
// 真实的旅游服务提供商
func (r *AlitripDaytoursProductUploadRequest) SetTouristServiceProvider(_touristServiceProvider string) error {
    r._touristServiceProvider = _touristServiceProvider
    r.Set("tourist_service_provider", _touristServiceProvider)
    return nil
}

// TouristServiceProvider Getter
func (r AlitripDaytoursProductUploadRequest) GetTouristServiceProvider() string {
    return r._touristServiceProvider
}
// RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
func (r *AlitripDaytoursProductUploadRequest) SetRefundType(_refundType int64) error {
    r._refundType = _refundType
    r.Set("refund_type", _refundType)
    return nil
}

// RefundType Getter
func (r AlitripDaytoursProductUploadRequest) GetRefundType() int64 {
    return r._refundType
}
// RefundRegulationsJson Setter
// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
func (r *AlitripDaytoursProductUploadRequest) SetRefundRegulationsJson(_refundRegulationsJson string) error {
    r._refundRegulationsJson = _refundRegulationsJson
    r.Set("refund_regulations_json", _refundRegulationsJson)
    return nil
}

// RefundRegulationsJson Getter
func (r AlitripDaytoursProductUploadRequest) GetRefundRegulationsJson() string {
    return r._refundRegulationsJson
}
// ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripDaytoursProductUploadRequest) SetReserveLimit(_reserveLimit string) error {
    r._reserveLimit = _reserveLimit
    r.Set("reserve_limit", _reserveLimit)
    return nil
}

// ReserveLimit Getter
func (r AlitripDaytoursProductUploadRequest) GetReserveLimit() string {
    return r._reserveLimit
}
// ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripDaytoursProductUploadRequest) SetConfirmType(_confirmType int64) error {
    r._confirmType = _confirmType
    r.Set("confirm_type", _confirmType)
    return nil
}

// ConfirmType Getter
func (r AlitripDaytoursProductUploadRequest) GetConfirmType() int64 {
    return r._confirmType
}
// ConfirmTime Setter
// 特殊可选（confirm_type为2时必填），资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripDaytoursProductUploadRequest) SetConfirmTime(_confirmTime int64) error {
    r._confirmTime = _confirmTime
    r.Set("confirm_time", _confirmTime)
    return nil
}

// ConfirmTime Getter
func (r AlitripDaytoursProductUploadRequest) GetConfirmTime() int64 {
    return r._confirmTime
}
// SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripDaytoursProductUploadRequest) SetSubStock(_subStock int64) error {
    r._subStock = _subStock
    r.Set("sub_stock", _subStock)
    return nil
}

// SubStock Getter
func (r AlitripDaytoursProductUploadRequest) GetSubStock() int64 {
    return r._subStock
}
// TravellerTemplateId Setter
// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
func (r *AlitripDaytoursProductUploadRequest) SetTravellerTemplateId(_travellerTemplateId int64) error {
    r._travellerTemplateId = _travellerTemplateId
    r.Set("traveller_template_id", _travellerTemplateId)
    return nil
}

// TravellerTemplateId Getter
func (r AlitripDaytoursProductUploadRequest) GetTravellerTemplateId() int64 {
    return r._travellerTemplateId
}
// ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripDaytoursProductUploadRequest) SetItemCustomTag(_itemCustomTag string) error {
    r._itemCustomTag = _itemCustomTag
    r.Set("item_custom_tag", _itemCustomTag)
    return nil
}

// ItemCustomTag Getter
func (r AlitripDaytoursProductUploadRequest) GetItemCustomTag() string {
    return r._itemCustomTag
}
// HighLights Setter
// 一日游 产品亮点
func (r *AlitripDaytoursProductUploadRequest) SetHighLights(_highLights []ProductHighLights) error {
    r._highLights = _highLights
    r.Set("high_lights", _highLights)
    return nil
}

// HighLights Getter
func (r AlitripDaytoursProductUploadRequest) GetHighLights() []ProductHighLights {
    return r._highLights
}
// BusinessLicense Setter
// 必填，营业执照图片路径。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在3M以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间。
func (r *AlitripDaytoursProductUploadRequest) SetBusinessLicense(_businessLicense string) error {
    r._businessLicense = _businessLicense
    r.Set("business_license", _businessLicense)
    return nil
}

// BusinessLicense Getter
func (r AlitripDaytoursProductUploadRequest) GetBusinessLicense() string {
    return r._businessLicense
}
// SellerCids Setter
// 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
func (r *AlitripDaytoursProductUploadRequest) SetSellerCids(_sellerCids []string) error {
    r._sellerCids = _sellerCids
    r.Set("seller_cids", _sellerCids)
    return nil
}

// SellerCids Getter
func (r AlitripDaytoursProductUploadRequest) GetSellerCids() []string {
    return r._sellerCids
}
// SecondKill Setter
// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
func (r *AlitripDaytoursProductUploadRequest) SetSecondKill(_secondKill string) error {
    r._secondKill = _secondKill
    r.Set("second_kill", _secondKill)
    return nil
}

// SecondKill Getter
func (r AlitripDaytoursProductUploadRequest) GetSecondKill() string {
    return r._secondKill
}
// HasDiscount Setter
// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
func (r *AlitripDaytoursProductUploadRequest) SetHasDiscount(_hasDiscount bool) error {
    r._hasDiscount = _hasDiscount
    r.Set("has_discount", _hasDiscount)
    return nil
}

// HasDiscount Getter
func (r AlitripDaytoursProductUploadRequest) GetHasDiscount() bool {
    return r._hasDiscount
}
