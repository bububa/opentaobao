package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniitemCategoryGetAPIRequest
全渠道商品轻发布类目信息 API请求
taobao.omniitem.category.get

全渠道商品轻发布类目信息 */
type TaobaoOmniitemCategoryGetAPIRequest struct {
	model.Params
	// 全渠道商品类目ID，不填表示获取所有全渠道商品类目信息
	_categoryId int64
}

// New
