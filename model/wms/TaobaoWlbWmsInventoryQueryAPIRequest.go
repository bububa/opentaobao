package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsInventoryQueryAPIRequest
菜鸟商品库存查询 API请求
taobao.wlb.wms.inventory.query

支持按汇总（不分批次和渠道的总的库存数量）、渠道、批次三类方式查询商品实时库存 */
type TaobaoWlbWmsInventoryQueryAPIRequest struct {
	model.Params
	// 菜鸟商品ID
	_itemId string
	// 仓库编码
	_storeCode string
	// 库存类型。 (1 正品 101 残次 102 机损 103 箱损 201 冻结库存 301 在途库存 )
	_inventoryType int64
	// 库存查询类型 1-	汇总库存，不区分批次和渠道 2-	批次库存，库存按商品批次维度划分 3-	渠道库存，库存按渠道维度划分 （当前业务不支持批次库存和渠道库存共存，批次库存无渠道属性，渠道库存无批次属性）
	_type int64
	// 库存批次号，type=2时字段传值有效。 商品设置为批次管理时，商品产生批次库存。当商品为批次管理时，此字段不传值，返回所有批次库存信息。
	_batchCode string
	// 生产日期，type=2时字段传值有效。
	_produceDate string
	// 失效日期，type=2时字段传值有效。
	_dueDate string
	// 渠道编码，type=3时字段传值有效。（TB 淘系， OTHERS 其他）
	_channelCode string
	// 分页的第几页
	_pageNo int64
	// 每页多少条，最大50条
	_pageSize int64
}

// New
