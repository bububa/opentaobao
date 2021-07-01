package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
民宿房型信息添加 API请求
taobao.xhotel.house.roomtype.add

房型添加或更新
*/
type TaobaoXhotelHouseRoomtypeAddAPIRequest struct {
    model.Params
    // 卖家房型ID，不能重复建议格式是:酒店code_房型code
    _outerId   string
    // （已废弃）请使用outHid
    _hid   int64
    // 房型名称。不能超过30字
    _name   string
    // 最大入住人数，默认2（1-99）
    _maxOccupancy   int64
    // 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
    _area   string
    // 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
    _floor   string
    // 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
    _internet   string
    // 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
    _service   string
    // 不要使用
    _extend   string
    // 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
    _windowType   int64
    // 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
    _srid   int64
    // （必传）商家酒店ID，指明该房型属于哪家酒店
    _outHid   string
    // 系统商，无申请不可使用
    _vendor   string
    // 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
    _pics   string
    // 卖家房型英文名称
    _nameE   string
    // 操作人信息
    _operator   string
    // 属性值为1: 含义是非直连房型
    _connectionType   int64
    // 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房
    _houseModel   string
    // 房屋面积
    _houseSize   int64
    // 出租类型:1.整租;2.单间;3.床位
    _rentType   int64
    // 出租面积,单位平方米
    _rentSize   int64
    // 是否和房东合住:0.不和房东合住;1.和房东合住;
    _hasLandlord   int64
    // 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
    _bedInfo   string
    // 数据状态 0:正常，-1:删除，-2:停售
    _status   int64
}

// 初始化TaobaoXhotelHouseRoomtypeAddAPIRequest对象
func NewTaobaoXhotelHouseRoomtypeAddRequest() *TaobaoXhotelHouseRoomtypeAddAPIRequest{
    return &TaobaoXhotelHouseRoomtypeAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.house.roomtype.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 卖家房型ID，不能重复建议格式是:酒店code_房型code
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetOuterId(_outerId string) error {
    r._outerId = _outerId
    r.Set("outer_id", _outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetOuterId() string {
    return r._outerId
}
// Hid Setter
// （已废弃）请使用outHid
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetHid(_hid int64) error {
    r._hid = _hid
    r.Set("hid", _hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetHid() int64 {
    return r._hid
}
// Name Setter
// 房型名称。不能超过30字
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetName() string {
    return r._name
}
// MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetMaxOccupancy(_maxOccupancy int64) error {
    r._maxOccupancy = _maxOccupancy
    r.Set("max_occupancy", _maxOccupancy)
    return nil
}

// MaxOccupancy Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetMaxOccupancy() int64 {
    return r._maxOccupancy
}
// Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetArea(_area string) error {
    r._area = _area
    r.Set("area", _area)
    return nil
}

// Area Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetArea() string {
    return r._area
}
// Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetFloor(_floor string) error {
    r._floor = _floor
    r.Set("floor", _floor)
    return nil
}

// Floor Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetFloor() string {
    return r._floor
}
// Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetInternet(_internet string) error {
    r._internet = _internet
    r.Set("internet", _internet)
    return nil
}

// Internet Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetInternet() string {
    return r._internet
}
// Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetService(_service string) error {
    r._service = _service
    r.Set("service", _service)
    return nil
}

// Service Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetService() string {
    return r._service
}
// Extend Setter
// 不要使用
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetExtend(_extend string) error {
    r._extend = _extend
    r.Set("extend", _extend)
    return nil
}

// Extend Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetExtend() string {
    return r._extend
}
// WindowType Setter
// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetWindowType(_windowType int64) error {
    r._windowType = _windowType
    r.Set("window_type", _windowType)
    return nil
}

// WindowType Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetWindowType() int64 {
    return r._windowType
}
// Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetSrid(_srid int64) error {
    r._srid = _srid
    r.Set("srid", _srid)
    return nil
}

// Srid Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetSrid() int64 {
    return r._srid
}
// OutHid Setter
// （必传）商家酒店ID，指明该房型属于哪家酒店
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetOutHid(_outHid string) error {
    r._outHid = _outHid
    r.Set("out_hid", _outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetOutHid() string {
    return r._outHid
}
// Vendor Setter
// 系统商，无申请不可使用
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetVendor() string {
    return r._vendor
}
// Pics Setter
// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetPics(_pics string) error {
    r._pics = _pics
    r.Set("pics", _pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetPics() string {
    return r._pics
}
// NameE Setter
// 卖家房型英文名称
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetNameE(_nameE string) error {
    r._nameE = _nameE
    r.Set("name_e", _nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetNameE() string {
    return r._nameE
}
// Operator Setter
// 操作人信息
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetOperator(_operator string) error {
    r._operator = _operator
    r.Set("operator", _operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetOperator() string {
    return r._operator
}
// ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetConnectionType(_connectionType int64) error {
    r._connectionType = _connectionType
    r.Set("connection_type", _connectionType)
    return nil
}

// ConnectionType Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetConnectionType() int64 {
    return r._connectionType
}
// HouseModel Setter
// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetHouseModel(_houseModel string) error {
    r._houseModel = _houseModel
    r.Set("house_model", _houseModel)
    return nil
}

// HouseModel Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetHouseModel() string {
    return r._houseModel
}
// HouseSize Setter
// 房屋面积
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetHouseSize(_houseSize int64) error {
    r._houseSize = _houseSize
    r.Set("house_size", _houseSize)
    return nil
}

// HouseSize Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetHouseSize() int64 {
    return r._houseSize
}
// RentType Setter
// 出租类型:1.整租;2.单间;3.床位
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetRentType(_rentType int64) error {
    r._rentType = _rentType
    r.Set("rent_type", _rentType)
    return nil
}

// RentType Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetRentType() int64 {
    return r._rentType
}
// RentSize Setter
// 出租面积,单位平方米
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetRentSize(_rentSize int64) error {
    r._rentSize = _rentSize
    r.Set("rent_size", _rentSize)
    return nil
}

// RentSize Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetRentSize() int64 {
    return r._rentSize
}
// HasLandlord Setter
// 是否和房东合住:0.不和房东合住;1.和房东合住;
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetHasLandlord(_hasLandlord int64) error {
    r._hasLandlord = _hasLandlord
    r.Set("has_landlord", _hasLandlord)
    return nil
}

// HasLandlord Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetHasLandlord() int64 {
    return r._hasLandlord
}
// BedInfo Setter
// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetBedInfo(_bedInfo string) error {
    r._bedInfo = _bedInfo
    r.Set("bed_info", _bedInfo)
    return nil
}

// BedInfo Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetBedInfo() string {
    return r._bedInfo
}
// Status Setter
// 数据状态 0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelHouseRoomtypeAddAPIRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoXhotelHouseRoomtypeAddAPIRequest) GetStatus() int64 {
    return r._status
}
