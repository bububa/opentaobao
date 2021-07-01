package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoXhotelQuotaUpdateAPIRequest
库存更新接口 API请求
taobao.xhotel.quota.update

库存更新接口 */
type TaobaoXhotelQuotaUpdateAPIRequest struct {
	model.Params
	// 库存类型,0: 普通库存, 1: 普通保留房库存, 2:协议保留房库存
	_quotaType int64
	// 是否使用room库存,true使用，false不使用
	_useRoomInventory bool
	// room的gid
	_gid int64
	// 增减的值
	_quota int64
	// 数量类型, 2:增加房量,3:减少房量
	_quotaNumType int64
	// 修改日期列表
	_dates []string
	// rate的id, rate库存时必填
	_rateId int64
}

// New
