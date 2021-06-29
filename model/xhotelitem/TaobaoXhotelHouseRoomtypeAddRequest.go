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
type TaobaoXhotelHouseRoomtypeAddRequest struct {
    model.Params
    // 卖家房型ID，不能重复建议格式是:酒店code_房型code
    outerId   string
    // （已废弃）请使用outHid
    hid   int64
    // 房型名称。不能超过30字
    name   string
    // 最大入住人数，默认2（1-99）
    maxOccupancy   int64
    // 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
    area   string
    // 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
    floor   string
    // 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
    internet   string
    // 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
    service   string
    // 不要使用
    extend   string
    // 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
    windowType   int64
    // 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
    srid   int64
    // （必传）商家酒店ID，指明该房型属于哪家酒店
    outHid   string
    // 系统商，无申请不可使用
    vendor   string
    // 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
    pics   string
    // 卖家房型英文名称
    nameE   string
    // 操作人信息
    operator   string
    // 属性值为1: 含义是非直连房型
    connectionType   int64
    // 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房
    houseModel   string
    // 房屋面积
    houseSize   int64
    // 出租类型:1.整租;2.单间;3.床位
    rentType   int64
    // 出租面积,单位平方米
    rentSize   int64
    // 是否和房东合住:0.不和房东合住;1.和房东合住;
    hasLandlord   int64
    // 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
    bedInfo   string
    // 数据状态 0:正常，-1:删除，-2:停售
    status   int64
}

// 初始化TaobaoXhotelHouseRoomtypeAddRequest对象
func NewTaobaoXhotelHouseRoomtypeAddRequest() *TaobaoXhotelHouseRoomtypeAddRequest{
    return &TaobaoXhotelHouseRoomtypeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.house.roomtype.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 卖家房型ID，不能重复建议格式是:酒店code_房型code
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetOuterId() string {
    return r.outerId
}
// Hid Setter
// （已废弃）请使用outHid
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetHid() int64 {
    return r.hid
}
// Name Setter
// 房型名称。不能超过30字
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetName() string {
    return r.name
}
// MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetMaxOccupancy(maxOccupancy int64) error {
    r.maxOccupancy = maxOccupancy
    r.Set("max_occupancy", maxOccupancy)
    return nil
}

// MaxOccupancy Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetMaxOccupancy() int64 {
    return r.maxOccupancy
}
// Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

// Area Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetArea() string {
    return r.area
}
// Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetFloor(floor string) error {
    r.floor = floor
    r.Set("floor", floor)
    return nil
}

// Floor Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetFloor() string {
    return r.floor
}
// Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetInternet(internet string) error {
    r.internet = internet
    r.Set("internet", internet)
    return nil
}

// Internet Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetInternet() string {
    return r.internet
}
// Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetService(service string) error {
    r.service = service
    r.Set("service", service)
    return nil
}

// Service Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetService() string {
    return r.service
}
// Extend Setter
// 不要使用
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetExtend() string {
    return r.extend
}
// WindowType Setter
// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetWindowType(windowType int64) error {
    r.windowType = windowType
    r.Set("window_type", windowType)
    return nil
}

// WindowType Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetWindowType() int64 {
    return r.windowType
}
// Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetSrid(srid int64) error {
    r.srid = srid
    r.Set("srid", srid)
    return nil
}

// Srid Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetSrid() int64 {
    return r.srid
}
// OutHid Setter
// （必传）商家酒店ID，指明该房型属于哪家酒店
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetOutHid() string {
    return r.outHid
}
// Vendor Setter
// 系统商，无申请不可使用
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetVendor() string {
    return r.vendor
}
// Pics Setter
// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetPics(pics string) error {
    r.pics = pics
    r.Set("pics", pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetPics() string {
    return r.pics
}
// NameE Setter
// 卖家房型英文名称
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetNameE(nameE string) error {
    r.nameE = nameE
    r.Set("name_e", nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetNameE() string {
    return r.nameE
}
// Operator Setter
// 操作人信息
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetOperator() string {
    return r.operator
}
// ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetConnectionType(connectionType int64) error {
    r.connectionType = connectionType
    r.Set("connection_type", connectionType)
    return nil
}

// ConnectionType Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetConnectionType() int64 {
    return r.connectionType
}
// HouseModel Setter
// 房屋户型， bedroom: 室, bathroom: 卫, livingroom: 厅, study: 书房, balcony: 阳台,kitchen: 厨房
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetHouseModel(houseModel string) error {
    r.houseModel = houseModel
    r.Set("house_model", houseModel)
    return nil
}

// HouseModel Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetHouseModel() string {
    return r.houseModel
}
// HouseSize Setter
// 房屋面积
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetHouseSize(houseSize int64) error {
    r.houseSize = houseSize
    r.Set("house_size", houseSize)
    return nil
}

// HouseSize Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetHouseSize() int64 {
    return r.houseSize
}
// RentType Setter
// 出租类型:1.整租;2.单间;3.床位
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetRentType(rentType int64) error {
    r.rentType = rentType
    r.Set("rent_type", rentType)
    return nil
}

// RentType Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetRentType() int64 {
    return r.rentType
}
// RentSize Setter
// 出租面积,单位平方米
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetRentSize(rentSize int64) error {
    r.rentSize = rentSize
    r.Set("rent_size", rentSize)
    return nil
}

// RentSize Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetRentSize() int64 {
    return r.rentSize
}
// HasLandlord Setter
// 是否和房东合住:0.不和房东合住;1.和房东合住;
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetHasLandlord(hasLandlord int64) error {
    r.hasLandlord = hasLandlord
    r.Set("has_landlord", hasLandlord)
    return nil
}

// HasLandlord Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetHasLandlord() int64 {
    return r.hasLandlord
}
// BedInfo Setter
// 床信息: bedType:床型, desc: 床型名, width:床宽, length：床长, bedNum: 床数。床型取值见链接https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.4zBOVn&docType=1&articleId=108347
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetBedInfo(bedInfo string) error {
    r.bedInfo = bedInfo
    r.Set("bed_info", bedInfo)
    return nil
}

// BedInfo Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetBedInfo() string {
    return r.bedInfo
}
// Status Setter
// 数据状态 0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelHouseRoomtypeAddRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoXhotelHouseRoomtypeAddRequest) GetStatus() int64 {
    return r.status
}
