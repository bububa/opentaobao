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
    _gid   int64
    // 废弃，宝贝名称展示在店铺里
    _title   string
    // 废弃，房型购买须知展示在PC购物路径
    _guide   string
    // 废弃，宝贝描述展示在宝贝详情页面
    _desc   string
    // 废弃，宝贝图片，没有默认使用标准酒店房型图片
    _pic   *model.File
    // 废弃，房型是否提供发票
    _hasReceipt   bool
    // 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
    _receiptType   string
    // 废弃，房型发票类型为其他时的发票描述,不能超过30个字
    _receiptOtherTypeDesc   string
    // 废弃，房型发票说明，不能超过100个字
    _receiptInfo   string
    // 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]
    _inventory   string
    // 系统商，一般不填写，使用须申请
    _vendor   string
    // 卖家房型ID
    _outRid   string
    // 房型库存开关。 1，开；2，关
    _roomSwitchCal   string
    // 超预定库存截止时间
    _superbookEndTime   string
    // 超预定库存开始时间
    _superbookStartTime   string
    // 保留房库存截止时间
    _allotmentEndTime   string
    // 保留房库存截止时间
    _allotmentStartTime   string
    // 宝贝状态,1上架。
    _status   int64
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
func (r *TaobaoXhotelRoomUpdateRequest) SetGid(_gid int64) error {
    r._gid = _gid
    r.Set("gid", _gid)
    return nil
}

// Gid Getter
func (r TaobaoXhotelRoomUpdateRequest) GetGid() int64 {
    return r._gid
}
// Title Setter
// 废弃，宝贝名称展示在店铺里
func (r *TaobaoXhotelRoomUpdateRequest) SetTitle(_title string) error {
    r._title = _title
    r.Set("title", _title)
    return nil
}

// Title Getter
func (r TaobaoXhotelRoomUpdateRequest) GetTitle() string {
    return r._title
}
// Guide Setter
// 废弃，房型购买须知展示在PC购物路径
func (r *TaobaoXhotelRoomUpdateRequest) SetGuide(_guide string) error {
    r._guide = _guide
    r.Set("guide", _guide)
    return nil
}

// Guide Getter
func (r TaobaoXhotelRoomUpdateRequest) GetGuide() string {
    return r._guide
}
// Desc Setter
// 废弃，宝贝描述展示在宝贝详情页面
func (r *TaobaoXhotelRoomUpdateRequest) SetDesc(_desc string) error {
    r._desc = _desc
    r.Set("desc", _desc)
    return nil
}

// Desc Getter
func (r TaobaoXhotelRoomUpdateRequest) GetDesc() string {
    return r._desc
}
// Pic Setter
// 废弃，宝贝图片，没有默认使用标准酒店房型图片
func (r *TaobaoXhotelRoomUpdateRequest) SetPic(_pic *model.File) error {
    r._pic = _pic
    r.Set("pic", _pic)
    return nil
}

// Pic Getter
func (r TaobaoXhotelRoomUpdateRequest) GetPic() *model.File {
    return r._pic
}
// HasReceipt Setter
// 废弃，房型是否提供发票
func (r *TaobaoXhotelRoomUpdateRequest) SetHasReceipt(_hasReceipt bool) error {
    r._hasReceipt = _hasReceipt
    r.Set("has_receipt", _hasReceipt)
    return nil
}

// HasReceipt Getter
func (r TaobaoXhotelRoomUpdateRequest) GetHasReceipt() bool {
    return r._hasReceipt
}
// ReceiptType Setter
// 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
func (r *TaobaoXhotelRoomUpdateRequest) SetReceiptType(_receiptType string) error {
    r._receiptType = _receiptType
    r.Set("receipt_type", _receiptType)
    return nil
}

// ReceiptType Getter
func (r TaobaoXhotelRoomUpdateRequest) GetReceiptType() string {
    return r._receiptType
}
// ReceiptOtherTypeDesc Setter
// 废弃，房型发票类型为其他时的发票描述,不能超过30个字
func (r *TaobaoXhotelRoomUpdateRequest) SetReceiptOtherTypeDesc(_receiptOtherTypeDesc string) error {
    r._receiptOtherTypeDesc = _receiptOtherTypeDesc
    r.Set("receipt_other_type_desc", _receiptOtherTypeDesc)
    return nil
}

// ReceiptOtherTypeDesc Getter
func (r TaobaoXhotelRoomUpdateRequest) GetReceiptOtherTypeDesc() string {
    return r._receiptOtherTypeDesc
}
// ReceiptInfo Setter
// 废弃，房型发票说明，不能超过100个字
func (r *TaobaoXhotelRoomUpdateRequest) SetReceiptInfo(_receiptInfo string) error {
    r._receiptInfo = _receiptInfo
    r.Set("receipt_info", _receiptInfo)
    return nil
}

// ReceiptInfo Getter
func (r TaobaoXhotelRoomUpdateRequest) GetReceiptInfo() string {
    return r._receiptInfo
}
// Inventory Setter
// 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]
func (r *TaobaoXhotelRoomUpdateRequest) SetInventory(_inventory string) error {
    r._inventory = _inventory
    r.Set("inventory", _inventory)
    return nil
}

// Inventory Getter
func (r TaobaoXhotelRoomUpdateRequest) GetInventory() string {
    return r._inventory
}
// Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRoomUpdateRequest) SetVendor(_vendor string) error {
    r._vendor = _vendor
    r.Set("vendor", _vendor)
    return nil
}

// Vendor Getter
func (r TaobaoXhotelRoomUpdateRequest) GetVendor() string {
    return r._vendor
}
// OutRid Setter
// 卖家房型ID
func (r *TaobaoXhotelRoomUpdateRequest) SetOutRid(_outRid string) error {
    r._outRid = _outRid
    r.Set("out_rid", _outRid)
    return nil
}

// OutRid Getter
func (r TaobaoXhotelRoomUpdateRequest) GetOutRid() string {
    return r._outRid
}
// RoomSwitchCal Setter
// 房型库存开关。 1，开；2，关
func (r *TaobaoXhotelRoomUpdateRequest) SetRoomSwitchCal(_roomSwitchCal string) error {
    r._roomSwitchCal = _roomSwitchCal
    r.Set("room_switch_cal", _roomSwitchCal)
    return nil
}

// RoomSwitchCal Getter
func (r TaobaoXhotelRoomUpdateRequest) GetRoomSwitchCal() string {
    return r._roomSwitchCal
}
// SuperbookEndTime Setter
// 超预定库存截止时间
func (r *TaobaoXhotelRoomUpdateRequest) SetSuperbookEndTime(_superbookEndTime string) error {
    r._superbookEndTime = _superbookEndTime
    r.Set("superbook_end_time", _superbookEndTime)
    return nil
}

// SuperbookEndTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetSuperbookEndTime() string {
    return r._superbookEndTime
}
// SuperbookStartTime Setter
// 超预定库存开始时间
func (r *TaobaoXhotelRoomUpdateRequest) SetSuperbookStartTime(_superbookStartTime string) error {
    r._superbookStartTime = _superbookStartTime
    r.Set("superbook_start_time", _superbookStartTime)
    return nil
}

// SuperbookStartTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetSuperbookStartTime() string {
    return r._superbookStartTime
}
// AllotmentEndTime Setter
// 保留房库存截止时间
func (r *TaobaoXhotelRoomUpdateRequest) SetAllotmentEndTime(_allotmentEndTime string) error {
    r._allotmentEndTime = _allotmentEndTime
    r.Set("allotment_end_time", _allotmentEndTime)
    return nil
}

// AllotmentEndTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetAllotmentEndTime() string {
    return r._allotmentEndTime
}
// AllotmentStartTime Setter
// 保留房库存截止时间
func (r *TaobaoXhotelRoomUpdateRequest) SetAllotmentStartTime(_allotmentStartTime string) error {
    r._allotmentStartTime = _allotmentStartTime
    r.Set("allotment_start_time", _allotmentStartTime)
    return nil
}

// AllotmentStartTime Getter
func (r TaobaoXhotelRoomUpdateRequest) GetAllotmentStartTime() string {
    return r._allotmentStartTime
}
// Status Setter
// 宝贝状态,1上架。
func (r *TaobaoXhotelRoomUpdateRequest) SetStatus(_status int64) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r TaobaoXhotelRoomUpdateRequest) GetStatus() int64 {
    return r._status
}
