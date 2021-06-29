package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿新增房源 API请求
taobao.xhotel.bnbroomtype.add

添加民宿房源
*/
type TaobaoXhotelBnbroomtypeAddRequest struct {
    model.Params
    // 销售渠道,默认taobao
    _vendor   string
    // 房型id, 这是卖家自己系统中的ID
    _outerId   string
    // 外部门店id
    _outHid   string
    // 房型名
    _name   string
    // 房型英文名
    _nameE   string
    // 房型类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    _productType   int64
    // 有无资质执照 0 没有 1有
    _hasLicense   int64
    // 单间面积，单位平方米
    _houseSize   int64
    // 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
    _floor   string
    // 房型介绍
    _introduction   string
    // 亮点描述
    _brightspot   string
    // 位置描述
    _localInfo   string
    // 周边描述
    _surroundInfo   string
    // 品牌名称
    _brand   string
    // 开业时间，格式为2015-01-01
    _openingTime   string
    // 装修时间，格式为2015-01-01装修时间
    _decorateTime   string
    // 装修等级 1 精装；2普通；3简装
    _decorateLevel   int64
    // 装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    _decorateStyle   int64
    // 风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    _scenicFeature   int64
    // 出租类型，1整租；2分租。3床位 默认整租，该字段不能更新
    _rentType   int64
    // 单间面积，单位平方米
    _rentSize   int64
    // 是否与房东同住 0 不同住 1同住
    _hasLandlord   int64
    // 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
    _tel   string
    // 真实联系方式
    _realTel   string
    // 状态 0：在线 -1：不在线 -2:停售
    _status   *model.File
    // 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
    _settlementCurrency   string
    // 是否支持IM聊天 0不支持 1支持
    _supportIm   int64
    // 是否开启闪订 0不开启 1开启
    _quickOrder   int64
    // 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
    _bedInfo   string
    // 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空
    _houseModel   string
    // 窗型-1.有窗；2.无窗；3.部分有窗
    _windowType   int64
    // 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
    _pics   []BnbPictureDTO
    // 房型外部标签 标签信息，逗号(,)分隔
    _outerTags   string
    // 是否使用实拍图片 0不使用 1使用
    _isUseShootImage   int64
    // 视频地址
    _videoUrl   string
    // 民宿房源位置信息
    _location   *BnbLocationDTO
    // 最大入住人数 1-50
    _maxOccupancy   int64
    // 民宿入住要求&附加信息
    _bnbBookingTime   *BnbBookingTimeDTO
    // 入住须知
    _checkInNotes   string
    // 0：不限制，1：只限男性，2：只限女性'
    _guestGender   int64
    // 是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    _guestAge   int64
    // 是否可接待外宾 0：否 1：是
    _receiveForeigners   int64
    // “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    _cleaningFrequency   int64
    // 详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    _activitiesAllowed   string
    // 0-n；若不可加床，值为0
    _extraBedsNum   int64
    // 押金类型0.线下；1.线上
    _depositType   int64
    // 是否信用免押金0：否 1：是
    _supportcredit   int64
    // 押金金额
    _depositAmount   int64
    // 加人收费信息
    _charge   *BnbChargeDTO
    // 清洁费是否收取 0：否 1：是
    _cleaningCharge   int64
    // 清洁费类型 0.线下；1.线上
    _cleaningType   int64
    // 清洁费金额；整数[1，9999999]
    _extraCleaningCharge   int64
    // 发票，0：卖家提供发票，1：房东提供发票
    _invoice   int64
    // 可提供发票类型，1.专票 2.纸质普票 3.电子普票
    _invoiceType   int64
    // 是否有前台 0没有 1有
    _hasFrontDesk   int64
    // 如果要变更商品房型编码请使用该字段。
    _newOuterId   string
    // 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}，见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
    _service   string
}

// 初始化TaobaoXhotelBnbroomtypeAddRequest对象
func NewTaobaoXhotelBnbroomtypeAddRequest() *TaobaoXhotelBnbroomtypeAddRequest{
    return &TaobaoXhotelBnbroomtypeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbroomtypeAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.bnbroomtype.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbroomtypeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Vendor Setter
// 销售渠道,默认taobao
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetVendor() string {
    return r._vendor
}
// OuterId Setter
// 房型id, 这是卖家自己系统中的ID
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOuterId() string {
    return r._outerId
}
// OutHid Setter
// 外部门店id
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOutHid(_outHid string) error {
    r._outHid = _outHid
    r.Set("out_hid", _outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOutHid() string {
    return r._outHid
}
// Name Setter
// 房型名
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetName() string {
    return r._name
}
// NameE Setter
// 房型英文名
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetNameE(_nameE string) error {
    r._nameE = _nameE
    r.Set("name_e", _nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetNameE() string {
    return r._nameE
}
// ProductType Setter
// 房型类型,见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetProductType(_productType int64) error {
    r._productType = _productType
    r.Set("product_type", _productType)
    return nil
}

// ProductType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetProductType() int64 {
    return r._productType
}
// HasLicense Setter
// 有无资质执照 0 没有 1有
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHasLicense(_hasLicense int64) error {
    r._hasLicense = _hasLicense
    r.Set("has_license", _hasLicense)
    return nil
}

// HasLicense Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHasLicense() int64 {
    return r._hasLicense
}
// HouseSize Setter
// 单间面积，单位平方米
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHouseSize(_houseSize int64) error {
    r._houseSize = _houseSize
    r.Set("house_size", _houseSize)
    return nil
}

// HouseSize Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHouseSize() int64 {
    return r._houseSize
}
// Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetFloor(_floor string) error {
    r._floor = _floor
    r.Set("floor", _floor)
    return nil
}

// Floor Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetFloor() string {
    return r._floor
}
// Introduction Setter
// 房型介绍
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetIntroduction(_introduction string) error {
    r._introduction = _introduction
    r.Set("introduction", _introduction)
    return nil
}

// Introduction Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetIntroduction() string {
    return r._introduction
}
// Brightspot Setter
// 亮点描述
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBrightspot(_brightspot string) error {
    r._brightspot = _brightspot
    r.Set("brightspot", _brightspot)
    return nil
}

// Brightspot Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBrightspot() string {
    return r._brightspot
}
// LocalInfo Setter
// 位置描述
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetLocalInfo(_localInfo string) error {
    r._localInfo = _localInfo
    r.Set("local_info", _localInfo)
    return nil
}

// LocalInfo Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetLocalInfo() string {
    return r._localInfo
}
// SurroundInfo Setter
// 周边描述
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSurroundInfo(_surroundInfo string) error {
    r._surroundInfo = _surroundInfo
    r.Set("surround_info", _surroundInfo)
    return nil
}

// SurroundInfo Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSurroundInfo() string {
    return r._surroundInfo
}
// Brand Setter
// 品牌名称
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBrand(_brand string) error {
    r._brand = _brand
    r.Set("brand", _brand)
    return nil
}

// Brand Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBrand() string {
    return r._brand
}
// OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOpeningTime(_openingTime string) error {
    r._openingTime = _openingTime
    r.Set("opening_time", _openingTime)
    return nil
}

// OpeningTime Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOpeningTime() string {
    return r._openingTime
}
// DecorateTime Setter
// 装修时间，格式为2015-01-01装修时间
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateTime(_decorateTime string) error {
    r._decorateTime = _decorateTime
    r.Set("decorate_time", _decorateTime)
    return nil
}

// DecorateTime Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDecorateTime() string {
    return r._decorateTime
}
// DecorateLevel Setter
// 装修等级 1 精装；2普通；3简装
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateLevel(_decorateLevel int64) error {
    r._decorateLevel = _decorateLevel
    r.Set("decorate_level", _decorateLevel)
    return nil
}

// DecorateLevel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDecorateLevel() int64 {
    return r._decorateLevel
}
// DecorateStyle Setter
// 装修风格https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDecorateStyle(_decorateStyle int64) error {
    r._decorateStyle = _decorateStyle
    r.Set("decorate_style", _decorateStyle)
    return nil
}

// DecorateStyle Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDecorateStyle() int64 {
    return r._decorateStyle
}
// ScenicFeature Setter
// 风景类型(枚举)https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetScenicFeature(_scenicFeature int64) error {
    r._scenicFeature = _scenicFeature
    r.Set("scenic_feature", _scenicFeature)
    return nil
}

// ScenicFeature Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetScenicFeature() int64 {
    return r._scenicFeature
}
// RentType Setter
// 出租类型，1整租；2分租。3床位 默认整租，该字段不能更新
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetRentType(_rentType int64) error {
    r._rentType = _rentType
    r.Set("rent_type", _rentType)
    return nil
}

// RentType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetRentType() int64 {
    return r._rentType
}
// RentSize Setter
// 单间面积，单位平方米
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetRentSize(_rentSize int64) error {
    r._rentSize = _rentSize
    r.Set("rent_size", _rentSize)
    return nil
}

// RentSize Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetRentSize() int64 {
    return r._rentSize
}
// HasLandlord Setter
// 是否与房东同住 0 不同住 1同住
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHasLandlord(_hasLandlord int64) error {
    r._hasLandlord = _hasLandlord
    r.Set("has_landlord", _hasLandlord)
    return nil
}

// HasLandlord Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHasLandlord() int64 {
    return r._hasLandlord
}
// Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetTel(_tel string) error {
    r._tel = _tel
    r.Set("tel", _tel)
    return nil
}

// Tel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetTel() string {
    return r._tel
}
// RealTel Setter
// 真实联系方式
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetRealTel(_realTel string) error {
    r._realTel = _realTel
    r.Set("real_tel", _realTel)
    return nil
}

// RealTel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetRealTel() string {
    return r._realTel
}
// Status Setter
// 状态 0：在线 -1：不在线 -2:停售
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetStatus(_status *model.File) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetStatus() *model.File {
    return r._status
}
// SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSettlementCurrency(_settlementCurrency string) error {
    r._settlementCurrency = _settlementCurrency
    r.Set("settlement_currency", _settlementCurrency)
    return nil
}

// SettlementCurrency Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSettlementCurrency() string {
    return r._settlementCurrency
}
// SupportIm Setter
// 是否支持IM聊天 0不支持 1支持
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSupportIm(_supportIm int64) error {
    r._supportIm = _supportIm
    r.Set("support_im", _supportIm)
    return nil
}

// SupportIm Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSupportIm() int64 {
    return r._supportIm
}
// QuickOrder Setter
// 是否开启闪订 0不开启 1开启
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetQuickOrder(_quickOrder int64) error {
    r._quickOrder = _quickOrder
    r.Set("quick_order", _quickOrder)
    return nil
}

// QuickOrder Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetQuickOrder() int64 {
    return r._quickOrder
}
// BedInfo Setter
// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBedInfo(_bedInfo string) error {
    r._bedInfo = _bedInfo
    r.Set("bed_info", _bedInfo)
    return nil
}

// BedInfo Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBedInfo() string {
    return r._bedInfo
}
// HouseModel Setter
// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房,bedroom和livingroom不能为空
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHouseModel(_houseModel string) error {
    r._houseModel = _houseModel
    r.Set("house_model", _houseModel)
    return nil
}

// HouseModel Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHouseModel() string {
    return r._houseModel
}
// WindowType Setter
// 窗型-1.有窗；2.无窗；3.部分有窗
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetWindowType(_windowType int64) error {
    r._windowType = _windowType
    r.Set("window_type", _windowType)
    return nil
}

// WindowType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetWindowType() int64 {
    return r._windowType
}
// Pics Setter
// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetPics(_pics []BnbPictureDTO) error {
    r._pics = _pics
    r.Set("pics", _pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetPics() []BnbPictureDTO {
    return r._pics
}
// OuterTags Setter
// 房型外部标签 标签信息，逗号(,)分隔
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetOuterTags(_outerTags string) error {
    r._outerTags = _outerTags
    r.Set("outer_tags", _outerTags)
    return nil
}

// OuterTags Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetOuterTags() string {
    return r._outerTags
}
// IsUseShootImage Setter
// 是否使用实拍图片 0不使用 1使用
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetIsUseShootImage(_isUseShootImage int64) error {
    r._isUseShootImage = _isUseShootImage
    r.Set("is_use_shoot_image", _isUseShootImage)
    return nil
}

// IsUseShootImage Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetIsUseShootImage() int64 {
    return r._isUseShootImage
}
// VideoUrl Setter
// 视频地址
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetVideoUrl(_videoUrl string) error {
    r._videoUrl = _videoUrl
    r.Set("video_url", _videoUrl)
    return nil
}

// VideoUrl Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetVideoUrl() string {
    return r._videoUrl
}
// Location Setter
// 民宿房源位置信息
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetLocation(_location *BnbLocationDTO) error {
    r._location = _location
    r.Set("location", _location)
    return nil
}

// Location Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetLocation() *BnbLocationDTO {
    return r._location
}
// MaxOccupancy Setter
// 最大入住人数 1-50
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetMaxOccupancy(_maxOccupancy int64) error {
    r._maxOccupancy = _maxOccupancy
    r.Set("max_occupancy", _maxOccupancy)
    return nil
}

// MaxOccupancy Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetMaxOccupancy() int64 {
    return r._maxOccupancy
}
// BnbBookingTime Setter
// 民宿入住要求&附加信息
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetBnbBookingTime(_bnbBookingTime *BnbBookingTimeDTO) error {
    r._bnbBookingTime = _bnbBookingTime
    r.Set("bnb_booking_time", _bnbBookingTime)
    return nil
}

// BnbBookingTime Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetBnbBookingTime() *BnbBookingTimeDTO {
    return r._bnbBookingTime
}
// CheckInNotes Setter
// 入住须知
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCheckInNotes(_checkInNotes string) error {
    r._checkInNotes = _checkInNotes
    r.Set("check_in_notes", _checkInNotes)
    return nil
}

// CheckInNotes Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCheckInNotes() string {
    return r._checkInNotes
}
// GuestGender Setter
// 0：不限制，1：只限男性，2：只限女性'
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetGuestGender(_guestGender int64) error {
    r._guestGender = _guestGender
    r.Set("guest_gender", _guestGender)
    return nil
}

// GuestGender Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetGuestGender() int64 {
    return r._guestGender
}
// GuestAge Setter
// 是否接待儿童、老人；成年人必接待，详见“可接待客人”https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetGuestAge(_guestAge int64) error {
    r._guestAge = _guestAge
    r.Set("guest_age", _guestAge)
    return nil
}

// GuestAge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetGuestAge() int64 {
    return r._guestAge
}
// ReceiveForeigners Setter
// 是否可接待外宾 0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetReceiveForeigners(_receiveForeigners int64) error {
    r._receiveForeigners = _receiveForeigners
    r.Set("receive_foreigners", _receiveForeigners)
    return nil
}

// ReceiveForeigners Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetReceiveForeigners() int64 {
    return r._receiveForeigners
}
// CleaningFrequency Setter
// “打扫类型1(1客1扫/换),2(1天1扫/换),https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningFrequency(_cleaningFrequency int64) error {
    r._cleaningFrequency = _cleaningFrequency
    r.Set("cleaning_frequency", _cleaningFrequency)
    return nil
}

// CleaningFrequency Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCleaningFrequency() int64 {
    return r._cleaningFrequency
}
// ActivitiesAllowed Setter
// 详见“允许活动”：https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetActivitiesAllowed(_activitiesAllowed string) error {
    r._activitiesAllowed = _activitiesAllowed
    r.Set("activities_allowed", _activitiesAllowed)
    return nil
}

// ActivitiesAllowed Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetActivitiesAllowed() string {
    return r._activitiesAllowed
}
// ExtraBedsNum Setter
// 0-n；若不可加床，值为0
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetExtraBedsNum(_extraBedsNum int64) error {
    r._extraBedsNum = _extraBedsNum
    r.Set("extra_beds_num", _extraBedsNum)
    return nil
}

// ExtraBedsNum Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetExtraBedsNum() int64 {
    return r._extraBedsNum
}
// DepositType Setter
// 押金类型0.线下；1.线上
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDepositType(_depositType int64) error {
    r._depositType = _depositType
    r.Set("deposit_type", _depositType)
    return nil
}

// DepositType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDepositType() int64 {
    return r._depositType
}
// Supportcredit Setter
// 是否信用免押金0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetSupportcredit(_supportcredit int64) error {
    r._supportcredit = _supportcredit
    r.Set("supportcredit", _supportcredit)
    return nil
}

// Supportcredit Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetSupportcredit() int64 {
    return r._supportcredit
}
// DepositAmount Setter
// 押金金额
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetDepositAmount(_depositAmount int64) error {
    r._depositAmount = _depositAmount
    r.Set("deposit_amount", _depositAmount)
    return nil
}

// DepositAmount Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetDepositAmount() int64 {
    return r._depositAmount
}
// Charge Setter
// 加人收费信息
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCharge(_charge *BnbChargeDTO) error {
    r._charge = _charge
    r.Set("charge", _charge)
    return nil
}

// Charge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCharge() *BnbChargeDTO {
    return r._charge
}
// CleaningCharge Setter
// 清洁费是否收取 0：否 1：是
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningCharge(_cleaningCharge int64) error {
    r._cleaningCharge = _cleaningCharge
    r.Set("cleaning_charge", _cleaningCharge)
    return nil
}

// CleaningCharge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCleaningCharge() int64 {
    return r._cleaningCharge
}
// CleaningType Setter
// 清洁费类型 0.线下；1.线上
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetCleaningType(_cleaningType int64) error {
    r._cleaningType = _cleaningType
    r.Set("cleaning_type", _cleaningType)
    return nil
}

// CleaningType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetCleaningType() int64 {
    return r._cleaningType
}
// ExtraCleaningCharge Setter
// 清洁费金额；整数[1，9999999]
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetExtraCleaningCharge(_extraCleaningCharge int64) error {
    r._extraCleaningCharge = _extraCleaningCharge
    r.Set("extra_cleaning_charge", _extraCleaningCharge)
    return nil
}

// ExtraCleaningCharge Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetExtraCleaningCharge() int64 {
    return r._extraCleaningCharge
}
// Invoice Setter
// 发票，0：卖家提供发票，1：房东提供发票
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetInvoice(_invoice int64) error {
    r._invoice = _invoice
    r.Set("invoice", _invoice)
    return nil
}

// Invoice Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetInvoice() int64 {
    return r._invoice
}
// InvoiceType Setter
// 可提供发票类型，1.专票 2.纸质普票 3.电子普票
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetInvoiceType(_invoiceType int64) error {
    r._invoiceType = _invoiceType
    r.Set("invoice_type", _invoiceType)
    return nil
}

// InvoiceType Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetInvoiceType() int64 {
    return r._invoiceType
}
// HasFrontDesk Setter
// 是否有前台 0没有 1有
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetHasFrontDesk(_hasFrontDesk int64) error {
    r._hasFrontDesk = _hasFrontDesk
    r.Set("has_front_desk", _hasFrontDesk)
    return nil
}

// HasFrontDesk Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetHasFrontDesk() int64 {
    return r._hasFrontDesk
}
// NewOuterId Setter
// 如果要变更商品房型编码请使用该字段。
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetNewOuterId(_newOuterId string) error {
    r._newOuterId = _newOuterId
    r.Set("new_outer_id", _newOuterId)
    return nil
}

// NewOuterId Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetNewOuterId() string {
    return r._newOuterId
}
// Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}，见https://fliggy.open.taobao.com/doc.htm?docId=120148&docType=1
func (r *TaobaoXhotelBnbroomtypeAddRequest) SetService(_service string) error {
    r._service = _service
    r.Set("service", _service)
    return nil
}

// Service Getter
func (r TaobaoXhotelBnbroomtypeAddRequest) GetService() string {
    return r._service
}
