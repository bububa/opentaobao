package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelRoomUpdateAPIRequest
房型库存推送接口（全量推送） API请求
taobao.xhotel.room.update

此接口用于更新一个酒店商品，根据传入的gid更新商品信息，该商品必须为对应的发布者才能执行更新操作。如果对应的商品在酒店系统中不存在，则会返回错误提示。 */
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

// New
