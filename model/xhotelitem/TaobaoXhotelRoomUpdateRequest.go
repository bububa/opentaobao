package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
房型库存推送接口（全量推送） API请求
taobao.xhotel.room.update

此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。
*/
type TaobaoXhotelRoomUpdateRequest struct {
    model.Params
    // 废弃，使用out_rid
    gid   int64
    // 废弃，宝贝名称展示在店铺里
    title   string
    // 废弃，房型购买须知展示在PC购物路径
    guide   string
    // 废弃，宝贝描述展示在宝贝详情页面
    desc   string
    // 废弃，宝贝图片，没有默认使用标准酒店房型图片
    pic   []*model.File
    // 废弃，房型是否提供发票
    hasReceipt   bool
    // 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
    receiptType   string
    // 废弃，房型发票类型为其他时的发票描述,不能超过30个字
    receiptOtherTypeDesc   string
    // 废弃，房型发票说明，不能超过100个字
    receiptInfo   string
    // 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]
    inventory   string
    // 系统商，一般不填写，使用须申请
    vendor   string
    // 卖家房型ID
    outRid   string
    // 房型库存开关。 1，开；2，关
    roomSwitchCal   string
    // 超预定库存截止时间
    superbookEndTime   string
    // 超预定库存开始时间
    superbookStartTime   string
    // 保留房库存截止时间
    allotmentEndTime   string
    // 保留房库存截止时间
    allotmentStartTime   string
    // 宝贝状态,1上架。
    status   int64
}

// 初始化TaobaoXhotelRoomUpdateRequest对象
func NewTaobaoXhotelRoomUpdateRequest() *TaobaoXhotelRoomUpdateRequest{
    return &TaobaoXhotelRoomUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomUpdateRequest) GetApiMethodName() string {
    return "taobao.xhotel.room.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Gid Setter
// 废弃，使用out_rid
func (r *TaobaoXhotelRoomUpdateRequest) SetGid(gid int64) error {
    r.gid = gid
    r.Set("gid", gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelRoomUpdateRequest) GetGid() int64 {
    return r.gid
}
// Title Setter
// 废弃，宝贝名称展示在店铺里
func (r *TaobaoXhotelRoomUpdateRequest) SetTitle(title string) error {
    r.title = title
    r.Set("title", title)
    return nil
}

// Title Getter
func (r TaobaoXhotelRoomUpdateRequest) GetTitle() string {
    return r.title
}
// Guide Setter
// 废弃，房型购买须知展示在PC购物路径
func (r *TaobaoXhotelRoomUpdateRequest) SetGuide(guide string) error {
    r.guide = guide
    r.Set("guide", guide)
    return nil
}

// Guide Getter
func (r TaobaoXhotelRoomUpdateRequest) GetGuide() string {
    return r.guide
}
// Desc Setter
// 废弃，宝贝描述展示在宝贝详情页面
func (r *TaobaoXhotelRoomUpdateRequest) SetDesc(desc string) error {
    r.desc = desc
    r.Set("desc", desc)
    return nil
}

// Desc Getter
func (r TaobaoXhotelRoomUpdateRequest) GetDesc() string {
    return r.desc
}
// Pic Setter
// 废弃，宝贝图片，没有默认使用标准酒店房型图片
func (r *TaobaoXhotelRoomUpdateRequest) SetPic(pic []*model.File) error {
    r.pic = pic
    r.Set("pic", pic)
    return nil
}

// Pic Getter
func (r TaobaoXhotelRoomUpdateRequest) GetPic() []*model.File {
    return r.pic
}
// HasReceipt Setter
// 废弃，房型是否提供发票
func (r *TaobaoXhotelRoomUpdateRequest) SetHasReceipt(hasReceipt bool) error {
    r.hasReceipt = hasReceipt
    r.Set("has_receipt", hasReceipt)
    return nil
}

// HasReceipt Getter
func (r TaobaoXhotelRoomUpdateRequest) GetHasReceipt() bool {
    return r.hasReceipt
}
// ReceiptType Setter
// 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
func (r *TaobaoXhotelRoomUpdateRequest) SetReceiptType(receiptType string) error {
    r.receiptType = receiptType
    r.Set("receipt_type", receiptType)
    return nil
}

// ReceiptType Getter
func (r TaobaoXhotelRoomUpdateRequest) GetReceiptType() string {
    return r.receiptType
}
// ReceiptOtherTypeDesc Setter
// 废弃，房型发票类型为其他时的发票描述,不能超过30个字
func (r *TaobaoXhotelRoomUpdateRequest) SetReceiptOtherTypeDesc(receiptOtherTypeDesc string) error {
    r.receiptOtherTypeDesc = receiptOtherTypeDesc
    r.Set("receipt_other_type_desc", receiptOtherTypeDesc)
    return nil
}

// ReceiptOtherTypeDesc Getter
func (r TaobaoXhotelRoomUpdateRequest) GetReceiptOtherTypeDesc() string {
    return r.receiptOtherTypeDesc
}
// ReceiptInfo Setter
// 废弃，房型发票说明，不能超过100个字
func (r *TaobaoXhotelRoomUpdateRequest) SetReceiptInfo(receiptInfo string) error {
    r.receiptInfo = receiptInfo
    r.Set("receipt_info", receiptInfo)
    return nil
}

// ReceiptInfo Getter
func (r TaobaoXhotelRoomUpdateRequest) GetReceiptInfo() string {
    return r.receiptInfo
}
// Inventory Setter
// 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]
func (r *TaobaoXhotelRoomUpdateRequest) SetInventory(inventory string) error {
    r.inventory = inventory
    r.Set("inventory", inventory)
    return nil
}

// Inventory Getter
func (r TaobaoXhotelRoomUpdateRequest) GetInventory() string {
    return r.inventory
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRoomUpdateRequest) SetVendor(vendor string) error {
    r.vendor = vendor
    r.Set("vendor", vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomUpdateRequest) GetVendor() string {
    return r.vendor
}
// OutRid Setter
// 卖家房型ID
func (r *TaobaoXhotelRoomUpdateRequest) SetOutRid(outRid string) error {
    r.outRid = outRid
    r.Set("out_rid", outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelRoomUpdateRequest) GetOutRid() string {
    return r.outRid
}
// RoomSwitchCal Setter
// 房型库存开关。 1，开；2，关
func (r *TaobaoXhotelRoomUpdateRequest) SetRoomSwitchCal(roomSwitchCal string) error {
    r.roomSwitchCal = roomSwitchCal
    r.Set("room_switch_cal", roomSwitchCal)
    return nil
}

// RoomSwitchCal Getter
func (r TaobaoXhotelRoomUpdateRequest) GetRoomSwitchCal() string {
    return r.roomSwitchCal
}
// SuperbookEndTime Setter
// 超预定库存截止时间
func (r *TaobaoXhotelRoomUpdateRequest) SetSuperbookEndTime(superbookEndTime string) error {
    r.superbookEndTime = superbookEndTime
    r.Set("superbook_end_time", superbookEndTime)
    return nil
}

// SuperbookEndTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetSuperbookEndTime() string {
    return r.superbookEndTime
}
// SuperbookStartTime Setter
// 超预定库存开始时间
func (r *TaobaoXhotelRoomUpdateRequest) SetSuperbookStartTime(superbookStartTime string) error {
    r.superbookStartTime = superbookStartTime
    r.Set("superbook_start_time", superbookStartTime)
    return nil
}

// SuperbookStartTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetSuperbookStartTime() string {
    return r.superbookStartTime
}
// AllotmentEndTime Setter
// 保留房库存截止时间
func (r *TaobaoXhotelRoomUpdateRequest) SetAllotmentEndTime(allotmentEndTime string) error {
    r.allotmentEndTime = allotmentEndTime
    r.Set("allotment_end_time", allotmentEndTime)
    return nil
}

// AllotmentEndTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetAllotmentEndTime() string {
    return r.allotmentEndTime
}
// AllotmentStartTime Setter
// 保留房库存截止时间
func (r *TaobaoXhotelRoomUpdateRequest) SetAllotmentStartTime(allotmentStartTime string) error {
    r.allotmentStartTime = allotmentStartTime
    r.Set("allotment_start_time", allotmentStartTime)
    return nil
}

// AllotmentStartTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetAllotmentStartTime() string {
    return r.allotmentStartTime
}
// Status Setter
// 宝贝状态,1上架。
func (r *TaobaoXhotelRoomUpdateRequest) SetStatus(status int64) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r TaobaoXhotelRoomUpdateRequest) GetStatus() int64 {
    return r.status
}
