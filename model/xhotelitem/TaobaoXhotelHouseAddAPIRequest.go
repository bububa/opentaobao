package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelHouseAddAPIRequest 非标准民宿房源添加 API请求
// taobao.xhotel.house.add
//
// 添加酒店或更新酒店
type TaobaoXhotelHouseAddAPIRequest struct {
	model.Params
	// 外部酒店ID, 这是卖家自己系统中的ID
	_outerId string
	// 酒店名称
	_name string
	// 酒店曾用名
	_usedName string
	// 是否国内酒店。0:国内;1:国外。默认是国内
	_domestic int64
	// domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm
	_country string
	// 省份编码。选填，不填入的时候已city字段为准.参见：http://hotel.alitrip.com/area.htm，domestic为false时默认为0
	_province int64
	// 城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新酒店时为可选）
	_city int64
	// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm
	_district int64
	// 商业区（圈）长度不超过20字
	_business string
	// 酒店地址。长度不能超过120。不填入会导致不能自动匹配,此地址为下单前展示给用户地址
	_address string
	// 经度
	_longitude string
	// 纬度
	_latitude string
	// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
	_positionType string
	// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
	_tel string
	// 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
	_extend string
	// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
	_shid int64
	// 对接系统商名称：可为空不要乱填，需要申请后使用
	_vendor string
	// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
	_star string
	// 开业时间，格式为2015-01-01
	_openingTime string
	// 装修时间，格式为2015-01-01装修时间
	_decorateTime string
	// 楼层信息。
	_floors string
	// 房间数 0~9999之内的数字
	_rooms int64
	// 酒店描述
	_description string
	// 酒店入住政策，{"10003":"仅2岁以上儿童可随同入住"}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10、code=10003,value为文字描述
	_hotelPolicies string
	// 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
	_hotelFacilities string
	// 酒店服务。json格式示例值：{"24058":"可以接待外宾","24198":"叫醒服务","24200":"洗衣服务"}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
	_service string
	// 房间设施。json格式示例值：{"24101": true,"24091": true,"24095": true}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
	_roomFacilities string
	// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多是能是10张。
	_pics string
	// 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）：    ruJia("1", "rujiakuaijie", "如家快捷", 1),    qiTian("2", "7 days", "7天连锁", 1),    hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1),    geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1),    jinJiang("5", "Jinjiang Inn", "锦江之星", 1),    su8("6", "Super 8", "速8", 1),    moTai("7", "Motel", "莫泰", 1),    zhouji("8", "InterContinental", "洲际", 4),    budint("9", "Pod Inn", "布丁", 1),    jiuJiu("10", "jiujiuliansuo", "99连锁", 1),    piaoHome("11", "Piao Home Inn", "飘HOME", 1),    juzi("12", "Orange Hotels", "桔子酒店", 1),    yibai("13", "yibai", "易佰", 1),    weiyena("14","weiyena","维也纳",2),    huangguanjiari("15", "huangguanjiari", "皇冠假日", 4),    xidawu("16", "xidawu", "喜达屋", 3),    chengshiBJ("17", "chengshibianjie", "城市便捷", 1),    shagnKeYou("18", "shagnkeyou", "尚客优", 1),    jinjiang("19", "jinjiang", "锦江酒店", 3),    wendemu("20", "Hawthorn Suites", "温德姆", 4),    yibisi("21", "Ibis Hotels", "宜必思", 1),    wanhao("22", "JM Hoteles", "万豪", 4),    yijia365("23", "yijia365", "驿家365", 1),    shoulv("24", "shoulvjituan", "首旅建国", 3),    kaiyuan("25", "New Century Hotel", "开元大酒店", 4),    yagao("26", "yagao", "雅高", 3),    daisi("27", "daisi", "戴斯", 3),    jinling("28", "jinlingliansuo", "金陵", 4),    xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4),    xierdun("30", "Hilton", "希尔顿", 4),
	_brand string
	// 邮政编码。
	_postalCode string
	// 预订须知。json格式，示例:{"10001":"14:00","10002":"12:00","10005":"清洁福50元","10006":"请准备好您的身份证件，我需要登记 不允许吸烟"},预订须知，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10的分类
	_bookingNotice string
	// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
	_creditCardTypes string
	// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
	_orbitTrack string
	// 卖家酒店英文名称
	_nameE string
	// 供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用
	_supplier string
	// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
	_settlementCurrency string
	// 房东信息,{"outerId: 外部房东ID,": "xxxx", "nickName": "张三", "avatarUrl": "http://test.com/1.jpg", "telephone": "0571-1234567", "mobilePhone": "12334567678", "email":"test@test.com", "gender": "F", "avgConfirmTime": 30, "responseRate": 100, "description": "房东太懒,什么也没有填", "birthday":"2018-01-01", "qualifacation": 1, "bloodType": 1, "profession":"交互设计师", "country":"CN", "province":"420000", "city":"421200", "real_name_status": true, "validate":"1,2,4,8","confirmRate": 98} JSON字段描述: outerId: 商家房东ID, nickName: 房东昵称, avatarUrl: 房东头像地址, telephone: 固定电话, mobilePhone: 移动电话, email: 邮箱地址, gender: 性别 M男性， F女性， avgConfirmTime: 平均确认时间, 单位分钟, responseRate: 房东回复率, description: 房东介绍, birthday:生日，格式yyyy-MM-dd, qualifacation:学历,1:小学,2:初中,3:高中,4:本科,5:硕士,6:博士,7:博士后,0:其他, profession: 职业 country: 国家code province: 省code city: 城市code realNameStatus: 实名认证状态, true已认证 validate: 认证情况:1:身份验证,2:头像验证,4:手机验证,8:邮箱验证,二进制各位代表含义, bloodType: 血型: 0未知,1:A型,2:B型,3:AB型,4:O型;confirmRate: 订单接单率，0-100
	_ownerInfo string
	// 描述信息，inside： 内部描述,traffic:交通情况,arround:周边情况
	_arroundDesc string
	// 用户下单之后展示的详细地址
	_realAddress string
	// 数据状态 0:正常，-1:删除，-2:停售
	_status int64
}

// NewTaobaoXhotelHouseAddRequest 初始化TaobaoXhotelHouseAddAPIRequest对象
func NewTaobaoXhotelHouseAddRequest() *TaobaoXhotelHouseAddAPIRequest {
	return &TaobaoXhotelHouseAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelHouseAddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.house.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelHouseAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuterId is OuterId Setter
// 外部酒店ID, 这是卖家自己系统中的ID
func (r *TaobaoXhotelHouseAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetName is Name Setter
// 酒店名称
func (r *TaobaoXhotelHouseAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetName() string {
	return r._name
}

// SetUsedName is UsedName Setter
// 酒店曾用名
func (r *TaobaoXhotelHouseAddAPIRequest) SetUsedName(_usedName string) error {
	r._usedName = _usedName
	r.Set("used_name", _usedName)
	return nil
}

// GetUsedName UsedName Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetUsedName() string {
	return r._usedName
}

// SetDomestic is Domestic Setter
// 是否国内酒店。0:国内;1:国外。默认是国内
func (r *TaobaoXhotelHouseAddAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// GetDomestic Domestic Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetDomestic() int64 {
	return r._domestic
}

// SetCountry is Country Setter
// domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm
func (r *TaobaoXhotelHouseAddAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetCountry() string {
	return r._country
}

// SetProvince is Province Setter
// 省份编码。选填，不填入的时候已city字段为准.参见：http://hotel.alitrip.com/area.htm，domestic为false时默认为0
func (r *TaobaoXhotelHouseAddAPIRequest) SetProvince(_province int64) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// GetProvince Province Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetProvince() int64 {
	return r._province
}

// SetCity is City Setter
// 城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新酒店时为可选）
func (r *TaobaoXhotelHouseAddAPIRequest) SetCity(_city int64) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// GetCity City Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetCity() int64 {
	return r._city
}

// SetDistrict is District Setter
// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm
func (r *TaobaoXhotelHouseAddAPIRequest) SetDistrict(_district int64) error {
	r._district = _district
	r.Set("district", _district)
	return nil
}

// GetDistrict District Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetDistrict() int64 {
	return r._district
}

// SetBusiness is Business Setter
// 商业区（圈）长度不超过20字
func (r *TaobaoXhotelHouseAddAPIRequest) SetBusiness(_business string) error {
	r._business = _business
	r.Set("business", _business)
	return nil
}

// GetBusiness Business Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetBusiness() string {
	return r._business
}

// SetAddress is Address Setter
// 酒店地址。长度不能超过120。不填入会导致不能自动匹配,此地址为下单前展示给用户地址
func (r *TaobaoXhotelHouseAddAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// GetAddress Address Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetAddress() string {
	return r._address
}

// SetLongitude is Longitude Setter
// 经度
func (r *TaobaoXhotelHouseAddAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度
func (r *TaobaoXhotelHouseAddAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetPositionType is PositionType Setter
// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
func (r *TaobaoXhotelHouseAddAPIRequest) SetPositionType(_positionType string) error {
	r._positionType = _positionType
	r.Set("position_type", _positionType)
	return nil
}

// GetPositionType PositionType Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetPositionType() string {
	return r._positionType
}

// SetTel is Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelHouseAddAPIRequest) SetTel(_tel string) error {
	r._tel = _tel
	r.Set("tel", _tel)
	return nil
}

// GetTel Tel Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetTel() string {
	return r._tel
}

// SetExtend is Extend Setter
// 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
func (r *TaobaoXhotelHouseAddAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetExtend() string {
	return r._extend
}

// SetShid is Shid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelHouseAddAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// GetShid Shid Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetShid() int64 {
	return r._shid
}

// SetVendor is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoXhotelHouseAddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetVendor() string {
	return r._vendor
}

// SetStar is Star Setter
// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
func (r *TaobaoXhotelHouseAddAPIRequest) SetStar(_star string) error {
	r._star = _star
	r.Set("star", _star)
	return nil
}

// GetStar Star Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetStar() string {
	return r._star
}

// SetOpeningTime is OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelHouseAddAPIRequest) SetOpeningTime(_openingTime string) error {
	r._openingTime = _openingTime
	r.Set("opening_time", _openingTime)
	return nil
}

// GetOpeningTime OpeningTime Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetOpeningTime() string {
	return r._openingTime
}

// SetDecorateTime is DecorateTime Setter
// 装修时间，格式为2015-01-01装修时间
func (r *TaobaoXhotelHouseAddAPIRequest) SetDecorateTime(_decorateTime string) error {
	r._decorateTime = _decorateTime
	r.Set("decorate_time", _decorateTime)
	return nil
}

// GetDecorateTime DecorateTime Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetDecorateTime() string {
	return r._decorateTime
}

// SetFloors is Floors Setter
// 楼层信息。
func (r *TaobaoXhotelHouseAddAPIRequest) SetFloors(_floors string) error {
	r._floors = _floors
	r.Set("floors", _floors)
	return nil
}

// GetFloors Floors Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetFloors() string {
	return r._floors
}

// SetRooms is Rooms Setter
// 房间数 0~9999之内的数字
func (r *TaobaoXhotelHouseAddAPIRequest) SetRooms(_rooms int64) error {
	r._rooms = _rooms
	r.Set("rooms", _rooms)
	return nil
}

// GetRooms Rooms Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetRooms() int64 {
	return r._rooms
}

// SetDescription is Description Setter
// 酒店描述
func (r *TaobaoXhotelHouseAddAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetDescription() string {
	return r._description
}

// SetHotelPolicies is HotelPolicies Setter
// 酒店入住政策，{"10003":"仅2岁以上儿童可随同入住"}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10、code=10003,value为文字描述
func (r *TaobaoXhotelHouseAddAPIRequest) SetHotelPolicies(_hotelPolicies string) error {
	r._hotelPolicies = _hotelPolicies
	r.Set("hotel_policies", _hotelPolicies)
	return nil
}

// GetHotelPolicies HotelPolicies Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetHotelPolicies() string {
	return r._hotelPolicies
}

// SetHotelFacilities is HotelFacilities Setter
// 酒店设施。json格式示例值：{"24152":true,"24149":true,"24150":true}，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=20的分类
func (r *TaobaoXhotelHouseAddAPIRequest) SetHotelFacilities(_hotelFacilities string) error {
	r._hotelFacilities = _hotelFacilities
	r.Set("hotel_facilities", _hotelFacilities)
	return nil
}

// GetHotelFacilities HotelFacilities Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetHotelFacilities() string {
	return r._hotelFacilities
}

// SetService is Service Setter
// 酒店服务。json格式示例值：{"24058":"可以接待外宾","24198":"叫醒服务","24200":"洗衣服务"}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
func (r *TaobaoXhotelHouseAddAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetService() string {
	return r._service
}

// SetRoomFacilities is RoomFacilities Setter
// 房间设施。json格式示例值：{"24101": true,"24091": true,"24095": true}，key-24101为属性编码，value-为"true"时表示有该属性，为文字时表示具体描述，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=40的分类
func (r *TaobaoXhotelHouseAddAPIRequest) SetRoomFacilities(_roomFacilities string) error {
	r._roomFacilities = _roomFacilities
	r.Set("room_facilities", _roomFacilities)
	return nil
}

// GetRoomFacilities RoomFacilities Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetRoomFacilities() string {
	return r._roomFacilities
}

// SetPics is Pics Setter
// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多是能是10张。
func (r *TaobaoXhotelHouseAddAPIRequest) SetPics(_pics string) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// GetPics Pics Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetPics() string {
	return r._pics
}

// SetBrand is Brand Setter
// 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）：    ruJia("1", "rujiakuaijie", "如家快捷", 1),    qiTian("2", "7 days", "7天连锁", 1),    hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1),    geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1),    jinJiang("5", "Jinjiang Inn", "锦江之星", 1),    su8("6", "Super 8", "速8", 1),    moTai("7", "Motel", "莫泰", 1),    zhouji("8", "InterContinental", "洲际", 4),    budint("9", "Pod Inn", "布丁", 1),    jiuJiu("10", "jiujiuliansuo", "99连锁", 1),    piaoHome("11", "Piao Home Inn", "飘HOME", 1),    juzi("12", "Orange Hotels", "桔子酒店", 1),    yibai("13", "yibai", "易佰", 1),    weiyena("14","weiyena","维也纳",2),    huangguanjiari("15", "huangguanjiari", "皇冠假日", 4),    xidawu("16", "xidawu", "喜达屋", 3),    chengshiBJ("17", "chengshibianjie", "城市便捷", 1),    shagnKeYou("18", "shagnkeyou", "尚客优", 1),    jinjiang("19", "jinjiang", "锦江酒店", 3),    wendemu("20", "Hawthorn Suites", "温德姆", 4),    yibisi("21", "Ibis Hotels", "宜必思", 1),    wanhao("22", "JM Hoteles", "万豪", 4),    yijia365("23", "yijia365", "驿家365", 1),    shoulv("24", "shoulvjituan", "首旅建国", 3),    kaiyuan("25", "New Century Hotel", "开元大酒店", 4),    yagao("26", "yagao", "雅高", 3),    daisi("27", "daisi", "戴斯", 3),    jinling("28", "jinlingliansuo", "金陵", 4),    xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4),    xierdun("30", "Hilton", "希尔顿", 4),
func (r *TaobaoXhotelHouseAddAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// GetBrand Brand Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetBrand() string {
	return r._brand
}

// SetPostalCode is PostalCode Setter
// 邮政编码。
func (r *TaobaoXhotelHouseAddAPIRequest) SetPostalCode(_postalCode string) error {
	r._postalCode = _postalCode
	r.Set("postal_code", _postalCode)
	return nil
}

// GetPostalCode PostalCode Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetPostalCode() string {
	return r._postalCode
}

// SetBookingNotice is BookingNotice Setter
// 预订须知。json格式，示例:{"10001":"14:00","10002":"12:00","10005":"清洁福50元","10006":"请准备好您的身份证件，我需要登记 不允许吸烟"},预订须知，key取值参见 https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.vSVPks&docType=1&articleId=108416&previewCode=987A11324A278EF679E24102BA30D426 中type=10的分类
func (r *TaobaoXhotelHouseAddAPIRequest) SetBookingNotice(_bookingNotice string) error {
	r._bookingNotice = _bookingNotice
	r.Set("booking_notice", _bookingNotice)
	return nil
}

// GetBookingNotice BookingNotice Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetBookingNotice() string {
	return r._bookingNotice
}

// SetCreditCardTypes is CreditCardTypes Setter
// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
func (r *TaobaoXhotelHouseAddAPIRequest) SetCreditCardTypes(_creditCardTypes string) error {
	r._creditCardTypes = _creditCardTypes
	r.Set("credit_card_types", _creditCardTypes)
	return nil
}

// GetCreditCardTypes CreditCardTypes Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetCreditCardTypes() string {
	return r._creditCardTypes
}

// SetOrbitTrack is OrbitTrack Setter
// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
func (r *TaobaoXhotelHouseAddAPIRequest) SetOrbitTrack(_orbitTrack string) error {
	r._orbitTrack = _orbitTrack
	r.Set("orbit_track", _orbitTrack)
	return nil
}

// GetOrbitTrack OrbitTrack Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetOrbitTrack() string {
	return r._orbitTrack
}

// SetNameE is NameE Setter
// 卖家酒店英文名称
func (r *TaobaoXhotelHouseAddAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// GetNameE NameE Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetNameE() string {
	return r._nameE
}

// SetSupplier is Supplier Setter
// 供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用
func (r *TaobaoXhotelHouseAddAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// GetSupplier Supplier Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetSupplier() string {
	return r._supplier
}

// SetSettlementCurrency is SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelHouseAddAPIRequest) SetSettlementCurrency(_settlementCurrency string) error {
	r._settlementCurrency = _settlementCurrency
	r.Set("settlement_currency", _settlementCurrency)
	return nil
}

// GetSettlementCurrency SettlementCurrency Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetSettlementCurrency() string {
	return r._settlementCurrency
}

// SetOwnerInfo is OwnerInfo Setter
// 房东信息,{"outerId: 外部房东ID,": "xxxx", "nickName": "张三", "avatarUrl": "http://test.com/1.jpg", "telephone": "0571-1234567", "mobilePhone": "12334567678", "email":"test@test.com", "gender": "F", "avgConfirmTime": 30, "responseRate": 100, "description": "房东太懒,什么也没有填", "birthday":"2018-01-01", "qualifacation": 1, "bloodType": 1, "profession":"交互设计师", "country":"CN", "province":"420000", "city":"421200", "real_name_status": true, "validate":"1,2,4,8","confirmRate": 98} JSON字段描述: outerId: 商家房东ID, nickName: 房东昵称, avatarUrl: 房东头像地址, telephone: 固定电话, mobilePhone: 移动电话, email: 邮箱地址, gender: 性别 M男性， F女性， avgConfirmTime: 平均确认时间, 单位分钟, responseRate: 房东回复率, description: 房东介绍, birthday:生日，格式yyyy-MM-dd, qualifacation:学历,1:小学,2:初中,3:高中,4:本科,5:硕士,6:博士,7:博士后,0:其他, profession: 职业 country: 国家code province: 省code city: 城市code realNameStatus: 实名认证状态, true已认证 validate: 认证情况:1:身份验证,2:头像验证,4:手机验证,8:邮箱验证,二进制各位代表含义, bloodType: 血型: 0未知,1:A型,2:B型,3:AB型,4:O型;confirmRate: 订单接单率，0-100
func (r *TaobaoXhotelHouseAddAPIRequest) SetOwnerInfo(_ownerInfo string) error {
	r._ownerInfo = _ownerInfo
	r.Set("owner_info", _ownerInfo)
	return nil
}

// GetOwnerInfo OwnerInfo Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetOwnerInfo() string {
	return r._ownerInfo
}

// SetArroundDesc is ArroundDesc Setter
// 描述信息，inside： 内部描述,traffic:交通情况,arround:周边情况
func (r *TaobaoXhotelHouseAddAPIRequest) SetArroundDesc(_arroundDesc string) error {
	r._arroundDesc = _arroundDesc
	r.Set("arround_desc", _arroundDesc)
	return nil
}

// GetArroundDesc ArroundDesc Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetArroundDesc() string {
	return r._arroundDesc
}

// SetRealAddress is RealAddress Setter
// 用户下单之后展示的详细地址
func (r *TaobaoXhotelHouseAddAPIRequest) SetRealAddress(_realAddress string) error {
	r._realAddress = _realAddress
	r.Set("real_address", _realAddress)
	return nil
}

// GetRealAddress RealAddress Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetRealAddress() string {
	return r._realAddress
}

// SetStatus is Status Setter
// 数据状态 0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelHouseAddAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoXhotelHouseAddAPIRequest) GetStatus() int64 {
	return r._status
}
