package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿门店信息添加 APIRequest
taobao.xhotel.bnbhouse.add

添加和更新民宿门店的信息
*/
type TaobaoXhotelBnbhouseAddRequest struct {
    model.Params

    // 外部房东id
    outOwnerId   string 

    // 对接系统商名称：可为空不要乱填，需要申请后使用
    vendor   string 

    // 供应商渠道门店ID
    outerId   string 

    // 门店名称
    name   string 

    // 门店英文名称
    nameE   string 

    // 门店属性 1 个人房源 2 城市公寓 3 乡村民宿等
    attributes   int64 

    // 门店类型，详见“房源类型list
    productType   int64 

    // 有无资质执照 0 无资质 1有资质
    hasLicense   int64 

    // 面积大小
    houseSize   int64 

    // 楼层
    floors   string 

    // 门店简介
    description   string 

    // 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
    facilities   string 

    // 品牌名称
    brand   string 

    // 开业时间，格式为2015-01-01
    openingTime   string 

    // 装修等级 1 精装 2普通 3简装
    decorateLevel   int64 

    // 装修时间，格式为2015-01-01
    decorateTime   string 

    // 装修风格，详见装修风格枚举表
    decorateStyle   int64 

    // 风景类型，详见风景类型枚举表
    scenicFeature   int64 

    // 房型状态。0:正常，-1:删除，-2:停售
    status   *model.File 

    // 联系方式。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
    tel   string 

    // 真实联系方式
    realTel   string 

    // 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
    settlementCurrency   string 

    // 民宿门店添加
    pictures   []BnbPictureDTO 

    // 门店标签 标签信息，逗号(,)分隔
    tags   string 

    // 是否使用实拍图片 0不使用 1使用
    isUseShootImage   int64 

    // 视频地址
    videoUrl   string 

    // 是否有前台 0没有 1有
    hasFrontDesk   int64 

    // 位置信息
    location   *BnbLocationDto 

    // 入住要求&附加信息
    bnbBookingTime   *BnbBookingTimeDto 

    // 入住须知
    checkInNotes   string 

    // 可接待客人性别 0：不限制，1：只限男性，2：只限女性'
    guestGender   int64 

    // 可接待客人年龄情况：是否接待儿童、老人；成年人必接待，详见“可接待客人”list
    guestAge   int64 

    // 是否可接待外宾 0不接待 1接待
    receiveForeigners   int64 

    // 详见“允许活动”list 12,32,33,
    activitiesAllowed   string 

    // 可加床数
    extraBedsNum   int64 

    // 加人收费信息
    charge   *BnbChargeDto 

}

func NewTaobaoXhotelBnbhouseAddRequest() *TaobaoXhotelBnbhouseAddRequest{
    return &TaobaoXhotelBnbhouseAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelBnbhouseAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.bnbhouse.add"
}

func (r TaobaoXhotelBnbhouseAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelBnbhouseAddRequest) SetOutOwnerId(outOwnerId string) error {
    r.outOwnerId = outOwnerId
    r.Set("out_owner_id", outOwnerId)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetOutOwnerId() string {
    return r.outOwnerId
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetVendor() string {
    return r.vendor
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetOuterId() string {
    return r.outerId
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetName() string {
    return r.name
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetNameE(nameE string) error {
    r.nameE = nameE
    r.Set("name_e", nameE)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetNameE() string {
    return r.nameE
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetAttributes(attributes int64) error {
    r.attributes = attributes
    r.Set("attributes", attributes)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetAttributes() int64 {
    return r.attributes
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetProductType(productType int64) error {
    r.productType = productType
    r.Set("product_type", productType)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetProductType() int64 {
    return r.productType
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetHasLicense(hasLicense int64) error {
    r.hasLicense = hasLicense
    r.Set("has_license", hasLicense)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetHasLicense() int64 {
    return r.hasLicense
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetHouseSize(houseSize int64) error {
    r.houseSize = houseSize
    r.Set("house_size", houseSize)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetHouseSize() int64 {
    return r.houseSize
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetFloors(floors string) error {
    r.floors = floors
    r.Set("floors", floors)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetFloors() string {
    return r.floors
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetDescription() string {
    return r.description
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetFacilities(facilities string) error {
    r.facilities = facilities
    r.Set("facilities", facilities)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetFacilities() string {
    return r.facilities
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetBrand() string {
    return r.brand
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetOpeningTime(openingTime string) error {
    r.openingTime = openingTime
    r.Set("opening_time", openingTime)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetOpeningTime() string {
    return r.openingTime
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetDecorateLevel(decorateLevel int64) error {
    r.decorateLevel = decorateLevel
    r.Set("decorate_level", decorateLevel)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetDecorateLevel() int64 {
    return r.decorateLevel
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetDecorateTime(decorateTime string) error {
    r.decorateTime = decorateTime
    r.Set("decorate_time", decorateTime)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetDecorateTime() string {
    return r.decorateTime
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetDecorateStyle(decorateStyle int64) error {
    r.decorateStyle = decorateStyle
    r.Set("decorate_style", decorateStyle)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetDecorateStyle() int64 {
    return r.decorateStyle
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetScenicFeature(scenicFeature int64) error {
    r.scenicFeature = scenicFeature
    r.Set("scenic_feature", scenicFeature)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetScenicFeature() int64 {
    return r.scenicFeature
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetStatus(status *model.File) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetStatus() *model.File {
    return r.status
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetTel(tel string) error {
    r.tel = tel
    r.Set("tel", tel)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetTel() string {
    return r.tel
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetRealTel(realTel string) error {
    r.realTel = realTel
    r.Set("real_tel", realTel)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetRealTel() string {
    return r.realTel
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetSettlementCurrency(settlementCurrency string) error {
    r.settlementCurrency = settlementCurrency
    r.Set("settlement_currency", settlementCurrency)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetSettlementCurrency() string {
    return r.settlementCurrency
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetPictures(pictures []BnbPictureDTO) error {
    r.pictures = pictures
    r.Set("pictures", pictures)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetPictures() []BnbPictureDTO {
    return r.pictures
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetTags(tags string) error {
    r.tags = tags
    r.Set("tags", tags)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetTags() string {
    return r.tags
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetIsUseShootImage(isUseShootImage int64) error {
    r.isUseShootImage = isUseShootImage
    r.Set("is_use_shoot_image", isUseShootImage)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetIsUseShootImage() int64 {
    return r.isUseShootImage
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetVideoUrl(videoUrl string) error {
    r.videoUrl = videoUrl
    r.Set("video_url", videoUrl)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetVideoUrl() string {
    return r.videoUrl
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetHasFrontDesk(hasFrontDesk int64) error {
    r.hasFrontDesk = hasFrontDesk
    r.Set("has_front_desk", hasFrontDesk)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetHasFrontDesk() int64 {
    return r.hasFrontDesk
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetLocation(location *BnbLocationDto) error {
    r.location = location
    r.Set("location", location)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetLocation() *BnbLocationDto {
    return r.location
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetBnbBookingTime(bnbBookingTime *BnbBookingTimeDto) error {
    r.bnbBookingTime = bnbBookingTime
    r.Set("bnb_booking_time", bnbBookingTime)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetBnbBookingTime() *BnbBookingTimeDto {
    return r.bnbBookingTime
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetCheckInNotes(checkInNotes string) error {
    r.checkInNotes = checkInNotes
    r.Set("check_in_notes", checkInNotes)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetCheckInNotes() string {
    return r.checkInNotes
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetGuestGender(guestGender int64) error {
    r.guestGender = guestGender
    r.Set("guest_gender", guestGender)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetGuestGender() int64 {
    return r.guestGender
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetGuestAge(guestAge int64) error {
    r.guestAge = guestAge
    r.Set("guest_age", guestAge)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetGuestAge() int64 {
    return r.guestAge
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetReceiveForeigners(receiveForeigners int64) error {
    r.receiveForeigners = receiveForeigners
    r.Set("receive_foreigners", receiveForeigners)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetReceiveForeigners() int64 {
    return r.receiveForeigners
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetActivitiesAllowed(activitiesAllowed string) error {
    r.activitiesAllowed = activitiesAllowed
    r.Set("activities_allowed", activitiesAllowed)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetActivitiesAllowed() string {
    return r.activitiesAllowed
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetExtraBedsNum(extraBedsNum int64) error {
    r.extraBedsNum = extraBedsNum
    r.Set("extra_beds_num", extraBedsNum)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetExtraBedsNum() int64 {
    return r.extraBedsNum
}

func (r *TaobaoXhotelBnbhouseAddRequest) SetCharge(charge *BnbChargeDto) error {
    r.charge = charge
    r.Set("charge", charge)
    return nil
}

func (r TaobaoXhotelBnbhouseAddRequest) GetCharge() *BnbChargeDto {
    return r.charge
}

