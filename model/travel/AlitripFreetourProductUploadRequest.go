package travel

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自由行商品发布及编辑接口 APIRequest
alitrip.freetour.product.upload

自由行 产品维护接口。
接口同时支持新商品发布 和 现有商品编辑：
1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
*/
type AlitripFreetourProductUploadRequest struct {
    model.Params

    // 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船，100-其他
    goTrafficType   int64 

    // 套餐信息,数组类型，支持上传多个套餐信息
    freeTourPackageInfo   []FreeTourPackageInfo 

    // 新发布商品时必填。旅游天数
    tripDay   int64 

    // 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
    subStock   int64 

    // 可选，手机端详情描述，xml格式，格式详见示例。
    wapDesc   string 

    // 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他
    backTrafficType   int64 

    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
    descXml   string 

    // 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
    confirmTime   int64 

    // 新发布商品时必填。商品标题，30个中文字符以内
    title   string 

    // 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
    confirmType   int64 

    // 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
    itemCustomTag   string 

    // 商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
    outProductId   string 

    // 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
    toLocations   string 

    // 可选，旅游晚数，不传默认旅游天数-1
    tripNight   int64 

    // 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
    picUrls   []string 

    // 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
    itemId   int64 

    // 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
    reserveLimit   string 

    // 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
    refundType   int64 

    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
    descHtml   string 

    // 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
    fromLocations   string 

    // 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
    subTitles   []string 

    // 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
    travellerTemplateId   int64 

    // 新发布商品时必填。是否出境游，0-不是，1-是。
    isOverseasTour   int64 

    // 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
    refundRegulationsJson   string 

    // 0：使用上传的套餐信息（free_tour_package_info）覆盖商品上原有的套餐信息（此时free_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（free_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0===
    packageOperation   int64 

    // 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
    sellerCids   []string 

    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    secondKill   string 

    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    hasDiscount   bool 

}

func NewAlitripFreetourProductUploadRequest() *AlitripFreetourProductUploadRequest{
    return &AlitripFreetourProductUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripFreetourProductUploadRequest) GetApiMethodName() string {
    return "alitrip.freetour.product.upload"
}

func (r AlitripFreetourProductUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripFreetourProductUploadRequest) SetGoTrafficType(goTrafficType int64) error {
    r.goTrafficType = goTrafficType
    r.Set("go_traffic_type", goTrafficType)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetGoTrafficType() int64 {
    return r.goTrafficType
}

func (r *AlitripFreetourProductUploadRequest) SetFreeTourPackageInfo(freeTourPackageInfo []FreeTourPackageInfo) error {
    r.freeTourPackageInfo = freeTourPackageInfo
    r.Set("free_tour_package_info", freeTourPackageInfo)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetFreeTourPackageInfo() []FreeTourPackageInfo {
    return r.freeTourPackageInfo
}

func (r *AlitripFreetourProductUploadRequest) SetTripDay(tripDay int64) error {
    r.tripDay = tripDay
    r.Set("trip_day", tripDay)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetTripDay() int64 {
    return r.tripDay
}

func (r *AlitripFreetourProductUploadRequest) SetSubStock(subStock int64) error {
    r.subStock = subStock
    r.Set("sub_stock", subStock)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetSubStock() int64 {
    return r.subStock
}

func (r *AlitripFreetourProductUploadRequest) SetWapDesc(wapDesc string) error {
    r.wapDesc = wapDesc
    r.Set("wap_desc", wapDesc)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetWapDesc() string {
    return r.wapDesc
}

func (r *AlitripFreetourProductUploadRequest) SetBackTrafficType(backTrafficType int64) error {
    r.backTrafficType = backTrafficType
    r.Set("back_traffic_type", backTrafficType)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetBackTrafficType() int64 {
    return r.backTrafficType
}

func (r *AlitripFreetourProductUploadRequest) SetDescXml(descXml string) error {
    r.descXml = descXml
    r.Set("desc_xml", descXml)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetDescXml() string {
    return r.descXml
}

func (r *AlitripFreetourProductUploadRequest) SetConfirmTime(confirmTime int64) error {
    r.confirmTime = confirmTime
    r.Set("confirm_time", confirmTime)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetConfirmTime() int64 {
    return r.confirmTime
}

func (r *AlitripFreetourProductUploadRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetTitle() string {
    return r.title
}

func (r *AlitripFreetourProductUploadRequest) SetConfirmType(confirmType int64) error {
    r.confirmType = confirmType
    r.Set("confirm_type", confirmType)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetConfirmType() int64 {
    return r.confirmType
}

func (r *AlitripFreetourProductUploadRequest) SetItemCustomTag(itemCustomTag string) error {
    r.itemCustomTag = itemCustomTag
    r.Set("item_custom_tag", itemCustomTag)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetItemCustomTag() string {
    return r.itemCustomTag
}

func (r *AlitripFreetourProductUploadRequest) SetOutProductId(outProductId string) error {
    r.outProductId = outProductId
    r.Set("out_product_id", outProductId)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetOutProductId() string {
    return r.outProductId
}

func (r *AlitripFreetourProductUploadRequest) SetToLocations(toLocations string) error {
    r.toLocations = toLocations
    r.Set("to_locations", toLocations)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetToLocations() string {
    return r.toLocations
}

func (r *AlitripFreetourProductUploadRequest) SetTripNight(tripNight int64) error {
    r.tripNight = tripNight
    r.Set("trip_night", tripNight)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetTripNight() int64 {
    return r.tripNight
}

func (r *AlitripFreetourProductUploadRequest) SetPicUrls(picUrls []string) error {
    r.picUrls = picUrls
    r.Set("pic_urls", picUrls)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetPicUrls() []string {
    return r.picUrls
}

func (r *AlitripFreetourProductUploadRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlitripFreetourProductUploadRequest) SetReserveLimit(reserveLimit string) error {
    r.reserveLimit = reserveLimit
    r.Set("reserve_limit", reserveLimit)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetReserveLimit() string {
    return r.reserveLimit
}

func (r *AlitripFreetourProductUploadRequest) SetRefundType(refundType int64) error {
    r.refundType = refundType
    r.Set("refund_type", refundType)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetRefundType() int64 {
    return r.refundType
}

func (r *AlitripFreetourProductUploadRequest) SetDescHtml(descHtml string) error {
    r.descHtml = descHtml
    r.Set("desc_html", descHtml)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetDescHtml() string {
    return r.descHtml
}

func (r *AlitripFreetourProductUploadRequest) SetFromLocations(fromLocations string) error {
    r.fromLocations = fromLocations
    r.Set("from_locations", fromLocations)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetFromLocations() string {
    return r.fromLocations
}

func (r *AlitripFreetourProductUploadRequest) SetSubTitles(subTitles []string) error {
    r.subTitles = subTitles
    r.Set("sub_titles", subTitles)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetSubTitles() []string {
    return r.subTitles
}

func (r *AlitripFreetourProductUploadRequest) SetTravellerTemplateId(travellerTemplateId int64) error {
    r.travellerTemplateId = travellerTemplateId
    r.Set("traveller_template_id", travellerTemplateId)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetTravellerTemplateId() int64 {
    return r.travellerTemplateId
}

func (r *AlitripFreetourProductUploadRequest) SetIsOverseasTour(isOverseasTour int64) error {
    r.isOverseasTour = isOverseasTour
    r.Set("is_overseas_tour", isOverseasTour)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetIsOverseasTour() int64 {
    return r.isOverseasTour
}

func (r *AlitripFreetourProductUploadRequest) SetRefundRegulationsJson(refundRegulationsJson string) error {
    r.refundRegulationsJson = refundRegulationsJson
    r.Set("refund_regulations_json", refundRegulationsJson)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetRefundRegulationsJson() string {
    return r.refundRegulationsJson
}

func (r *AlitripFreetourProductUploadRequest) SetPackageOperation(packageOperation int64) error {
    r.packageOperation = packageOperation
    r.Set("package_operation", packageOperation)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetPackageOperation() int64 {
    return r.packageOperation
}

func (r *AlitripFreetourProductUploadRequest) SetSellerCids(sellerCids []string) error {
    r.sellerCids = sellerCids
    r.Set("seller_cids", sellerCids)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetSellerCids() []string {
    return r.sellerCids
}

func (r *AlitripFreetourProductUploadRequest) SetSecondKill(secondKill string) error {
    r.secondKill = secondKill
    r.Set("second_kill", secondKill)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetSecondKill() string {
    return r.secondKill
}

func (r *AlitripFreetourProductUploadRequest) SetHasDiscount(hasDiscount bool) error {
    r.hasDiscount = hasDiscount
    r.Set("has_discount", hasDiscount)
    return nil
}

func (r AlitripFreetourProductUploadRequest) GetHasDiscount() bool {
    return r.hasDiscount
}

