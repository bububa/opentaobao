package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
跟团游 产品维护接口 API请求
alitrip.grouptours.product.upload

跟团游 产品维护接口。
接口同时支持新商品发布 和 现有商品编辑：
1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
*/
type AlitripGrouptoursProductUploadRequest struct {
    model.Params
    // 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
    outProductId   string
    // 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
    itemId   int64
    // 新发布商品时必填。商品标题，30个中文字符以内
    title   string
    // 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
    subTitles   []string
    // 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
    picUrls   []string
    // 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
    fromLocations   string
    // 可选，跟团时的集合地点，列表中每一个元素对应一个集合地点
    gatherPlaces   []GatherPlaceInfo
    // 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
    toLocations   string
    // 新发布商品时必填。是否出境游，0-不是，1-是。
    isOverseasTour   int64
    // 新发布商品时必填。旅游天数
    tripDay   int64
    // 可选，旅游晚数，不传默认旅游天数-1
    tripNight   int64
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
    descXml   string
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
    descHtml   string
    // 可选，手机端详情描述，xml格式，格式详见示例。
    wapDesc   string
    // 可选，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
    itineraries   []string
    // 行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
    structItineraries   []StructItinerary
    // 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    feeInclude   []string
    // 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    feeExclude   []string
    // 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    orderInfo   []string
    // 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
    refundType   int64
    // 特殊可选，当refund_type=1或7时，需要上传自定义退改内容。自定义退改规则，最多可含5组规则，每组规则间以英文逗号分隔。 1）当refund_type为1时格式为：a_b_num,b-1_c_num。含义：提前a天至提前b天发起退款，买家需支付num比例违约费。 2）当refund_type为7时格式为：a_b_num1_num2_0,b-1_c_num1_num2_0。含义：提前a天至提前b天发起退款，买家需支付num1比例违约费，卖家需支付num2比例违约费，最后一个数字代表是否节假日规则（0-不是，1-是）。特别注意，当refund_type为7时，自定义退改规则必须设置 n天以上违约规则 以及 行程当日违约规则，即第一组规则需要以-1_a_num1_num2_0或-1_a_num1_num2_1开头，且最后一组规则需要以0_0_num1_num2_0或0_0_num1_num2_1结尾。
    refundRegulations   []string
    // 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船
    goTrafficType   int64
    // 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船
    backTrafficType   int64
    // 新发布商品时必填。参团线路类型。0 -目的地参团，1-为出发地参团
    routeType   int64
    // 可选，是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
    purePlay   int64
    // 可选，电子合同信息设置。
    electronContract   *ElectronContract
    // 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
    reserveLimit   string
    // 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
    confirmType   int64
    // 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
    confirmTime   int64
    // 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
    subStock   int64
    // 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
    itemCustomTag   string
    // 可选，出行人模板id，预留，暂不支持
    travellerTemplateId   int64
}

// 初始化AlitripGrouptoursProductUploadRequest对象
func NewAlitripGrouptoursProductUploadRequest() *AlitripGrouptoursProductUploadRequest{
    return &AlitripGrouptoursProductUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripGrouptoursProductUploadRequest) GetApiMethodName() string {
    return "alitrip.grouptours.product.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlitripGrouptoursProductUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutProductId Setter
// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripGrouptoursProductUploadRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r AlitripGrouptoursProductUploadRequest) GetOutProductId() string {
    return r.outProductId
}
// ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripGrouptoursProductUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlitripGrouptoursProductUploadRequest) GetItemId() int64 {
    return r.itemId
}
// Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripGrouptoursProductUploadRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlitripGrouptoursProductUploadRequest) GetTitle() string {
    return r.title
}
// SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadRequest) SetSubTitles(subTitles []string) error {
    r.subTitles = subTitles
    r.Set("sub_titles", subTitles)
    return nil
}

// SubTitles Getter
func (r AlitripGrouptoursProductUploadRequest) GetSubTitles() []string {
    return r.subTitles
}
// PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadRequest) SetPicUrls(picUrls []string) error {
    r.picUrls = picUrls
    r.Set("pic_urls", picUrls)
    return nil
}

// PicUrls Getter
func (r AlitripGrouptoursProductUploadRequest) GetPicUrls() []string {
    return r.picUrls
}
// FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripGrouptoursProductUploadRequest) SetFromLocations(fromLocations string) error {
    r.fromLocations = fromLocations
    r.Set("from_locations", fromLocations)
    return nil
}

// FromLocations Getter
func (r AlitripGrouptoursProductUploadRequest) GetFromLocations() string {
    return r.fromLocations
}
// GatherPlaces Setter
// 可选，跟团时的集合地点，列表中每一个元素对应一个集合地点
func (r *AlitripGrouptoursProductUploadRequest) SetGatherPlaces(gatherPlaces []GatherPlaceInfo) error {
    r.gatherPlaces = gatherPlaces
    r.Set("gather_places", gatherPlaces)
    return nil
}

// GatherPlaces Getter
func (r AlitripGrouptoursProductUploadRequest) GetGatherPlaces() []GatherPlaceInfo {
    return r.gatherPlaces
}
// ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripGrouptoursProductUploadRequest) SetToLocations(toLocations string) error {
    r.toLocations = toLocations
    r.Set("to_locations", toLocations)
    return nil
}

// ToLocations Getter
func (r AlitripGrouptoursProductUploadRequest) GetToLocations() string {
    return r.toLocations
}
// IsOverseasTour Setter
// 新发布商品时必填。是否出境游，0-不是，1-是。
func (r *AlitripGrouptoursProductUploadRequest) SetIsOverseasTour(isOverseasTour int64) error {
    r.isOverseasTour = isOverseasTour
    r.Set("is_overseas_tour", isOverseasTour)
    return nil
}

// IsOverseasTour Getter
func (r AlitripGrouptoursProductUploadRequest) GetIsOverseasTour() int64 {
    return r.isOverseasTour
}
// TripDay Setter
// 新发布商品时必填。旅游天数
func (r *AlitripGrouptoursProductUploadRequest) SetTripDay(tripDay int64) error {
    r.tripDay = tripDay
    r.Set("trip_day", tripDay)
    return nil
}

// TripDay Getter
func (r AlitripGrouptoursProductUploadRequest) GetTripDay() int64 {
    return r.tripDay
}
// TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1
func (r *AlitripGrouptoursProductUploadRequest) SetTripNight(tripNight int64) error {
    r.tripNight = tripNight
    r.Set("trip_night", tripNight)
    return nil
}

// TripNight Getter
func (r AlitripGrouptoursProductUploadRequest) GetTripNight() int64 {
    return r.tripNight
}
// DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripGrouptoursProductUploadRequest) SetDescXml(descXml string) error {
    r.descXml = descXml
    r.Set("desc_xml", descXml)
    return nil
}

// DescXml Getter
func (r AlitripGrouptoursProductUploadRequest) GetDescXml() string {
    return r.descXml
}
// DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripGrouptoursProductUploadRequest) SetDescHtml(descHtml string) error {
    r.descHtml = descHtml
    r.Set("desc_html", descHtml)
    return nil
}

// DescHtml Getter
func (r AlitripGrouptoursProductUploadRequest) GetDescHtml() string {
    return r.descHtml
}
// WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripGrouptoursProductUploadRequest) SetWapDesc(wapDesc string) error {
    r.wapDesc = wapDesc
    r.Set("wap_desc", wapDesc)
    return nil
}

// WapDesc Getter
func (r AlitripGrouptoursProductUploadRequest) GetWapDesc() string {
    return r.wapDesc
}
// Itineraries Setter
// 可选，（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。itineraries数组的元素个数必须与旅游天数trip_day一致。行程描述，每一天行程都是xml格式（数组的每个元素对应每一天的行程）： ITINERARY根标签必须有，每个活动（ACTIVITY）支持1个title子标签，1个txt子标签和多个img子标签。注：在SDK中多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadRequest) SetItineraries(itineraries []string) error {
    r.itineraries = itineraries
    r.Set("itineraries", itineraries)
    return nil
}

// Itineraries Getter
func (r AlitripGrouptoursProductUploadRequest) GetItineraries() []string {
    return r.itineraries
}
// StructItineraries Setter
// 行程描述（struct_itineraries与itineraries二者选填一个即可，如果两个都上传则以struct_itineraries为准）。列表中每一个元素对应一天的行程，每天行程由多个活动组成。行程描述是可选项，如果上传了行程内容 则行程序号和行程活动信息必填。
func (r *AlitripGrouptoursProductUploadRequest) SetStructItineraries(structItineraries []StructItinerary) error {
    r.structItineraries = structItineraries
    r.Set("struct_itineraries", structItineraries)
    return nil
}

// StructItineraries Getter
func (r AlitripGrouptoursProductUploadRequest) GetStructItineraries() []StructItinerary {
    return r.structItineraries
}
// FeeInclude Setter
// 新发布商品时必填。费用包含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadRequest) SetFeeInclude(feeInclude []string) error {
    r.feeInclude = feeInclude
    r.Set("fee_include", feeInclude)
    return nil
}

// FeeInclude Getter
func (r AlitripGrouptoursProductUploadRequest) GetFeeInclude() []string {
    return r.feeInclude
}
// FeeExclude Setter
// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadRequest) SetFeeExclude(feeExclude []string) error {
    r.feeExclude = feeExclude
    r.Set("fee_exclude", feeExclude)
    return nil
}

// FeeExclude Getter
func (r AlitripGrouptoursProductUploadRequest) GetFeeExclude() []string {
    return r.feeExclude
}
// OrderInfo Setter
// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptoursProductUploadRequest) SetOrderInfo(orderInfo []string) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlitripGrouptoursProductUploadRequest) GetOrderInfo() []string {
    return r.orderInfo
}
// RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
func (r *AlitripGrouptoursProductUploadRequest) SetRefundType(refundType int64) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

// RefundType Getter
func (r AlitripGrouptoursProductUploadRequest) GetRefundType() int64 {
    return r.refundType
}
// RefundRegulations Setter
// 特殊可选，当refund_type=1或7时，需要上传自定义退改内容。自定义退改规则，最多可含5组规则，每组规则间以英文逗号分隔。 1）当refund_type为1时格式为：a_b_num,b-1_c_num。含义：提前a天至提前b天发起退款，买家需支付num比例违约费。 2）当refund_type为7时格式为：a_b_num1_num2_0,b-1_c_num1_num2_0。含义：提前a天至提前b天发起退款，买家需支付num1比例违约费，卖家需支付num2比例违约费，最后一个数字代表是否节假日规则（0-不是，1-是）。特别注意，当refund_type为7时，自定义退改规则必须设置 n天以上违约规则 以及 行程当日违约规则，即第一组规则需要以-1_a_num1_num2_0或-1_a_num1_num2_1开头，且最后一组规则需要以0_0_num1_num2_0或0_0_num1_num2_1结尾。
func (r *AlitripGrouptoursProductUploadRequest) SetRefundRegulations(refundRegulations []string) error {
    r.refundRegulations = refundRegulations
    r.Set("refund_regulations", refundRegulations)
    return nil
}

// RefundRegulations Getter
func (r AlitripGrouptoursProductUploadRequest) GetRefundRegulations() []string {
    return r.refundRegulations
}
// GoTrafficType Setter
// 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船
func (r *AlitripGrouptoursProductUploadRequest) SetGoTrafficType(goTrafficType int64) error {
    r.goTrafficType = goTrafficType
    r.Set("go_traffic_type", goTrafficType)
    return nil
}

// GoTrafficType Getter
func (r AlitripGrouptoursProductUploadRequest) GetGoTrafficType() int64 {
    return r.goTrafficType
}
// BackTrafficType Setter
// 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船
func (r *AlitripGrouptoursProductUploadRequest) SetBackTrafficType(backTrafficType int64) error {
    r.backTrafficType = backTrafficType
    r.Set("back_traffic_type", backTrafficType)
    return nil
}

// BackTrafficType Getter
func (r AlitripGrouptoursProductUploadRequest) GetBackTrafficType() int64 {
    return r.backTrafficType
}
// RouteType Setter
// 新发布商品时必填。参团线路类型。0 -目的地参团，1-为出发地参团
func (r *AlitripGrouptoursProductUploadRequest) SetRouteType(routeType int64) error {
    r.routeType = routeType
    r.Set("route_type", routeType)
    return nil
}

// RouteType Getter
func (r AlitripGrouptoursProductUploadRequest) GetRouteType() int64 {
    return r.routeType
}
// PurePlay Setter
// 可选，是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
func (r *AlitripGrouptoursProductUploadRequest) SetPurePlay(purePlay int64) error {
    r.purePlay = purePlay
    r.Set("pure_play", purePlay)
    return nil
}

// PurePlay Getter
func (r AlitripGrouptoursProductUploadRequest) GetPurePlay() int64 {
    return r.purePlay
}
// ElectronContract Setter
// 可选，电子合同信息设置。
func (r *AlitripGrouptoursProductUploadRequest) SetElectronContract(electronContract *ElectronContract) error {
    r.electronContract = electronContract
    r.Set("electron_contract", electronContract)
    return nil
}

// ElectronContract Getter
func (r AlitripGrouptoursProductUploadRequest) GetElectronContract() *ElectronContract {
    return r.electronContract
}
// ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripGrouptoursProductUploadRequest) SetReserveLimit(reserveLimit string) error {
    r.reserveLimit = reserveLimit
    r.Set("reserve_limit", reserveLimit)
    return nil
}

// ReserveLimit Getter
func (r AlitripGrouptoursProductUploadRequest) GetReserveLimit() string {
    return r.reserveLimit
}
// ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripGrouptoursProductUploadRequest) SetConfirmType(confirmType int64) error {
    r.confirmType = confirmType
    r.Set("confirm_type", confirmType)
    return nil
}

// ConfirmType Getter
func (r AlitripGrouptoursProductUploadRequest) GetConfirmType() int64 {
    return r.confirmType
}
// ConfirmTime Setter
// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripGrouptoursProductUploadRequest) SetConfirmTime(confirmTime int64) error {
    r.confirmTime = confirmTime
    r.Set("confirm_time", confirmTime)
    return nil
}

// ConfirmTime Getter
func (r AlitripGrouptoursProductUploadRequest) GetConfirmTime() int64 {
    return r.confirmTime
}
// SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripGrouptoursProductUploadRequest) SetSubStock(subStock int64) error {
    r.subStock = subStock
    r.Set("sub_stock", subStock)
    return nil
}

// SubStock Getter
func (r AlitripGrouptoursProductUploadRequest) GetSubStock() int64 {
    return r.subStock
}
// ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripGrouptoursProductUploadRequest) SetItemCustomTag(itemCustomTag string) error {
    r.itemCustomTag = itemCustomTag
    r.Set("item_custom_tag", itemCustomTag)
    return nil
}

// ItemCustomTag Getter
func (r AlitripGrouptoursProductUploadRequest) GetItemCustomTag() string {
    return r.itemCustomTag
}
// TravellerTemplateId Setter
// 可选，出行人模板id，预留，暂不支持
func (r *AlitripGrouptoursProductUploadRequest) SetTravellerTemplateId(travellerTemplateId int64) error {
    r.travellerTemplateId = travellerTemplateId
    r.Set("traveller_template_id", travellerTemplateId)
    return nil
}

// TravellerTemplateId Getter
func (r AlitripGrouptoursProductUploadRequest) GetTravellerTemplateId() int64 {
    return r.travellerTemplateId
}
