package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房型新增接口（ID重复变更新） API请求
taobao.xhotel.roomtype.add

房型添加或更新
*/
type TaobaoXhotelRoomtypeAddRequest struct {
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
    // 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
    bedType   string
    // 床宽。按自己定义存储，比如：2.1米
    bedSize   string
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
    // main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
    bedInfo   string
    // 酒店房型设施
    standardRoomFacilities   string
}

// 初始化TaobaoXhotelRoomtypeAddRequest对象
func NewTaobaoXhotelRoomtypeAddRequest() *TaobaoXhotelRoomtypeAddRequest{
    return &TaobaoXhotelRoomtypeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeAddRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OuterId Setter
// 卖家房型ID，不能重复建议格式是:酒店code_房型code
func (r *TaobaoXhotelRoomtypeAddRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetOuterId() string {
    return r.outerId
}
// Hid Setter
// （已废弃）请使用outHid
func (r *TaobaoXhotelRoomtypeAddRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetHid() int64 {
    return r.hid
}
// Name Setter
// 房型名称。不能超过30字
func (r *TaobaoXhotelRoomtypeAddRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetName() string {
    return r.name
}
// MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoXhotelRoomtypeAddRequest) SetMaxOccupancy(maxOccupancy int64) error {
    r.maxOccupancy = maxOccupancy
    r.Set("max_occupancy", maxOccupancy)
    return nil
}

// MaxOccupancy Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetMaxOccupancy() int64 {
    return r.maxOccupancy
}
// Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40或者 10-20
func (r *TaobaoXhotelRoomtypeAddRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

// Area Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetArea() string {
    return r.area
}
// Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelRoomtypeAddRequest) SetFloor(floor string) error {
    r.floor = floor
    r.Set("floor", floor)
    return nil
}

// Floor Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetFloor() string {
    return r.floor
}
// BedType Setter
// 床型。按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
func (r *TaobaoXhotelRoomtypeAddRequest) SetBedType(bedType string) error {
    r.bedType = bedType
    r.Set("bed_type", bedType)
    return nil
}

// BedType Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetBedType() string {
    return r.bedType
}
// BedSize Setter
// 床宽。按自己定义存储，比如：2.1米
func (r *TaobaoXhotelRoomtypeAddRequest) SetBedSize(bedSize string) error {
    r.bedSize = bedSize
    r.Set("bed_size", bedSize)
    return nil
}

// BedSize Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetBedSize() string {
    return r.bedSize
}
// Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoXhotelRoomtypeAddRequest) SetInternet(internet string) error {
    r.internet = internet
    r.Set("internet", internet)
    return nil
}

// Internet Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetInternet() string {
    return r.internet
}
// Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
func (r *TaobaoXhotelRoomtypeAddRequest) SetService(service string) error {
    r.service = service
    r.Set("service", service)
    return nil
}

// Service Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetService() string {
    return r.service
}
// Extend Setter
// 不要使用
func (r *TaobaoXhotelRoomtypeAddRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetExtend() string {
    return r.extend
}
// WindowType Setter
// 0:无窗/1:有窗/2:部分有窗/3:暗窗/4:部分暗窗
func (r *TaobaoXhotelRoomtypeAddRequest) SetWindowType(windowType int64) error {
    r.windowType = windowType
    r.Set("window_type", windowType)
    return nil
}

// WindowType Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetWindowType() int64 {
    return r.windowType
}
// Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelRoomtypeAddRequest) SetSrid(srid int64) error {
    r.srid = srid
    r.Set("srid", srid)
    return nil
}

// Srid Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetSrid() int64 {
    return r.srid
}
// OutHid Setter
// （必传）商家酒店ID，指明该房型属于哪家酒店
func (r *TaobaoXhotelRoomtypeAddRequest) SetOutHid(outHid string) error {
    r.outHid = outHid
    r.Set("out_hid", outHid)
    return nil
}

// OutHid Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetOutHid() string {
    return r.outHid
}
// Vendor Setter
// 系统商，无申请不可使用
func (r *TaobaoXhotelRoomtypeAddRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetVendor() string {
    return r.vendor
}
// Pics Setter
// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。
func (r *TaobaoXhotelRoomtypeAddRequest) SetPics(pics string) error {
    r.pics = pics
    r.Set("pics", pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetPics() string {
    return r.pics
}
// NameE Setter
// 卖家房型英文名称
func (r *TaobaoXhotelRoomtypeAddRequest) SetNameE(nameE string) error {
    r.nameE = nameE
    r.Set("name_e", nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetNameE() string {
    return r.nameE
}
// Operator Setter
// 操作人信息
func (r *TaobaoXhotelRoomtypeAddRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetOperator() string {
    return r.operator
}
// ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoXhotelRoomtypeAddRequest) SetConnectionType(connectionType int64) error {
    r.connectionType = connectionType
    r.Set("connection_type", connectionType)
    return nil
}

// ConnectionType Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetConnectionType() int64 {
    return r.connectionType
}
// BedInfo Setter
// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
func (r *TaobaoXhotelRoomtypeAddRequest) SetBedInfo(bedInfo string) error {
    r.bedInfo = bedInfo
    r.Set("bed_info", bedInfo)
    return nil
}

// BedInfo Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetBedInfo() string {
    return r.bedInfo
}
// StandardRoomFacilities Setter
// 酒店房型设施
func (r *TaobaoXhotelRoomtypeAddRequest) SetStandardRoomFacilities(standardRoomFacilities string) error {
    r.standardRoomFacilities = standardRoomFacilities
    r.Set("standard_room_facilities", standardRoomFacilities)
    return nil
}

// StandardRoomFacilities Getter
func (r TaobaoXhotelRoomtypeAddRequest) GetStandardRoomFacilities() string {
    return r.standardRoomFacilities
}
