package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
非标准民宿房源添加 API请求
taobao.xhotel.house.add

添加酒店或更新酒店
*/
type TaobaoXhotelHouseAddRequest struct {
    model.Params
    // 外部酒店ID, 这是卖家自己系统中的ID
    outerId   string
    // 酒店名称
    name   string
    // 酒店曾用名
    usedName   string
    // 是否国内酒店。0:国内;1:国外。默认是国内
    domestic   int64
    // domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm
    country   string
    // 省份编码。选填，不填入的时候已city字段为准.参见：http://hotel.alitrip.com/area.htm，domestic为false时默认为0
    province   int64
    // 城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新酒店时为可选）
    city   int64
    // 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm
    district   int64
    // 商业区（圈）长度不超过20字
    business   string
    // 酒店地址。长度不能超过120。不填入会导致不能自动匹配,此地址为下单前展示给用户地址
    address   string
    // 经度
    longitude   string
    // 纬度
    latitude   string
    // 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
    positionType   string
    // 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
    tel   string
    // 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
    extend   string
    // 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
    shid   int64
    // 对接系统商名称：可为空不要乱填，需要申请后使用
    vendor   string
    // 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
    star   string
    // 开业时间，格式为2015-01-01
    openingTime   string
    // 装修时间，格式为2015-01-01装修时间
    decorateTime   string
    // 楼层信息。
    floors   string
    // 房间数 0~9999之内的数字
    rooms   int64
    // 酒店描述
    description   string
    // 酒店入住政策，{"10003":"仅2岁以上儿童可随同入住"}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10、code=10003,value为文字描述
    hotelPolicies   string
    // 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
    hotelFacilities   string
    // 酒店服务。json格式示例值：{"24058":"可以接待外宾","24198":"叫醒服务","24200":"洗衣服务"}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
    service   string
    // 房间设施。json格式示例值：{"24101": true,"24091": true,"24095": true}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
    roomFacilities   string
    // 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多是能是10张。
    pics   string
    // 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）：    ruJia("1", "rujiakuaijie", "如家快捷", 1),    qiTian("2", "7 days", "7天连锁", 1),    hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1),    geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1),    jinJiang("5", "Jinjiang Inn", "锦江之星", 1),    su8("6", "Super 8", "速8", 1),    moTai("7", "Motel", "莫泰", 1),    zhouji("8", "InterContinental", "洲际", 4),    budint("9", "Pod Inn", "布丁", 1),    jiuJiu("10", "jiujiuliansuo", "99连锁", 1),    piaoHome("11", "Piao Home Inn", "飘HOME", 1),    juzi("12", "Orange Hotels", "桔子酒店", 1),    yibai("13", "yibai", "易佰", 1),    weiyena("14","weiyena","维也纳",2),    huangguanjiari("15", "huangguanjiari", "皇冠假日", 4),    xidawu("16", "xidawu", "喜达屋", 3),    chengshiBJ("17", "chengshibianjie", "城市便捷", 1),    shagnKeYou("18", "shagnkeyou", "尚客优", 1),    jinjiang("19", "jinjiang", "锦江酒店", 3),    wendemu("20", "Hawthorn Suites", "温德姆", 4),    yibisi("21", "Ibis Hotels", "宜必思", 1),    wanhao("22", "JM Hoteles", "万豪", 4),    yijia365("23", "yijia365", "驿家365", 1),    shoulv("24", "shoulvjituan", "首旅建国", 3),    kaiyuan("25", "New Century Hotel", "开元大酒店", 4),    yagao("26", "yagao", "雅高", 3),    daisi("27", "daisi", "戴斯", 3),    jinling("28", "jinlingliansuo", "金陵", 4),    xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4),    xierdun("30", "Hilton", "希尔顿", 4),
    brand   string
    // 邮政编码。
    postalCode   string
    // 预订须知。json格式，示例:{"10001":"14:00","10002":"12:00","10005":"清洁福50元","10006":"请准备好您的身份证件，我需要登记 不允许吸烟"},预订须知，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10的分类
    bookingNotice   string
    // 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
    creditCardTypes   string
    // 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
    orbitTrack   string
    // 卖家酒店英文名称
    nameE   string
    // 供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用
    supplier   string
    // 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
    settlementCurrency   string
    // 房东信息,{"outerId: 外部房东ID,": "xxxx", "nickName": "张三", "avatarUrl": "http://test.com/1.jpg", "telephone": "0571-1234567", "mobilePhone": "12334567678", "email":"test@test.com", "gender": "F", "avgConfirmTime": 30, "responseRate": 100, "description": "房东太懒,什么也没有填", "birthday":"2018-01-01", "qualifacation": 1, "bloodType": 1, "profession":"交互设计师", "country":"CN", "province":"420000", "city":"421200", "real_name_status": true, "validate":"1,2,4,8","confirmRate": 98} JSON字段描述: outerId: 商家房东ID, nickName: 房东昵称, avatarUrl: 房东头像地址, telephone: 固定电话, mobilePhone: 移动电话, email: 邮箱地址, gender: 性别 M男性， F女性， avgConfirmTime: 平均确认时间, 单位分钟, responseRate: 房东回复率, description: 房东介绍, birthday:生日，格式yyyy-MM-dd, qualifacation:学历,1:小学,2:初中,3:高中,4:本科,5:硕士,6:博士,7:博士后,0:其他, profession: 职业 country: 国家code province: 省code city: 城市code realNameStatus: 实名认证状态, true已认证 validate: 认证情况:1:身份验证,2:头像验证,4:手机验证,8:邮箱验证,二进制各位代表含义, bloodType: 血型: 0未知,1:A型,2:B型,3:AB型,4:O型;confirmRate: 订单接单率，0-100
    ownerInfo   string
    // 描述信息，inside： 内部描述,traffic:交通情况,arround:周边情况
    arroundDesc   string
    // 用户下单之后展示的详细地址
    realAddress   string
    // 数据状态 0:正常，-1:删除，-2:停售
    status   int64
}

// 初始化TaobaoXhotelHouseAddRequest对象
func NewTaobaoXhotelHouseAddRequest() *TaobaoXhotelHouseAddRequest{
    return &TaobaoXhotelHouseAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelHouseAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.house.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelHouseAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 外部酒店ID, 这是卖家自己系统中的ID
func (r *TaobaoXhotelHouseAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelHouseAddRequest) GetOuterId() string {
    return r.outerId
}
// Name Setter
// 酒店名称
func (r *TaobaoXhotelHouseAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoXhotelHouseAddRequest) GetName() string {
    return r.name
}
// UsedName Setter
// 酒店曾用名
func (r *TaobaoXhotelHouseAddRequest) SetUsedName(usedName string) error {
    r.usedName = usedName
    r.Set("used_name", usedName)
    return nil
}

// UsedName Getter
func (r TaobaoXhotelHouseAddRequest) GetUsedName() string {
    return r.usedName
}
// Domestic Setter
// 是否国内酒店。0:国内;1:国外。默认是国内
func (r *TaobaoXhotelHouseAddRequest) SetDomestic(domestic int64) error {
    r.domestic = domestic
    r.Set("domestic", domestic)
    return nil
}

// Domestic Getter
func (r TaobaoXhotelHouseAddRequest) GetDomestic() int64 {
    return r.domestic
}
// Country Setter
// domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm
func (r *TaobaoXhotelHouseAddRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

// Country Getter
func (r TaobaoXhotelHouseAddRequest) GetCountry() string {
    return r.country
}
// Province Setter
// 省份编码。选填，不填入的时候已city字段为准.参见：http://hotel.alitrip.com/area.htm，domestic为false时默认为0
func (r *TaobaoXhotelHouseAddRequest) SetProvince(province int64) error {
    r.province = province
    r.Set("province", province)
    return nil
}

// Province Getter
func (r TaobaoXhotelHouseAddRequest) GetProvince() int64 {
    return r.province
}
// City Setter
// 城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新酒店时为可选）
func (r *TaobaoXhotelHouseAddRequest) SetCity(city int64) error {
    r.city = city
    r.Set("city", city)
    return nil
}

// City Getter
func (r TaobaoXhotelHouseAddRequest) GetCity() int64 {
    return r.city
}
// District Setter
// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm
func (r *TaobaoXhotelHouseAddRequest) SetDistrict(district int64) error {
    r.district = district
    r.Set("district", district)
    return nil
}

// District Getter
func (r TaobaoXhotelHouseAddRequest) GetDistrict() int64 {
    return r.district
}
// Business Setter
// 商业区（圈）长度不超过20字
func (r *TaobaoXhotelHouseAddRequest) SetBusiness(business string) error {
    r.business = business
    r.Set("business", business)
    return nil
}

// Business Getter
func (r TaobaoXhotelHouseAddRequest) GetBusiness() string {
    return r.business
}
// Address Setter
// 酒店地址。长度不能超过120。不填入会导致不能自动匹配,此地址为下单前展示给用户地址
func (r *TaobaoXhotelHouseAddRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r TaobaoXhotelHouseAddRequest) GetAddress() string {
    return r.address
}
// Longitude Setter
// 经度
func (r *TaobaoXhotelHouseAddRequest) SetLongitude(longitude string) error {
    r.longitude = longitude
    r.Set("longitude", longitude)
    return nil
}

// Longitude Getter
func (r TaobaoXhotelHouseAddRequest) GetLongitude() string {
    return r.longitude
}
// Latitude Setter
// 纬度
func (r *TaobaoXhotelHouseAddRequest) SetLatitude(latitude string) error {
    r.latitude = latitude
    r.Set("latitude", latitude)
    return nil
}

// Latitude Getter
func (r TaobaoXhotelHouseAddRequest) GetLatitude() string {
    return r.latitude
}
// PositionType Setter
// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
func (r *TaobaoXhotelHouseAddRequest) SetPositionType(positionType string) error {
    r.positionType = positionType
    r.Set("position_type", positionType)
    return nil
}

// PositionType Getter
func (r TaobaoXhotelHouseAddRequest) GetPositionType() string {
    return r.positionType
}
// Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelHouseAddRequest) SetTel(tel string) error {
    r.tel = tel
    r.Set("tel", tel)
    return nil
}

// Tel Getter
func (r TaobaoXhotelHouseAddRequest) GetTel() string {
    return r.tel
}
// Extend Setter
// 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
func (r *TaobaoXhotelHouseAddRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r TaobaoXhotelHouseAddRequest) GetExtend() string {
    return r.extend
}
// Shid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelHouseAddRequest) SetShid(shid int64) error {
    r.shid = shid
    r.Set("shid", shid)
    return nil
}

// Shid Getter
func (r TaobaoXhotelHouseAddRequest) GetShid() int64 {
    return r.shid
}
// Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoXhotelHouseAddRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelHouseAddRequest) GetVendor() string {
    return r.vendor
}
// Star Setter
// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
func (r *TaobaoXhotelHouseAddRequest) SetStar(star string) error {
    r.star = star
    r.Set("star", star)
    return nil
}

// Star Getter
func (r TaobaoXhotelHouseAddRequest) GetStar() string {
    return r.star
}
// OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelHouseAddRequest) SetOpeningTime(openingTime string) error {
    r.openingTime = openingTime
    r.Set("opening_time", openingTime)
    return nil
}

// OpeningTime Getter
func (r TaobaoXhotelHouseAddRequest) GetOpeningTime() string {
    return r.openingTime
}
// DecorateTime Setter
// 装修时间，格式为2015-01-01装修时间
func (r *TaobaoXhotelHouseAddRequest) SetDecorateTime(decorateTime string) error {
    r.decorateTime = decorateTime
    r.Set("decorate_time", decorateTime)
    return nil
}

// DecorateTime Getter
func (r TaobaoXhotelHouseAddRequest) GetDecorateTime() string {
    return r.decorateTime
}
// Floors Setter
// 楼层信息。
func (r *TaobaoXhotelHouseAddRequest) SetFloors(floors string) error {
    r.floors = floors
    r.Set("floors", floors)
    return nil
}

// Floors Getter
func (r TaobaoXhotelHouseAddRequest) GetFloors() string {
    return r.floors
}
// Rooms Setter
// 房间数 0~9999之内的数字
func (r *TaobaoXhotelHouseAddRequest) SetRooms(rooms int64) error {
    r.rooms = rooms
    r.Set("rooms", rooms)
    return nil
}

// Rooms Getter
func (r TaobaoXhotelHouseAddRequest) GetRooms() int64 {
    return r.rooms
}
// Description Setter
// 酒店描述
func (r *TaobaoXhotelHouseAddRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r TaobaoXhotelHouseAddRequest) GetDescription() string {
    return r.description
}
// HotelPolicies Setter
// 酒店入住政策，{"10003":"仅2岁以上儿童可随同入住"}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10、code=10003,value为文字描述
func (r *TaobaoXhotelHouseAddRequest) SetHotelPolicies(hotelPolicies string) error {
    r.hotelPolicies = hotelPolicies
    r.Set("hotel_policies", hotelPolicies)
    return nil
}

// HotelPolicies Getter
func (r TaobaoXhotelHouseAddRequest) GetHotelPolicies() string {
    return r.hotelPolicies
}
// HotelFacilities Setter
// 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
func (r *TaobaoXhotelHouseAddRequest) SetHotelFacilities(hotelFacilities string) error {
    r.hotelFacilities = hotelFacilities
    r.Set("hotel_facilities", hotelFacilities)
    return nil
}

// HotelFacilities Getter
func (r TaobaoXhotelHouseAddRequest) GetHotelFacilities() string {
    return r.hotelFacilities
}
// Service Setter
// 酒店服务。json格式示例值：{"24058":"可以接待外宾","24198":"叫醒服务","24200":"洗衣服务"}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
func (r *TaobaoXhotelHouseAddRequest) SetService(service string) error {
    r.service = service
    r.Set("service", service)
    return nil
}

// Service Getter
func (r TaobaoXhotelHouseAddRequest) GetService() string {
    return r.service
}
// RoomFacilities Setter
// 房间设施。json格式示例值：{"24101": true,"24091": true,"24095": true}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
func (r *TaobaoXhotelHouseAddRequest) SetRoomFacilities(roomFacilities string) error {
    r.roomFacilities = roomFacilities
    r.Set("room_facilities", roomFacilities)
    return nil
}

// RoomFacilities Getter
func (r TaobaoXhotelHouseAddRequest) GetRoomFacilities() string {
    return r.roomFacilities
}
// Pics Setter
// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多是能是10张。
func (r *TaobaoXhotelHouseAddRequest) SetPics(pics string) error {
    r.pics = pics
    r.Set("pics", pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelHouseAddRequest) GetPics() string {
    return r.pics
}
// Brand Setter
// 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）：    ruJia("1", "rujiakuaijie", "如家快捷", 1),    qiTian("2", "7 days", "7天连锁", 1),    hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1),    geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1),    jinJiang("5", "Jinjiang Inn", "锦江之星", 1),    su8("6", "Super 8", "速8", 1),    moTai("7", "Motel", "莫泰", 1),    zhouji("8", "InterContinental", "洲际", 4),    budint("9", "Pod Inn", "布丁", 1),    jiuJiu("10", "jiujiuliansuo", "99连锁", 1),    piaoHome("11", "Piao Home Inn", "飘HOME", 1),    juzi("12", "Orange Hotels", "桔子酒店", 1),    yibai("13", "yibai", "易佰", 1),    weiyena("14","weiyena","维也纳",2),    huangguanjiari("15", "huangguanjiari", "皇冠假日", 4),    xidawu("16", "xidawu", "喜达屋", 3),    chengshiBJ("17", "chengshibianjie", "城市便捷", 1),    shagnKeYou("18", "shagnkeyou", "尚客优", 1),    jinjiang("19", "jinjiang", "锦江酒店", 3),    wendemu("20", "Hawthorn Suites", "温德姆", 4),    yibisi("21", "Ibis Hotels", "宜必思", 1),    wanhao("22", "JM Hoteles", "万豪", 4),    yijia365("23", "yijia365", "驿家365", 1),    shoulv("24", "shoulvjituan", "首旅建国", 3),    kaiyuan("25", "New Century Hotel", "开元大酒店", 4),    yagao("26", "yagao", "雅高", 3),    daisi("27", "daisi", "戴斯", 3),    jinling("28", "jinlingliansuo", "金陵", 4),    xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4),    xierdun("30", "Hilton", "希尔顿", 4),
func (r *TaobaoXhotelHouseAddRequest) SetBrand(brand string) error {
    r.brand = brand
    r.Set("brand", brand)
    return nil
}

// Brand Getter
func (r TaobaoXhotelHouseAddRequest) GetBrand() string {
    return r.brand
}
// PostalCode Setter
// 邮政编码。
func (r *TaobaoXhotelHouseAddRequest) SetPostalCode(postalCode string) error {
    r.postalCode = postalCode
    r.Set("postal_code", postalCode)
    return nil
}

// PostalCode Getter
func (r TaobaoXhotelHouseAddRequest) GetPostalCode() string {
    return r.postalCode
}
// BookingNotice Setter
// 预订须知。json格式，示例:{"10001":"14:00","10002":"12:00","10005":"清洁福50元","10006":"请准备好您的身份证件，我需要登记 不允许吸烟"},预订须知，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10的分类
func (r *TaobaoXhotelHouseAddRequest) SetBookingNotice(bookingNotice string) error {
    r.bookingNotice = bookingNotice
    r.Set("booking_notice", bookingNotice)
    return nil
}

// BookingNotice Getter
func (r TaobaoXhotelHouseAddRequest) GetBookingNotice() string {
    return r.bookingNotice
}
// CreditCardTypes Setter
// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
func (r *TaobaoXhotelHouseAddRequest) SetCreditCardTypes(creditCardTypes string) error {
    r.creditCardTypes = creditCardTypes
    r.Set("credit_card_types", creditCardTypes)
    return nil
}

// CreditCardTypes Getter
func (r TaobaoXhotelHouseAddRequest) GetCreditCardTypes() string {
    return r.creditCardTypes
}
// OrbitTrack Setter
// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
func (r *TaobaoXhotelHouseAddRequest) SetOrbitTrack(orbitTrack string) error {
    r.orbitTrack = orbitTrack
    r.Set("orbit_track", orbitTrack)
    return nil
}

// OrbitTrack Getter
func (r TaobaoXhotelHouseAddRequest) GetOrbitTrack() string {
    return r.orbitTrack
}
// NameE Setter
// 卖家酒店英文名称
func (r *TaobaoXhotelHouseAddRequest) SetNameE(nameE string) error {
    r.nameE = nameE
    r.Set("name_e", nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelHouseAddRequest) GetNameE() string {
    return r.nameE
}
// Supplier Setter
// 供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用
func (r *TaobaoXhotelHouseAddRequest) SetSupplier(supplier string) error {
    r.supplier = supplier
    r.Set("supplier", supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelHouseAddRequest) GetSupplier() string {
    return r.supplier
}
// SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelHouseAddRequest) SetSettlementCurrency(settlementCurrency string) error {
    r.settlementCurrency = settlementCurrency
    r.Set("settlement_currency", settlementCurrency)
    return nil
}

// SettlementCurrency Getter
func (r TaobaoXhotelHouseAddRequest) GetSettlementCurrency() string {
    return r.settlementCurrency
}
// OwnerInfo Setter
// 房东信息,{"outerId: 外部房东ID,": "xxxx", "nickName": "张三", "avatarUrl": "http://test.com/1.jpg", "telephone": "0571-1234567", "mobilePhone": "12334567678", "email":"test@test.com", "gender": "F", "avgConfirmTime": 30, "responseRate": 100, "description": "房东太懒,什么也没有填", "birthday":"2018-01-01", "qualifacation": 1, "bloodType": 1, "profession":"交互设计师", "country":"CN", "province":"420000", "city":"421200", "real_name_status": true, "validate":"1,2,4,8","confirmRate": 98} JSON字段描述: outerId: 商家房东ID, nickName: 房东昵称, avatarUrl: 房东头像地址, telephone: 固定电话, mobilePhone: 移动电话, email: 邮箱地址, gender: 性别 M男性， F女性， avgConfirmTime: 平均确认时间, 单位分钟, responseRate: 房东回复率, description: 房东介绍, birthday:生日，格式yyyy-MM-dd, qualifacation:学历,1:小学,2:初中,3:高中,4:本科,5:硕士,6:博士,7:博士后,0:其他, profession: 职业 country: 国家code province: 省code city: 城市code realNameStatus: 实名认证状态, true已认证 validate: 认证情况:1:身份验证,2:头像验证,4:手机验证,8:邮箱验证,二进制各位代表含义, bloodType: 血型: 0未知,1:A型,2:B型,3:AB型,4:O型;confirmRate: 订单接单率，0-100
func (r *TaobaoXhotelHouseAddRequest) SetOwnerInfo(ownerInfo string) error {
    r.ownerInfo = ownerInfo
    r.Set("owner_info", ownerInfo)
    return nil
}

// OwnerInfo Getter
func (r TaobaoXhotelHouseAddRequest) GetOwnerInfo() string {
    return r.ownerInfo
}
// ArroundDesc Setter
// 描述信息，inside： 内部描述,traffic:交通情况,arround:周边情况
func (r *TaobaoXhotelHouseAddRequest) SetArroundDesc(arroundDesc string) error {
    r.arroundDesc = arroundDesc
    r.Set("arround_desc", arroundDesc)
    return nil
}

// ArroundDesc Getter
func (r TaobaoXhotelHouseAddRequest) GetArroundDesc() string {
    return r.arroundDesc
}
// RealAddress Setter
// 用户下单之后展示的详细地址
func (r *TaobaoXhotelHouseAddRequest) SetRealAddress(realAddress string) error {
    r.realAddress = realAddress
    r.Set("real_address", realAddress)
    return nil
}

// RealAddress Getter
func (r TaobaoXhotelHouseAddRequest) GetRealAddress() string {
    return r.realAddress
}
// Status Setter
// 数据状态 0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelHouseAddRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoXhotelHouseAddRequest) GetStatus() int64 {
    return r.status
}
