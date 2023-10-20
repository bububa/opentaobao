package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelbnbhouseaddAPIRequest 民宿门店信息添加 API请求
// taobao.xhotel.bnbhouse.add
//
// 添加和更新民宿门店的信息
type TaobaoxhotelbnbhouseaddAPIRequest struct {
	model.Params
	// 民宿门店添加
	_pictures []BnbPictureDto
	// 外部房东id
	_outOwnerId string
	// 联系方式。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
	_tel string
	// 门店英文名称
	_nameE string
	// 装修时间，格式为2015-01-01
	_decorateTime string
	// 门店标签 标签信息，逗号(,)分隔
	_tags string
	// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
	_settlementCurrency string
	// 门店名称
	_name string
	// 详见“允许活动”list 12,32,33,
	_activitiesAllowed string
	// 开业时间，格式为2015-01-01
	_openingTime string
	// 门店简介
	_description string
	// 楼层
	_floors string
	// 视频地址
	_videoUrl string
	// 对接系统商名称：可为空不要乱填，需要申请后使用
	_vendor string
	// 入住须知
	_checkInNotes string
	// 真实联系方式
	_realTel string
	// 供应商渠道门店ID
	_outerId string
	// 品牌名称
	_brand string
	// 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
	_facilities string
	// 标准酒店服务,参考文档https://fliggy.open.taobao.com/doc.htm?docId=120362&docType=1
	_standardHotelFacilities string
	// 入住要求&附加信息
	_bnbBookingTime *BnbBookingTimeDto
	// 装修等级 1 精装 2普通 3简装
	_decorateLevel int64
	// 是否可接待外宾 0不接待 1接待,默认值: 0
	_receiveForeigners int64
	// 可加床数
	_extraBedsNum int64
	// 是否有前台 0没有 1有
	_hasFrontDesk int64
	// 可接待客人年龄情况：是否接待儿童、老人；成年人必接待，详见“可接待客人”list
	_guestAge int64
	// 是否使用实拍图片 0不使用 1使用
	_isUseShootImage int64
	// 房型状态。0:正常，-1:删除，-2:停售
	_status *model.File
	// 可接待客人性别 0：不限制，1：只限男性，2：只限女性'
	_guestGender int64
	// 有无资质执照 0 无资质 1有资质
	_hasLicense int64
	// 门店类型，详见“房源类型list
	_productType int64
	// 加人收费信息
	_charge *BnbChargeDto
	// 装修风格，详见装修风格枚举表
	_decorateStyle int64
	// 位置信息
	_location *BnbLocationDto
	// 门店属性 1 个人房源 2 城市公寓 3 乡村民宿等
	_attributes int64
	// 风景类型，详见风景类型枚举表
	_scenicFeature int64
	// 面积大小
	_houseSize int64
	// 匹配的标准门店
	_shid int64
	// 作为菲住门店签约的佣金比率,范围为10-50
	_commissionRate int64
	// 传入是或者否，是表明为菲住合作模式，hid打标；“否”表示为正常合作模式，hid去标；不传保持原有的合作模式不变
	_isFeizhuHotel bool
}

// NewTaobaoxhotelbnbhouseaddRequest 初始化TaobaoxhotelbnbhouseaddAPIRequest对象
func NewTaobaoxhotelbnbhouseaddRequest() *TaobaoxhotelbnbhouseaddAPIRequest {
	return &TaobaoxhotelbnbhouseaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbhouse.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPictures is Pictures Setter
// 民宿门店添加
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetPictures(_pictures []BnbPictureDto) error {
	r._pictures = _pictures
	r.Set("pictures", _pictures)
	return nil
}

// GetPictures Pictures Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetPictures() []BnbPictureDto {
	return r._pictures
}

// SetOutOwnerId is OutOwnerId Setter
// 外部房东id
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetOutOwnerId(_outOwnerId string) error {
	r._outOwnerId = _outOwnerId
	r.Set("out_owner_id", _outOwnerId)
	return nil
}

// GetOutOwnerId OutOwnerId Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetOutOwnerId() string {
	return r._outOwnerId
}

// SetTel is Tel Setter
// 联系方式。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetTel(_tel string) error {
	r._tel = _tel
	r.Set("tel", _tel)
	return nil
}

// GetTel Tel Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetTel() string {
	return r._tel
}

// SetNameE is NameE Setter
// 门店英文名称
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// GetNameE NameE Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetNameE() string {
	return r._nameE
}

// SetDecorateTime is DecorateTime Setter
// 装修时间，格式为2015-01-01
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetDecorateTime(_decorateTime string) error {
	r._decorateTime = _decorateTime
	r.Set("decorate_time", _decorateTime)
	return nil
}

// GetDecorateTime DecorateTime Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetDecorateTime() string {
	return r._decorateTime
}

// SetTags is Tags Setter
// 门店标签 标签信息，逗号(,)分隔
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetTags(_tags string) error {
	r._tags = _tags
	r.Set("tags", _tags)
	return nil
}

// GetTags Tags Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetTags() string {
	return r._tags
}

// SetSettlementCurrency is SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetSettlementCurrency(_settlementCurrency string) error {
	r._settlementCurrency = _settlementCurrency
	r.Set("settlement_currency", _settlementCurrency)
	return nil
}

// GetSettlementCurrency SettlementCurrency Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetSettlementCurrency() string {
	return r._settlementCurrency
}

// SetName is Name Setter
// 门店名称
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetName() string {
	return r._name
}

// SetActivitiesAllowed is ActivitiesAllowed Setter
// 详见“允许活动”list 12,32,33,
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetActivitiesAllowed(_activitiesAllowed string) error {
	r._activitiesAllowed = _activitiesAllowed
	r.Set("activities_allowed", _activitiesAllowed)
	return nil
}

// GetActivitiesAllowed ActivitiesAllowed Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetActivitiesAllowed() string {
	return r._activitiesAllowed
}

// SetOpeningTime is OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetOpeningTime(_openingTime string) error {
	r._openingTime = _openingTime
	r.Set("opening_time", _openingTime)
	return nil
}

// GetOpeningTime OpeningTime Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetOpeningTime() string {
	return r._openingTime
}

// SetDescription is Description Setter
// 门店简介
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetDescription() string {
	return r._description
}

// SetFloors is Floors Setter
// 楼层
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetFloors(_floors string) error {
	r._floors = _floors
	r.Set("floors", _floors)
	return nil
}

// GetFloors Floors Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetFloors() string {
	return r._floors
}

// SetVideoUrl is VideoUrl Setter
// 视频地址
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetVideoUrl(_videoUrl string) error {
	r._videoUrl = _videoUrl
	r.Set("video_url", _videoUrl)
	return nil
}

// GetVideoUrl VideoUrl Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetVideoUrl() string {
	return r._videoUrl
}

// SetVendor is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetVendor() string {
	return r._vendor
}

// SetCheckInNotes is CheckInNotes Setter
// 入住须知
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetCheckInNotes(_checkInNotes string) error {
	r._checkInNotes = _checkInNotes
	r.Set("check_in_notes", _checkInNotes)
	return nil
}

// GetCheckInNotes CheckInNotes Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetCheckInNotes() string {
	return r._checkInNotes
}

// SetRealTel is RealTel Setter
// 真实联系方式
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetRealTel(_realTel string) error {
	r._realTel = _realTel
	r.Set("real_tel", _realTel)
	return nil
}

// GetRealTel RealTel Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetRealTel() string {
	return r._realTel
}

// SetOuterId is OuterId Setter
// 供应商渠道门店ID
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetBrand is Brand Setter
// 品牌名称
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// GetBrand Brand Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetBrand() string {
	return r._brand
}

// SetFacilities is Facilities Setter
// 酒店设施。json格式示例值：{&#34;24152&#34;:true,&#34;24149&#34;:true,&#34;24150&#34;:true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&amp;docType=1&amp;articleId=108416&amp;previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetFacilities(_facilities string) error {
	r._facilities = _facilities
	r.Set("facilities", _facilities)
	return nil
}

// GetFacilities Facilities Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetFacilities() string {
	return r._facilities
}

// SetStandardHotelFacilities is StandardHotelFacilities Setter
// 标准酒店服务,参考文档https://fliggy.open.taobao.com/doc.htm?docId=120362&amp;docType=1
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetStandardHotelFacilities(_standardHotelFacilities string) error {
	r._standardHotelFacilities = _standardHotelFacilities
	r.Set("standard_hotel_facilities", _standardHotelFacilities)
	return nil
}

// GetStandardHotelFacilities StandardHotelFacilities Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetStandardHotelFacilities() string {
	return r._standardHotelFacilities
}

// SetBnbBookingTime is BnbBookingTime Setter
// 入住要求&amp;附加信息
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetBnbBookingTime(_bnbBookingTime *BnbBookingTimeDto) error {
	r._bnbBookingTime = _bnbBookingTime
	r.Set("bnb_booking_time", _bnbBookingTime)
	return nil
}

// GetBnbBookingTime BnbBookingTime Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetBnbBookingTime() *BnbBookingTimeDto {
	return r._bnbBookingTime
}

// SetDecorateLevel is DecorateLevel Setter
// 装修等级 1 精装 2普通 3简装
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetDecorateLevel(_decorateLevel int64) error {
	r._decorateLevel = _decorateLevel
	r.Set("decorate_level", _decorateLevel)
	return nil
}

// GetDecorateLevel DecorateLevel Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetDecorateLevel() int64 {
	return r._decorateLevel
}

// SetReceiveForeigners is ReceiveForeigners Setter
// 是否可接待外宾 0不接待 1接待,默认值: 0
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetReceiveForeigners(_receiveForeigners int64) error {
	r._receiveForeigners = _receiveForeigners
	r.Set("receive_foreigners", _receiveForeigners)
	return nil
}

// GetReceiveForeigners ReceiveForeigners Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetReceiveForeigners() int64 {
	return r._receiveForeigners
}

// SetExtraBedsNum is ExtraBedsNum Setter
// 可加床数
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetExtraBedsNum(_extraBedsNum int64) error {
	r._extraBedsNum = _extraBedsNum
	r.Set("extra_beds_num", _extraBedsNum)
	return nil
}

// GetExtraBedsNum ExtraBedsNum Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetExtraBedsNum() int64 {
	return r._extraBedsNum
}

// SetHasFrontDesk is HasFrontDesk Setter
// 是否有前台 0没有 1有
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetHasFrontDesk(_hasFrontDesk int64) error {
	r._hasFrontDesk = _hasFrontDesk
	r.Set("has_front_desk", _hasFrontDesk)
	return nil
}

// GetHasFrontDesk HasFrontDesk Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetHasFrontDesk() int64 {
	return r._hasFrontDesk
}

// SetGuestAge is GuestAge Setter
// 可接待客人年龄情况：是否接待儿童、老人；成年人必接待，详见“可接待客人”list
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetGuestAge(_guestAge int64) error {
	r._guestAge = _guestAge
	r.Set("guest_age", _guestAge)
	return nil
}

// GetGuestAge GuestAge Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetGuestAge() int64 {
	return r._guestAge
}

// SetIsUseShootImage is IsUseShootImage Setter
// 是否使用实拍图片 0不使用 1使用
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetIsUseShootImage(_isUseShootImage int64) error {
	r._isUseShootImage = _isUseShootImage
	r.Set("is_use_shoot_image", _isUseShootImage)
	return nil
}

// GetIsUseShootImage IsUseShootImage Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetIsUseShootImage() int64 {
	return r._isUseShootImage
}

// SetStatus is Status Setter
// 房型状态。0:正常，-1:删除，-2:停售
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetStatus(_status *model.File) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetStatus() *model.File {
	return r._status
}

// SetGuestGender is GuestGender Setter
// 可接待客人性别 0：不限制，1：只限男性，2：只限女性&#39;
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetGuestGender(_guestGender int64) error {
	r._guestGender = _guestGender
	r.Set("guest_gender", _guestGender)
	return nil
}

// GetGuestGender GuestGender Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetGuestGender() int64 {
	return r._guestGender
}

// SetHasLicense is HasLicense Setter
// 有无资质执照 0 无资质 1有资质
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetHasLicense(_hasLicense int64) error {
	r._hasLicense = _hasLicense
	r.Set("has_license", _hasLicense)
	return nil
}

// GetHasLicense HasLicense Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetHasLicense() int64 {
	return r._hasLicense
}

// SetProductType is ProductType Setter
// 门店类型，详见“房源类型list
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetProductType(_productType int64) error {
	r._productType = _productType
	r.Set("product_type", _productType)
	return nil
}

// GetProductType ProductType Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetProductType() int64 {
	return r._productType
}

// SetCharge is Charge Setter
// 加人收费信息
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetCharge(_charge *BnbChargeDto) error {
	r._charge = _charge
	r.Set("charge", _charge)
	return nil
}

// GetCharge Charge Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetCharge() *BnbChargeDto {
	return r._charge
}

// SetDecorateStyle is DecorateStyle Setter
// 装修风格，详见装修风格枚举表
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetDecorateStyle(_decorateStyle int64) error {
	r._decorateStyle = _decorateStyle
	r.Set("decorate_style", _decorateStyle)
	return nil
}

// GetDecorateStyle DecorateStyle Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetDecorateStyle() int64 {
	return r._decorateStyle
}

// SetLocation is Location Setter
// 位置信息
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetLocation(_location *BnbLocationDto) error {
	r._location = _location
	r.Set("location", _location)
	return nil
}

// GetLocation Location Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetLocation() *BnbLocationDto {
	return r._location
}

// SetAttributes is Attributes Setter
// 门店属性 1 个人房源 2 城市公寓 3 乡村民宿等
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetAttributes(_attributes int64) error {
	r._attributes = _attributes
	r.Set("attributes", _attributes)
	return nil
}

// GetAttributes Attributes Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetAttributes() int64 {
	return r._attributes
}

// SetScenicFeature is ScenicFeature Setter
// 风景类型，详见风景类型枚举表
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetScenicFeature(_scenicFeature int64) error {
	r._scenicFeature = _scenicFeature
	r.Set("scenic_feature", _scenicFeature)
	return nil
}

// GetScenicFeature ScenicFeature Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetScenicFeature() int64 {
	return r._scenicFeature
}

// SetHouseSize is HouseSize Setter
// 面积大小
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetHouseSize(_houseSize int64) error {
	r._houseSize = _houseSize
	r.Set("house_size", _houseSize)
	return nil
}

// GetHouseSize HouseSize Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetHouseSize() int64 {
	return r._houseSize
}

// SetShid is Shid Setter
// 匹配的标准门店
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetShid() int64 {
	return r._shid
}

// SetCommissionRate is CommissionRate Setter
// 作为菲住门店签约的佣金比率,范围为10-50
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetCommissionRate(_commissionRate int64) error {
	r._commissionRate = _commissionRate
	r.Set("commission_rate", _commissionRate)
	return nil
}

// GetCommissionRate CommissionRate Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetCommissionRate() int64 {
	return r._commissionRate
}

// SetIsFeizhuHotel is IsFeizhuHotel Setter
// 传入是或者否，是表明为菲住合作模式，hid打标；“否”表示为正常合作模式，hid去标；不传保持原有的合作模式不变
func (r *TaobaoxhotelbnbhouseaddAPIRequest) SetIsFeizhuHotel(_isFeizhuHotel bool) error {
	r._isFeizhuHotel = _isFeizhuHotel
	r.Set("is_feizhu_hotel", _isFeizhuHotel)
	return nil
}

// GetIsFeizhuHotel IsFeizhuHotel Getter
func (r TaobaoxhotelbnbhouseaddAPIRequest) GetIsFeizhuHotel() bool {
	return r._isFeizhuHotel
}
