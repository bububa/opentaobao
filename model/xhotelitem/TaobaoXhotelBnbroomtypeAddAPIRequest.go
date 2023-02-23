package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbroomtypeAddAPIRequest 民宿新增房源 API请求
// taobao.xhotel.bnbroomtype.add
//
// 添加民宿房源
type TaobaoXhotelBnbroomtypeAddAPIRequest struct {
	model.Params
	// 房源图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
	_pics []BnbPictureDto
	// 房源名
	_name []string
	// 民宿名称，默认取bnbName
	_bnbName []string
	// 房源外部标签 标签信息，逗号(,)分隔，最多1000字符
	_outerTags string
	// 外部门店id
	_outHid string
	// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
	_tel string
	// 位置描述
	_localInfo string
	// 品牌名称，最多100字符
	_brand string
	// 房源英文名
	_nameE string
	// 装修时间，格式为2015-01-01装修时间
	_decorateTime string
	// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
	_settlementCurrency string
	// 周边描述
	_surroundInfo string
	// 详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1  ，最多500字符
	_activitiesAllowed string
	// 开业时间，格式为2015-01-01
	_openingTime string
	// 如果要变更商品房源编码请使用该字段。
	_newOuterId string
	// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空
	_houseModel string
	// 视频地址，最多1000字符
	_videoUrl string
	// 销售渠道,默认taobao
	_vendor string
	// 亮点描述，最多1000字符
	_brightspot string
	// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
	_floor string
	// 房源介绍,最多2000字符
	_introduction string
	// 入住须知，最多2000字符
	_checkInNotes string
	// 真实联系方式
	_realTel string
	// 设施服务。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
	_service string
	// 房源id, 这是卖家自己系统中的ID
	_outerId string
	// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
	_bedInfo string
	// 额外收费
	_extraFee string
	// 标准酒店服务,参考文档https://fliggy.open.taobao.com/doc.htm?docId=120362&docType=1
	_standardRoomFacilities string
	// 民宿扩展信息
	_bnbExtend string
	// 单间面积，单位平方米
	_rentSize int64
	// 是否支持IM聊天 0不支持 1支持
	_supportIm int64
	// 清洁费是否收取 0：否 1：是
	_cleaningCharge int64
	// 发票，0：卖家提供发票，1：房东提供发票
	_invoice int64
	// 装修等级 1 精装；2普通；3简装
	_decorateLevel int64
	// 民宿入住要求&附加信息
	_bnbBookingTime *BnbBookingTimeDto
	// 是否可接待外宾 0：否 1：是；默认值: 0
	_receiveForeigners int64
	// 清洁费类型 0.线下；1.线上
	_cleaningType int64
	// 押金金额
	_depositAmount int64
	// 0-n；若不可加床，值为0
	_extraBedsNum int64
	// 可提供发票类型，1.专票 2.纸质普票 3.电子普票
	_invoiceType int64
	// 是否有前台 0没有 1有
	_hasFrontDesk int64
	// 是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_guestAge int64
	// 最大入住人数 1-99
	_maxOccupancy int64
	// 是否使用实拍图片 0不使用 1使用
	_isUseShootImage int64
	// 状态 0：在线 -1：不在线 -2:停售
	_status *model.File
	// 0：不限制，1：只限男性，2：只限女性'
	_guestGender int64
	// 清洁费金额；整数[1，9999999]
	_extraCleaningCharge int64
	// 出租类型，1整租；2分租。3床位 默认整租，该字段不能更新
	_rentType int64
	// 0-无窗  1-有窗  2-部分有窗 3-暗窗 4-部分暗窗  5-落地窗
	_windowType int64
	// 有无资质执照 0 没有 1有
	_hasLicense int64
	// 是否开启闪订 0不开启 1开启
	_quickOrder int64
	// 单间面积，单位平方米
	_houseSize int64
	// 房源类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_productType int64
	// 是否与房东同住 0 不同住 1同住
	_hasLandlord int64
	// 加人收费信息
	_charge *BnbChargeDto
	// 装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_decorateStyle int64
	// 是否信用免押金0：否 1：是
	_supportcredit int64
	// “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_cleaningFrequency int64
	// 民宿房源位置信息
	_location *BnbLocationDto
	// 风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
	_scenicFeature int64
	// 押金类型0.线下；1.线上
	_depositType int64
	// 加床费,分为单位
	_extraBedsFee int64
	// 添加标准房源匹配
	_srid int64
}

// NewTaobaoXhotelBnbroomtypeAddRequest 初始化TaobaoXhotelBnbroomtypeAddAPIRequest对象
func NewTaobaoXhotelBnbroomtypeAddRequest() *TaobaoXhotelBnbroomtypeAddAPIRequest {
	return &TaobaoXhotelBnbroomtypeAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbroomtype.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPics is Pics Setter
// 房源图片只支持远程图片，格式如下：[{&#34;url&#34;:&#34;http://taobao.com/123.jpg&#34;,&#34;ismain&#34;:&#34;true&#34;},{&#34;url&#34;:&#34;http://taobao.com/456.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;},{&#34;url&#34;:&#34;http://taobao.com/789.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetPics(_pics []BnbPictureDto) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// GetPics Pics Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetPics() []BnbPictureDto {
	return r._pics
}

// SetName is Name Setter
// 房源名
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetName(_name []string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetName() []string {
	return r._name
}

// SetBnbName is BnbName Setter
// 民宿名称，默认取bnbName
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetBnbName(_bnbName []string) error {
	r._bnbName = _bnbName
	r.Set("bnb_name", _bnbName)
	return nil
}

// GetBnbName BnbName Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetBnbName() []string {
	return r._bnbName
}

// SetOuterTags is OuterTags Setter
// 房源外部标签 标签信息，逗号(,)分隔，最多1000字符
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetOuterTags(_outerTags string) error {
	r._outerTags = _outerTags
	r.Set("outer_tags", _outerTags)
	return nil
}

// GetOuterTags OuterTags Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetOuterTags() string {
	return r._outerTags
}

// SetOutHid is OutHid Setter
// 外部门店id
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetTel is Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetTel(_tel string) error {
	r._tel = _tel
	r.Set("tel", _tel)
	return nil
}

// GetTel Tel Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetTel() string {
	return r._tel
}

// SetLocalInfo is LocalInfo Setter
// 位置描述
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetLocalInfo(_localInfo string) error {
	r._localInfo = _localInfo
	r.Set("local_info", _localInfo)
	return nil
}

// GetLocalInfo LocalInfo Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetLocalInfo() string {
	return r._localInfo
}

// SetBrand is Brand Setter
// 品牌名称，最多100字符
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// GetBrand Brand Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetBrand() string {
	return r._brand
}

// SetNameE is NameE Setter
// 房源英文名
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// GetNameE NameE Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetNameE() string {
	return r._nameE
}

// SetDecorateTime is DecorateTime Setter
// 装修时间，格式为2015-01-01装修时间
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetDecorateTime(_decorateTime string) error {
	r._decorateTime = _decorateTime
	r.Set("decorate_time", _decorateTime)
	return nil
}

// GetDecorateTime DecorateTime Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetDecorateTime() string {
	return r._decorateTime
}

// SetSettlementCurrency is SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetSettlementCurrency(_settlementCurrency string) error {
	r._settlementCurrency = _settlementCurrency
	r.Set("settlement_currency", _settlementCurrency)
	return nil
}

// GetSettlementCurrency SettlementCurrency Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetSettlementCurrency() string {
	return r._settlementCurrency
}

// SetSurroundInfo is SurroundInfo Setter
// 周边描述
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetSurroundInfo(_surroundInfo string) error {
	r._surroundInfo = _surroundInfo
	r.Set("surround_info", _surroundInfo)
	return nil
}

// GetSurroundInfo SurroundInfo Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetSurroundInfo() string {
	return r._surroundInfo
}

// SetActivitiesAllowed is ActivitiesAllowed Setter
// 详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&amp;docType=1  ，最多500字符
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetActivitiesAllowed(_activitiesAllowed string) error {
	r._activitiesAllowed = _activitiesAllowed
	r.Set("activities_allowed", _activitiesAllowed)
	return nil
}

// GetActivitiesAllowed ActivitiesAllowed Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetActivitiesAllowed() string {
	return r._activitiesAllowed
}

// SetOpeningTime is OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetOpeningTime(_openingTime string) error {
	r._openingTime = _openingTime
	r.Set("opening_time", _openingTime)
	return nil
}

// GetOpeningTime OpeningTime Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetOpeningTime() string {
	return r._openingTime
}

// SetNewOuterId is NewOuterId Setter
// 如果要变更商品房源编码请使用该字段。
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetNewOuterId(_newOuterId string) error {
	r._newOuterId = _newOuterId
	r.Set("new_outer_id", _newOuterId)
	return nil
}

// GetNewOuterId NewOuterId Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetNewOuterId() string {
	return r._newOuterId
}

// SetHouseModel is HouseModel Setter
// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetHouseModel(_houseModel string) error {
	r._houseModel = _houseModel
	r.Set("house_model", _houseModel)
	return nil
}

// GetHouseModel HouseModel Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetHouseModel() string {
	return r._houseModel
}

// SetVideoUrl is VideoUrl Setter
// 视频地址，最多1000字符
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetVideoUrl(_videoUrl string) error {
	r._videoUrl = _videoUrl
	r.Set("video_url", _videoUrl)
	return nil
}

// GetVideoUrl VideoUrl Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetVideoUrl() string {
	return r._videoUrl
}

// SetVendor is Vendor Setter
// 销售渠道,默认taobao
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetVendor() string {
	return r._vendor
}

// SetBrightspot is Brightspot Setter
// 亮点描述，最多1000字符
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetBrightspot(_brightspot string) error {
	r._brightspot = _brightspot
	r.Set("brightspot", _brightspot)
	return nil
}

// GetBrightspot Brightspot Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetBrightspot() string {
	return r._brightspot
}

// SetFloor is Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetFloor(_floor string) error {
	r._floor = _floor
	r.Set("floor", _floor)
	return nil
}

// GetFloor Floor Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetFloor() string {
	return r._floor
}

// SetIntroduction is Introduction Setter
// 房源介绍,最多2000字符
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetIntroduction(_introduction string) error {
	r._introduction = _introduction
	r.Set("introduction", _introduction)
	return nil
}

// GetIntroduction Introduction Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetIntroduction() string {
	return r._introduction
}

// SetCheckInNotes is CheckInNotes Setter
// 入住须知，最多2000字符
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetCheckInNotes(_checkInNotes string) error {
	r._checkInNotes = _checkInNotes
	r.Set("check_in_notes", _checkInNotes)
	return nil
}

// GetCheckInNotes CheckInNotes Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetCheckInNotes() string {
	return r._checkInNotes
}

// SetRealTel is RealTel Setter
// 真实联系方式
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetRealTel(_realTel string) error {
	r._realTel = _realTel
	r.Set("real_tel", _realTel)
	return nil
}

// GetRealTel RealTel Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetRealTel() string {
	return r._realTel
}

// SetService is Service Setter
// 设施服务。json格式示例值：{&#34;24152&#34;:true,&#34;24149&#34;:true,&#34;24150&#34;:true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&amp;docType=1&amp;articleId=108416&amp;previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetService() string {
	return r._service
}

// SetOuterId is OuterId Setter
// 房源id, 这是卖家自己系统中的ID
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetBedInfo is BedInfo Setter
// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&amp;docType=1&amp;articleId=108347
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetBedInfo(_bedInfo string) error {
	r._bedInfo = _bedInfo
	r.Set("bed_info", _bedInfo)
	return nil
}

// GetBedInfo BedInfo Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetBedInfo() string {
	return r._bedInfo
}

// SetExtraFee is ExtraFee Setter
// 额外收费
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetExtraFee(_extraFee string) error {
	r._extraFee = _extraFee
	r.Set("extra_fee", _extraFee)
	return nil
}

// GetExtraFee ExtraFee Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetExtraFee() string {
	return r._extraFee
}

// SetStandardRoomFacilities is StandardRoomFacilities Setter
// 标准酒店服务,参考文档https://fliggy.open.taobao.com/doc.htm?docId=120362&amp;docType=1
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetStandardRoomFacilities(_standardRoomFacilities string) error {
	r._standardRoomFacilities = _standardRoomFacilities
	r.Set("standard_room_facilities", _standardRoomFacilities)
	return nil
}

// GetStandardRoomFacilities StandardRoomFacilities Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetStandardRoomFacilities() string {
	return r._standardRoomFacilities
}

// SetBnbExtend is BnbExtend Setter
// 民宿扩展信息
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetBnbExtend(_bnbExtend string) error {
	r._bnbExtend = _bnbExtend
	r.Set("bnb_extend", _bnbExtend)
	return nil
}

// GetBnbExtend BnbExtend Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetBnbExtend() string {
	return r._bnbExtend
}

// SetRentSize is RentSize Setter
// 单间面积，单位平方米
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetRentSize(_rentSize int64) error {
	r._rentSize = _rentSize
	r.Set("rent_size", _rentSize)
	return nil
}

// GetRentSize RentSize Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetRentSize() int64 {
	return r._rentSize
}

// SetSupportIm is SupportIm Setter
// 是否支持IM聊天 0不支持 1支持
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetSupportIm(_supportIm int64) error {
	r._supportIm = _supportIm
	r.Set("support_im", _supportIm)
	return nil
}

// GetSupportIm SupportIm Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetSupportIm() int64 {
	return r._supportIm
}

// SetCleaningCharge is CleaningCharge Setter
// 清洁费是否收取 0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetCleaningCharge(_cleaningCharge int64) error {
	r._cleaningCharge = _cleaningCharge
	r.Set("cleaning_charge", _cleaningCharge)
	return nil
}

// GetCleaningCharge CleaningCharge Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetCleaningCharge() int64 {
	return r._cleaningCharge
}

// SetInvoice is Invoice Setter
// 发票，0：卖家提供发票，1：房东提供发票
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetInvoice(_invoice int64) error {
	r._invoice = _invoice
	r.Set("invoice", _invoice)
	return nil
}

// GetInvoice Invoice Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetInvoice() int64 {
	return r._invoice
}

// SetDecorateLevel is DecorateLevel Setter
// 装修等级 1 精装；2普通；3简装
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetDecorateLevel(_decorateLevel int64) error {
	r._decorateLevel = _decorateLevel
	r.Set("decorate_level", _decorateLevel)
	return nil
}

// GetDecorateLevel DecorateLevel Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetDecorateLevel() int64 {
	return r._decorateLevel
}

// SetBnbBookingTime is BnbBookingTime Setter
// 民宿入住要求&amp;附加信息
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetBnbBookingTime(_bnbBookingTime *BnbBookingTimeDto) error {
	r._bnbBookingTime = _bnbBookingTime
	r.Set("bnb_booking_time", _bnbBookingTime)
	return nil
}

// GetBnbBookingTime BnbBookingTime Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetBnbBookingTime() *BnbBookingTimeDto {
	return r._bnbBookingTime
}

// SetReceiveForeigners is ReceiveForeigners Setter
// 是否可接待外宾 0：否 1：是；默认值: 0
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetReceiveForeigners(_receiveForeigners int64) error {
	r._receiveForeigners = _receiveForeigners
	r.Set("receive_foreigners", _receiveForeigners)
	return nil
}

// GetReceiveForeigners ReceiveForeigners Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetReceiveForeigners() int64 {
	return r._receiveForeigners
}

// SetCleaningType is CleaningType Setter
// 清洁费类型 0.线下；1.线上
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetCleaningType(_cleaningType int64) error {
	r._cleaningType = _cleaningType
	r.Set("cleaning_type", _cleaningType)
	return nil
}

// GetCleaningType CleaningType Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetCleaningType() int64 {
	return r._cleaningType
}

// SetDepositAmount is DepositAmount Setter
// 押金金额
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetDepositAmount(_depositAmount int64) error {
	r._depositAmount = _depositAmount
	r.Set("deposit_amount", _depositAmount)
	return nil
}

// GetDepositAmount DepositAmount Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetDepositAmount() int64 {
	return r._depositAmount
}

// SetExtraBedsNum is ExtraBedsNum Setter
// 0-n；若不可加床，值为0
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetExtraBedsNum(_extraBedsNum int64) error {
	r._extraBedsNum = _extraBedsNum
	r.Set("extra_beds_num", _extraBedsNum)
	return nil
}

// GetExtraBedsNum ExtraBedsNum Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetExtraBedsNum() int64 {
	return r._extraBedsNum
}

// SetInvoiceType is InvoiceType Setter
// 可提供发票类型，1.专票 2.纸质普票 3.电子普票
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetInvoiceType(_invoiceType int64) error {
	r._invoiceType = _invoiceType
	r.Set("invoice_type", _invoiceType)
	return nil
}

// GetInvoiceType InvoiceType Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetInvoiceType() int64 {
	return r._invoiceType
}

// SetHasFrontDesk is HasFrontDesk Setter
// 是否有前台 0没有 1有
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetHasFrontDesk(_hasFrontDesk int64) error {
	r._hasFrontDesk = _hasFrontDesk
	r.Set("has_front_desk", _hasFrontDesk)
	return nil
}

// GetHasFrontDesk HasFrontDesk Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetHasFrontDesk() int64 {
	return r._hasFrontDesk
}

// SetGuestAge is GuestAge Setter
// 是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&amp;docType=1
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetGuestAge(_guestAge int64) error {
	r._guestAge = _guestAge
	r.Set("guest_age", _guestAge)
	return nil
}

// GetGuestAge GuestAge Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetGuestAge() int64 {
	return r._guestAge
}

// SetMaxOccupancy is MaxOccupancy Setter
// 最大入住人数 1-99
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetMaxOccupancy(_maxOccupancy int64) error {
	r._maxOccupancy = _maxOccupancy
	r.Set("max_occupancy", _maxOccupancy)
	return nil
}

// GetMaxOccupancy MaxOccupancy Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetMaxOccupancy() int64 {
	return r._maxOccupancy
}

// SetIsUseShootImage is IsUseShootImage Setter
// 是否使用实拍图片 0不使用 1使用
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetIsUseShootImage(_isUseShootImage int64) error {
	r._isUseShootImage = _isUseShootImage
	r.Set("is_use_shoot_image", _isUseShootImage)
	return nil
}

// GetIsUseShootImage IsUseShootImage Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetIsUseShootImage() int64 {
	return r._isUseShootImage
}

// SetStatus is Status Setter
// 状态 0：在线 -1：不在线 -2:停售
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetStatus(_status *model.File) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetStatus() *model.File {
	return r._status
}

// SetGuestGender is GuestGender Setter
// 0：不限制，1：只限男性，2：只限女性&#39;
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetGuestGender(_guestGender int64) error {
	r._guestGender = _guestGender
	r.Set("guest_gender", _guestGender)
	return nil
}

// GetGuestGender GuestGender Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetGuestGender() int64 {
	return r._guestGender
}

// SetExtraCleaningCharge is ExtraCleaningCharge Setter
// 清洁费金额；整数[1，9999999]
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetExtraCleaningCharge(_extraCleaningCharge int64) error {
	r._extraCleaningCharge = _extraCleaningCharge
	r.Set("extra_cleaning_charge", _extraCleaningCharge)
	return nil
}

// GetExtraCleaningCharge ExtraCleaningCharge Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetExtraCleaningCharge() int64 {
	return r._extraCleaningCharge
}

// SetRentType is RentType Setter
// 出租类型，1整租；2分租。3床位 默认整租，该字段不能更新
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetRentType(_rentType int64) error {
	r._rentType = _rentType
	r.Set("rent_type", _rentType)
	return nil
}

// GetRentType RentType Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetRentType() int64 {
	return r._rentType
}

// SetWindowType is WindowType Setter
// 0-无窗  1-有窗  2-部分有窗 3-暗窗 4-部分暗窗  5-落地窗
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetWindowType(_windowType int64) error {
	r._windowType = _windowType
	r.Set("window_type", _windowType)
	return nil
}

// GetWindowType WindowType Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetWindowType() int64 {
	return r._windowType
}

// SetHasLicense is HasLicense Setter
// 有无资质执照 0 没有 1有
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetHasLicense(_hasLicense int64) error {
	r._hasLicense = _hasLicense
	r.Set("has_license", _hasLicense)
	return nil
}

// GetHasLicense HasLicense Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetHasLicense() int64 {
	return r._hasLicense
}

// SetQuickOrder is QuickOrder Setter
// 是否开启闪订 0不开启 1开启
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetQuickOrder(_quickOrder int64) error {
	r._quickOrder = _quickOrder
	r.Set("quick_order", _quickOrder)
	return nil
}

// GetQuickOrder QuickOrder Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetQuickOrder() int64 {
	return r._quickOrder
}

// SetHouseSize is HouseSize Setter
// 单间面积，单位平方米
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetHouseSize(_houseSize int64) error {
	r._houseSize = _houseSize
	r.Set("house_size", _houseSize)
	return nil
}

// GetHouseSize HouseSize Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetHouseSize() int64 {
	return r._houseSize
}

// SetProductType is ProductType Setter
// 房源类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&amp;docType=1
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetProductType(_productType int64) error {
	r._productType = _productType
	r.Set("product_type", _productType)
	return nil
}

// GetProductType ProductType Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetProductType() int64 {
	return r._productType
}

// SetHasLandlord is HasLandlord Setter
// 是否与房东同住 0 不同住 1同住
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetHasLandlord(_hasLandlord int64) error {
	r._hasLandlord = _hasLandlord
	r.Set("has_landlord", _hasLandlord)
	return nil
}

// GetHasLandlord HasLandlord Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetHasLandlord() int64 {
	return r._hasLandlord
}

// SetCharge is Charge Setter
// 加人收费信息
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetCharge(_charge *BnbChargeDto) error {
	r._charge = _charge
	r.Set("charge", _charge)
	return nil
}

// GetCharge Charge Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetCharge() *BnbChargeDto {
	return r._charge
}

// SetDecorateStyle is DecorateStyle Setter
// 装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&amp;docType=1
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetDecorateStyle(_decorateStyle int64) error {
	r._decorateStyle = _decorateStyle
	r.Set("decorate_style", _decorateStyle)
	return nil
}

// GetDecorateStyle DecorateStyle Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetDecorateStyle() int64 {
	return r._decorateStyle
}

// SetSupportcredit is Supportcredit Setter
// 是否信用免押金0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetSupportcredit(_supportcredit int64) error {
	r._supportcredit = _supportcredit
	r.Set("supportcredit", _supportcredit)
	return nil
}

// GetSupportcredit Supportcredit Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetSupportcredit() int64 {
	return r._supportcredit
}

// SetCleaningFrequency is CleaningFrequency Setter
// “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&amp;docType=1
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetCleaningFrequency(_cleaningFrequency int64) error {
	r._cleaningFrequency = _cleaningFrequency
	r.Set("cleaning_frequency", _cleaningFrequency)
	return nil
}

// GetCleaningFrequency CleaningFrequency Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetCleaningFrequency() int64 {
	return r._cleaningFrequency
}

// SetLocation is Location Setter
// 民宿房源位置信息
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetLocation(_location *BnbLocationDto) error {
	r._location = _location
	r.Set("location", _location)
	return nil
}

// GetLocation Location Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetLocation() *BnbLocationDto {
	return r._location
}

// SetScenicFeature is ScenicFeature Setter
// 风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&amp;docType=1
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetScenicFeature(_scenicFeature int64) error {
	r._scenicFeature = _scenicFeature
	r.Set("scenic_feature", _scenicFeature)
	return nil
}

// GetScenicFeature ScenicFeature Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetScenicFeature() int64 {
	return r._scenicFeature
}

// SetDepositType is DepositType Setter
// 押金类型0.线下；1.线上
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetDepositType(_depositType int64) error {
	r._depositType = _depositType
	r.Set("deposit_type", _depositType)
	return nil
}

// GetDepositType DepositType Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetDepositType() int64 {
	return r._depositType
}

// SetExtraBedsFee is ExtraBedsFee Setter
// 加床费,分为单位
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetExtraBedsFee(_extraBedsFee int64) error {
	r._extraBedsFee = _extraBedsFee
	r.Set("extra_beds_fee", _extraBedsFee)
	return nil
}

// GetExtraBedsFee ExtraBedsFee Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetExtraBedsFee() int64 {
	return r._extraBedsFee
}

// SetSrid is Srid Setter
// 添加标准房源匹配
func (r *TaobaoXhotelBnbroomtypeAddAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// GetSrid Srid Getter
func (r TaobaoXhotelBnbroomtypeAddAPIRequest) GetSrid() int64 {
	return r._srid
}
