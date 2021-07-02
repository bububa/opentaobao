package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelRoomUpdateAPIRequest 房型库存推送接口（全量推送） API请求
// taobao.xhotel.room.update
//
// 此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。
type TaobaoXhotelRoomUpdateAPIRequest struct {
	model.Params
	// 废弃，使用out_rid
	_gid int64
	// 废弃，宝贝名称展示在店铺里
	_title string
	// 废弃，房型购买须知展示在PC购物路径
	_guide string
	// 废弃，宝贝描述展示在宝贝详情页面
	_desc string
	// 废弃，宝贝图片，没有默认使用标准酒店房型图片
	_pic *model.File
	// 废弃，房型是否提供发票
	_hasReceipt bool
	// 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
	_receiptType string
	// 废弃，房型发票类型为其他时的发票描述,不能超过30个字
	_receiptOtherTypeDesc string
	// 废弃，房型发票说明，不能超过100个字
	_receiptInfo string
	// 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]
	_inventory string
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 卖家房型ID
	_outRid string
	// 房型库存开关。 1，开；2，关
	_roomSwitchCal string
	// 超预定库存截止时间
	_superbookEndTime string
	// 超预定库存开始时间
	_superbookStartTime string
	// 保留房库存截止时间
	_allotmentEndTime string
	// 保留房库存截止时间
	_allotmentStartTime string
	// 宝贝状态,1上架。
	_status int64
}

// NewTaobaoXhotelRoomUpdateRequest 初始化TaobaoXhotelRoomUpdateAPIRequest对象
func NewTaobaoXhotelRoomUpdateRequest() *TaobaoXhotelRoomUpdateAPIRequest {
	return &TaobaoXhotelRoomUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelRoomUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.room.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelRoomUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Gid Setter
// 废弃，使用out_rid
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// Get Gid Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetGid() int64 {
	return r._gid
}

// Set is Title Setter
// 废弃，宝贝名称展示在店铺里
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetTitle() string {
	return r._title
}

// Set is Guide Setter
// 废弃，房型购买须知展示在PC购物路径
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetGuide(_guide string) error {
	r._guide = _guide
	r.Set("guide", _guide)
	return nil
}

// Get Guide Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetGuide() string {
	return r._guide
}

// Set is Desc Setter
// 废弃，宝贝描述展示在宝贝详情页面
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// Get Desc Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetDesc() string {
	return r._desc
}

// Set is Pic Setter
// 废弃，宝贝图片，没有默认使用标准酒店房型图片
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetPic(_pic *model.File) error {
	r._pic = _pic
	r.Set("pic", _pic)
	return nil
}

// Get Pic Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetPic() *model.File {
	return r._pic
}

// Set is HasReceipt Setter
// 废弃，房型是否提供发票
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetHasReceipt(_hasReceipt bool) error {
	r._hasReceipt = _hasReceipt
	r.Set("has_receipt", _hasReceipt)
	return nil
}

// Get HasReceipt Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetHasReceipt() bool {
	return r._hasReceipt
}

// Set is ReceiptType Setter
// 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetReceiptType(_receiptType string) error {
	r._receiptType = _receiptType
	r.Set("receipt_type", _receiptType)
	return nil
}

// Get ReceiptType Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetReceiptType() string {
	return r._receiptType
}

// Set is ReceiptOtherTypeDesc Setter
// 废弃，房型发票类型为其他时的发票描述,不能超过30个字
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetReceiptOtherTypeDesc(_receiptOtherTypeDesc string) error {
	r._receiptOtherTypeDesc = _receiptOtherTypeDesc
	r.Set("receipt_other_type_desc", _receiptOtherTypeDesc)
	return nil
}

// Get ReceiptOtherTypeDesc Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetReceiptOtherTypeDesc() string {
	return r._receiptOtherTypeDesc
}

// Set is ReceiptInfo Setter
// 废弃，房型发票说明，不能超过100个字
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetReceiptInfo(_receiptInfo string) error {
	r._receiptInfo = _receiptInfo
	r.Set("receipt_info", _receiptInfo)
	return nil
}

// Get ReceiptInfo Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetReceiptInfo() string {
	return r._receiptInfo
}

// Set is Inventory Setter
// 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetInventory(_inventory string) error {
	r._inventory = _inventory
	r.Set("inventory", _inventory)
	return nil
}

// Get Inventory Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetInventory() string {
	return r._inventory
}

// Set is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// Get Vendor Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetVendor() string {
	return r._vendor
}

// Set is OutRid Setter
// 卖家房型ID
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// Get OutRid Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetOutRid() string {
	return r._outRid
}

// Set is RoomSwitchCal Setter
// 房型库存开关。 1，开；2，关
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetRoomSwitchCal(_roomSwitchCal string) error {
	r._roomSwitchCal = _roomSwitchCal
	r.Set("room_switch_cal", _roomSwitchCal)
	return nil
}

// Get RoomSwitchCal Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetRoomSwitchCal() string {
	return r._roomSwitchCal
}

// Set is SuperbookEndTime Setter
// 超预定库存截止时间
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetSuperbookEndTime(_superbookEndTime string) error {
	r._superbookEndTime = _superbookEndTime
	r.Set("superbook_end_time", _superbookEndTime)
	return nil
}

// Get SuperbookEndTime Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetSuperbookEndTime() string {
	return r._superbookEndTime
}

// Set is SuperbookStartTime Setter
// 超预定库存开始时间
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetSuperbookStartTime(_superbookStartTime string) error {
	r._superbookStartTime = _superbookStartTime
	r.Set("superbook_start_time", _superbookStartTime)
	return nil
}

// Get SuperbookStartTime Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetSuperbookStartTime() string {
	return r._superbookStartTime
}

// Set is AllotmentEndTime Setter
// 保留房库存截止时间
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetAllotmentEndTime(_allotmentEndTime string) error {
	r._allotmentEndTime = _allotmentEndTime
	r.Set("allotment_end_time", _allotmentEndTime)
	return nil
}

// Get AllotmentEndTime Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetAllotmentEndTime() string {
	return r._allotmentEndTime
}

// Set is AllotmentStartTime Setter
// 保留房库存截止时间
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetAllotmentStartTime(_allotmentStartTime string) error {
	r._allotmentStartTime = _allotmentStartTime
	r.Set("allotment_start_time", _allotmentStartTime)
	return nil
}

// Get AllotmentStartTime Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetAllotmentStartTime() string {
	return r._allotmentStartTime
}

// Set is Status Setter
// 宝贝状态,1上架。
func (r *TaobaoXhotelRoomUpdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r TaobaoXhotelRoomUpdateAPIRequest) GetStatus() int64 {
	return r._status
}
