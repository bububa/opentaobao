package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新版跟团游商品维护接口 API请求
alitrip.grouptour.product.upload

新版跟团游商品维护接口
*/
type AlitripGrouptourProductUploadRequest struct {
    model.Params
    // 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船，100-其他
    goTrafficType   int64
    // 新发布商品时必填。旅游天数。已废弃，以套餐维度行程天数为准。
    tripDay   int64
    // 可选，手机端详情描述，xml格式，格式详见示例。
    wapDesc   string
    // 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
    subStock   int64
    // 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他
    backTrafficType   int64
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
    descXml   string
    // 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
    confirmTime   int64
    // 新发布商品时必填。商品标题，30个中文字符以内
    title   string
    // 参团线路类型。0 -目的地参团，1-为出发地参团
    routeType   int64
    // 套餐信息,数组类型，支持上传多个套餐信息
    groupTourPackageInfo   []GroupTourPackageInfo
    // 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
    confirmType   int64
    // 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
    itemCustomTag   string
    // 必填，商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
    outProductId   string
    // 可选，旅游晚数，不传默认旅游天数-1。已废弃，以套餐维度行程晚数为准。
    tripNight   int64
    // 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
    toLocations   string
    // 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
    picUrls   []string
    // 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
    reserveLimit   string
    // 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
    itemId   int64
    // 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
    refundType   int64
    // 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
    subTitles   []string
    // 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
    fromLocations   string
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
    descHtml   string
    // 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
    travellerTemplateId   int64
    // 新发布商品时必填。是否出境游，0-不是，1-是。
    isOverseasTour   int64
    // 是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
    purePlay   int64
    // 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
    refundRegulationsJson   string
    // 0：使用上传的套餐信息（group_tour_package_info）覆盖商品上原有的套餐信息（此时group_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（group_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0===
    packageOperation   int64
    // 必填，线路的“细分类型”属性：1-普通跟团游、2-半自由行、3-私家团；不填默认值设置为"1-普通跟团游"。
    groupTourType   int64
    // 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
    sellerCids   []string
    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    secondKill   string
    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    hasDiscount   bool
}

// 初始化AlitripGrouptourProductUploadRequest对象
func NewAlitripGrouptourProductUploadRequest() *AlitripGrouptourProductUploadRequest{
    return &AlitripGrouptourProductUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripGrouptourProductUploadRequest) GetApiMethodName() string {
    return "alitrip.grouptour.product.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlitripGrouptourProductUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoTrafficType Setter
// 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船，100-其他
func (r *AlitripGrouptourProductUploadRequest) SetGoTrafficType(goTrafficType int64) error {
    r.goTrafficType = goTrafficType
    r.Set("go_traffic_type", goTrafficType)
    return nil
}

// GoTrafficType Getter
func (r AlitripGrouptourProductUploadRequest) GetGoTrafficType() int64 {
    return r.goTrafficType
}
// TripDay Setter
// 新发布商品时必填。旅游天数。已废弃，以套餐维度行程天数为准。
func (r *AlitripGrouptourProductUploadRequest) SetTripDay(tripDay int64) error {
    r.tripDay = tripDay
    r.Set("trip_day", tripDay)
    return nil
}

// TripDay Getter
func (r AlitripGrouptourProductUploadRequest) GetTripDay() int64 {
    return r.tripDay
}
// WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripGrouptourProductUploadRequest) SetWapDesc(wapDesc string) error {
    r.wapDesc = wapDesc
    r.Set("wap_desc", wapDesc)
    return nil
}

// WapDesc Getter
func (r AlitripGrouptourProductUploadRequest) GetWapDesc() string {
    return r.wapDesc
}
// SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripGrouptourProductUploadRequest) SetSubStock(subStock int64) error {
    r.subStock = subStock
    r.Set("sub_stock", subStock)
    return nil
}

// SubStock Getter
func (r AlitripGrouptourProductUploadRequest) GetSubStock() int64 {
    return r.subStock
}
// BackTrafficType Setter
// 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他
func (r *AlitripGrouptourProductUploadRequest) SetBackTrafficType(backTrafficType int64) error {
    r.backTrafficType = backTrafficType
    r.Set("back_traffic_type", backTrafficType)
    return nil
}

// BackTrafficType Getter
func (r AlitripGrouptourProductUploadRequest) GetBackTrafficType() int64 {
    return r.backTrafficType
}
// DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripGrouptourProductUploadRequest) SetDescXml(descXml string) error {
    r.descXml = descXml
    r.Set("desc_xml", descXml)
    return nil
}

// DescXml Getter
func (r AlitripGrouptourProductUploadRequest) GetDescXml() string {
    return r.descXml
}
// ConfirmTime Setter
// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripGrouptourProductUploadRequest) SetConfirmTime(confirmTime int64) error {
    r.confirmTime = confirmTime
    r.Set("confirm_time", confirmTime)
    return nil
}

// ConfirmTime Getter
func (r AlitripGrouptourProductUploadRequest) GetConfirmTime() int64 {
    return r.confirmTime
}
// Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripGrouptourProductUploadRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r AlitripGrouptourProductUploadRequest) GetTitle() string {
    return r.title
}
// RouteType Setter
// 参团线路类型。0 -目的地参团，1-为出发地参团
func (r *AlitripGrouptourProductUploadRequest) SetRouteType(routeType int64) error {
    r.routeType = routeType
    r.Set("route_type", routeType)
    return nil
}

// RouteType Getter
func (r AlitripGrouptourProductUploadRequest) GetRouteType() int64 {
    return r.routeType
}
// GroupTourPackageInfo Setter
// 套餐信息,数组类型，支持上传多个套餐信息
func (r *AlitripGrouptourProductUploadRequest) SetGroupTourPackageInfo(groupTourPackageInfo []GroupTourPackageInfo) error {
    r.groupTourPackageInfo = groupTourPackageInfo
    r.Set("group_tour_package_info", groupTourPackageInfo)
    return nil
}

// GroupTourPackageInfo Getter
func (r AlitripGrouptourProductUploadRequest) GetGroupTourPackageInfo() []GroupTourPackageInfo {
    return r.groupTourPackageInfo
}
// ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripGrouptourProductUploadRequest) SetConfirmType(confirmType int64) error {
    r.confirmType = confirmType
    r.Set("confirm_type", confirmType)
    return nil
}

// ConfirmType Getter
func (r AlitripGrouptourProductUploadRequest) GetConfirmType() int64 {
    return r.confirmType
}
// ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripGrouptourProductUploadRequest) SetItemCustomTag(itemCustomTag string) error {
    r.itemCustomTag = itemCustomTag
    r.Set("item_custom_tag", itemCustomTag)
    return nil
}

// ItemCustomTag Getter
func (r AlitripGrouptourProductUploadRequest) GetItemCustomTag() string {
    return r.itemCustomTag
}
// OutProductId Setter
// 必填，商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripGrouptourProductUploadRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

// OutProductId Getter
func (r AlitripGrouptourProductUploadRequest) GetOutProductId() string {
    return r.outProductId
}
// TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1。已废弃，以套餐维度行程晚数为准。
func (r *AlitripGrouptourProductUploadRequest) SetTripNight(tripNight int64) error {
    r.tripNight = tripNight
    r.Set("trip_night", tripNight)
    return nil
}

// TripNight Getter
func (r AlitripGrouptourProductUploadRequest) GetTripNight() int64 {
    return r.tripNight
}
// ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripGrouptourProductUploadRequest) SetToLocations(toLocations string) error {
    r.toLocations = toLocations
    r.Set("to_locations", toLocations)
    return nil
}

// ToLocations Getter
func (r AlitripGrouptourProductUploadRequest) GetToLocations() string {
    return r.toLocations
}
// PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptourProductUploadRequest) SetPicUrls(picUrls []string) error {
    r.picUrls = picUrls
    r.Set("pic_urls", picUrls)
    return nil
}

// PicUrls Getter
func (r AlitripGrouptourProductUploadRequest) GetPicUrls() []string {
    return r.picUrls
}
// ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripGrouptourProductUploadRequest) SetReserveLimit(reserveLimit string) error {
    r.reserveLimit = reserveLimit
    r.Set("reserve_limit", reserveLimit)
    return nil
}

// ReserveLimit Getter
func (r AlitripGrouptourProductUploadRequest) GetReserveLimit() string {
    return r.reserveLimit
}
// ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripGrouptourProductUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlitripGrouptourProductUploadRequest) GetItemId() int64 {
    return r.itemId
}
// RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
func (r *AlitripGrouptourProductUploadRequest) SetRefundType(refundType int64) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

// RefundType Getter
func (r AlitripGrouptourProductUploadRequest) GetRefundType() int64 {
    return r.refundType
}
// SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptourProductUploadRequest) SetSubTitles(subTitles []string) error {
    r.subTitles = subTitles
    r.Set("sub_titles", subTitles)
    return nil
}

// SubTitles Getter
func (r AlitripGrouptourProductUploadRequest) GetSubTitles() []string {
    return r.subTitles
}
// FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripGrouptourProductUploadRequest) SetFromLocations(fromLocations string) error {
    r.fromLocations = fromLocations
    r.Set("from_locations", fromLocations)
    return nil
}

// FromLocations Getter
func (r AlitripGrouptourProductUploadRequest) GetFromLocations() string {
    return r.fromLocations
}
// DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripGrouptourProductUploadRequest) SetDescHtml(descHtml string) error {
    r.descHtml = descHtml
    r.Set("desc_html", descHtml)
    return nil
}

// DescHtml Getter
func (r AlitripGrouptourProductUploadRequest) GetDescHtml() string {
    return r.descHtml
}
// TravellerTemplateId Setter
// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
func (r *AlitripGrouptourProductUploadRequest) SetTravellerTemplateId(travellerTemplateId int64) error {
    r.travellerTemplateId = travellerTemplateId
    r.Set("traveller_template_id", travellerTemplateId)
    return nil
}

// TravellerTemplateId Getter
func (r AlitripGrouptourProductUploadRequest) GetTravellerTemplateId() int64 {
    return r.travellerTemplateId
}
// IsOverseasTour Setter
// 新发布商品时必填。是否出境游，0-不是，1-是。
func (r *AlitripGrouptourProductUploadRequest) SetIsOverseasTour(isOverseasTour int64) error {
    r.isOverseasTour = isOverseasTour
    r.Set("is_overseas_tour", isOverseasTour)
    return nil
}

// IsOverseasTour Getter
func (r AlitripGrouptourProductUploadRequest) GetIsOverseasTour() int64 {
    return r.isOverseasTour
}
// PurePlay Setter
// 是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
func (r *AlitripGrouptourProductUploadRequest) SetPurePlay(purePlay int64) error {
    r.purePlay = purePlay
    r.Set("pure_play", purePlay)
    return nil
}

// PurePlay Getter
func (r AlitripGrouptourProductUploadRequest) GetPurePlay() int64 {
    return r.purePlay
}
// RefundRegulationsJson Setter
// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
func (r *AlitripGrouptourProductUploadRequest) SetRefundRegulationsJson(refundRegulationsJson string) error {
    r.refundRegulationsJson = refundRegulationsJson
    r.Set("refund_regulations_json", refundRegulationsJson)
    return nil
}

// RefundRegulationsJson Getter
func (r AlitripGrouptourProductUploadRequest) GetRefundRegulationsJson() string {
    return r.refundRegulationsJson
}
// PackageOperation Setter
// 0：使用上传的套餐信息（group_tour_package_info）覆盖商品上原有的套餐信息（此时group_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（group_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0===
func (r *AlitripGrouptourProductUploadRequest) SetPackageOperation(packageOperation int64) error {
    r.packageOperation = packageOperation
    r.Set("package_operation", packageOperation)
    return nil
}

// PackageOperation Getter
func (r AlitripGrouptourProductUploadRequest) GetPackageOperation() int64 {
    return r.packageOperation
}
// GroupTourType Setter
// 必填，线路的“细分类型”属性：1-普通跟团游、2-半自由行、3-私家团；不填默认值设置为"1-普通跟团游"。
func (r *AlitripGrouptourProductUploadRequest) SetGroupTourType(groupTourType int64) error {
    r.groupTourType = groupTourType
    r.Set("group_tour_type", groupTourType)
    return nil
}

// GroupTourType Getter
func (r AlitripGrouptourProductUploadRequest) GetGroupTourType() int64 {
    return r.groupTourType
}
// SellerCids Setter
// 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
func (r *AlitripGrouptourProductUploadRequest) SetSellerCids(sellerCids []string) error {
    r.sellerCids = sellerCids
    r.Set("seller_cids", sellerCids)
    return nil
}

// SellerCids Getter
func (r AlitripGrouptourProductUploadRequest) GetSellerCids() []string {
    return r.sellerCids
}
// SecondKill Setter
// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
func (r *AlitripGrouptourProductUploadRequest) SetSecondKill(secondKill string) error {
    r.secondKill = secondKill
    r.Set("second_kill", secondKill)
    return nil
}

// SecondKill Getter
func (r AlitripGrouptourProductUploadRequest) GetSecondKill() string {
    return r.secondKill
}
// HasDiscount Setter
// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
func (r *AlitripGrouptourProductUploadRequest) SetHasDiscount(hasDiscount bool) error {
    r.hasDiscount = hasDiscount
    r.Set("has_discount", hasDiscount)
    return nil
}

// HasDiscount Getter
func (r AlitripGrouptourProductUploadRequest) GetHasDiscount() bool {
    return r.hasDiscount
}
