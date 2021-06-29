package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
当地玩乐 产品维护接口 API请求
alitrip.localplay.product.upload

当地玩乐（境内当地玩乐/境外玩乐套餐） 产品维护接口。
接口同时支持新商品发布 和 现有商品编辑：
1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
*/
type AlitripLocalplayProductUploadRequest struct {
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
    // 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    feeExclude   []string
    // 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
    orderInfo   []string
    // 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-线路商品（跟团、自由行等）新版自定义退改规则。不传默认为0
    refundType   int64
    // 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
    refundRegulationsJson   string
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
    // 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
    travellerTemplateId   int64
    // 代订服务说明（请填写真实的旅游服务提供商）
    touristServiceProvider   string
    // 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
    sellerCids   []string
    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    secondKill   string
    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    hasDiscount   bool
}

// 初始化AlitripLocalplayProductUploadRequest对象
func NewAlitripLocalplayProductUploadRequest() *AlitripLocalplayProductUploadRequest{
    return &AlitripLocalplayProductUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripLocalplayProductUploadRequest) GetApiMethodName() string {
    return "alitrip.localplay.product.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlitripLocalplayProductUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutProductId Setter
// 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripLocalplayProductUploadRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r AlitripLocalplayProductUploadRequest) GetOutProductId() string {
    return r.outProductId
}
// ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripLocalplayProductUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlitripLocalplayProductUploadRequest) GetItemId() int64 {
    return r.itemId
}
// Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripLocalplayProductUploadRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlitripLocalplayProductUploadRequest) GetTitle() string {
    return r.title
}
// SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadRequest) SetSubTitles(subTitles []string) error {
    r.subTitles = subTitles
    r.Set("sub_titles", subTitles)
    return nil
}

// SubTitles Getter
func (r AlitripLocalplayProductUploadRequest) GetSubTitles() []string {
    return r.subTitles
}
// PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadRequest) SetPicUrls(picUrls []string) error {
    r.picUrls = picUrls
    r.Set("pic_urls", picUrls)
    return nil
}

// PicUrls Getter
func (r AlitripLocalplayProductUploadRequest) GetPicUrls() []string {
    return r.picUrls
}
// FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripLocalplayProductUploadRequest) SetFromLocations(fromLocations string) error {
    r.fromLocations = fromLocations
    r.Set("from_locations", fromLocations)
    return nil
}

// FromLocations Getter
func (r AlitripLocalplayProductUploadRequest) GetFromLocations() string {
    return r.fromLocations
}
// ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripLocalplayProductUploadRequest) SetToLocations(toLocations string) error {
    r.toLocations = toLocations
    r.Set("to_locations", toLocations)
    return nil
}

// ToLocations Getter
func (r AlitripLocalplayProductUploadRequest) GetToLocations() string {
    return r.toLocations
}
// IsOverseasTour Setter
// 新发布商品时必填。是否出境游，0-不是，1-是。
func (r *AlitripLocalplayProductUploadRequest) SetIsOverseasTour(isOverseasTour int64) error {
    r.isOverseasTour = isOverseasTour
    r.Set("is_overseas_tour", isOverseasTour)
    return nil
}

// IsOverseasTour Getter
func (r AlitripLocalplayProductUploadRequest) GetIsOverseasTour() int64 {
    return r.isOverseasTour
}
// TripDay Setter
// 新发布商品时必填。旅游天数
func (r *AlitripLocalplayProductUploadRequest) SetTripDay(tripDay int64) error {
    r.tripDay = tripDay
    r.Set("trip_day", tripDay)
    return nil
}

// TripDay Getter
func (r AlitripLocalplayProductUploadRequest) GetTripDay() int64 {
    return r.tripDay
}
// TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1
func (r *AlitripLocalplayProductUploadRequest) SetTripNight(tripNight int64) error {
    r.tripNight = tripNight
    r.Set("trip_night", tripNight)
    return nil
}

// TripNight Getter
func (r AlitripLocalplayProductUploadRequest) GetTripNight() int64 {
    return r.tripNight
}
// DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripLocalplayProductUploadRequest) SetDescXml(descXml string) error {
    r.descXml = descXml
    r.Set("desc_xml", descXml)
    return nil
}

// DescXml Getter
func (r AlitripLocalplayProductUploadRequest) GetDescXml() string {
    return r.descXml
}
// DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripLocalplayProductUploadRequest) SetDescHtml(descHtml string) error {
    r.descHtml = descHtml
    r.Set("desc_html", descHtml)
    return nil
}

// DescHtml Getter
func (r AlitripLocalplayProductUploadRequest) GetDescHtml() string {
    return r.descHtml
}
// WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripLocalplayProductUploadRequest) SetWapDesc(wapDesc string) error {
    r.wapDesc = wapDesc
    r.Set("wap_desc", wapDesc)
    return nil
}

// WapDesc Getter
func (r AlitripLocalplayProductUploadRequest) GetWapDesc() string {
    return r.wapDesc
}
// FeeExclude Setter
// 新发布商品时必填。费用不含。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadRequest) SetFeeExclude(feeExclude []string) error {
    r.feeExclude = feeExclude
    r.Set("fee_exclude", feeExclude)
    return nil
}

// FeeExclude Getter
func (r AlitripLocalplayProductUploadRequest) GetFeeExclude() []string {
    return r.feeExclude
}
// OrderInfo Setter
// 新发布商品时必填。预定须知。列表中每一个元素 对应一点描述，所有描述合起来必须小于1500个中文字符。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripLocalplayProductUploadRequest) SetOrderInfo(orderInfo []string) error {
    r.orderInfo = orderInfo
    r.Set("order_info", orderInfo)
    return nil
}

// OrderInfo Getter
func (r AlitripLocalplayProductUploadRequest) GetOrderInfo() []string {
    return r.orderInfo
}
// RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-线路商品（跟团、自由行等）新版自定义退改规则。不传默认为0
func (r *AlitripLocalplayProductUploadRequest) SetRefundType(refundType int64) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

// RefundType Getter
func (r AlitripLocalplayProductUploadRequest) GetRefundType() int64 {
    return r.refundType
}
// RefundRegulationsJson Setter
// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
func (r *AlitripLocalplayProductUploadRequest) SetRefundRegulationsJson(refundRegulationsJson string) error {
    r.refundRegulationsJson = refundRegulationsJson
    r.Set("refund_regulations_json", refundRegulationsJson)
    return nil
}

// RefundRegulationsJson Getter
func (r AlitripLocalplayProductUploadRequest) GetRefundRegulationsJson() string {
    return r.refundRegulationsJson
}
// ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripLocalplayProductUploadRequest) SetReserveLimit(reserveLimit string) error {
    r.reserveLimit = reserveLimit
    r.Set("reserve_limit", reserveLimit)
    return nil
}

// ReserveLimit Getter
func (r AlitripLocalplayProductUploadRequest) GetReserveLimit() string {
    return r.reserveLimit
}
// ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripLocalplayProductUploadRequest) SetConfirmType(confirmType int64) error {
    r.confirmType = confirmType
    r.Set("confirm_type", confirmType)
    return nil
}

// ConfirmType Getter
func (r AlitripLocalplayProductUploadRequest) GetConfirmType() int64 {
    return r.confirmType
}
// ConfirmTime Setter
// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripLocalplayProductUploadRequest) SetConfirmTime(confirmTime int64) error {
    r.confirmTime = confirmTime
    r.Set("confirm_time", confirmTime)
    return nil
}

// ConfirmTime Getter
func (r AlitripLocalplayProductUploadRequest) GetConfirmTime() int64 {
    return r.confirmTime
}
// SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripLocalplayProductUploadRequest) SetSubStock(subStock int64) error {
    r.subStock = subStock
    r.Set("sub_stock", subStock)
    return nil
}

// SubStock Getter
func (r AlitripLocalplayProductUploadRequest) GetSubStock() int64 {
    return r.subStock
}
// ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripLocalplayProductUploadRequest) SetItemCustomTag(itemCustomTag string) error {
    r.itemCustomTag = itemCustomTag
    r.Set("item_custom_tag", itemCustomTag)
    return nil
}

// ItemCustomTag Getter
func (r AlitripLocalplayProductUploadRequest) GetItemCustomTag() string {
    return r.itemCustomTag
}
// TravellerTemplateId Setter
// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
func (r *AlitripLocalplayProductUploadRequest) SetTravellerTemplateId(travellerTemplateId int64) error {
    r.travellerTemplateId = travellerTemplateId
    r.Set("traveller_template_id", travellerTemplateId)
    return nil
}

// TravellerTemplateId Getter
func (r AlitripLocalplayProductUploadRequest) GetTravellerTemplateId() int64 {
    return r.travellerTemplateId
}
// TouristServiceProvider Setter
// 代订服务说明（请填写真实的旅游服务提供商）
func (r *AlitripLocalplayProductUploadRequest) SetTouristServiceProvider(touristServiceProvider string) error {
    r.touristServiceProvider = touristServiceProvider
    r.Set("tourist_service_provider", touristServiceProvider)
    return nil
}

// TouristServiceProvider Getter
func (r AlitripLocalplayProductUploadRequest) GetTouristServiceProvider() string {
    return r.touristServiceProvider
}
// SellerCids Setter
// 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
func (r *AlitripLocalplayProductUploadRequest) SetSellerCids(sellerCids []string) error {
    r.sellerCids = sellerCids
    r.Set("seller_cids", sellerCids)
    return nil
}

// SellerCids Getter
func (r AlitripLocalplayProductUploadRequest) GetSellerCids() []string {
    return r.sellerCids
}
// SecondKill Setter
// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
func (r *AlitripLocalplayProductUploadRequest) SetSecondKill(secondKill string) error {
    r.secondKill = secondKill
    r.Set("second_kill", secondKill)
    return nil
}

// SecondKill Getter
func (r AlitripLocalplayProductUploadRequest) GetSecondKill() string {
    return r.secondKill
}
// HasDiscount Setter
// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
func (r *AlitripLocalplayProductUploadRequest) SetHasDiscount(hasDiscount bool) error {
    r.hasDiscount = hasDiscount
    r.Set("has_discount", hasDiscount)
    return nil
}

// HasDiscount Getter
func (r AlitripLocalplayProductUploadRequest) GetHasDiscount() bool {
    return r.hasDiscount
}
