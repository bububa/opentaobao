package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemUpdateListingAPIRequest
一口价商品上架 API请求
taobao.item.update.listing

* 单个商品上架<br/>* 输入的num_iid必须属于当前会话用户 */
type TaobaoItemUpdateListingAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
	// 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
	_num int64
}

// New
