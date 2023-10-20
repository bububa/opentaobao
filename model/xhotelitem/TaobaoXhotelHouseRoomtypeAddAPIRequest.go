package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelhouseroomtypeaddAPIRequest 民宿房型信息添加 API请求
// taobao.xhotel.house.roomtype.add
//
// 房型添加或更新
type TaobaoxhotelhouseroomtypeaddAPIRequest struct {
	model.Params
	// 房型名称。不能超过30字
	_name string
	// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
	_area string
	// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
	_floor string
	// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
	_internet string
	// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
	_service string
	// 不要使用
	_extend string
	// 卖家房型ID，不能重复建议格式是:酒店code_房型code
	_outerId string
	// 系统商，无申请不可使用
	_vendor string
	// （必传）商家酒店ID，指明该房型属于哪家酒店
	_outHid string
	// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
	_pics string
	// 卖家房型英文名称
	_nameE string
	// 操作人信息
	_operator string
	// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房
	_houseModel string
	// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
	_bedInfo string
	// （已废弃）请使用outHid
	_hid int64
	// 最大入住人数，默认2（1-99）
	_maxOccupancy int64
	// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
	_windowType int64
	// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
	_srid int64
	// 属性值为1: 含义是非直连房型
	_connectionType int64
	// 房屋面积
	_houseSize int64
	// 出租类型:1.整租;2.单间;3.床位
	_rentType int64
	// 出租面积,单位平方米
	_rentSize int64
	// 是否和房东合住:0.不和房东合住;1.和房东合住;
	_hasLandlord int64
	// 数据状态 0:正常，-1:删除，-2:停售
	_status int64
}

// NewTaobaoxhotelhouseroomtypeaddRequest 初始化TaobaoxhotelhouseroomtypeaddAPIRequest对象
func NewTaobaoxhotelhouseroomtypeaddRequest() *TaobaoxhotelhouseroomtypeaddAPIRequest {
	return &TaobaoxhotelhouseroomtypeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.house.roomtype.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 房型名称。不能超过30字
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetName() string {
	return r._name
}

// SetArea is Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetArea() string {
	return r._area
}

// SetFloor is Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetFloor(_floor string) error {
	r._floor = _floor
	r.Set("floor", _floor)
	return nil
}

// GetFloor Floor Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetFloor() string {
	return r._floor
}

// SetInternet is Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetInternet(_internet string) error {
	r._internet = _internet
	r.Set("internet", _internet)
	return nil
}

// GetInternet Internet Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetInternet() string {
	return r._internet
}

// SetService is Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&#34;bar&#34;:false,&#34;catv&#34;:false,&#34;ddd&#34;:false,&#34;idd&#34;:false,&#34;pubtoilet&#34;:false,&#34;toilet&#34;:false}
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetService() string {
	return r._service
}

// SetExtend is Extend Setter
// 不要使用
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetExtend() string {
	return r._extend
}

// SetOuterId is OuterId Setter
// 卖家房型ID，不能重复建议格式是:酒店code_房型code
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 系统商，无申请不可使用
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetVendor() string {
	return r._vendor
}

// SetOutHid is OutHid Setter
// （必传）商家酒店ID，指明该房型属于哪家酒店
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetOutHid(_outHid string) error {
	r._outHid = _outHid
	r.Set("out_hid", _outHid)
	return nil
}

// GetOutHid OutHid Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetOutHid() string {
	return r._outHid
}

// SetPics is Pics Setter
// 房型图片只支持远程图片，格式如下：[{&#34;url&#34;:&#34;http://taobao.com/123.jpg&#34;,&#34;ismain&#34;:&#34;true&#34;},{&#34;url&#34;:&#34;http://taobao.com/456.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;},{&#34;url&#34;:&#34;http://taobao.com/789.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetPics(_pics string) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// GetPics Pics Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetPics() string {
	return r._pics
}

// SetNameE is NameE Setter
// 卖家房型英文名称
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// GetNameE NameE Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetNameE() string {
	return r._nameE
}

// SetOperator is Operator Setter
// 操作人信息
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetOperator() string {
	return r._operator
}

// SetHouseModel is HouseModel Setter
// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetHouseModel(_houseModel string) error {
	r._houseModel = _houseModel
	r.Set("house_model", _houseModel)
	return nil
}

// GetHouseModel HouseModel Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetHouseModel() string {
	return r._houseModel
}

// SetBedInfo is BedInfo Setter
// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&amp;docType=1&amp;articleId=108347
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetBedInfo(_bedInfo string) error {
	r._bedInfo = _bedInfo
	r.Set("bed_info", _bedInfo)
	return nil
}

// GetBedInfo BedInfo Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetBedInfo() string {
	return r._bedInfo
}

// SetHid is Hid Setter
// （已废弃）请使用outHid
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetHid() int64 {
	return r._hid
}

// SetMaxOccupancy is MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetMaxOccupancy(_maxOccupancy int64) error {
	r._maxOccupancy = _maxOccupancy
	r.Set("max_occupancy", _maxOccupancy)
	return nil
}

// GetMaxOccupancy MaxOccupancy Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetMaxOccupancy() int64 {
	return r._maxOccupancy
}

// SetWindowType is WindowType Setter
// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetWindowType(_windowType int64) error {
	r._windowType = _windowType
	r.Set("window_type", _windowType)
	return nil
}

// GetWindowType WindowType Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetWindowType() int64 {
	return r._windowType
}

// SetSrid is Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// GetSrid Srid Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetSrid() int64 {
	return r._srid
}

// SetConnectionType is ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetConnectionType(_connectionType int64) error {
	r._connectionType = _connectionType
	r.Set("connection_type", _connectionType)
	return nil
}

// GetConnectionType ConnectionType Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetConnectionType() int64 {
	return r._connectionType
}

// SetHouseSize is HouseSize Setter
// 房屋面积
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetHouseSize(_houseSize int64) error {
	r._houseSize = _houseSize
	r.Set("house_size", _houseSize)
	return nil
}

// GetHouseSize HouseSize Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetHouseSize() int64 {
	return r._houseSize
}

// SetRentType is RentType Setter
// 出租类型:1.整租;2.单间;3.床位
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetRentType(_rentType int64) error {
	r._rentType = _rentType
	r.Set("rent_type", _rentType)
	return nil
}

// GetRentType RentType Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetRentType() int64 {
	return r._rentType
}

// SetRentSize is RentSize Setter
// 出租面积,单位平方米
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetRentSize(_rentSize int64) error {
	r._rentSize = _rentSize
	r.Set("rent_size", _rentSize)
	return nil
}

// GetRentSize RentSize Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetRentSize() int64 {
	return r._rentSize
}

// SetHasLandlord is HasLandlord Setter
// 是否和房东合住:0.不和房东合住;1.和房东合住;
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetHasLandlord(_hasLandlord int64) error {
	r._hasLandlord = _hasLandlord
	r.Set("has_landlord", _hasLandlord)
	return nil
}

// GetHasLandlord HasLandlord Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetHasLandlord() int64 {
	return r._hasLandlord
}

// SetStatus is Status Setter
// 数据状态 0:正常，-1:删除，-2:停售
func (r *TaobaoxhotelhouseroomtypeaddAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoxhotelhouseroomtypeaddAPIRequest) GetStatus() int64 {
	return r._status
}
