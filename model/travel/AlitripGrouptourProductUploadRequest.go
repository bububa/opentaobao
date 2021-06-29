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
    _goTrafficType   int64
    // 新发布商品时必填。旅游天数。已废弃，以套餐维度行程天数为准。
    _tripDay   int64
    // 可选，手机端详情描述，xml格式，格式详见示例。
    _wapDesc   string
    // 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
    _subStock   int64
    // 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他
    _backTrafficType   int64
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
    _descXml   string
    // 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
    _confirmTime   int64
    // 新发布商品时必填。商品标题，30个中文字符以内
    _title   string
    // 参团线路类型。0 -目的地参团，1-为出发地参团
    _routeType   int64
    // 套餐信息,数组类型，支持上传多个套餐信息
    _groupTourPackageInfo   []GroupTourPackageInfo
    // 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
    _confirmType   int64
    // 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
    _itemCustomTag   string
    // 必填，商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
    _outProductId   string
    // 可选，旅游晚数，不传默认旅游天数-1。已废弃，以套餐维度行程晚数为准。
    _tripNight   int64
    // 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
    _toLocations   string
    // 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
    _picUrls   []string
    // 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
    _reserveLimit   string
    // 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
    _itemId   int64
    // 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
    _refundType   int64
    // 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
    _subTitles   []string
    // 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
    _fromLocations   string
    // PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
    _descHtml   string
    // 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
    _travellerTemplateId   int64
    // 新发布商品时必填。是否出境游，0-不是，1-是。
    _isOverseasTour   int64
    // 是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
    _purePlay   int64
    // 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
    _refundRegulationsJson   string
    // 0：使用上传的套餐信息（group_tour_package_info）覆盖商品上原有的套餐信息（此时group_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（group_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0===
    _packageOperation   int64
    // 必填，线路的“细分类型”属性：1-普通跟团游、2-半自由行、3-私家团；不填默认值设置为"1-普通跟团游"。
    _groupTourType   int64
    // 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
    _sellerCids   []string
    // 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
    _secondKill   string
    // 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
    _hasDiscount   bool
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
func (r *AlitripGrouptourProductUploadRequest) SetGoTrafficType(_goTrafficType int64) error {
    r._goTrafficType = _goTrafficType
    r.Set("go_traffic_type", _goTrafficType)
    return nil
}

// GoTrafficType Getter
func (r AlitripGrouptourProductUploadRequest) GetGoTrafficType() int64 {
    return r._goTrafficType
}
// TripDay Setter
// 新发布商品时必填。旅游天数。已废弃，以套餐维度行程天数为准。
func (r *AlitripGrouptourProductUploadRequest) SetTripDay(_tripDay int64) error {
    r._tripDay = _tripDay
    r.Set("trip_day", _tripDay)
    return nil
}

// TripDay Getter
func (r AlitripGrouptourProductUploadRequest) GetTripDay() int64 {
    return r._tripDay
}
// WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripGrouptourProductUploadRequest) SetWapDesc(_wapDesc string) error {
    r._wapDesc = _wapDesc
    r.Set("wap_desc", _wapDesc)
    return nil
}

// WapDesc Getter
func (r AlitripGrouptourProductUploadRequest) GetWapDesc() string {
    return r._wapDesc
}
// SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripGrouptourProductUploadRequest) SetSubStock(_subStock int64) error {
    r._subStock = _subStock
    r.Set("sub_stock", _subStock)
    return nil
}

// SubStock Getter
func (r AlitripGrouptourProductUploadRequest) GetSubStock() int64 {
    return r._subStock
}
// BackTrafficType Setter
// 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他
func (r *AlitripGrouptourProductUploadRequest) SetBackTrafficType(_backTrafficType int64) error {
    r._backTrafficType = _backTrafficType
    r.Set("back_traffic_type", _backTrafficType)
    return nil
}

// BackTrafficType Getter
func (r AlitripGrouptourProductUploadRequest) GetBackTrafficType() int64 {
    return r._backTrafficType
}
// DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripGrouptourProductUploadRequest) SetDescXml(_descXml string) error {
    r._descXml = _descXml
    r.Set("desc_xml", _descXml)
    return nil
}

// DescXml Getter
func (r AlitripGrouptourProductUploadRequest) GetDescXml() string {
    return r._descXml
}
// ConfirmTime Setter
// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripGrouptourProductUploadRequest) SetConfirmTime(_confirmTime int64) error {
    r._confirmTime = _confirmTime
    r.Set("confirm_time", _confirmTime)
    return nil
}

// ConfirmTime Getter
func (r AlitripGrouptourProductUploadRequest) GetConfirmTime() int64 {
    return r._confirmTime
}
// Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripGrouptourProductUploadRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r AlitripGrouptourProductUploadRequest) GetTitle() string {
    return r._title
}
// RouteType Setter
// 参团线路类型。0 -目的地参团，1-为出发地参团
func (r *AlitripGrouptourProductUploadRequest) SetRouteType(_routeType int64) error {
    r._routeType = _routeType
    r.Set("route_type", _routeType)
    return nil
}

// RouteType Getter
func (r AlitripGrouptourProductUploadRequest) GetRouteType() int64 {
    return r._routeType
}
// GroupTourPackageInfo Setter
// 套餐信息,数组类型，支持上传多个套餐信息
func (r *AlitripGrouptourProductUploadRequest) SetGroupTourPackageInfo(_groupTourPackageInfo []GroupTourPackageInfo) error {
    r._groupTourPackageInfo = _groupTourPackageInfo
    r.Set("group_tour_package_info", _groupTourPackageInfo)
    return nil
}

// GroupTourPackageInfo Getter
func (r AlitripGrouptourProductUploadRequest) GetGroupTourPackageInfo() []GroupTourPackageInfo {
    return r._groupTourPackageInfo
}
// ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripGrouptourProductUploadRequest) SetConfirmType(_confirmType int64) error {
    r._confirmType = _confirmType
    r.Set("confirm_type", _confirmType)
    return nil
}

// ConfirmType Getter
func (r AlitripGrouptourProductUploadRequest) GetConfirmType() int64 {
    return r._confirmType
}
// ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripGrouptourProductUploadRequest) SetItemCustomTag(_itemCustomTag string) error {
    r._itemCustomTag = _itemCustomTag
    r.Set("item_custom_tag", _itemCustomTag)
    return nil
}

// ItemCustomTag Getter
func (r AlitripGrouptourProductUploadRequest) GetItemCustomTag() string {
    return r._itemCustomTag
}
// OutProductId Setter
// 必填，商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripGrouptourProductUploadRequest) SetOutProductId(_outProductId string) error {
    r._outProductId = _outProductId
    r.Set("out_product_id", _outProductId)
    return nil
}

// OutProductId Getter
func (r AlitripGrouptourProductUploadRequest) GetOutProductId() string {
    return r._outProductId
}
// TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1。已废弃，以套餐维度行程晚数为准。
func (r *AlitripGrouptourProductUploadRequest) SetTripNight(_tripNight int64) error {
    r._tripNight = _tripNight
    r.Set("trip_night", _tripNight)
    return nil
}

// TripNight Getter
func (r AlitripGrouptourProductUploadRequest) GetTripNight() int64 {
    return r._tripNight
}
// ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripGrouptourProductUploadRequest) SetToLocations(_toLocations string) error {
    r._toLocations = _toLocations
    r.Set("to_locations", _toLocations)
    return nil
}

// ToLocations Getter
func (r AlitripGrouptourProductUploadRequest) GetToLocations() string {
    return r._toLocations
}
// PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptourProductUploadRequest) SetPicUrls(_picUrls []string) error {
    r._picUrls = _picUrls
    r.Set("pic_urls", _picUrls)
    return nil
}

// PicUrls Getter
func (r AlitripGrouptourProductUploadRequest) GetPicUrls() []string {
    return r._picUrls
}
// ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripGrouptourProductUploadRequest) SetReserveLimit(_reserveLimit string) error {
    r._reserveLimit = _reserveLimit
    r.Set("reserve_limit", _reserveLimit)
    return nil
}

// ReserveLimit Getter
func (r AlitripGrouptourProductUploadRequest) GetReserveLimit() string {
    return r._reserveLimit
}
// ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripGrouptourProductUploadRequest) SetItemId(_itemId int64) error {
    r._itemId = _itemId
    r.Set("item_id", _itemId)
    return nil
}

// ItemId Getter
func (r AlitripGrouptourProductUploadRequest) GetItemId() int64 {
    return r._itemId
}
// RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
func (r *AlitripGrouptourProductUploadRequest) SetRefundType(_refundType int64) error {
    r._refundType = _refundType
    r.Set("refund_type", _refundType)
    return nil
}

// RefundType Getter
func (r AlitripGrouptourProductUploadRequest) GetRefundType() int64 {
    return r._refundType
}
// SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptourProductUploadRequest) SetSubTitles(_subTitles []string) error {
    r._subTitles = _subTitles
    r.Set("sub_titles", _subTitles)
    return nil
}

// SubTitles Getter
func (r AlitripGrouptourProductUploadRequest) GetSubTitles() []string {
    return r._subTitles
}
// FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripGrouptourProductUploadRequest) SetFromLocations(_fromLocations string) error {
    r._fromLocations = _fromLocations
    r.Set("from_locations", _fromLocations)
    return nil
}

// FromLocations Getter
func (r AlitripGrouptourProductUploadRequest) GetFromLocations() string {
    return r._fromLocations
}
// DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripGrouptourProductUploadRequest) SetDescHtml(_descHtml string) error {
    r._descHtml = _descHtml
    r.Set("desc_html", _descHtml)
    return nil
}

// DescHtml Getter
func (r AlitripGrouptourProductUploadRequest) GetDescHtml() string {
    return r._descHtml
}
// TravellerTemplateId Setter
// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
func (r *AlitripGrouptourProductUploadRequest) SetTravellerTemplateId(_travellerTemplateId int64) error {
    r._travellerTemplateId = _travellerTemplateId
    r.Set("traveller_template_id", _travellerTemplateId)
    return nil
}

// TravellerTemplateId Getter
func (r AlitripGrouptourProductUploadRequest) GetTravellerTemplateId() int64 {
    return r._travellerTemplateId
}
// IsOverseasTour Setter
// 新发布商品时必填。是否出境游，0-不是，1-是。
func (r *AlitripGrouptourProductUploadRequest) SetIsOverseasTour(_isOverseasTour int64) error {
    r._isOverseasTour = _isOverseasTour
    r.Set("is_overseas_tour", _isOverseasTour)
    return nil
}

// IsOverseasTour Getter
func (r AlitripGrouptourProductUploadRequest) GetIsOverseasTour() int64 {
    return r._isOverseasTour
}
// PurePlay Setter
// 是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
func (r *AlitripGrouptourProductUploadRequest) SetPurePlay(_purePlay int64) error {
    r._purePlay = _purePlay
    r.Set("pure_play", _purePlay)
    return nil
}

// PurePlay Getter
func (r AlitripGrouptourProductUploadRequest) GetPurePlay() int64 {
    return r._purePlay
}
// RefundRegulationsJson Setter
// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
func (r *AlitripGrouptourProductUploadRequest) SetRefundRegulationsJson(_refundRegulationsJson string) error {
    r._refundRegulationsJson = _refundRegulationsJson
    r.Set("refund_regulations_json", _refundRegulationsJson)
    return nil
}

// RefundRegulationsJson Getter
func (r AlitripGrouptourProductUploadRequest) GetRefundRegulationsJson() string {
    return r._refundRegulationsJson
}
// PackageOperation Setter
// 0：使用上传的套餐信息（group_tour_package_info）覆盖商品上原有的套餐信息（此时group_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（group_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0===
func (r *AlitripGrouptourProductUploadRequest) SetPackageOperation(_packageOperation int64) error {
    r._packageOperation = _packageOperation
    r.Set("package_operation", _packageOperation)
    return nil
}

// PackageOperation Getter
func (r AlitripGrouptourProductUploadRequest) GetPackageOperation() int64 {
    return r._packageOperation
}
// GroupTourType Setter
// 必填，线路的“细分类型”属性：1-普通跟团游、2-半自由行、3-私家团；不填默认值设置为"1-普通跟团游"。
func (r *AlitripGrouptourProductUploadRequest) SetGroupTourType(_groupTourType int64) error {
    r._groupTourType = _groupTourType
    r.Set("group_tour_type", _groupTourType)
    return nil
}

// GroupTourType Getter
func (r AlitripGrouptourProductUploadRequest) GetGroupTourType() int64 {
    return r._groupTourType
}
// SellerCids Setter
// 关联商品与店铺类目 结构:"cid1,cid2,...,"。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
func (r *AlitripGrouptourProductUploadRequest) SetSellerCids(_sellerCids []string) error {
    r._sellerCids = _sellerCids
    r.Set("seller_cids", _sellerCids)
    return nil
}

// SellerCids Getter
func (r AlitripGrouptourProductUploadRequest) GetSellerCids() []string {
    return r._sellerCids
}
// SecondKill Setter
// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
func (r *AlitripGrouptourProductUploadRequest) SetSecondKill(_secondKill string) error {
    r._secondKill = _secondKill
    r.Set("second_kill", _secondKill)
    return nil
}

// SecondKill Getter
func (r AlitripGrouptourProductUploadRequest) GetSecondKill() string {
    return r._secondKill
}
// HasDiscount Setter
// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
func (r *AlitripGrouptourProductUploadRequest) SetHasDiscount(_hasDiscount bool) error {
    r._hasDiscount = _hasDiscount
    r.Set("has_discount", _hasDiscount)
    return nil
}

// HasDiscount Getter
func (r AlitripGrouptourProductUploadRequest) GetHasDiscount() bool {
    return r._hasDiscount
}
