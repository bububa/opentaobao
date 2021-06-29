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
    _hid   int64
    // 酒店名称；（新增酒店时为必须）,国内酒店请传中文名称
    _name   string
    // 酒店曾用名
    _usedName   string
    // 是否国内酒店。0:国内;1:国外
    _domestic   int64
    // domestic为true时，固定China； domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html
    _country   string
    // 省份编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时默认为0
    _province   int64
    // 城市编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（新增酒店时为必须）
    _city   int64
    // 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3
    _district   int64
    // 商业区（圈）长度不超过20字
    _business   string
    // 酒店地址。长度不能超过255
    _address   string
    // 经度
    _longitude   string
    // 纬度
    _latitude   string
    // 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
    _positionType   string
    // 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
    _tel   string
    // 不要使用
    _extend   string
    // 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
    _shid   int64
    // 必传，酒店标识，商家酒店ID
    _outerId   string
    // 系统商，一般情况不用，需申请使用
    _vendor   string
    // 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
    _star   string
    // 开业时间，格式为2015-01-01
    _openingTime   string
    // 装修时间，格式为2015-10-01
    _decorateTime   string
    // 楼层信息
    _floors   string
    // 房间数 0~9999之内的数字
    _rooms   int64
    // 酒店描述
    _description   string
    // 酒店设施。json格式示例值：{"free Wi-Fi in all rooms":"true","massage":"true","meetingRoom":"true"}目前支持维护的设施枚举有：free Wi-Fi in all rooms 所有房间设有免费无线网络;meetingRoom 会议室;massage  按摩室;fitnessClub 健身房;bar 酒吧;cafe 咖啡厅;frontDeskSafe 前台贵重物品保险柜wifi 无线上网公共区域;casino 娱乐场/棋牌室;restaurant 餐厅;smoking area 吸烟区;Business Facilities 商务设施
    _hotelFacilities   string
    // 酒店基础服务。json格式示例值：{"receiveForeignGuests":"true","morningCall":"true","breakfast":"true"}目前支持维护的设施枚举有：receiveForeignGuests 接待外宾;morningCall 叫醒服务; breakfast  早餐服务; airportShuttle 接机服务; luggageClaim 行李寄存; rentCar 租车; HourRoomService24 24小时客房服务; airportTransfer 酒店/机场接送; dryCleaning 干洗; expressCheckInCheckOut 快速入住/退房登记; custodyServices 保管服务
    _service   string
    // 房间的基础设施。json格式示例值：{"bathtub":"true","bathPub":"true"}目前支持维护的设施枚举有：bathtub 独立卫浴;bathPub 公共卫浴
    _roomFacilities   string
    // 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址，main是否为主图（主图只能有一个）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多10张。要求：无logo、水印、边框、人物，不模糊、重复、歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
    _pics   string
    // 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）： ruJia("1", "rujiakuaijie", "如家快捷", 1), qiTian("2", "7 days", "7天连锁", 1), hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1), geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1), jinJiang("5", "Jinjiang Inn", "锦江之星", 1), su8("6", "Super 8", "速8", 1), moTai("7", "Motel", "莫泰", 1), zhouji("8", "InterContinental", "洲际", 4), budint("9", "Pod Inn", "布丁", 1), jiuJiu("10", "jiujiuliansuo", "99连锁", 1), piaoHome("11", "Piao Home Inn", "飘HOME", 1), juzi("12", "Orange Hotels", "桔子酒店", 1), yibai("13", "yibai", "易佰", 1), weiyena("14","weiyena","维也纳",2), huangguanjiari("15", "huangguanjiari", "皇冠假日", 4), xidawu("16", "xidawu", "喜达屋", 3), chengshiBJ("17", "chengshibianjie", "城市便捷", 1), shagnKeYou("18", "shagnkeyou", "尚客优", 1), jinjiang("19", "jinjiang", "锦江酒店", 3), wendemu("20", "Hawthorn Suites", "温德姆", 4), yibisi("21", "Ibis Hotels", "宜必思", 1), wanhao("22", "JM Hoteles", "万豪", 4), yijia365("23", "yijia365", "驿家365", 1), shoulv("24", "shoulvjituan", "首旅建国", 3), kaiyuan("25", "New Century Hotel", "开元大酒店", 4), yagao("26", "yagao", "雅高", 3), daisi("27", "daisi", "戴斯", 3), jinling("28", "jinlingliansuo", "金陵", 4), xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4), xierdun("30", "Hilton", "希尔顿", 4)
    _brand   string
    // 邮编
    _postalCode   string
    // 酒店入住政策(针对国际酒店，儿童及加床信息)格式：{"children_age_from":"2","children_age_to":"3","children_stay_free":"True","infant_age":"1","min_guest_age":"4"}
    _hotelPolicies   string
    // 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
    _bookingNotice   string
    // 酒店状态 0:正常，-1:删除，-2:停售
    _status   *model.File
    // 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
    _creditCardTypes   string
    // 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
    _orbitTrack   string
    // 卖家酒店英文名称
    _nameE   string
    // 打标去标使用的 tagJson 字段
    _tagJson   string
    // 旺旺昵称
    _aliNick   string
    // 供应商标识，如果确实需要修改原来的供应商标识才需要填写，否则不需要填写，请谨慎使用。
    _supplier   string
    // 结算流程中的结算币种，如需对接请联系飞猪技术支持，请谨慎使用
    _settlementCurrency   string
    // 资源方酒店预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    _standardBookingNotice   string
    // 资源方酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    _standardHotelFacilities   string
    // 资源方酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    _standardHotelService   string
    // 资源方房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    _standardRoomFacilities   string
    // 资源方娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
    _standardAmuseFacilities   string
    // 0:酒店；1:客栈
    _hotelType   int64
    // 0:可以接待外宾；1:仅内宾
    _serviceType   int64
    // 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
    _coordinateSystem   string
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
func (r *TaobaoXhotelUpdateRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelUpdateRequest) GetHid() int64 {
    return r._hid
}
// Name Setter
// 酒店名称；（新增酒店时为必须）,国内酒店请传中文名称
func (r *TaobaoXhotelUpdateRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoXhotelUpdateRequest) GetName() string {
    return r._name
}
// UsedName Setter
// 酒店曾用名
func (r *TaobaoXhotelUpdateRequest) SetUsedName(_usedName string) error {
    r._usedName = _usedName
    r.Set("used_name", _usedName)
    return nil
}

// UsedName Getter
func (r TaobaoXhotelUpdateRequest) GetUsedName() string {
    return r._usedName
}
// Domestic Setter
// 是否国内酒店。0:国内;1:国外
func (r *TaobaoXhotelUpdateRequest) SetDomestic(_domestic int64) error {
    r._domestic = _domestic
    r.Set("domestic", _domestic)
    return nil
}

// Domestic Getter
func (r TaobaoXhotelUpdateRequest) GetDomestic() int64 {
    return r._domestic
}
// Country Setter
// domestic为true时，固定China； domestic为false时，必须传定义的海外国家编码值。参见：http://kezhan.trip.taobao.com/countrys.html
func (r *TaobaoXhotelUpdateRequest) SetCountry(_country string) error {
    r._country = _country
    r.Set("country", _country)
    return nil
}

// Country Getter
func (r TaobaoXhotelUpdateRequest) GetCountry() string {
    return r._country
}
// Province Setter
// 省份编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时默认为0
func (r *TaobaoXhotelUpdateRequest) SetProvince(_province int64) error {
    r._province = _province
    r.Set("province", _province)
    return nil
}

// Province Getter
func (r TaobaoXhotelUpdateRequest) GetProvince() int64 {
    return r._province
}
// City Setter
// 城市编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（新增酒店时为必须）
func (r *TaobaoXhotelUpdateRequest) SetCity(_city int64) error {
    r._city = _city
    r.Set("city", _city)
    return nil
}

// City Getter
func (r TaobaoXhotelUpdateRequest) GetCity() int64 {
    return r._city
}
// District Setter
// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm?tbpm=3
func (r *TaobaoXhotelUpdateRequest) SetDistrict(_district int64) error {
    r._district = _district
    r.Set("district", _district)
    return nil
}

// District Getter
func (r TaobaoXhotelUpdateRequest) GetDistrict() int64 {
    return r._district
}
// Business Setter
// 商业区（圈）长度不超过20字
func (r *TaobaoXhotelUpdateRequest) SetBusiness(_business string) error {
    r._business = _business
    r.Set("business", _business)
    return nil
}

// Business Getter
func (r TaobaoXhotelUpdateRequest) GetBusiness() string {
    return r._business
}
// Address Setter
// 酒店地址。长度不能超过255
func (r *TaobaoXhotelUpdateRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoXhotelUpdateRequest) GetAddress() string {
    return r._address
}
// Longitude Setter
// 经度
func (r *TaobaoXhotelUpdateRequest) SetLongitude(_longitude string) error {
    r._longitude = _longitude
    r.Set("longitude", _longitude)
    return nil
}

// Longitude Getter
func (r TaobaoXhotelUpdateRequest) GetLongitude() string {
    return r._longitude
}
// Latitude Setter
// 纬度
func (r *TaobaoXhotelUpdateRequest) SetLatitude(_latitude string) error {
    r._latitude = _latitude
    r.Set("latitude", _latitude)
    return nil
}

// Latitude Getter
func (r TaobaoXhotelUpdateRequest) GetLatitude() string {
    return r._latitude
}
// PositionType Setter
// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
func (r *TaobaoXhotelUpdateRequest) SetPositionType(_positionType string) error {
    r._positionType = _positionType
    r.Set("position_type", _positionType)
    return nil
}

// PositionType Getter
func (r TaobaoXhotelUpdateRequest) GetPositionType() string {
    return r._positionType
}
// Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelUpdateRequest) SetTel(_tel string) error {
    r._tel = _tel
    r.Set("tel", _tel)
    return nil
}

// Tel Getter
func (r TaobaoXhotelUpdateRequest) GetTel() string {
    return r._tel
}
// Extend Setter
// 不要使用
func (r *TaobaoXhotelUpdateRequest) SetExtend(_extend string) error {
    r._extend = _extend
    r.Set("extend", _extend)
    return nil
}

// Extend Getter
func (r TaobaoXhotelUpdateRequest) GetExtend() string {
    return r._extend
}
// Shid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelUpdateRequest) SetShid(_shid int64) error {
    r._shid = _shid
    r.Set("shid", _shid)
    return nil
}

// Shid Getter
func (r TaobaoXhotelUpdateRequest) GetShid() int64 {
    return r._shid
}
// OuterId Setter
// 必传，酒店标识，商家酒店ID
func (r *TaobaoXhotelUpdateRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelUpdateRequest) GetOuterId() string {
    return r._outerId
}
// Vendor Setter
// 系统商，一般情况不用，需申请使用
func (r *TaobaoXhotelUpdateRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelUpdateRequest) GetVendor() string {
    return r._vendor
}
// Star Setter
// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
func (r *TaobaoXhotelUpdateRequest) SetStar(_star string) error {
    r._star = _star
    r.Set("star", _star)
    return nil
}

// Star Getter
func (r TaobaoXhotelUpdateRequest) GetStar() string {
    return r._star
}
// OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelUpdateRequest) SetOpeningTime(_openingTime string) error {
    r._openingTime = _openingTime
    r.Set("opening_time", _openingTime)
    return nil
}

// OpeningTime Getter
func (r TaobaoXhotelUpdateRequest) GetOpeningTime() string {
    return r._openingTime
}
// DecorateTime Setter
// 装修时间，格式为2015-10-01
func (r *TaobaoXhotelUpdateRequest) SetDecorateTime(_decorateTime string) error {
    r._decorateTime = _decorateTime
    r.Set("decorate_time", _decorateTime)
    return nil
}

// DecorateTime Getter
func (r TaobaoXhotelUpdateRequest) GetDecorateTime() string {
    return r._decorateTime
}
// Floors Setter
// 楼层信息
func (r *TaobaoXhotelUpdateRequest) SetFloors(_floors string) error {
    r._floors = _floors
    r.Set("floors", _floors)
    return nil
}

// Floors Getter
func (r TaobaoXhotelUpdateRequest) GetFloors() string {
    return r._floors
}
// Rooms Setter
// 房间数 0~9999之内的数字
func (r *TaobaoXhotelUpdateRequest) SetRooms(_rooms int64) error {
    r._rooms = _rooms
    r.Set("rooms", _rooms)
    return nil
}

// Rooms Getter
func (r TaobaoXhotelUpdateRequest) GetRooms() int64 {
    return r._rooms
}
// Description Setter
// 酒店描述
func (r *TaobaoXhotelUpdateRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r TaobaoXhotelUpdateRequest) GetDescription() string {
    return r._description
}
// HotelFacilities Setter
// 酒店设施。json格式示例值：{"free Wi-Fi in all rooms":"true","massage":"true","meetingRoom":"true"}目前支持维护的设施枚举有：free Wi-Fi in all rooms 所有房间设有免费无线网络;meetingRoom 会议室;massage  按摩室;fitnessClub 健身房;bar 酒吧;cafe 咖啡厅;frontDeskSafe 前台贵重物品保险柜wifi 无线上网公共区域;casino 娱乐场/棋牌室;restaurant 餐厅;smoking area 吸烟区;Business Facilities 商务设施
func (r *TaobaoXhotelUpdateRequest) SetHotelFacilities(_hotelFacilities string) error {
    r._hotelFacilities = _hotelFacilities
    r.Set("hotel_facilities", _hotelFacilities)
    return nil
}

// HotelFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetHotelFacilities() string {
    return r._hotelFacilities
}
// Service Setter
// 酒店基础服务。json格式示例值：{"receiveForeignGuests":"true","morningCall":"true","breakfast":"true"}目前支持维护的设施枚举有：receiveForeignGuests 接待外宾;morningCall 叫醒服务; breakfast  早餐服务; airportShuttle 接机服务; luggageClaim 行李寄存; rentCar 租车; HourRoomService24 24小时客房服务; airportTransfer 酒店/机场接送; dryCleaning 干洗; expressCheckInCheckOut 快速入住/退房登记; custodyServices 保管服务
func (r *TaobaoXhotelUpdateRequest) SetService(_service string) error {
    r._service = _service
    r.Set("service", _service)
    return nil
}

// Service Getter
func (r TaobaoXhotelUpdateRequest) GetService() string {
    return r._service
}
// RoomFacilities Setter
// 房间的基础设施。json格式示例值：{"bathtub":"true","bathPub":"true"}目前支持维护的设施枚举有：bathtub 独立卫浴;bathPub 公共卫浴
func (r *TaobaoXhotelUpdateRequest) SetRoomFacilities(_roomFacilities string) error {
    r._roomFacilities = _roomFacilities
    r.Set("room_facilities", _roomFacilities)
    return nil
}

// RoomFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetRoomFacilities() string {
    return r._roomFacilities
}
// Pics Setter
// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址，main是否为主图（主图只能有一个）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多10张。要求：无logo、水印、边框、人物，不模糊、重复、歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
func (r *TaobaoXhotelUpdateRequest) SetPics(_pics string) error {
    r._pics = _pics
    r.Set("pics", _pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelUpdateRequest) GetPics() string {
    return r._pics
}
// Brand Setter
// 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）： ruJia("1", "rujiakuaijie", "如家快捷", 1), qiTian("2", "7 days", "7天连锁", 1), hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1), geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1), jinJiang("5", "Jinjiang Inn", "锦江之星", 1), su8("6", "Super 8", "速8", 1), moTai("7", "Motel", "莫泰", 1), zhouji("8", "InterContinental", "洲际", 4), budint("9", "Pod Inn", "布丁", 1), jiuJiu("10", "jiujiuliansuo", "99连锁", 1), piaoHome("11", "Piao Home Inn", "飘HOME", 1), juzi("12", "Orange Hotels", "桔子酒店", 1), yibai("13", "yibai", "易佰", 1), weiyena("14","weiyena","维也纳",2), huangguanjiari("15", "huangguanjiari", "皇冠假日", 4), xidawu("16", "xidawu", "喜达屋", 3), chengshiBJ("17", "chengshibianjie", "城市便捷", 1), shagnKeYou("18", "shagnkeyou", "尚客优", 1), jinjiang("19", "jinjiang", "锦江酒店", 3), wendemu("20", "Hawthorn Suites", "温德姆", 4), yibisi("21", "Ibis Hotels", "宜必思", 1), wanhao("22", "JM Hoteles", "万豪", 4), yijia365("23", "yijia365", "驿家365", 1), shoulv("24", "shoulvjituan", "首旅建国", 3), kaiyuan("25", "New Century Hotel", "开元大酒店", 4), yagao("26", "yagao", "雅高", 3), daisi("27", "daisi", "戴斯", 3), jinling("28", "jinlingliansuo", "金陵", 4), xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4), xierdun("30", "Hilton", "希尔顿", 4)
func (r *TaobaoXhotelUpdateRequest) SetBrand(_brand string) error {
    r._brand = _brand
    r.Set("brand", _brand)
    return nil
}

// Brand Getter
func (r TaobaoXhotelUpdateRequest) GetBrand() string {
    return r._brand
}
// PostalCode Setter
// 邮编
func (r *TaobaoXhotelUpdateRequest) SetPostalCode(_postalCode string) error {
    r._postalCode = _postalCode
    r.Set("postal_code", _postalCode)
    return nil
}

// PostalCode Getter
func (r TaobaoXhotelUpdateRequest) GetPostalCode() string {
    return r._postalCode
}
// HotelPolicies Setter
// 酒店入住政策(针对国际酒店，儿童及加床信息)格式：{"children_age_from":"2","children_age_to":"3","children_stay_free":"True","infant_age":"1","min_guest_age":"4"}
func (r *TaobaoXhotelUpdateRequest) SetHotelPolicies(_hotelPolicies string) error {
    r._hotelPolicies = _hotelPolicies
    r.Set("hotel_policies", _hotelPolicies)
    return nil
}

// HotelPolicies Getter
func (r TaobaoXhotelUpdateRequest) GetHotelPolicies() string {
    return r._hotelPolicies
}
// BookingNotice Setter
// 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
func (r *TaobaoXhotelUpdateRequest) SetBookingNotice(_bookingNotice string) error {
    r._bookingNotice = _bookingNotice
    r.Set("booking_notice", _bookingNotice)
    return nil
}

// BookingNotice Getter
func (r TaobaoXhotelUpdateRequest) GetBookingNotice() string {
    return r._bookingNotice
}
// Status Setter
// 酒店状态 0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelUpdateRequest) SetStatus(_status *model.File) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoXhotelUpdateRequest) GetStatus() *model.File {
    return r._status
}
// CreditCardTypes Setter
// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
func (r *TaobaoXhotelUpdateRequest) SetCreditCardTypes(_creditCardTypes string) error {
    r._creditCardTypes = _creditCardTypes
    r.Set("credit_card_types", _creditCardTypes)
    return nil
}

// CreditCardTypes Getter
func (r TaobaoXhotelUpdateRequest) GetCreditCardTypes() string {
    return r._creditCardTypes
}
// OrbitTrack Setter
// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
func (r *TaobaoXhotelUpdateRequest) SetOrbitTrack(_orbitTrack string) error {
    r._orbitTrack = _orbitTrack
    r.Set("orbit_track", _orbitTrack)
    return nil
}

// OrbitTrack Getter
func (r TaobaoXhotelUpdateRequest) GetOrbitTrack() string {
    return r._orbitTrack
}
// NameE Setter
// 卖家酒店英文名称
func (r *TaobaoXhotelUpdateRequest) SetNameE(_nameE string) error {
    r._nameE = _nameE
    r.Set("name_e", _nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelUpdateRequest) GetNameE() string {
    return r._nameE
}
// TagJson Setter
// 打标去标使用的 tagJson 字段
func (r *TaobaoXhotelUpdateRequest) SetTagJson(_tagJson string) error {
    r._tagJson = _tagJson
    r.Set("tag_json", _tagJson)
    return nil
}

// TagJson Getter
func (r TaobaoXhotelUpdateRequest) GetTagJson() string {
    return r._tagJson
}
// AliNick Setter
// 旺旺昵称
func (r *TaobaoXhotelUpdateRequest) SetAliNick(_aliNick string) error {
    r._aliNick = _aliNick
    r.Set("ali_nick", _aliNick)
    return nil
}

// AliNick Getter
func (r TaobaoXhotelUpdateRequest) GetAliNick() string {
    return r._aliNick
}
// Supplier Setter
// 供应商标识，如果确实需要修改原来的供应商标识才需要填写，否则不需要填写，请谨慎使用。
func (r *TaobaoXhotelUpdateRequest) SetSupplier(_supplier string) error {
    r._supplier = _supplier
    r.Set("supplier", _supplier)
    return nil
}

// Supplier Getter
func (r TaobaoXhotelUpdateRequest) GetSupplier() string {
    return r._supplier
}
// SettlementCurrency Setter
// 结算流程中的结算币种，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelUpdateRequest) SetSettlementCurrency(_settlementCurrency string) error {
    r._settlementCurrency = _settlementCurrency
    r.Set("settlement_currency", _settlementCurrency)
    return nil
}

// SettlementCurrency Getter
func (r TaobaoXhotelUpdateRequest) GetSettlementCurrency() string {
    return r._settlementCurrency
}
// StandardBookingNotice Setter
// 资源方酒店预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardBookingNotice(_standardBookingNotice string) error {
    r._standardBookingNotice = _standardBookingNotice
    r.Set("standard_booking_notice", _standardBookingNotice)
    return nil
}

// StandardBookingNotice Getter
func (r TaobaoXhotelUpdateRequest) GetStandardBookingNotice() string {
    return r._standardBookingNotice
}
// StandardHotelFacilities Setter
// 资源方酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardHotelFacilities(_standardHotelFacilities string) error {
    r._standardHotelFacilities = _standardHotelFacilities
    r.Set("standard_hotel_facilities", _standardHotelFacilities)
    return nil
}

// StandardHotelFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetStandardHotelFacilities() string {
    return r._standardHotelFacilities
}
// StandardHotelService Setter
// 资源方酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardHotelService(_standardHotelService string) error {
    r._standardHotelService = _standardHotelService
    r.Set("standard_hotel_service", _standardHotelService)
    return nil
}

// StandardHotelService Getter
func (r TaobaoXhotelUpdateRequest) GetStandardHotelService() string {
    return r._standardHotelService
}
// StandardRoomFacilities Setter
// 资源方房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardRoomFacilities(_standardRoomFacilities string) error {
    r._standardRoomFacilities = _standardRoomFacilities
    r.Set("standard_room_facilities", _standardRoomFacilities)
    return nil
}

// StandardRoomFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetStandardRoomFacilities() string {
    return r._standardRoomFacilities
}
// StandardAmuseFacilities Setter
// 资源方娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelUpdateRequest) SetStandardAmuseFacilities(_standardAmuseFacilities string) error {
    r._standardAmuseFacilities = _standardAmuseFacilities
    r.Set("standard_amuse_facilities", _standardAmuseFacilities)
    return nil
}

// StandardAmuseFacilities Getter
func (r TaobaoXhotelUpdateRequest) GetStandardAmuseFacilities() string {
    return r._standardAmuseFacilities
}
// HotelType Setter
// 0:酒店；1:客栈
func (r *TaobaoXhotelUpdateRequest) SetHotelType(_hotelType int64) error {
    r._hotelType = _hotelType
    r.Set("hotel_type", _hotelType)
    return nil
}

// HotelType Getter
func (r TaobaoXhotelUpdateRequest) GetHotelType() int64 {
    return r._hotelType
}
// ServiceType Setter
// 0:可以接待外宾；1:仅内宾
func (r *TaobaoXhotelUpdateRequest) SetServiceType(_serviceType int64) error {
    r._serviceType = _serviceType
    r.Set("service_type", _serviceType)
    return nil
}

// ServiceType Getter
func (r TaobaoXhotelUpdateRequest) GetServiceType() int64 {
    return r._serviceType
}
// CoordinateSystem Setter
// 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
func (r *TaobaoXhotelUpdateRequest) SetCoordinateSystem(_coordinateSystem string) error {
    r._coordinateSystem = _coordinateSystem
    r.Set("coordinate_system", _coordinateSystem)
    return nil
}

// CoordinateSystem Getter
func (r TaobaoXhotelUpdateRequest) GetCoordinateSystem() string {
    return r._coordinateSystem
}
