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

// New
