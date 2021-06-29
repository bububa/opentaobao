package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
酒店更新接口（ID不存在自动新增） API请求
taobao.xhotel.update

酒店更新接口
*/
type TaobaoXhotelUpdateRequest struct {
    model.Params
    // （已废弃）请使用outer_id来标识要修改的酒店
    hid   int64
    // 酒店名称；（新增酒店时为必须）,国内酒店请传中文名称
    name   string
    // 酒店曾用名
    usedName   string
    // 是否国内酒店。0:国内;1:国外
    domestic   int64
    // domestic为true时，固定China； domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html
    country   string
    // 省份编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时默认为0
    province   int64
    // 城市编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（新增酒店时为必须）
    city   int64
    // 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3
    district   int64
    // 商业区（圈）长度不超过20字
    business   string
    // 酒店地址。长度不能超过255
    address   string
    // 经度
    longitude   string
    // 纬度
    latitude   string
    // 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
    positionType   string
    // 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
    tel   string
    // 不要使用
    extend   string
    // 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
    shid   int64
    // 必传，酒店标识，商家酒店ID
    outerId   string
    // 系统商，一般情况不用，需申请使用
    vendor   string
    // 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
    star   string
    // 开业时间，格式为2015-01-01
    openingTime   string
    // 装修时间，格式为2015-10-01
    decorateTime   string
    // 楼层信息
    floors   string
    // 房间数 0~9999之内的数字
    rooms   int64
    // 酒店描述
    description   string
    // 酒店设施。json格式示例值：{"free Wi-Fi in all rooms":"true","massage":"true","meetingRoom":"true"}目前支持维护的设施枚举有：free Wi-Fi in all rooms 所有房间设有免费无线网络;meetingRoom 会议室;massage  按摩室;fitnessClub 健身房;bar 酒吧;cafe 咖啡厅;frontDeskSafe 前台贵重物品保险柜wifi 无线上网公共区域;casino 娱乐场/棋牌室;restaurant 餐厅;smoking area 吸烟区;Business Facilities 商务设施
    hotelFacilities   string
    // 酒店基础服务。json格式示例值：{"receiveForeignGuests":"true","morningCall":"true","breakfast":"true"}目前支持维护的设施枚举有：receiveForeignGuests 接待外宾;morningCall 叫醒服务; breakfast  早餐服务; airportShuttle 接机服务; luggageClaim 行李寄存; rentCar 租车; HourRoomService24 24小时客房服务; airportTransfer 酒店/机场接送; dryCleaning 干洗; expressCheckInCheckOut 快速入住/退房登记; custodyServices 保管服务
    service   string
    // 房间的基础设施。json格式示例值：{"bathtub":"true","bathPub":"true"}目前支持维护的设施枚举有：bathtub 独立卫浴;bathPub 公共卫浴
    roomFacilities   string
    // 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址，main是否为主图（主图只能有一个）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多10张。要求：无logo、水印、边框、人物，不模糊、重复、歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
    pics   string
    // 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）： ruJia("1", "rujiakuaijie", "如家快捷", 1), qiTian("2", "7 days", "7天连锁", 1), hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1), geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1), jinJiang("5", "Jinjiang Inn", "锦江之星", 1), su8("6", "Super 8", "速8", 1), moTai("7", "Motel", "莫泰", 1), zhouji("8", "InterContinental", "洲际", 4), budint("9", "Pod Inn", "布丁", 1), jiuJiu("10", "jiujiuliansuo", "99连锁", 1), piaoHome("11", "Piao Home Inn", "飘HOME", 1), juzi("12", "Orange Hotels", "桔子酒店", 1), yibai("13", "yibai", "易佰", 1), weiyena("14","weiyena","维也纳",2), huangguanjiari("15", "huangguanjiari", "皇冠假日", 4), xidawu("16", "xidawu", "喜达屋", 3), chengshiBJ("17", "chengshibianjie", "城市便捷", 1), shagnKeYou("18", "shagnkeyou", "尚客优", 1), jinjiang("19", "jinjiang", "锦江酒店", 3), wendemu("20", "Hawthorn Suites", "温德姆", 4), yibisi("21", "Ibis Hotels", "宜必思", 1), wanhao("22", "JM Hoteles", "万豪", 4), yijia365("23", "yijia365", "驿家365", 1), shoulv("24", "shoulvjituan", "首旅建国", 3), kaiyuan("25", "New Century Hotel", "开元大酒店", 4), yagao("26", "yagao", "雅高", 3), daisi("27", "daisi", "戴斯", 3), jinling("28", "jinlingliansuo", "金陵", 4), xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4), xierdun("30", "Hilton", "希尔顿", 4)
    brand   string
    // 邮编
    postalCode   string
    // 酒店入住政策(针对国际酒店，儿童及加床信息)格式：{"children_age_from":"2","children_age_to":"3","children_stay_free":"True","infant_age":"1","min_guest_age":"4"}
    hotelPolicies   string
    // 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
    bookingNotice   string
    // 酒店状态 0:正常，-1:删除，-2:停售
    status   *model.File
    // 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
    creditCardTypes   string
    // 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
    orbitTrack   string
    // 卖家酒店英文名称
    nameE   string
    // 打标去标使用的 tagJson 字段
    tagJson   string
    // 旺旺昵称
    aliNick   string
    // 供应商标识，如果确实需要修改原来的供应商标识才需要填写，否则不需要填写，请谨慎使用。
    supplier   string
    // 结算流程中的结算币种，如需对接请联系飞猪技术支持，请谨慎使用
    settlementCurrency   string
    // 资源方酒店预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    standardBookingNotice   string
    // 资源方酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    standardHotelFacilities   string
    // 资源方酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    standardHotelService   string
    // 资源方房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    standardRoomFacilities   string
    // 资源方娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    standardAmuseFacilities   string
    // 0:酒店；1:客栈
    hotelType   int64
    // 0:可以接待外宾；1:仅内宾
    serviceType   int64
    // 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
    coordinateSystem   string
}

// 初始化TaobaoXhotelUpdateRequest对象
func NewTaobaoXhotelUpdateRequest() *TaobaoXhotelUpdateRequest{
    return &TaobaoXhotelUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Hid Setter
// （已废弃）请使用outer_id来标识要修改的酒店
func (r *TaobaoXhotelUpdateRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelUpdateRequest) GetHid() int64 {
    return r.hid
}
// Name Setter
// 酒店名称；（新增酒店时为必须）,国内酒店请传中文名称
func (r *TaobaoXhotelUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoXhotelUpdateRequest) GetName() string {
    return r.name
}
// UsedName Setter
// 酒店曾用名
func (r *TaobaoXhotelUpdateRequest) SetUsedName(usedName string) error {
    r.usedName = usedName
    r.Set("used_name", usedName)
    return nil
}

// UsedName Getter
func (r TaobaoXhotelUpdateRequest) GetUsedName() string {
    return r.usedName
}
// Domestic Setter
// 是否国内酒店。0:国内;1:国外
func (r *TaobaoXhotelUpdateRequest) SetDomestic(domestic int64) error {
    r.domestic = domestic
    r.Set("domestic", domestic)
    return nil
}

// Domestic Getter
func (r TaobaoXhotelUpdateRequest) GetDomestic() int64 {
    return r.domestic
}
// Country Setter
// domestic为true时，固定China； domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html
func (r *TaobaoXhotelUpdateRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

// Country Getter
func (r TaobaoXhotelUpdateRequest) GetCountry() string {
    return r.country
}
// Province Setter
// 省份编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时默认为0
func (r *TaobaoXhotelUpdateRequest) SetProvince(province int64) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r TaobaoXhotelUpdateRequest) GetProvince() int64 {
    return r.province
}
// City Setter
// 城市编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（新增酒店时为必须）
func (r *TaobaoXhotelUpdateRequest) SetCity(city int64) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r TaobaoXhotelUpdateRequest) GetCity() int64 {
    return r.city
}
// District Setter
// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3
func (r *TaobaoXhotelUpdateRequest) SetDistrict(district int64) error {
    r.district = district
    r.Set("district", district)
    return nil
}

// District Getter
func (r TaobaoXhotelUpdateRequest) GetDistrict() int64 {
    return r.district
}
// Business Setter
// 商业区（圈）长度不超过20字
func (r *TaobaoXhotelUpdateRequest) SetBusiness(business string) error {
    r.business = business
    r.Set("business", business)
    return nil
}

// Business Getter
func (r TaobaoXhotelUpdateRequest) GetBusiness() string {
    return r.business
}
// Address Setter
// 酒店地址。长度不能超过255
func (r *TaobaoXhotelUpdateRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r TaobaoXhotelUpdateRequest) GetAddress() string {
    return r.address
}
// Longitude Setter
// 经度
func (r *TaobaoXhotelUpdateRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r TaobaoXhotelUpdateRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度
func (r *TaobaoXhotelUpdateRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r TaobaoXhotelUpdateRequest) GetLatitude() string {
    return r.latitude
}
// PositionType Setter
// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
func (r *TaobaoXhotelUpdateRequest) SetPositionType(positionType string) error {
    r.positionType = positionType
    r.Set("position_type", positionType)
    return nil
}

// PositionType Getter
func (r TaobaoXhotelUpdateRequest) GetPositionType() string {
    return r.positionType
}
// Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelUpdateRequest) SetTel(tel string) error {
    r.tel = tel
    r.Set("tel", tel)
    return nil
}

// Tel Getter
func (r TaobaoXhotelUpdateRequest) GetTel() string {
    return r.tel
}
// Extend Setter
// 不要使用
func (r *TaobaoXhotelUpdateRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r TaobaoXhotelUpdateRequest) GetExtend() string {
    return r.extend
}
// Shid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelUpdateRequest) SetShid(shid int64) error {
    r.shid = shid
    r.Set("shid", shid)
    return nil
}

// Shid Getter
func (r TaobaoXhotelUpdateRequest) GetShid() int64 {
    return r.shid
}
// OuterId Setter
// 必传，酒店标识，商家酒店ID
func (r *TaobaoXhotelUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelUpdateRequest) GetOuterId() string {
    return r.outerId
}
// Vendor Setter
// 系统商，一般情况不用，需申请使用
func (r *TaobaoXhotelUpdateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelUpdateRequest) GetVendor() string {
    return r.vendor
}
// Star Setter
// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
func (r *TaobaoXhotelUpdateRequest) SetStar(star string) error {
    r.star = star
    r.Set("star", star)
    return nil
}

// Star Getter
func (r TaobaoXhotelUpdateRequest) GetStar() string {
    return r.star
}
// OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelUpdateRequest) SetOpeningTime(openingTime string) error {
    r.openingTime = openingTime
    r.Set("opening_time", openingTime)
    return nil
}

// OpeningTime Getter
func (r TaobaoXhotelUpdateRequest) GetOpeningTime() string {
    return r.openingTime
}
// DecorateTime Setter
// 装修时间，格式为2015-10-01
func (r *TaobaoXhotelUpdateRequest) SetDecorateTime(decorateTime string) error {
    r.decorateTime = decorateTime
    r.Set("decorate_time", decorateTime)
    return nil
}

// DecorateTime Getter
func (r TaobaoXhotelUpdateRequest) GetDecorateTime() string {
    return r.decorateTime
}
// Floors Setter
// 楼层信息
func (r *TaobaoXhotelUpdateRequest) SetFloors(floors string) error {
    r.floors = floors
    r.Set("floors", floors)
    return nil
}

// Floors Getter
func (r TaobaoXhotelUpdateRequest) GetFloors() string {
    return r.floors
}
// Rooms Setter
// 房间数 0~9999之内的数字
func (r *TaobaoXhotelUpdateRequest) SetRooms(rooms int64) error {
    r.rooms = rooms
    r.Set("rooms", rooms)
    return nil
}

// Rooms Getter
func (r TaobaoXhotelUpdateRequest) GetRooms() int64 {
    return r.rooms
}
// Description Setter
// 酒店描述
func (r *TaobaoXhotelUpdateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoXhotelUpdateRequest) GetDescription() string {
    return r.description
}
// HotelFacilities Setter
// 酒店设施。json格式示例值：{"free Wi-Fi in all rooms":"true","massage":"true","meetingRoom":"true"}目前支持维护的设施枚举有：free Wi-Fi in all rooms 所有房间设有免费无线网络;meetingRoom 会议室;massage  按摩室;fitnessClub 健身房;bar 酒吧;cafe 咖啡厅;frontDeskSafe 前台贵重物品保险柜wifi 无线上网公共区域;casino 娱乐场/棋牌室;restaurant 餐厅;smoking area 吸烟区;Business Facilities 商务设施
func (r *TaobaoXhotelUpdateRequest) SetHotelFacilities(hotelFacilities string) error {
    r.hotelFacilities = hotelFacilities
    r.Set("hotel_facilities", hotelFacilities)
    return nil
}

// HotelFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetHotelFacilities() string {
    return r.hotelFacilities
}
// Service Setter
// 酒店基础服务。json格式示例值：{"receiveForeignGuests":"true","morningCall":"true","breakfast":"true"}目前支持维护的设施枚举有：receiveForeignGuests 接待外宾;morningCall 叫醒服务; breakfast  早餐服务; airportShuttle 接机服务; luggageClaim 行李寄存; rentCar 租车; HourRoomService24 24小时客房服务; airportTransfer 酒店/机场接送; dryCleaning 干洗; expressCheckInCheckOut 快速入住/退房登记; custodyServices 保管服务
func (r *TaobaoXhotelUpdateRequest) SetService(service string) error {
    r.service = service
    r.Set("service", service)
    return nil
}

// Service Getter
func (r TaobaoXhotelUpdateRequest) GetService() string {
    return r.service
}
// RoomFacilities Setter
// 房间的基础设施。json格式示例值：{"bathtub":"true","bathPub":"true"}目前支持维护的设施枚举有：bathtub 独立卫浴;bathPub 公共卫浴
func (r *TaobaoXhotelUpdateRequest) SetRoomFacilities(roomFacilities string) error {
    r.roomFacilities = roomFacilities
    r.Set("room_facilities", roomFacilities)
    return nil
}

// RoomFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetRoomFacilities() string {
    return r.roomFacilities
}
// Pics Setter
// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址，main是否为主图（主图只能有一个）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多10张。要求：无logo、水印、边框、人物，不模糊、重复、歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
func (r *TaobaoXhotelUpdateRequest) SetPics(pics string) error {
    r.pics = pics
    r.Set("pics", pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelUpdateRequest) GetPics() string {
    return r.pics
}
// Brand Setter
// 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）： ruJia("1", "rujiakuaijie", "如家快捷", 1), qiTian("2", "7 days", "7天连锁", 1), hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1), geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1), jinJiang("5", "Jinjiang Inn", "锦江之星", 1), su8("6", "Super 8", "速8", 1), moTai("7", "Motel", "莫泰", 1), zhouji("8", "InterContinental", "洲际", 4), budint("9", "Pod Inn", "布丁", 1), jiuJiu("10", "jiujiuliansuo", "99连锁", 1), piaoHome("11", "Piao Home Inn", "飘HOME", 1), juzi("12", "Orange Hotels", "桔子酒店", 1), yibai("13", "yibai", "易佰", 1), weiyena("14","weiyena","维也纳",2), huangguanjiari("15", "huangguanjiari", "皇冠假日", 4), xidawu("16", "xidawu", "喜达屋", 3), chengshiBJ("17", "chengshibianjie", "城市便捷", 1), shagnKeYou("18", "shagnkeyou", "尚客优", 1), jinjiang("19", "jinjiang", "锦江酒店", 3), wendemu("20", "Hawthorn Suites", "温德姆", 4), yibisi("21", "Ibis Hotels", "宜必思", 1), wanhao("22", "JM Hoteles", "万豪", 4), yijia365("23", "yijia365", "驿家365", 1), shoulv("24", "shoulvjituan", "首旅建国", 3), kaiyuan("25", "New Century Hotel", "开元大酒店", 4), yagao("26", "yagao", "雅高", 3), daisi("27", "daisi", "戴斯", 3), jinling("28", "jinlingliansuo", "金陵", 4), xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4), xierdun("30", "Hilton", "希尔顿", 4)
func (r *TaobaoXhotelUpdateRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

// Brand Getter
func (r TaobaoXhotelUpdateRequest) GetBrand() string {
    return r.brand
}
// PostalCode Setter
// 邮编
func (r *TaobaoXhotelUpdateRequest) SetPostalCode(postalCode string) error {
    r.postalCode = postalCode
    r.Set("postal_code", postalCode)
    return nil
}

// PostalCode Getter
func (r TaobaoXhotelUpdateRequest) GetPostalCode() string {
    return r.postalCode
}
// HotelPolicies Setter
// 酒店入住政策(针对国际酒店，儿童及加床信息)格式：{"children_age_from":"2","children_age_to":"3","children_stay_free":"True","infant_age":"1","min_guest_age":"4"}
func (r *TaobaoXhotelUpdateRequest) SetHotelPolicies(hotelPolicies string) error {
    r.hotelPolicies = hotelPolicies
    r.Set("hotel_policies", hotelPolicies)
    return nil
}

// HotelPolicies Getter
func (r TaobaoXhotelUpdateRequest) GetHotelPolicies() string {
    return r.hotelPolicies
}
// BookingNotice Setter
// 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
func (r *TaobaoXhotelUpdateRequest) SetBookingNotice(bookingNotice string) error {
    r.bookingNotice = bookingNotice
    r.Set("booking_notice", bookingNotice)
    return nil
}

// BookingNotice Getter
func (r TaobaoXhotelUpdateRequest) GetBookingNotice() string {
    return r.bookingNotice
}
// Status Setter
// 酒店状态 0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelUpdateRequest) SetStatus(status *model.File) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoXhotelUpdateRequest) GetStatus() *model.File {
    return r.status
}
// CreditCardTypes Setter
// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
func (r *TaobaoXhotelUpdateRequest) SetCreditCardTypes(creditCardTypes string) error {
    r.creditCardTypes = creditCardTypes
    r.Set("credit_card_types", creditCardTypes)
    return nil
}

// CreditCardTypes Getter
func (r TaobaoXhotelUpdateRequest) GetCreditCardTypes() string {
    return r.creditCardTypes
}
// OrbitTrack Setter
// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
func (r *TaobaoXhotelUpdateRequest) SetOrbitTrack(orbitTrack string) error {
    r.orbitTrack = orbitTrack
    r.Set("orbit_track", orbitTrack)
    return nil
}

// OrbitTrack Getter
func (r TaobaoXhotelUpdateRequest) GetOrbitTrack() string {
    return r.orbitTrack
}
// NameE Setter
// 卖家酒店英文名称
func (r *TaobaoXhotelUpdateRequest) SetNameE(nameE string) error {
    r.nameE = nameE
    r.Set("name_e", nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelUpdateRequest) GetNameE() string {
    return r.nameE
}
// TagJson Setter
// 打标去标使用的 tagJson 字段
func (r *TaobaoXhotelUpdateRequest) SetTagJson(tagJson string) error {
    r.tagJson = tagJson
    r.Set("tag_json", tagJson)
    return nil
}

// TagJson Getter
func (r TaobaoXhotelUpdateRequest) GetTagJson() string {
    return r.tagJson
}
// AliNick Setter
// 旺旺昵称
func (r *TaobaoXhotelUpdateRequest) SetAliNick(aliNick string) error {
    r.aliNick = aliNick
    r.Set("ali_nick", aliNick)
    return nil
}

// AliNick Getter
func (r TaobaoXhotelUpdateRequest) GetAliNick() string {
    return r.aliNick
}
// Supplier Setter
// 供应商标识，如果确实需要修改原来的供应商标识才需要填写，否则不需要填写，请谨慎使用。
func (r *TaobaoXhotelUpdateRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelUpdateRequest) GetSupplier() string {
    return r.supplier
}
// SettlementCurrency Setter
// 结算流程中的结算币种，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelUpdateRequest) SetSettlementCurrency(settlementCurrency string) error {
    r.settlementCurrency = settlementCurrency
    r.Set("settlement_currency", settlementCurrency)
    return nil
}

// SettlementCurrency Getter
func (r TaobaoXhotelUpdateRequest) GetSettlementCurrency() string {
    return r.settlementCurrency
}
// StandardBookingNotice Setter
// 资源方酒店预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardBookingNotice(standardBookingNotice string) error {
    r.standardBookingNotice = standardBookingNotice
    r.Set("standard_booking_notice", standardBookingNotice)
    return nil
}

// StandardBookingNotice Getter
func (r TaobaoXhotelUpdateRequest) GetStandardBookingNotice() string {
    return r.standardBookingNotice
}
// StandardHotelFacilities Setter
// 资源方酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardHotelFacilities(standardHotelFacilities string) error {
    r.standardHotelFacilities = standardHotelFacilities
    r.Set("standard_hotel_facilities", standardHotelFacilities)
    return nil
}

// StandardHotelFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetStandardHotelFacilities() string {
    return r.standardHotelFacilities
}
// StandardHotelService Setter
// 资源方酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardHotelService(standardHotelService string) error {
    r.standardHotelService = standardHotelService
    r.Set("standard_hotel_service", standardHotelService)
    return nil
}

// StandardHotelService Getter
func (r TaobaoXhotelUpdateRequest) GetStandardHotelService() string {
    return r.standardHotelService
}
// StandardRoomFacilities Setter
// 资源方房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardRoomFacilities(standardRoomFacilities string) error {
    r.standardRoomFacilities = standardRoomFacilities
    r.Set("standard_room_facilities", standardRoomFacilities)
    return nil
}

// StandardRoomFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetStandardRoomFacilities() string {
    return r.standardRoomFacilities
}
// StandardAmuseFacilities Setter
// 资源方娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardAmuseFacilities(standardAmuseFacilities string) error {
    r.standardAmuseFacilities = standardAmuseFacilities
    r.Set("standard_amuse_facilities", standardAmuseFacilities)
    return nil
}

// StandardAmuseFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetStandardAmuseFacilities() string {
    return r.standardAmuseFacilities
}
// HotelType Setter
// 0:酒店；1:客栈
func (r *TaobaoXhotelUpdateRequest) SetHotelType(hotelType int64) error {
    r.hotelType = hotelType
    r.Set("hotel_type", hotelType)
    return nil
}

// HotelType Getter
func (r TaobaoXhotelUpdateRequest) GetHotelType() int64 {
    return r.hotelType
}
// ServiceType Setter
// 0:可以接待外宾；1:仅内宾
func (r *TaobaoXhotelUpdateRequest) SetServiceType(serviceType int64) error {
    r.serviceType = serviceType
    r.Set("service_type", serviceType)
    return nil
}

// ServiceType Getter
func (r TaobaoXhotelUpdateRequest) GetServiceType() int64 {
    return r.serviceType
}
// CoordinateSystem Setter
// 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
func (r *TaobaoXhotelUpdateRequest) SetCoordinateSystem(coordinateSystem string) error {
    r.coordinateSystem = coordinateSystem
    r.Set("coordinate_system", coordinateSystem)
    return nil
}

// CoordinateSystem Getter
func (r TaobaoXhotelUpdateRequest) GetCoordinateSystem() string {
    return r.coordinateSystem
}
