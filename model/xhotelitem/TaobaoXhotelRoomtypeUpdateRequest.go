package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房型更新接口（ID不存在自动新增） API请求
taobao.xhotel.roomtype.update

酒店房型更新或添加
*/
type TaobaoXhotelRoomtypeUpdateRequest struct {
    model.Params
    // 房型名称。不能超过30字；添加房型时为必须
    name   string
    // 最大入住人数，默认2（1-99）
    maxOccupancy   int64
    // 具体面积大小，请按照正确格式填写。两种格式，例如：40 或者 10-20
    area   string
    // 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
    floor   string
    // 床型。必填，按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
    bedType   string
    // 床宽。按自己定义存储，比如：2.1米
    bedSize   string
    // 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
    internet   string
    // 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
    service   string
    // 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
    extend   string
    // （已废弃）
    rid   int64
    // 0:无窗/1:有窗
    windowType   int64
    // 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
    srid   int64
    // （必传）商家房型ID
    outerId   string
    // 系统商，不要使用，只有申请才可用
    vendor   string
    // （已废弃）
    hid   int64
    // 商家酒店ID(如果更新房型的时候房型不存在，会拿该code去新增房型)
    hotelCode   string
    // 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。要求：无logo、水印、边框、人物，不模糊，不重复，不歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
    pics   string
    // 房型状态。0:正常，-1:删除，-2:停售
    status   *model.File
    // 卖家房型英文名称
    nameE   string
    // 操作人信息
    operator   string
    // 属性值为1: 含义是非直连房型
    connectionType   int64
    // main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
    bedInfo   string
    // 新的房型编码，请确实需要修改某个房型的编码的时候才使用，如需使用，请联系飞猪技术支持开通权限，否则会更新失败
    newOuterId   string
    // 房间设施
    standardRoomFacilities   string
}

// 初始化TaobaoXhotelRoomtypeUpdateRequest对象
func NewTaobaoXhotelRoomtypeUpdateRequest() *TaobaoXhotelRoomtypeUpdateRequest{
    return &TaobaoXhotelRoomtypeUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomtypeUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.roomtype.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomtypeUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Name Setter
// 房型名称。不能超过30字；添加房型时为必须
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetName() string {
    return r.name
}
// MaxOccupancy Setter
// 最大入住人数，默认2（1-99）
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetMaxOccupancy(maxOccupancy int64) error {
    r.maxOccupancy = maxOccupancy
    r.Set("max_occupancy", maxOccupancy)
    return nil
}

// MaxOccupancy Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetMaxOccupancy() int64 {
    return r.maxOccupancy
}
// Area Setter
// 具体面积大小，请按照正确格式填写。两种格式，例如：40 或者 10-20
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetArea(area string) error {
    r.area = area
    r.Set("area", area)
    return nil
}

// Area Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetArea() string {
    return r.area
}
// Floor Setter
// 客房在建筑的第几层，隔层为1-2层，4-5层，7-8层
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetFloor(floor string) error {
    r.floor = floor
    r.Set("floor", floor)
    return nil
}

// Floor Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetFloor() string {
    return r.floor
}
// BedType Setter
// 床型。必填，按链接中床型列表定义值存储 http://open.taobao.com/docs/doc.htm?&docType=1&articleId=105610
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetBedType(bedType string) error {
    r.bedType = bedType
    r.Set("bed_type", bedType)
    return nil
}

// BedType Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetBedType() string {
    return r.bedType
}
// BedSize Setter
// 床宽。按自己定义存储，比如：2.1米
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetBedSize(bedSize string) error {
    r.bedSize = bedSize
    r.Set("bed_size", bedSize)
    return nil
}

// BedSize Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetBedSize() string {
    return r.bedSize
}
// Internet Setter
// 宽带服务。A,B,C,D。分别代表： A：无宽带，B：免费宽带，C：收费宽带，D：部分收费宽带
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetInternet(internet string) error {
    r.internet = internet
    r.Set("internet", internet)
    return nil
}

// Internet Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetInternet() string {
    return r.internet
}
// Service Setter
// 设施服务。JSON格式。 value值true有此服务，false没有。 bar：吧台，catv：有线电视，ddd：国内长途电话，idd：国际长途电话，toilet：独立卫生间，pubtoliet：公共卫生间。 如： {"bar":false,"catv":false,"ddd":false,"idd":false,"pubtoilet":false,"toilet":false}
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetService(service string) error {
    r.service = service
    r.Set("service", service)
    return nil
}

// Service Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetService() string {
    return r.service
}
// Extend Setter
// 扩展信息的JSON。注：此字段的值需要ISV在接入前与淘宝沟通，且确认能解析
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetExtend(extend string) error {
    r.extend = extend
    r.Set("extend", extend)
    return nil
}

// Extend Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetExtend() string {
    return r.extend
}
// Rid Setter
// （已废弃）
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetRid(rid int64) error {
    r.rid = rid
    r.Set("rid", rid)
    return nil
}

// Rid Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetRid() int64 {
    return r.rid
}
// WindowType Setter
// 0:无窗/1:有窗
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetWindowType(windowType int64) error {
    r.windowType = windowType
    r.Set("window_type", windowType)
    return nil
}

// WindowType Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetWindowType() int64 {
    return r.windowType
}
// Srid Setter
// 该字段只有确定的时候，才允许填入。用于标示和淘宝房型的匹配关系。目前尚未启动该字段。
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetSrid(srid int64) error {
    r.srid = srid
    r.Set("srid", srid)
    return nil
}

// Srid Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetSrid() int64 {
    return r.srid
}
// OuterId Setter
// （必传）商家房型ID
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetOuterId(outerId string) error {
    r.outerId = outerId
    r.Set("outer_id", outerId)
    return nil
}

// OuterId Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetOuterId() string {
    return r.outerId
}
// Vendor Setter
// 系统商，不要使用，只有申请才可用
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetVendor() string {
    return r.vendor
}
// Hid Setter
// （已废弃）
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetHid(hid int64) error {
    r.hid = hid
    r.Set("hid", hid)
    return nil
}

// Hid Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetHid() int64 {
    return r.hid
}
// HotelCode Setter
// 商家酒店ID(如果更新房型的时候房型不存在，会拿该code去新增房型)
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetHotelCode(hotelCode string) error {
    r.hotelCode = hotelCode
    r.Set("hotel_code", hotelCode)
    return nil
}

// HotelCode Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetHotelCode() string {
    return r.hotelCode
}
// Pics Setter
// 房型图片只支持远程图片，格式如下：[{"url":"http://taobao.com/123.jpg","ismain":"true"},{"url":"http://taobao.com/456.jpg","ismain":"false"},{"url":"http://taobao.com/789.jpg","ismain":"false"}]其中url是远程图片的访问地址（URL地址必须是合法的，否则会报错），main是是否为主图。只能设置一张图片为主图。要求：无logo、水印、边框、人物，不模糊，不重复，不歪斜，房间图清晰，图片尺寸不小于300*225，不小于5M
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetPics(pics string) error {
    r.pics = pics
    r.Set("pics", pics)
    return nil
}

// Pics Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetPics() string {
    return r.pics
}
// Status Setter
// 房型状态。0:正常，-1:删除，-2:停售
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetStatus(status *model.File) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetStatus() *model.File {
    return r.status
}
// NameE Setter
// 卖家房型英文名称
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetNameE(nameE string) error {
    r.nameE = nameE
    r.Set("name_e", nameE)
    return nil
}

// NameE Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetNameE() string {
    return r.nameE
}
// Operator Setter
// 操作人信息
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetOperator(operator string) error {
    r.operator = operator
    r.Set("operator", operator)
    return nil
}

// Operator Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetOperator() string {
    return r.operator
}
// ConnectionType Setter
// 属性值为1: 含义是非直连房型
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetConnectionType(connectionType int64) error {
    r.connectionType = connectionType
    r.Set("connection_type", connectionType)
    return nil
}

// ConnectionType Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetConnectionType() int64 {
    return r.connectionType
}
// BedInfo Setter
// main_bed_type母床型,sub_bed_type子床型。详情参见文档： https://open.alitrip.com/docs/doc.htm?spm=0.0.0.0.yN2mES&docType=1&articleId=118712&previewCode=1DABB73EA935608455E203BA06CF3566
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetBedInfo(bedInfo string) error {
    r.bedInfo = bedInfo
    r.Set("bed_info", bedInfo)
    return nil
}

// BedInfo Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetBedInfo() string {
    return r.bedInfo
}
// NewOuterId Setter
// 新的房型编码，请确实需要修改某个房型的编码的时候才使用，如需使用，请联系飞猪技术支持开通权限，否则会更新失败
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetNewOuterId(newOuterId string) error {
    r.newOuterId = newOuterId
    r.Set("new_outer_id", newOuterId)
    return nil
}

// NewOuterId Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetNewOuterId() string {
    return r.newOuterId
}
// StandardRoomFacilities Setter
// 房间设施
func (r *TaobaoXhotelRoomtypeUpdateRequest) SetStandardRoomFacilities(standardRoomFacilities string) error {
    r.standardRoomFacilities = standardRoomFacilities
    r.Set("standard_room_facilities", standardRoomFacilities)
    return nil
}

// StandardRoomFacilities Getter
func (r TaobaoXhotelRoomtypeUpdateRequest) GetStandardRoomFacilities() string {
    return r.standardRoomFacilities
}
