package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripGrouptourProductUploadAPIRequest 新版跟团游商品维护接口 API请求
// alitrip.grouptour.product.upload
//
// 新版跟团游商品维护接口
type AlitripGrouptourProductUploadAPIRequest struct {
	model.Params
	// 套餐信息,数组类型，支持上传多个套餐信息
	_groupTourPackageInfo []GroupTourPackageInfo
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
	// 必填，商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
	_outProductId string
	// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
	_toLocations string
	// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
	_reserveLimit string
	// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
	_fromLocations string
	// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
	_descHtml string
	// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
	_refundRegulationsJson string
	// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
	_secondKill string
	// 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船，100-其他
	_goTrafficType int64
	// 新发布商品时必填。旅游天数。已废弃，以套餐维度行程天数为准。
	_tripDay int64
	// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
	_subStock int64
	// 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他
	_backTrafficType int64
	// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
	_confirmTime int64
	// 参团线路类型。0 -目的地参团，1-为出发地参团
	_routeType int64
	// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
	_confirmType int64
	// 可选，旅游晚数，不传默认旅游天数-1。已废弃，以套餐维度行程晚数为准。
	_tripNight int64
	// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
	_itemId int64
	// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
	_refundType int64
	// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具->出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
	_travellerTemplateId int64
	// 新发布商品时必填。是否出境游，0-不是，1-是。
	_isOverseasTour int64
	// 是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
	_purePlay int64
	// 0：使用上传的套餐信息（group_tour_package_info）覆盖商品上原有的套餐信息（此时group_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（group_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0===
	_packageOperation int64
	// 必填，线路的“细分类型”属性：1-普通跟团游、2-半自由行、3-私家团；不填默认值设置为"1-普通跟团游"。
	_groupTourType int64
	// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
	_hasDiscount bool
}

// NewAlitripGrouptourProductUploadRequest 初始化AlitripGrouptourProductUploadAPIRequest对象
func NewAlitripGrouptourProductUploadRequest() *AlitripGrouptourProductUploadAPIRequest {
	return &AlitripGrouptourProductUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripGrouptourProductUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.grouptour.product.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripGrouptourProductUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripGrouptourProductUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGroupTourPackageInfo is GroupTourPackageInfo Setter
// 套餐信息,数组类型，支持上传多个套餐信息
func (r *AlitripGrouptourProductUploadAPIRequest) SetGroupTourPackageInfo(_groupTourPackageInfo []GroupTourPackageInfo) error {
	r._groupTourPackageInfo = _groupTourPackageInfo
	r.Set("group_tour_package_info", _groupTourPackageInfo)
	return nil
}

// GetGroupTourPackageInfo GroupTourPackageInfo Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetGroupTourPackageInfo() []GroupTourPackageInfo {
	return r._groupTourPackageInfo
}

// SetPicUrls is PicUrls Setter
// 新发布商品时必填。商品图片路径。最多支持5张，第一张为主图，必填，其余四张可选填。图片链接支持外链图片（即商家系统中图片链接，必须外网可访问，且格式为jpg或jpeg，大小在500k以内），或者用户淘宝空间内的图片链接。对于外链图片，将自动下载并上传用户淘宝图片空间，上传失败的外链图片将自动忽略不计。。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptourProductUploadAPIRequest) SetPicUrls(_picUrls []string) error {
	r._picUrls = _picUrls
	r.Set("pic_urls", _picUrls)
	return nil
}

// GetPicUrls PicUrls Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetPicUrls() []string {
	return r._picUrls
}

// SetSubTitles is SubTitles Setter
// 可选，商品亮点，最多支持4个亮点。注：在SDK中数组多个元素间以英文逗号分隔
func (r *AlitripGrouptourProductUploadAPIRequest) SetSubTitles(_subTitles []string) error {
	r._subTitles = _subTitles
	r.Set("sub_titles", _subTitles)
	return nil
}

// GetSubTitles SubTitles Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetSubTitles() []string {
	return r._subTitles
}

// SetSellerCids is SellerCids Setter
// 关联商品与店铺类目 结构:&#34;cid1,cid2,...,&#34;。如何获取卖家店铺类目具体参见：http://open.taobao.com/doc2/apiDetail.htm?apiId=65
func (r *AlitripGrouptourProductUploadAPIRequest) SetSellerCids(_sellerCids []string) error {
	r._sellerCids = _sellerCids
	r.Set("seller_cids", _sellerCids)
	return nil
}

// GetSellerCids SellerCids Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetSellerCids() []string {
	return r._sellerCids
}

// SetWapDesc is WapDesc Setter
// 可选，手机端详情描述，xml格式，格式详见示例。
func (r *AlitripGrouptourProductUploadAPIRequest) SetWapDesc(_wapDesc string) error {
	r._wapDesc = _wapDesc
	r.Set("wap_desc", _wapDesc)
	return nil
}

// GetWapDesc WapDesc Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetWapDesc() string {
	return r._wapDesc
}

// SetDescXml is DescXml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），xml格式：DESC根标签必须有，每一个亮点（HIGHLIGHT）支持1个title子标签，1个txt子标签和多个img子标签。
func (r *AlitripGrouptourProductUploadAPIRequest) SetDescXml(_descXml string) error {
	r._descXml = _descXml
	r.Set("desc_xml", _descXml)
	return nil
}

// GetDescXml DescXml Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetDescXml() string {
	return r._descXml
}

// SetTitle is Title Setter
// 新发布商品时必填。商品标题，30个中文字符以内
func (r *AlitripGrouptourProductUploadAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetTitle() string {
	return r._title
}

// SetItemCustomTag is ItemCustomTag Setter
// 可选，商家自定义标签（最多4个字，超长则自动截断，会进行违禁词校验）
func (r *AlitripGrouptourProductUploadAPIRequest) SetItemCustomTag(_itemCustomTag string) error {
	r._itemCustomTag = _itemCustomTag
	r.Set("item_custom_tag", _itemCustomTag)
	return nil
}

// GetItemCustomTag ItemCustomTag Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetItemCustomTag() string {
	return r._itemCustomTag
}

// SetOutProductId is OutProductId Setter
// 必填，商家自定义商品编码。注：商品基本信息维护、价格库存维护，商品查询都以该编码为主键。
func (r *AlitripGrouptourProductUploadAPIRequest) SetOutProductId(_outProductId string) error {
	r._outProductId = _outProductId
	r.Set("out_product_id", _outProductId)
	return nil
}

// GetOutProductId OutProductId Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetOutProductId() string {
	return r._outProductId
}

// SetToLocations is ToLocations Setter
// 新发布商品时必填。目的地，多个目的地用英文逗号分隔。地址可以使用飞猪标准地址名称，也可以使用商家系统中目的地地址（支持商家目的地id和商家目的地名称）。如果需要使用商家目的地地址，必须在目的地关联页（https://sell.alitrip.com/icenter/main.htm#/widgets/api-adaptor?_k=n61ii0）配置映射关系（一次性批量上传建立映射关系，之后度假所有类目、API接口共用该映射关系）。 商家目的地地址使用示例1：东京,大阪。示例2：123,124。说明：商家目的地id（123,124）会根据映射关系自动转换成飞猪标准地址
func (r *AlitripGrouptourProductUploadAPIRequest) SetToLocations(_toLocations string) error {
	r._toLocations = _toLocations
	r.Set("to_locations", _toLocations)
	return nil
}

// GetToLocations ToLocations Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetToLocations() string {
	return r._toLocations
}

// SetReserveLimit is ReserveLimit Setter
// 可选，该商品提前预定时间限制。格式：1_18_00，含义：该商品必须提前1天预定，且在18:00之前完成预定
func (r *AlitripGrouptourProductUploadAPIRequest) SetReserveLimit(_reserveLimit string) error {
	r._reserveLimit = _reserveLimit
	r.Set("reserve_limit", _reserveLimit)
	return nil
}

// GetReserveLimit ReserveLimit Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetReserveLimit() string {
	return r._reserveLimit
}

// SetFromLocations is FromLocations Setter
// 新发布商品时必填。出发地，多个出发地用英文逗号分隔。使用说明同“目的地”
func (r *AlitripGrouptourProductUploadAPIRequest) SetFromLocations(_fromLocations string) error {
	r._fromLocations = _fromLocations
	r.Set("from_locations", _fromLocations)
	return nil
}

// GetFromLocations FromLocations Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetFromLocations() string {
	return r._fromLocations
}

// SetDescHtml is DescHtml Setter
// PC端详情描述（新发布商品时，desc_xml和desc_html二者至少填写一个），Html格式。商家自定义Html格式描述。
func (r *AlitripGrouptourProductUploadAPIRequest) SetDescHtml(_descHtml string) error {
	r._descHtml = _descHtml
	r.Set("desc_html", _descHtml)
	return nil
}

// GetDescHtml DescHtml Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetDescHtml() string {
	return r._descHtml
}

// SetRefundRegulationsJson is RefundRegulationsJson Setter
// 特殊可选，退款规则（json数组格式）。自定义退改时需填写（与refund_regulations字段二选一）。示例中一共包含4条规则（3条平日规则，1条节假日规则），按照顺序每条规则含义如下：出行前5日及以上，买家违约收取总费用的50，卖家违约收取总费用的20；出行前4日至1日，买家违约收取总费用的80，卖家违约收取总费用的50；行程开始当天，买家违约收取总费用的100，卖家违约收取总费用的70；如果行程日期包含节假日，则节假日条款为买家违约收取总费用的100，卖家违约收取总费用的90
func (r *AlitripGrouptourProductUploadAPIRequest) SetRefundRegulationsJson(_refundRegulationsJson string) error {
	r._refundRegulationsJson = _refundRegulationsJson
	r.Set("refund_regulations_json", _refundRegulationsJson)
	return nil
}

// GetRefundRegulationsJson RefundRegulationsJson Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetRefundRegulationsJson() string {
	return r._refundRegulationsJson
}

// SetSecondKill is SecondKill Setter
// 商品秒杀，商品秒杀三个值：可选类型web_only(只能通过web网络秒杀)，wap_only(只能通过wap网络秒杀)，web_and_wap(既能通过web秒杀也能通过wap秒杀)
func (r *AlitripGrouptourProductUploadAPIRequest) SetSecondKill(_secondKill string) error {
	r._secondKill = _secondKill
	r.Set("second_kill", _secondKill)
	return nil
}

// GetSecondKill SecondKill Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetSecondKill() string {
	return r._secondKill
}

// SetGoTrafficType is GoTrafficType Setter
// 新发布商品时必填。去程交通。1-飞机，2-火车，3-汽，4-船，100-其他
func (r *AlitripGrouptourProductUploadAPIRequest) SetGoTrafficType(_goTrafficType int64) error {
	r._goTrafficType = _goTrafficType
	r.Set("go_traffic_type", _goTrafficType)
	return nil
}

// GetGoTrafficType GoTrafficType Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetGoTrafficType() int64 {
	return r._goTrafficType
}

// SetTripDay is TripDay Setter
// 新发布商品时必填。旅游天数。已废弃，以套餐维度行程天数为准。
func (r *AlitripGrouptourProductUploadAPIRequest) SetTripDay(_tripDay int64) error {
	r._tripDay = _tripDay
	r.Set("trip_day", _tripDay)
	return nil
}

// GetTripDay TripDay Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetTripDay() int64 {
	return r._tripDay
}

// SetSubStock is SubStock Setter
// 可选，减库存方式。0-拍下减库存。1-付款减库存。不传默认为0
func (r *AlitripGrouptourProductUploadAPIRequest) SetSubStock(_subStock int64) error {
	r._subStock = _subStock
	r.Set("sub_stock", _subStock)
	return nil
}

// GetSubStock SubStock Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetSubStock() int64 {
	return r._subStock
}

// SetBackTrafficType is BackTrafficType Setter
// 新发布商品时必填。回程交通。1-飞机，2-火车，3-汽，4-船，100-其他
func (r *AlitripGrouptourProductUploadAPIRequest) SetBackTrafficType(_backTrafficType int64) error {
	r._backTrafficType = _backTrafficType
	r.Set("back_traffic_type", _backTrafficType)
	return nil
}

// GetBackTrafficType BackTrafficType Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetBackTrafficType() int64 {
	return r._backTrafficType
}

// SetConfirmTime is ConfirmTime Setter
// 可选，资源确认时长，当confirm_type=2时必填。1：2个工作小时内确认，2：6个工作小时内确认，3：9个工作小时内确认，4：18个工作小时内确认
func (r *AlitripGrouptourProductUploadAPIRequest) SetConfirmTime(_confirmTime int64) error {
	r._confirmTime = _confirmTime
	r.Set("confirm_time", _confirmTime)
	return nil
}

// GetConfirmTime ConfirmTime Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetConfirmTime() int64 {
	return r._confirmTime
}

// SetRouteType is RouteType Setter
// 参团线路类型。0 -目的地参团，1-为出发地参团
func (r *AlitripGrouptourProductUploadAPIRequest) SetRouteType(_routeType int64) error {
	r._routeType = _routeType
	r.Set("route_type", _routeType)
	return nil
}

// GetRouteType RouteType Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetRouteType() int64 {
	return r._routeType
}

// SetConfirmType is ConfirmType Setter
// 可选，资源确认类型。1-即时确认，2-二次确认。不传默认1
func (r *AlitripGrouptourProductUploadAPIRequest) SetConfirmType(_confirmType int64) error {
	r._confirmType = _confirmType
	r.Set("confirm_type", _confirmType)
	return nil
}

// GetConfirmType ConfirmType Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetConfirmType() int64 {
	return r._confirmType
}

// SetTripNight is TripNight Setter
// 可选，旅游晚数，不传默认旅游天数-1。已废弃，以套餐维度行程晚数为准。
func (r *AlitripGrouptourProductUploadAPIRequest) SetTripNight(_tripNight int64) error {
	r._tripNight = _tripNight
	r.Set("trip_night", _tripNight)
	return nil
}

// GetTripNight TripNight Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetTripNight() int64 {
	return r._tripNight
}

// SetItemId is ItemId Setter
// 可选，淘系商品id，用于将out_product_id关联到已经存在的商品，并且修改该商品外部商家编码为out_product_id。
func (r *AlitripGrouptourProductUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetRefundType is RefundType Setter
// 可选，退改规则类型。0-平台标准退改规则，1-自定义退改规则，2-不支持退改（已废弃，勿用），7-新版自定义退改规则。不传默认为0
func (r *AlitripGrouptourProductUploadAPIRequest) SetRefundType(_refundType int64) error {
	r._refundType = _refundType
	r.Set("refund_type", _refundType)
	return nil
}

// GetRefundType RefundType Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetRefundType() int64 {
	return r._refundType
}

// SetTravellerTemplateId is TravellerTemplateId Setter
// 可选，出行人模板id。模板id需要商家以店铺账号身份登录飞猪商家工作台，从卖家工具-&gt;出行人管理中获取。注意：如果传0则代表设置为不需要出行人模板或使用飞猪平台默认的类目模板。
func (r *AlitripGrouptourProductUploadAPIRequest) SetTravellerTemplateId(_travellerTemplateId int64) error {
	r._travellerTemplateId = _travellerTemplateId
	r.Set("traveller_template_id", _travellerTemplateId)
	return nil
}

// GetTravellerTemplateId TravellerTemplateId Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetTravellerTemplateId() int64 {
	return r._travellerTemplateId
}

// SetIsOverseasTour is IsOverseasTour Setter
// 新发布商品时必填。是否出境游，0-不是，1-是。
func (r *AlitripGrouptourProductUploadAPIRequest) SetIsOverseasTour(_isOverseasTour int64) error {
	r._isOverseasTour = _isOverseasTour
	r.Set("is_overseas_tour", _isOverseasTour)
	return nil
}

// GetIsOverseasTour IsOverseasTour Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetIsOverseasTour() int64 {
	return r._isOverseasTour
}

// SetPurePlay is PurePlay Setter
// 是否纯玩团。0-纯玩团，1-含购物团。新发布商品时不传默认为“含购物团”
func (r *AlitripGrouptourProductUploadAPIRequest) SetPurePlay(_purePlay int64) error {
	r._purePlay = _purePlay
	r.Set("pure_play", _purePlay)
	return nil
}

// GetPurePlay PurePlay Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetPurePlay() int64 {
	return r._purePlay
}

// SetPackageOperation is PackageOperation Setter
// 0：使用上传的套餐信息（group_tour_package_info）覆盖商品上原有的套餐信息（此时group_tour_package_info中设置的packageOperation无效）；1：根据套餐信息（group_tour_package_info）中的packageOperation和outProductId增加，修改，删除指定套餐，====默认值为0===
func (r *AlitripGrouptourProductUploadAPIRequest) SetPackageOperation(_packageOperation int64) error {
	r._packageOperation = _packageOperation
	r.Set("package_operation", _packageOperation)
	return nil
}

// GetPackageOperation PackageOperation Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetPackageOperation() int64 {
	return r._packageOperation
}

// SetGroupTourType is GroupTourType Setter
// 必填，线路的“细分类型”属性：1-普通跟团游、2-半自由行、3-私家团；不填默认值设置为&#34;1-普通跟团游&#34;。
func (r *AlitripGrouptourProductUploadAPIRequest) SetGroupTourType(_groupTourType int64) error {
	r._groupTourType = _groupTourType
	r.Set("group_tour_type", _groupTourType)
	return nil
}

// GetGroupTourType GroupTourType Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetGroupTourType() int64 {
	return r._groupTourType
}

// SetHasDiscount is HasDiscount Setter
// 是否支持会员打折。可选值：true，false；默认值：false(不打折)。不传的话默认为false
func (r *AlitripGrouptourProductUploadAPIRequest) SetHasDiscount(_hasDiscount bool) error {
	r._hasDiscount = _hasDiscount
	r.Set("has_discount", _hasDiscount)
	return nil
}

// GetHasDiscount HasDiscount Getter
func (r AlitripGrouptourProductUploadAPIRequest) GetHasDiscount() bool {
	return r._hasDiscount
}
