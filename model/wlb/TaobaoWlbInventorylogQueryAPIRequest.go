package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbInventorylogQueryAPIRequest
根据商品ID查询所有库存变更记录 API请求
taobao.wlb.inventorylog.query

通过商品ID等几个条件来分页查询库存变更记录 */
type TaobaoWlbInventorylogQueryAPIRequest struct {
	model.Params
	// 商品ID
	_itemId int64
	// 仓库编码
	_storeCode string
	// 单号
	_orderCode string
	// 起始修改时间,大于等于该时间
	_gmtStart string
	// 结束修改时间,小于等于该时间
	_gmtEnd string
	// 当前页
	_pageNo int64
	// 分页记录个数
	_pageSize int64
	// 可指定授权的用户来查询
	_opUserId int64
	// 库存操作作类型(可以为空) CHU_KU 1-出库 RU_KU 2-入库 FREEZE 3-冻结 THAW 4-解冻 CHECK_FREEZE 5-冻结确认 CHANGE_KU 6-库存类型变更 若值不在范围内，则按CHU_KU处理
	_opType string
}

// New
