package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomtypeUpdateAPIRequest 房型更新接口（ID不存在自动新增） API请求
// taobao.xhotel.roomtype.update
//
// 酒店房型更新或添加
type TaobaoXhotelRoomtypeUpdateAPIRequest struct {
	model.Params
	// 房型名称。不能超过200字；添加房型时为必须
	_name string
	// 具体面积大小，请按照正确格式填写。两种格式，例如：40 或者 10-20
	_area string
	// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
	_floor string
	// 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
	_bedType string
	// 床宽。按自己定义存储，比如：2.1米
	_bedSize string
	// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
	_internet string
	// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
	_service string
	// 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
	_extend string
	// （必传）商家房型ID
	_outerId string
	// 系统商，不要使用，只有申请才可用
	_vendor string
	// 商家酒店ID(如果更新房型的时候房型不存在，会拿该code去新增房型)
	_hotelCode string
	// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。要求：无logo、水印、边框、人物，不模糊，不重复，不歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
	_pics string
	// 卖家房型英文名称
	_nameE string
	// 操作人信息
	_operator string
	// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
	_bedInfo string
	// 新的房型编码，请确实需要修改某个房型的编码的时候才使用，如需使用，请联系飞猪技术支持开通权限，否则会更新失败
	_newOuterId string
	// 房间设施
	_standardRoomFacilities string
	// 窗型，窗型（windowType）： 0=无窗 1=有窗 2=部分有窗  窗型缺陷（windowTypeDefect）： 0=窗户不可打开通风 1=窗外有遮挡 2=窗外为酒店内景观  特殊窗（windowTypeSpecial）： 0=落地窗 1=飘窗 2=天窗 3=小窗。  当为有窗或部分有窗时，窗型缺陷可多选，特殊窗型需单选
	_windowDesc string
	// 房型加床政策。 费用(fee) 说明(desc)
	_addBed string
	// 儿童政策
	_childrenPolicy string
	// （已废弃）
	_rid int64
	// 最大入住人数，默认2（1-99）
	_maxOccupancy int64
	// 0:无窗/1:有窗
	_windowType int64
	// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
	_srid int64
	// （已废弃）
	_hid int64
	// 房型状态。0:正常，-1:删除，-2:停售
	_status *model.File
	// 属性值为1: 含义是非直连房型
	_connectionType int64
}

// NewTaobaoXhotelRoomtypeUpdateRequest 初始化TaobaoXhotelRoomtypeUpdateAPIRequest对象
func NewTaobaoXhotelRoomtypeUpdateRequest() *TaobaoXhotelRoomtypeUpdateAPIRequest {
	return &TaobaoXhotelRoomtypeUpdateAPIRequest{
		Params: model.NewParams(27),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) Reset() {
	r._name = ""
	r._area = ""
	r._floor = ""
	r._bedType = ""
	r._bedSize = ""
	r._internet = ""
	r._service = ""
	r._extend = ""
	r._outerId = ""
	r._vendor = ""
	r._hotelCode = ""
	r._pics = ""
	r._nameE = ""
	r._operator = ""
	r._bedInfo = ""
	r._newOuterId = ""
	r._standardRoomFacilities = ""
	r._windowDesc = ""
	r._addBed = ""
	r._childrenPolicy = ""
	r._rid = 0
	r._maxOccupancy = 0
	r._windowType = 0
	r._srid = 0
	r._hid = 0
	r._status = nil
	r._connectionType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.roomtype.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 房型名称。不能超过200字；添加房型时为必须
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetName() string {
	return r._name
}

// SetArea is Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40 或者 10-20
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetArea(_area string) error {
	r._area = _area
	r.Set("area", _area)
	return nil
}

// GetArea Area Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetArea() string {
	return r._area
}

// SetFloor is Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetFloor(_floor string) error {
	r._floor = _floor
	r.Set("floor", _floor)
	return nil
}

// GetFloor Floor Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetFloor() string {
	return r._floor
}

// SetBedType is BedType Setter
// 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&amp;docType=1&amp;articleId=105610
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetBedType(_bedType string) error {
	r._bedType = _bedType
	r.Set("bed_type", _bedType)
	return nil
}

// GetBedType BedType Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetBedType() string {
	return r._bedType
}

// SetBedSize is BedSize Setter
// 床宽。按自己定义存储，比如：2.1米
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetBedSize(_bedSize string) error {
	r._bedSize = _bedSize
	r.Set("bed_size", _bedSize)
	return nil
}

// GetBedSize BedSize Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetBedSize() string {
	return r._bedSize
}

// SetInternet is Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetInternet(_internet string) error {
	r._internet = _internet
	r.Set("internet", _internet)
	return nil
}

// GetInternet Internet Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetInternet() string {
	return r._internet
}

// SetService is Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {&#34;bar&#34;:false,&#34;catv&#34;:false,&#34;ddd&#34;:false,&#34;idd&#34;:false,&#34;pubtoilet&#34;:false,&#34;toilet&#34;:false}
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetService(_service string) error {
	r._service = _service
	r.Set("service", _service)
	return nil
}

// GetService Service Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetService() string {
	return r._service
}

// SetExtend is Extend Setter
// 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetExtend(_extend string) error {
	r._extend = _extend
	r.Set("extend", _extend)
	return nil
}

// GetExtend Extend Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetExtend() string {
	return r._extend
}

// SetOuterId is OuterId Setter
// （必传）商家房型ID
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetVendor is Vendor Setter
// 系统商，不要使用，只有申请才可用
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetHotelCode is HotelCode Setter
// 商家酒店ID(如果更新房型的时候房型不存在，会拿该code去新增房型)
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetPics is Pics Setter
// 房型图片只支持远程图片，格式如下：[{&#34;url&#34;:&#34;http://taobao.com/123.jpg&#34;,&#34;ismain&#34;:&#34;true&#34;},{&#34;url&#34;:&#34;http://taobao.com/456.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;},{&#34;url&#34;:&#34;http://taobao.com/789.jpg&#34;,&#34;ismain&#34;:&#34;false&#34;}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。要求：无logo、水印、边框、人物，不模糊，不重复，不歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetPics(_pics string) error {
	r._pics = _pics
	r.Set("pics", _pics)
	return nil
}

// GetPics Pics Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetPics() string {
	return r._pics
}

// SetNameE is NameE Setter
// 卖家房型英文名称
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetNameE(_nameE string) error {
	r._nameE = _nameE
	r.Set("name_e", _nameE)
	return nil
}

// GetNameE NameE Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetNameE() string {
	return r._nameE
}

// SetOperator is Operator Setter
// 操作人信息
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetOperator() string {
	return r._operator
}

// SetBedInfo is BedInfo Setter
// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&amp;docType=1&amp;articleId=118712&amp;previewCode=1DABB73EA935608455E203BA06CF3566
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetBedInfo(_bedInfo string) error {
	r._bedInfo = _bedInfo
	r.Set("bed_info", _bedInfo)
	return nil
}

// GetBedInfo BedInfo Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetBedInfo() string {
	return r._bedInfo
}

// SetNewOuterId is NewOuterId Setter
// 新的房型编码，请确实需要修改某个房型的编码的时候才使用，如需使用，请联系飞猪技术支持开通权限，否则会更新失败
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetNewOuterId(_newOuterId string) error {
	r._newOuterId = _newOuterId
	r.Set("new_outer_id", _newOuterId)
	return nil
}

// GetNewOuterId NewOuterId Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetNewOuterId() string {
	return r._newOuterId
}

// SetStandardRoomFacilities is StandardRoomFacilities Setter
// 房间设施
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetStandardRoomFacilities(_standardRoomFacilities string) error {
	r._standardRoomFacilities = _standardRoomFacilities
	r.Set("standard_room_facilities", _standardRoomFacilities)
	return nil
}

// GetStandardRoomFacilities StandardRoomFacilities Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetStandardRoomFacilities() string {
	return r._standardRoomFacilities
}

// SetWindowDesc is WindowDesc Setter
// 窗型，窗型（windowType）： 0=无窗 1=有窗 2=部分有窗  窗型缺陷（windowTypeDefect）： 0=窗户不可打开通风 1=窗外有遮挡 2=窗外为酒店内景观  特殊窗（windowTypeSpecial）： 0=落地窗 1=飘窗 2=天窗 3=小窗。  当为有窗或部分有窗时，窗型缺陷可多选，特殊窗型需单选
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetWindowDesc(_windowDesc string) error {
	r._windowDesc = _windowDesc
	r.Set("window_desc", _windowDesc)
	return nil
}

// GetWindowDesc WindowDesc Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetWindowDesc() string {
	return r._windowDesc
}

// SetAddBed is AddBed Setter
// 房型加床政策。 费用(fee) 说明(desc)
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetAddBed(_addBed string) error {
	r._addBed = _addBed
	r.Set("add_bed", _addBed)
	return nil
}

// GetAddBed AddBed Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetAddBed() string {
	return r._addBed
}

// SetChildrenPolicy is ChildrenPolicy Setter
// 儿童政策
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetChildrenPolicy(_childrenPolicy string) error {
	r._childrenPolicy = _childrenPolicy
	r.Set("children_policy", _childrenPolicy)
	return nil
}

// GetChildrenPolicy ChildrenPolicy Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetChildrenPolicy() string {
	return r._childrenPolicy
}

// SetRid is Rid Setter
// （已废弃）
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetRid(_rid int64) error {
	r._rid = _rid
	r.Set("rid", _rid)
	return nil
}

// GetRid Rid Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetRid() int64 {
	return r._rid
}

// SetMaxOccupancy is MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetMaxOccupancy(_maxOccupancy int64) error {
	r._maxOccupancy = _maxOccupancy
	r.Set("max_occupancy", _maxOccupancy)
	return nil
}

// GetMaxOccupancy MaxOccupancy Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetMaxOccupancy() int64 {
	return r._maxOccupancy
}

// SetWindowType is WindowType Setter
// 0:无窗/1:有窗
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetWindowType(_windowType int64) error {
	r._windowType = _windowType
	r.Set("window_type", _windowType)
	return nil
}

// GetWindowType WindowType Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetWindowType() int64 {
	return r._windowType
}

// SetSrid is Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetSrid(_srid int64) error {
	r._srid = _srid
	r.Set("srid", _srid)
	return nil
}

// GetSrid Srid Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetSrid() int64 {
	return r._srid
}

// SetHid is Hid Setter
// （已废弃）
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetHid(_hid int64) error {
	r._hid = _hid
	r.Set("hid", _hid)
	return nil
}

// GetHid Hid Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetHid() int64 {
	return r._hid
}

// SetStatus is Status Setter
// 房型状态。0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetStatus(_status *model.File) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetStatus() *model.File {
	return r._status
}

// SetConnectionType is ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoXhotelRoomtypeUpdateAPIRequest) SetConnectionType(_connectionType int64) error {
	r._connectionType = _connectionType
	r.Set("connection_type", _connectionType)
	return nil
}

// GetConnectionType ConnectionType Getter
func (r TaobaoXhotelRoomtypeUpdateAPIRequest) GetConnectionType() int64 {
	return r._connectionType
}

var poolTaobaoXhotelRoomtypeUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelRoomtypeUpdateRequest()
	},
}

// GetTaobaoXhotelRoomtypeUpdateRequest 从 sync.Pool 获取 TaobaoXhotelRoomtypeUpdateAPIRequest
func GetTaobaoXhotelRoomtypeUpdateAPIRequest() *TaobaoXhotelRoomtypeUpdateAPIRequest {
	return poolTaobaoXhotelRoomtypeUpdateAPIRequest.Get().(*TaobaoXhotelRoomtypeUpdateAPIRequest)
}

// ReleaseTaobaoXhotelRoomtypeUpdateAPIRequest 将 TaobaoXhotelRoomtypeUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelRoomtypeUpdateAPIRequest(v *TaobaoXhotelRoomtypeUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelRoomtypeUpdateAPIRequest.Put(v)
}
