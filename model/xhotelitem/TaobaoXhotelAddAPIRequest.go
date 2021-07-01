package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelAddAPIRequest
酒店新增接口（ID重复自动更新） API请求
taobao.xhotel.add

添加酒店或更新酒店 */
type TaobaoXhotelAddAPIRequest struct {
	model.Params
	// 外部酒店ID, 这是卖家自己系统中的ID
	_outerId string
	// 酒店名称,国内酒店请传中文名称
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
	// 酒店地址。长度不能超过255。不填入会导致不能自动匹配。
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
	// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多是能是10张。
	_pics string
	// 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）：    ruJia("1", "rujiakuaijie", "如家快捷", 1),    qiTian("2", "7 days", "7天连锁", 1),    hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1),    geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1),    jinJiang("5", "Jinjiang Inn", "锦江之星", 1),    su8("6", "Super 8", "速8", 1),    moTai("7", "Motel", "莫泰", 1),    zhouji("8", "InterContinental", "洲际", 4),    budint("9", "Pod Inn", "布丁", 1),    jiuJiu("10", "jiujiuliansuo", "99连锁", 1),    piaoHome("11", "Piao Home Inn", "飘HOME", 1),    juzi("12", "Orange Hotels", "桔子酒店", 1),    yibai("13", "yibai", "易佰", 1),    weiyena("14","weiyena","维也纳",2),    huangguanjiari("15", "huangguanjiari", "皇冠假日", 4),    xidawu("16", "xidawu", "喜达屋", 3),    chengshiBJ("17", "chengshibianjie", "城市便捷", 1),    shagnKeYou("18", "shagnkeyou", "尚客优", 1),    jinjiang("19", "jinjiang", "锦江酒店", 3),    wendemu("20", "Hawthorn Suites", "温德姆", 4),    yibisi("21", "Ibis Hotels", "宜必思", 1),    wanhao("22", "JM Hoteles", "万豪", 4),    yijia365("23", "yijia365", "驿家365", 1),    shoulv("24", "shoulvjituan", "首旅建国", 3),    kaiyuan("25", "New Century Hotel", "开元大酒店", 4),    yagao("26", "yagao", "雅高", 3),    daisi("27", "daisi", "戴斯", 3),    jinling("28", "jinlingliansuo", "金陵", 4),    xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4),    xierdun("30", "Hilton", "希尔顿", 4),
	_brand string
	// 邮政编码。
	_postalCode string
	// 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
	_bookingNotice string
	// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
	_creditCardTypes string
	// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
	_orbitTrack string
	// 卖家酒店英文名称
	_nameE string
	// 供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用。注：如商家申请的应用类型为“飞猪-新业务”，此项则必填。
	_supplier string
	// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
	_settlementCurrency string
	// 标准娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardAmuseFacilities string
	// 标准房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardRoomFacilities string
	// 标准酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardHotelService string
	// 标准酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardHotelFacilities string
	// 标准预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
	_standardBookingNotice string
	// 0:可以接待外宾；1:仅内宾
	_serviceType int64
	// 0:酒店；1:客栈
	_hotelType int64
	// 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
	_coordinateSystem string
	// 废弃
	_roomFacilities string
	// 废弃
	_service string
	// 废弃
	_hotelFacilities string
	// 废弃
	_hotelPolicies string
}

// NewTaobaoXhotelAddRequest 初始化TaobaoXhotelAddAPIRequest对象
func NewTaobaoXhotelAddRequest() *TaobaoXhotelAddAPIRequest {
	return &TaobaoXhotelAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelAddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OuterId Setter
// 外部酒店ID, 这是卖家自己系统中的ID
func (r *TaobaoXhotelAddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// Get OuterId Getter
func (r TaobaoXhotelAddAPIRequest) GetOuterId() string {
	return r._outerId
}

// Set is Name Setter
// 酒店名称,国内酒店请传中文名称
func (r *TaobaoXhotelAddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// Get Name Getter
func (r TaobaoXhotelAddAPIRequest) GetName() string {
	return r._name
}

// Set is UsedName Setter
// 酒店曾用名
func (r *TaobaoXhotelAddAPIRequest) SetUsedName(_usedName string) error {
	r._usedName = _usedName
	r.Set("used_name", _usedName)
	return nil
}

// Get UsedName Getter
func (r TaobaoXhotelAddAPIRequest) GetUsedName() string {
	return r._usedName
}

// Set is Domestic Setter
// 是否国内酒店。0:国内;1:国外。默认是国内
func (r *TaobaoXhotelAddAPIRequest) SetDomestic(_domestic int64) error {
	r._domestic = _domestic
	r.Set("domestic", _domestic)
	return nil
}

// Get Domestic Getter
func (r TaobaoXhotelAddAPIRequest) GetDomestic() int64 {
	return r._domestic
}

// Set is Country Setter
// domestic为0时，固定China； domestic为1时，必须传定义的海外国家编码值。参见：http://hotel.alitrip.com/area.htm
func (r *TaobaoXhotelAddAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// Get Country Getter
func (r TaobaoXhotelAddAPIRequest) GetCountry() string {
	return r._country
}

// Set is Province Setter
// 省份编码。选填，不填入的时候已city字段为准.参见：http://hotel.alitrip.com/area.htm，domestic为false时默认为0
func (r *TaobaoXhotelAddAPIRequest) SetProvince(_province int64) error {
	r._province = _province
	r.Set("province", _province)
	return nil
}

// Get Province Getter
func (r TaobaoXhotelAddAPIRequest) GetProvince() int64 {
	return r._province
}

// Set is City Setter
// 城市编码。参见：http://hotel.alitrip.com/area.htm，domestic为false时，输入对应国家的海外城市编码，可调用海外城市查询接口获取；（更新酒店时为可选）
func (r *TaobaoXhotelAddAPIRequest) SetCity(_city int64) error {
	r._city = _city
	r.Set("city", _city)
	return nil
}

// Get City Getter
func (r TaobaoXhotelAddAPIRequest) GetCity() int64 {
	return r._city
}

// Set is District Setter
// 区域（县级市）编码。参见：http://hotel.alitrip.com/area.htm
func (r *TaobaoXhotelAddAPIRequest) SetDistrict(_district int64) error {
	r._district = _district
	r.Set("district", _district)
	return nil
}

// Get District Getter
func (r TaobaoXhotelAddAPIRequest) GetDistrict() int64 {
	return r._district
}

// Set is Business Setter
// 商业区（圈）长度不超过20字
func (r *TaobaoXhotelAddAPIRequest) SetBusiness(_business string) error {
	r._business = _business
	r.Set("business", _business)
	return nil
}

// Get Business Getter
func (r TaobaoXhotelAddAPIRequest) GetBusiness() string {
	return r._business
}

// Set is Address Setter
// 酒店地址。长度不能超过255。不填入会导致不能自动匹配。
func (r *TaobaoXhotelAddAPIRequest) SetAddress(_address string) error {
	r._address = _address
	r.Set("address", _address)
	return nil
}

// Get Address Getter
func (r TaobaoXhotelAddAPIRequest) GetAddress() string {
	return r._address
}

// Set is Longitude Setter
// 经度
func (r *TaobaoXhotelAddAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// Get Longitude Getter
func (r TaobaoXhotelAddAPIRequest) GetLongitude() string {
	return r._longitude
}

// Set is Latitude Setter
// 纬度
func (r *TaobaoXhotelAddAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// Get Latitude Getter
func (r TaobaoXhotelAddAPIRequest) GetLatitude() string {
	return r._latitude
}

// Set is PositionType Setter
// 坐标类型，现在支持：G – GoogleB – 百度A – 高德M – MapbarL – 灵图
func (r *TaobaoXhotelAddAPIRequest) SetPositionType(_positionType string) error {
	r._positionType = _positionType
	r.Set("position_type", _positionType)
	return nil
}

// Get PositionType Getter
func (r TaobaoXhotelAddAPIRequest) GetPositionType() string {
	return r._positionType
}

// Set is Tel Setter
// 酒店电话。格式：国家代码（最长6位）#区号（最长4位）#电话（最长20位）。国家代码提示：中国大陆0086、香港00852、澳门00853、台湾00886
func (r *TaobaoXhotelAddAPIRequest) SetTel(_tel string) error {
	r._tel = _tel
	r.Set("tel", _tel)
	return nil
}

// Get Tel Getter
func (r TaobaoXhotelAddAPIRequest) GetTel() string {
	return r._tel
}

// Set is Extend Setter
// 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
func (r *TaobaoXhotelAddAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// Get Extend Getter
func (r TaobaoXhotelAddAPIRequest) GetExtend() string {
	return r._extend
}

// Set is Shid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝酒店的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelAddAPIRequest) SetShid(_shid int64) error {
	r._shid = _shid
	r.Set("shid", _shid)
	return nil
}

// Get Shid Getter
func (r TaobaoXhotelAddAPIRequest) GetShid() int64 {
	return r._shid
}

// Set is Vendor Setter
// 对接系统商名称：可为空不要乱填，需要申请后使用
func (r *TaobaoXhotelAddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelAddAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is Star Setter
// 酒店档次，星级。取值范围为1,1.5,2,2.5,3,3.5,4,4.5,5
func (r *TaobaoXhotelAddAPIRequest) SetStar(_star string) error {
	r._star = _star
	r.Set("star", _star)
	return nil
}

// Get Star Getter
func (r TaobaoXhotelAddAPIRequest) GetStar() string {
	return r._star
}

// Set is OpeningTime Setter
// 开业时间，格式为2015-01-01
func (r *TaobaoXhotelAddAPIRequest) SetOpeningTime(_openingTime string) error {
	r._openingTime = _openingTime
	r.Set("opening_time", _openingTime)
	return nil
}

// Get OpeningTime Getter
func (r TaobaoXhotelAddAPIRequest) GetOpeningTime() string {
	return r._openingTime
}

// Set is DecorateTime Setter
// 装修时间，格式为2015-01-01装修时间
func (r *TaobaoXhotelAddAPIRequest) SetDecorateTime(_decorateTime string) error {
	r._decorateTime = _decorateTime
	r.Set("decorate_time", _decorateTime)
	return nil
}

// Get DecorateTime Getter
func (r TaobaoXhotelAddAPIRequest) GetDecorateTime() string {
	return r._decorateTime
}

// Set is Floors Setter
// 楼层信息。
func (r *TaobaoXhotelAddAPIRequest) SetFloors(_floors string) error {
	r._floors = _floors
	r.Set("floors", _floors)
	return nil
}

// Get Floors Getter
func (r TaobaoXhotelAddAPIRequest) GetFloors() string {
	return r._floors
}

// Set is Rooms Setter
// 房间数 0~9999之内的数字
func (r *TaobaoXhotelAddAPIRequest) SetRooms(_rooms int64) error {
	r._rooms = _rooms
	r.Set("rooms", _rooms)
	return nil
}

// Get Rooms Getter
func (r TaobaoXhotelAddAPIRequest) GetRooms() int64 {
	return r._rooms
}

// Set is Description Setter
// 酒店描述
func (r *TaobaoXhotelAddAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// Get Description Getter
func (r TaobaoXhotelAddAPIRequest) GetDescription() string {
	return r._description
}

// Set is Pics Setter
// 酒店图片只支持远程图片，格式如下：[{"url":"http://123.jpg","ismain":"false","type":"大堂","attribute":"普通图"},{"url":"http://456.jpg","ismain":"true","type":"公共区域","attribute":"全景图"},{"url":"http://789.jpg","ismain":"false","type":"大堂","attribute":"普通图"}] 其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图（主图只能有一个，如果有多个或者没有，则会报错）,attribute表示图片属性，取值范围只能是：[普通图, 平面图, 全景图] ,type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]，图片数量最多是能是10张。
func (r *TaobaoXhotelAddAPIRequest) SetPics(_pics string) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// Get Pics Getter
func (r TaobaoXhotelAddAPIRequest) GetPics() string {
	return r._pics
}

// Set is Brand Setter
// 酒店品牌。取值为数字。枚举如下（只给出top30，如果不满足，请联系去啊接口人）：    ruJia("1", "rujiakuaijie", "如家快捷", 1),    qiTian("2", "7 days", "7天连锁", 1),    hanTing("3", "Hanting Inns & Hotels", "汉庭酒店", 1),    geLinHaoTai("4", "Green Tree Inn", "格林豪泰", 1),    jinJiang("5", "Jinjiang Inn", "锦江之星", 1),    su8("6", "Super 8", "速8", 1),    moTai("7", "Motel", "莫泰", 1),    zhouji("8", "InterContinental", "洲际", 4),    budint("9", "Pod Inn", "布丁", 1),    jiuJiu("10", "jiujiuliansuo", "99连锁", 1),    piaoHome("11", "Piao Home Inn", "飘HOME", 1),    juzi("12", "Orange Hotels", "桔子酒店", 1),    yibai("13", "yibai", "易佰", 1),    weiyena("14","weiyena","维也纳",2),    huangguanjiari("15", "huangguanjiari", "皇冠假日", 4),    xidawu("16", "xidawu", "喜达屋", 3),    chengshiBJ("17", "chengshibianjie", "城市便捷", 1),    shagnKeYou("18", "shagnkeyou", "尚客优", 1),    jinjiang("19", "jinjiang", "锦江酒店", 3),    wendemu("20", "Hawthorn Suites", "温德姆", 4),    yibisi("21", "Ibis Hotels", "宜必思", 1),    wanhao("22", "JM Hoteles", "万豪", 4),    yijia365("23", "yijia365", "驿家365", 1),    shoulv("24", "shoulvjituan", "首旅建国", 3),    kaiyuan("25", "New Century Hotel", "开元大酒店", 4),    yagao("26", "yagao", "雅高", 3),    daisi("27", "daisi", "戴斯", 3),    jinling("28", "jinlingliansuo", "金陵", 4),    xianggelila("29", "Shangri-La City Hotels", "香格里拉", 4),    xierdun("30", "Hilton", "希尔顿", 4),
func (r *TaobaoXhotelAddAPIRequest) SetBrand(_brand string) error {
	r._brand = _brand
	r.Set("brand", _brand)
	return nil
}

// Get Brand Getter
func (r TaobaoXhotelAddAPIRequest) GetBrand() string {
	return r._brand
}

// Set is PostalCode Setter
// 邮政编码。
func (r *TaobaoXhotelAddAPIRequest) SetPostalCode(_postalCode string) error {
	r._postalCode = _postalCode
	r.Set("postal_code", _postalCode)
	return nil
}

// Get PostalCode Getter
func (r TaobaoXhotelAddAPIRequest) GetPostalCode() string {
	return r._postalCode
}

// Set is BookingNotice Setter
// 预订须知。json字段描述：hotelInMountaintop 酒店位于山顶 1在山顶、0不在；needBoat 酒店需要坐船前往 1需要、0不需要；酒店位于景区内 1在景区、0不在；extraBed 加床收费；extraCharge 额外收费；arrivalTime 到店时间；extend 其他补充项
func (r *TaobaoXhotelAddAPIRequest) SetBookingNotice(_bookingNotice string) error {
	r._bookingNotice = _bookingNotice
	r.Set("booking_notice", _bookingNotice)
	return nil
}

// Get BookingNotice Getter
func (r TaobaoXhotelAddAPIRequest) GetBookingNotice() string {
	return r._bookingNotice
}

// Set is CreditCardTypes Setter
// 逗号分隔的字符串 1visa；2万事达卡；3美国运通卡；4发现卡；5大来卡；6JCB卡；7银联卡
func (r *TaobaoXhotelAddAPIRequest) SetCreditCardTypes(_creditCardTypes string) error {
	r._creditCardTypes = _creditCardTypes
	r.Set("credit_card_types", _creditCardTypes)
	return nil
}

// Get CreditCardTypes Getter
func (r TaobaoXhotelAddAPIRequest) GetCreditCardTypes() string {
	return r._creditCardTypes
}

// Set is OrbitTrack Setter
// 扩展信息的JSON。 orbitTrack 业务字段是指从飞猪到酒店说经过平台名以及方式的一个数组，按顺序，从飞猪，再经过若干平台，最后到酒店， platform是指定当前平台名，ways 是指通过哪种方式到该平台 其中，飞猪到下一个平台里, ways 字段只能是【直连】、【人工】两个方式之一； 从最后一个平台到酒店的ways字段只能是【电话】、【传真】、【人工】、【系统】之一； 第一个 飞猪平台 和 最后具体酒店是至少得填的
func (r *TaobaoXhotelAddAPIRequest) SetOrbitTrack(_orbitTrack string) error {
	r._orbitTrack = _orbitTrack
	r.Set("orbit_track", _orbitTrack)
	return nil
}

// Get OrbitTrack Getter
func (r TaobaoXhotelAddAPIRequest) GetOrbitTrack() string {
	return r._orbitTrack
}

// Set is NameE Setter
// 卖家酒店英文名称
func (r *TaobaoXhotelAddAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// Get NameE Getter
func (r TaobaoXhotelAddAPIRequest) GetNameE() string {
	return r._nameE
}

// Set is Supplier Setter
// 供应商标识，需要提前开通权限，如果需要对接请联系技术支持，请谨慎使用。注：如商家申请的应用类型为“飞猪-新业务”，此项则必填。
func (r *TaobaoXhotelAddAPIRequest) SetSupplier(_supplier string) error {
	r._supplier = _supplier
	r.Set("supplier", _supplier)
	return nil
}

// Get Supplier Getter
func (r TaobaoXhotelAddAPIRequest) GetSupplier() string {
	return r._supplier
}

// Set is SettlementCurrency Setter
// 结算过程中的结算币种符合，如需对接请联系飞猪技术支持，请谨慎使用
func (r *TaobaoXhotelAddAPIRequest) SetSettlementCurrency(_settlementCurrency string) error {
	r._settlementCurrency = _settlementCurrency
	r.Set("settlement_currency", _settlementCurrency)
	return nil
}

// Get SettlementCurrency Getter
func (r TaobaoXhotelAddAPIRequest) GetSettlementCurrency() string {
	return r._settlementCurrency
}

// Set is StandardAmuseFacilities Setter
// 标准娱乐设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelAddAPIRequest) SetStandardAmuseFacilities(_standardAmuseFacilities string) error {
	r._standardAmuseFacilities = _standardAmuseFacilities
	r.Set("standard_amuse_facilities", _standardAmuseFacilities)
	return nil
}

// Get StandardAmuseFacilities Getter
func (r TaobaoXhotelAddAPIRequest) GetStandardAmuseFacilities() string {
	return r._standardAmuseFacilities
}

// Set is StandardRoomFacilities Setter
// 标准房型设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelAddAPIRequest) SetStandardRoomFacilities(_standardRoomFacilities string) error {
	r._standardRoomFacilities = _standardRoomFacilities
	r.Set("standard_room_facilities", _standardRoomFacilities)
	return nil
}

// Get StandardRoomFacilities Getter
func (r TaobaoXhotelAddAPIRequest) GetStandardRoomFacilities() string {
	return r._standardRoomFacilities
}

// Set is StandardHotelService Setter
// 标准酒店服务,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelAddAPIRequest) SetStandardHotelService(_standardHotelService string) error {
	r._standardHotelService = _standardHotelService
	r.Set("standard_hotel_service", _standardHotelService)
	return nil
}

// Get StandardHotelService Getter
func (r TaobaoXhotelAddAPIRequest) GetStandardHotelService() string {
	return r._standardHotelService
}

// Set is StandardHotelFacilities Setter
// 标准酒店设施,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelAddAPIRequest) SetStandardHotelFacilities(_standardHotelFacilities string) error {
	r._standardHotelFacilities = _standardHotelFacilities
	r.Set("standard_hotel_facilities", _standardHotelFacilities)
	return nil
}

// Get StandardHotelFacilities Getter
func (r TaobaoXhotelAddAPIRequest) GetStandardHotelFacilities() string {
	return r._standardHotelFacilities
}

// Set is StandardBookingNotice Setter
// 标准预订须知,参考文档https://open.alitrip.com/docs/doc.htm?docType=1&articleId=108891
func (r *TaobaoXhotelAddAPIRequest) SetStandardBookingNotice(_standardBookingNotice string) error {
	r._standardBookingNotice = _standardBookingNotice
	r.Set("standard_booking_notice", _standardBookingNotice)
	return nil
}

// Get StandardBookingNotice Getter
func (r TaobaoXhotelAddAPIRequest) GetStandardBookingNotice() string {
	return r._standardBookingNotice
}

// Set is ServiceType Setter
// 0:可以接待外宾；1:仅内宾
func (r *TaobaoXhotelAddAPIRequest) SetServiceType(_serviceType int64) error {
	r._serviceType = _serviceType
	r.Set("service_type", _serviceType)
	return nil
}

// Get ServiceType Getter
func (r TaobaoXhotelAddAPIRequest) GetServiceType() int64 {
	return r._serviceType
}

// Set is HotelType Setter
// 0:酒店；1:客栈
func (r *TaobaoXhotelAddAPIRequest) SetHotelType(_hotelType int64) error {
	r._hotelType = _hotelType
	r.Set("hotel_type", _hotelType)
	return nil
}

// Get HotelType Getter
func (r TaobaoXhotelAddAPIRequest) GetHotelType() int64 {
	return r._hotelType
}

// Set is CoordinateSystem Setter
// 标识坐标系类型。WGS84，表示地球坐标系 ；GCJ02，表示火星坐标系
func (r *TaobaoXhotelAddAPIRequest) SetCoordinateSystem(_coordinateSystem string) error {
	r._coordinateSystem = _coordinateSystem
	r.Set("coordinate_system", _coordinateSystem)
	return nil
}

// Get CoordinateSystem Getter
func (r TaobaoXhotelAddAPIRequest) GetCoordinateSystem() string {
	return r._coordinateSystem
}

// Set is RoomFacilities Setter
// 废弃
func (r *TaobaoXhotelAddAPIRequest) SetRoomFacilities(_roomFacilities string) error {
	r._roomFacilities = _roomFacilities
	r.Set("room_facilities", _roomFacilities)
	return nil
}

// Get RoomFacilities Getter
func (r TaobaoXhotelAddAPIRequest) GetRoomFacilities() string {
	return r._roomFacilities
}

// Set is Service Setter
// 废弃
func (r *TaobaoXhotelAddAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// Get Service Getter
func (r TaobaoXhotelAddAPIRequest) GetService() string {
	return r._service
}

// Set is HotelFacilities Setter
// 废弃
func (r *TaobaoXhotelAddAPIRequest) SetHotelFacilities(_hotelFacilities string) error {
	r._hotelFacilities = _hotelFacilities
	r.Set("hotel_facilities", _hotelFacilities)
	return nil
}

// Get HotelFacilities Getter
func (r TaobaoXhotelAddAPIRequest) GetHotelFacilities() string {
	return r._hotelFacilities
}

// Set is HotelPolicies Setter
// 废弃
func (r *TaobaoXhotelAddAPIRequest) SetHotelPolicies(_hotelPolicies string) error {
	r._hotelPolicies = _hotelPolicies
	r.Set("hotel_policies", _hotelPolicies)
	return nil
}

// Get HotelPolicies Getter
func (r TaobaoXhotelAddAPIRequest) GetHotelPolicies() string {
	return r._hotelPolicies
}
