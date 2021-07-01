package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbItemQueryAPIRequest
分页查询商品 API请求
taobao.wlb.item.query

根据状态、卖家、SKU等信息查询商品列表 */
type TaobaoWlbItemQueryAPIRequest struct {
	model.Params
	// 是否是最小库存单元，只有最小库存单元的商品才可以有库存,值只能给"true","false"来表示;  若值不在范围内，则按true处理
	_isSku string
	// 只能输入以下值或空：  ITEM_STATUS_VALID -- 1 表示 有效；  ITEM_STATUS_LOCK -- 2 表示锁住。  若值不在范围内，按ITEM_STATUS_VALID处理
	_status string
	// ITEM类型(只允许输入以下英文或空)  NORMAL 0:普通商品;  COMBINE 1:是否是组合商品  DISTRIBUTION 2:是否是分销商品(货主是别人)  若值不在范围内，则按NORMAL处理
	_itemType string
	// 商品名称
	_name string
	// 商品前台销售名字
	_title string
	// 商家编码
	_itemCode string
	// 父ID,只有is_sku=1时才能有父ID，商品也可以没有付商品
	_parentId int64
	// 当前页
	_pageNo int64
	// 分页记录个数，如果用户输入的记录数大于50，则一页显示50条记录
	_pageSize int64
}

// New
