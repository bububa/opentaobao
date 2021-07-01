package ju

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJuItemsSearchAPIRequest
聚划算商品搜索接口 API请求
taobao.ju.items.search

搜索聚划算商品 */
type TaobaoJuItemsSearchAPIRequest struct {
	model.Params
	// query
	_paramTopItemQuery *TopItemQuery
}

// New
