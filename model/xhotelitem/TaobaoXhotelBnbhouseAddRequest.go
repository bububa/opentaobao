package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿门店信息添加 API请求
taobao.xhotel.bnbhouse.add

添加和更新民宿门店的信息
*/
type TaobaoXhotelBnbhouseAddRequest struct {
    model.Params
    // 外部房东id
    _outOwnerId   string
    // 对接系统商名称：可为空不要乱填，需要申请后使用
    _vendor   string
    // 供应商渠道门店ID
    _outerId   string
    // 门店名称
    _name   string
    // 门店英文名称
    _nameE   string
    // 门店属性 1 个人房源 2 城市公寓 3 乡村民宿等
    _attributes   int64
    // 门店类型，详见“房源类型list
    _productType   int64
    // 有无资质执照 0 无资质 1有资质
    _hasLicense   int64
    // 面积大小
    _houseSize   int64
    // 楼层
    _floors   string
    // 门店简介
    _description   string
    // 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
    _facilities   string
    // 品牌名称
    _brand   string
    // 开业时间，格式为2015-01-01
    _openingTime   string
    // 装修等级 1 精装 2普通 3简装
    _decorateLevel   int64
    // 装修时间，格式为2015-01-01
    _decorateTime   string
    // 装修风格，详见装修风格枚举表
    _decorateStyle   int64
    // 风景类型，详见风景类型枚举表
    _scenicFeature   int64
    // 房型状态。0:正常，-1:删除，-2:停售
    _status   *model.File
    // 联系方式。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
    _tel   string
    // 真实联系方式
    _realTel   string
    // 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
    _settlementCurrency   string
    // 民宿门店添加
    _pictures   []BnbPictureDTO
    // 门店标签 标签信息，逗号(,)分隔
    _tags   string
    // 是否使用实拍图片 0不使用 1使用
    _isUseShootImage   int64
    // 视频地址
    _videoUrl   string
    // 是否有前台 0没有 1有
    _hasFrontDesk   int64
    // 位置信息
    _location   *BnbLocationDTO
    // 入住要求&附加信息
    _bnbBookingTime   *BnbBookingTimeDTO
    // 入住须知
    _checkInNotes   string
    // 可接待客人性别 0：不限制，1：只限男性，2：只限女性'
    _guestGender   int64
    // 可接待客人年龄情况：是否接待儿童、老人；成年人必接待，详见“可接待客人”list
    _guestAge   int64
    // 是否可接待外宾 0不接待 1接待
    _receiveForeigners   int64
    // 详见“允许活动”list 12,32,33,
    _activitiesAllowed   string
    // 可加床数
    _extraBedsNum   int64
    // 加人收费信息
    _charge   *BnbChargeDTO
}

// 初始化TaobaoXhotelBnbhouseAddRequest对象
func NewTaobaoXhotelBnbhouseAddRequest() *TaobaoXhotelBnbhouseAddRequest{
    return &TaobaoXhotelBnbhouseAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbhouseAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.bnbhouse.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbhouseAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutOwnerId Setter
// 外部房东id
func (r *TaobaoXhotelBnbhouseAddRequest) SetOutOwnerId(_outOwnerId string) error {
    r._outOwnerId = _outOwnerId
    r.Set("out_owner_id", _outOwnerId)
    return nil
}

// OutOwnerId Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetOutOwnerId() string {
    return r._outOwnerId
}
// Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoXhotelBnbhouseAddRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetVendor() string {
    return r._vendor
}
// OuterId Setter
// 供应商渠道门店ID
func (r *TaobaoXhotelBnbhouseAddRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetOuterId() string {
    return r._outerId
}
// Name Setter
// 门店名称
func (r *TaobaoXhotelBnbhouseAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetName() string {
    return r._name
}
// NameE Setter
// 门店英文名称
func (r *TaobaoXhotelBnbhouseAddRequest) SetNameE(_nameE string) error {
    r._nameE = _nameE
    r.Set("name_e", _nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetNameE() string {
    return r._nameE
}
// Attributes Setter
// 门店属性 1 个人房源 2 城市公寓 3 乡村民宿等
func (r *TaobaoXhotelBnbhouseAddRequest) SetAttributes(_attributes int64) error {
    r._attributes = _attributes
    r.Set("attributes", _attributes)
    return nil
}

// Attributes Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetAttributes() int64 {
    return r._attributes
}
// ProductType Setter
// 门店类型，详见“房源类型list
func (r *TaobaoXhotelBnbhouseAddRequest) SetProductType(_productType int64) error {
    r._productType = _productType
    r.Set("product_type", _productType)
    return nil
}

// ProductType Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetProductType() int64 {
    return r._productType
}
// HasLicense Setter
// 有无资质执照 0 无资质 1有资质
func (r *TaobaoXhotelBnbhouseAddRequest) SetHasLicense(_hasLicense int64) error {
    r._hasLicense = _hasLicense
    r.Set("has_license", _hasLicense)
    return nil
}

// HasLicense Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetHasLicense() int64 {
    return r._hasLicense
}
// HouseSize Setter
// 面积大小
func (r *TaobaoXhotelBnbhouseAddRequest) SetHouseSize(_houseSize int64) error {
    r._houseSize = _houseSize
    r.Set("house_size", _houseSize)
    return nil
}

// HouseSize Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetHouseSize() int64 {
    return r._houseSize
}
// Floors Setter
// 楼层
func (r *TaobaoXhotelBnbhouseAddRequest) SetFloors(_floors string) error {
    r._floors = _floors
    r.Set("floors", _floors)
    return nil
}

// Floors Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetFloors() string {
    return r._floors
}
// Description Setter
// 门店简介
func (r *TaobaoXhotelBnbhouseAddRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetDescription() string {
    return r._description
}
// Facilities Setter
// 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
func (r *TaobaoXhotelBnbhouseAddRequest) SetFacilities(_facilities string) error {
    r._facilities = _facilities
    r.Set("facilities", _facilities)
    return nil
}

// Facilities Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetFacilities() string {
    return r._facilities
}
// Brand Setter
// 品牌名称
func (r *TaobaoXhotelBnbhouseAddRequest) SetBrand(_brand string) error {
    r._brand = _brand
    r.Set("brand", _brand)
    return nil
}

// Brand Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetBrand() string {
    return r._brand
}
// OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelBnbhouseAddRequest) SetOpeningTime(_openingTime string) error {
    r._openingTime = _openingTime
    r.Set("opening_time", _openingTime)
    return nil
}

// OpeningTime Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetOpeningTime() string {
    return r._openingTime
}
// DecorateLevel Setter
// 装修等级 1 精装 2普通 3简装
func (r *TaobaoXhotelBnbhouseAddRequest) SetDecorateLevel(_decorateLevel int64) error {
    r._decorateLevel = _decorateLevel
    r.Set("decorate_level", _decorateLevel)
    return nil
}

// DecorateLevel Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetDecorateLevel() int64 {
    return r._decorateLevel
}
// DecorateTime Setter
// 装修时间，格式为2015-01-01
func (r *TaobaoXhotelBnbhouseAddRequest) SetDecorateTime(_decorateTime string) error {
    r._decorateTime = _decorateTime
    r.Set("decorate_time", _decorateTime)
    return nil
}

// DecorateTime Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetDecorateTime() string {
    return r._decorateTime
}
// DecorateStyle Setter
// 装修风格，详见装修风格枚举表
func (r *TaobaoXhotelBnbhouseAddRequest) SetDecorateStyle(_decorateStyle int64) error {
    r._decorateStyle = _decorateStyle
    r.Set("decorate_style", _decorateStyle)
    return nil
}

// DecorateStyle Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetDecorateStyle() int64 {
    return r._decorateStyle
}
// ScenicFeature Setter
// 风景类型，详见风景类型枚举表
func (r *TaobaoXhotelBnbhouseAddRequest) SetScenicFeature(_scenicFeature int64) error {
    r._scenicFeature = _scenicFeature
    r.Set("scenic_feature", _scenicFeature)
    return nil
}

// ScenicFeature Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetScenicFeature() int64 {
    return r._scenicFeature
}
// Status Setter
// 房型状态。0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelBnbhouseAddRequest) SetStatus(_status *model.File) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetStatus() *model.File {
    return r._status
}
// Tel Setter
// 联系方式。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelBnbhouseAddRequest) SetTel(_tel string) error {
    r._tel = _tel
    r.Set("tel", _tel)
    return nil
}

// Tel Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetTel() string {
    return r._tel
}
// RealTel Setter
// 真实联系方式
func (r *TaobaoXhotelBnbhouseAddRequest) SetRealTel(_realTel string) error {
    r._realTel = _realTel
    r.Set("real_tel", _realTel)
    return nil
}

// RealTel Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetRealTel() string {
    return r._realTel
}
// SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelBnbhouseAddRequest) SetSettlementCurrency(_settlementCurrency string) error {
    r._settlementCurrency = _settlementCurrency
    r.Set("settlement_currency", _settlementCurrency)
    return nil
}

// SettlementCurrency Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetSettlementCurrency() string {
    return r._settlementCurrency
}
// Pictures Setter
// 民宿门店添加
func (r *TaobaoXhotelBnbhouseAddRequest) SetPictures(_pictures []BnbPictureDTO) error {
    r._pictures = _pictures
    r.Set("pictures", _pictures)
    return nil
}

// Pictures Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetPictures() []BnbPictureDTO {
    return r._pictures
}
// Tags Setter
// 门店标签 标签信息，逗号(,)分隔
func (r *TaobaoXhotelBnbhouseAddRequest) SetTags(_tags string) error {
    r._tags = _tags
    r.Set("tags", _tags)
    return nil
}

// Tags Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetTags() string {
    return r._tags
}
// IsUseShootImage Setter
// 是否使用实拍图片 0不使用 1使用
func (r *TaobaoXhotelBnbhouseAddRequest) SetIsUseShootImage(_isUseShootImage int64) error {
    r._isUseShootImage = _isUseShootImage
    r.Set("is_use_shoot_image", _isUseShootImage)
    return nil
}

// IsUseShootImage Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetIsUseShootImage() int64 {
    return r._isUseShootImage
}
// VideoUrl Setter
// 视频地址
func (r *TaobaoXhotelBnbhouseAddRequest) SetVideoUrl(_videoUrl string) error {
    r._videoUrl = _videoUrl
    r.Set("video_url", _videoUrl)
    return nil
}

// VideoUrl Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetVideoUrl() string {
    return r._videoUrl
}
// HasFrontDesk Setter
// 是否有前台 0没有 1有
func (r *TaobaoXhotelBnbhouseAddRequest) SetHasFrontDesk(_hasFrontDesk int64) error {
    r._hasFrontDesk = _hasFrontDesk
    r.Set("has_front_desk", _hasFrontDesk)
    return nil
}

// HasFrontDesk Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetHasFrontDesk() int64 {
    return r._hasFrontDesk
}
// Location Setter
// 位置信息
func (r *TaobaoXhotelBnbhouseAddRequest) SetLocation(_location *BnbLocationDTO) error {
    r._location = _location
    r.Set("location", _location)
    return nil
}

// Location Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetLocation() *BnbLocationDTO {
    return r._location
}
// BnbBookingTime Setter
// 入住要求&附加信息
func (r *TaobaoXhotelBnbhouseAddRequest) SetBnbBookingTime(_bnbBookingTime *BnbBookingTimeDTO) error {
    r._bnbBookingTime = _bnbBookingTime
    r.Set("bnb_booking_time", _bnbBookingTime)
    return nil
}

// BnbBookingTime Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetBnbBookingTime() *BnbBookingTimeDTO {
    return r._bnbBookingTime
}
// CheckInNotes Setter
// 入住须知
func (r *TaobaoXhotelBnbhouseAddRequest) SetCheckInNotes(_checkInNotes string) error {
    r._checkInNotes = _checkInNotes
    r.Set("check_in_notes", _checkInNotes)
    return nil
}

// CheckInNotes Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetCheckInNotes() string {
    return r._checkInNotes
}
// GuestGender Setter
// 可接待客人性别 0：不限制，1：只限男性，2：只限女性'
func (r *TaobaoXhotelBnbhouseAddRequest) SetGuestGender(_guestGender int64) error {
    r._guestGender = _guestGender
    r.Set("guest_gender", _guestGender)
    return nil
}

// GuestGender Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetGuestGender() int64 {
    return r._guestGender
}
// GuestAge Setter
// 可接待客人年龄情况：是否接待儿童、老人；成年人必接待，详见“可接待客人”list
func (r *TaobaoXhotelBnbhouseAddRequest) SetGuestAge(_guestAge int64) error {
    r._guestAge = _guestAge
    r.Set("guest_age", _guestAge)
    return nil
}

// GuestAge Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetGuestAge() int64 {
    return r._guestAge
}
// ReceiveForeigners Setter
// 是否可接待外宾 0不接待 1接待
func (r *TaobaoXhotelBnbhouseAddRequest) SetReceiveForeigners(_receiveForeigners int64) error {
    r._receiveForeigners = _receiveForeigners
    r.Set("receive_foreigners", _receiveForeigners)
    return nil
}

// ReceiveForeigners Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetReceiveForeigners() int64 {
    return r._receiveForeigners
}
// ActivitiesAllowed Setter
// 详见“允许活动”list 12,32,33,
func (r *TaobaoXhotelBnbhouseAddRequest) SetActivitiesAllowed(_activitiesAllowed string) error {
    r._activitiesAllowed = _activitiesAllowed
    r.Set("activities_allowed", _activitiesAllowed)
    return nil
}

// ActivitiesAllowed Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetActivitiesAllowed() string {
    return r._activitiesAllowed
}
// ExtraBedsNum Setter
// 可加床数
func (r *TaobaoXhotelBnbhouseAddRequest) SetExtraBedsNum(_extraBedsNum int64) error {
    r._extraBedsNum = _extraBedsNum
    r.Set("extra_beds_num", _extraBedsNum)
    return nil
}

// ExtraBedsNum Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetExtraBedsNum() int64 {
    return r._extraBedsNum
}
// Charge Setter
// 加人收费信息
func (r *TaobaoXhotelBnbhouseAddRequest) SetCharge(_charge *BnbChargeDTO) error {
    r._charge = _charge
    r.Set("charge", _charge)
    return nil
}

// Charge Getter
func (r TaobaoXhotelBnbhouseAddRequest) GetCharge() *BnbChargeDTO {
    return r._charge
}
