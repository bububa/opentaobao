package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest
销量明星api相关接口 API请求
taobao.simba.salestar.keywords.recommend.get

取得一个推广组的推荐关键词列表 */
type TaobaoSimbaSalestarKeywordsRecommendGetAPIRequest struct {
	model.Params
	// 推广组ID
	_adgroupId int64
	// 产品类型101001005代表标准推广，101001014代表销量明星
	_productId int64
}

// New
