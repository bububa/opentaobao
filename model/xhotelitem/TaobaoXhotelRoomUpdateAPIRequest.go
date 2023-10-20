package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelroomupdateAPIRequest 房型库存推送接口（全量推送） API请求
// taobao.xhotel.room.update
//
// 此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。
type TaobaoxhotelroomupdateAPIRequest struct {
	model.Params
	// 卖家房型ID
	_outRid string
	// 系统商，一般不填写，使用须申请
	_vendor string
	// 废弃，宝贝名称展示在店铺里
	_title string
	// 废弃，房型购买须知展示在PC购物路径
	_guide string
	// 废弃，宝贝描述展示在宝贝详情页面
	_desc string
	// 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
	_receiptType string
	// 废弃，房型发票类型为其他时的发票描述,不能超过30个字
	_receiptOtherTypeDesc string
	// 废弃，房型发票说明，不能超过100个字
	_receiptInfo string
	// 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{"date":2011-01-28,"quota":10,"al_quota":2,"sp_quota":3}]
	_inventory string
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
	// 废弃，使用out_rid
	_gid int64
	// 废弃，宝贝图片，没有默认使用标准酒店房型图片
	_pic *model.File
	// 宝贝状态,1上架。
	_status int64
	// 废弃，房型是否提供发票
	_hasReceipt bool
}

// NewTaobaoxhotelroomupdateRequest 初始化TaobaoxhotelroomupdateAPIRequest对象
func NewTaobaoxhotelroomupdateRequest() *TaobaoxhotelroomupdateAPIRequest {
	return &TaobaoxhotelroomupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelroomupdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.room.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelroomupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelroomupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOutRid is OutRid Setter
// 卖家房型ID
func (r *TaobaoxhotelroomupdateAPIRequest) SetOutRid(_outRid string) error {
	r._outRid = _outRid
	r.Set("out_rid", _outRid)
	return nil
}

// GetOutRid OutRid Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetOutRid() string {
	return r._outRid
}

// SetVendor is Vendor Setter
// 系统商，一般不填写，使用须申请
func (r *TaobaoxhotelroomupdateAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetVendor() string {
	return r._vendor
}

// SetTitle is Title Setter
// 废弃，宝贝名称展示在店铺里
func (r *TaobaoxhotelroomupdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetTitle() string {
	return r._title
}

// SetGuide is Guide Setter
// 废弃，房型购买须知展示在PC购物路径
func (r *TaobaoxhotelroomupdateAPIRequest) SetGuide(_guide string) error {
	r._guide = _guide
	r.Set("guide", _guide)
	return nil
}

// GetGuide Guide Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetGuide() string {
	return r._guide
}

// SetDesc is Desc Setter
// 废弃，宝贝描述展示在宝贝详情页面
func (r *TaobaoxhotelroomupdateAPIRequest) SetDesc(_desc string) error {
	r._desc = _desc
	r.Set("desc", _desc)
	return nil
}

// GetDesc Desc Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetDesc() string {
	return r._desc
}

// SetReceiptType is ReceiptType Setter
// 废弃，房型发票类型。A,B。分别代表： A:酒店住宿发票,B:其他
func (r *TaobaoxhotelroomupdateAPIRequest) SetReceiptType(_receiptType string) error {
	r._receiptType = _receiptType
	r.Set("receipt_type", _receiptType)
	return nil
}

// GetReceiptType ReceiptType Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetReceiptType() string {
	return r._receiptType
}

// SetReceiptOtherTypeDesc is ReceiptOtherTypeDesc Setter
// 废弃，房型发票类型为其他时的发票描述,不能超过30个字
func (r *TaobaoxhotelroomupdateAPIRequest) SetReceiptOtherTypeDesc(_receiptOtherTypeDesc string) error {
	r._receiptOtherTypeDesc = _receiptOtherTypeDesc
	r.Set("receipt_other_type_desc", _receiptOtherTypeDesc)
	return nil
}

// GetReceiptOtherTypeDesc ReceiptOtherTypeDesc Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetReceiptOtherTypeDesc() string {
	return r._receiptOtherTypeDesc
}

// SetReceiptInfo is ReceiptInfo Setter
// 废弃，房型发票说明，不能超过100个字
func (r *TaobaoxhotelroomupdateAPIRequest) SetReceiptInfo(_receiptInfo string) error {
	r._receiptInfo = _receiptInfo
	r.Set("receipt_info", _receiptInfo)
	return nil
}

// GetReceiptInfo ReceiptInfo Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetReceiptInfo() string {
	return r._receiptInfo
}

// SetInventory is Inventory Setter
// 房型共享库存日历。quota物理库存；al_quota保留房库存；sp_quota超预定库存。其中保留房库存和超预定库存需要向运营申请权限示例：[{&#34;date&#34;:2011-01-28,&#34;quota&#34;:10,&#34;al_quota&#34;:2,&#34;sp_quota&#34;:3}]
func (r *TaobaoxhotelroomupdateAPIRequest) SetInventory(_inventory string) error {
	r._inventory = _inventory
	r.Set("inventory", _inventory)
	return nil
}

// GetInventory Inventory Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetInventory() string {
	return r._inventory
}

// SetRoomSwitchCal is RoomSwitchCal Setter
// 房型库存开关。 1，开；2，关
func (r *TaobaoxhotelroomupdateAPIRequest) SetRoomSwitchCal(_roomSwitchCal string) error {
	r._roomSwitchCal = _roomSwitchCal
	r.Set("room_switch_cal", _roomSwitchCal)
	return nil
}

// GetRoomSwitchCal RoomSwitchCal Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetRoomSwitchCal() string {
	return r._roomSwitchCal
}

// SetSuperbookEndTime is SuperbookEndTime Setter
// 超预定库存截止时间
func (r *TaobaoxhotelroomupdateAPIRequest) SetSuperbookEndTime(_superbookEndTime string) error {
	r._superbookEndTime = _superbookEndTime
	r.Set("superbook_end_time", _superbookEndTime)
	return nil
}

// GetSuperbookEndTime SuperbookEndTime Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetSuperbookEndTime() string {
	return r._superbookEndTime
}

// SetSuperbookStartTime is SuperbookStartTime Setter
// 超预定库存开始时间
func (r *TaobaoxhotelroomupdateAPIRequest) SetSuperbookStartTime(_superbookStartTime string) error {
	r._superbookStartTime = _superbookStartTime
	r.Set("superbook_start_time", _superbookStartTime)
	return nil
}

// GetSuperbookStartTime SuperbookStartTime Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetSuperbookStartTime() string {
	return r._superbookStartTime
}

// SetAllotmentEndTime is AllotmentEndTime Setter
// 保留房库存截止时间
func (r *TaobaoxhotelroomupdateAPIRequest) SetAllotmentEndTime(_allotmentEndTime string) error {
	r._allotmentEndTime = _allotmentEndTime
	r.Set("allotment_end_time", _allotmentEndTime)
	return nil
}

// GetAllotmentEndTime AllotmentEndTime Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetAllotmentEndTime() string {
	return r._allotmentEndTime
}

// SetAllotmentStartTime is AllotmentStartTime Setter
// 保留房库存截止时间
func (r *TaobaoxhotelroomupdateAPIRequest) SetAllotmentStartTime(_allotmentStartTime string) error {
	r._allotmentStartTime = _allotmentStartTime
	r.Set("allotment_start_time", _allotmentStartTime)
	return nil
}

// GetAllotmentStartTime AllotmentStartTime Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetAllotmentStartTime() string {
	return r._allotmentStartTime
}

// SetGid is Gid Setter
// 废弃，使用out_rid
func (r *TaobaoxhotelroomupdateAPIRequest) SetGid(_gid int64) error {
	r._gid = _gid
	r.Set("gid", _gid)
	return nil
}

// GetGid Gid Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetGid() int64 {
	return r._gid
}

// SetPic is Pic Setter
// 废弃，宝贝图片，没有默认使用标准酒店房型图片
func (r *TaobaoxhotelroomupdateAPIRequest) SetPic(_pic *model.File) error {
	r._pic = _pic
	r.Set("pic", _pic)
	return nil
}

// GetPic Pic Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetPic() *model.File {
	return r._pic
}

// SetStatus is Status Setter
// 宝贝状态,1上架。
func (r *TaobaoxhotelroomupdateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetHasReceipt is HasReceipt Setter
// 废弃，房型是否提供发票
func (r *TaobaoxhotelroomupdateAPIRequest) SetHasReceipt(_hasReceipt bool) error {
	r._hasReceipt = _hasReceipt
	r.Set("has_receipt", _hasReceipt)
	return nil
}

// GetHasReceipt HasReceipt Getter
func (r TaobaoxhotelroomupdateAPIRequest) GetHasReceipt() bool {
	return r._hasReceipt
}
