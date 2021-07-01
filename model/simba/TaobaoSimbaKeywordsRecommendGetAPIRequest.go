package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsRecommendGetAPIRequest
取得一个推广组的推荐关键词列表 API请求
taobao.simba.keywords.recommend.get

取得一个推广组的推荐关键词列表 */
type TaobaoSimbaKeywordsRecommendGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组ID
	_adgroupId int64
	// 搜索量,设置此值后返回的就是大于此搜索量的词列表
	_search int64
	// 相关度
	_pertinence string
	// 返回的每页数据量大小,最大200
	_pageSize int64
	// 返回的第几页数据，默认为1
	_pageNo int64
	// 排序方式: 搜索量 search_volume 市场平均价格 average_price 相关度 relevance 不排序 non 默认为 non
	_orderBy string
}

// New
