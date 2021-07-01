package category

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemCatpropsModificationGetAPIRequest
查询商品类目属性变更 API请求
taobao.item.catprops.modification.get

查询商品类目属性变更信息 */
type TaobaoItemCatpropsModificationGetAPIRequest struct {
	model.Params
	// 类目Id（与商品Id二选一即可）
	_categoryId int64
	// 商品Id（与类目Id二选一即可。若同时传入商品Id和类目Id，则优先使用商品Id。若填写商品Id，则起始时间设为该商品最近修改时间）
	_itemId string
	// 起始请求时间（建议传入，默认为90天内）
	_startTime string
}

// New
