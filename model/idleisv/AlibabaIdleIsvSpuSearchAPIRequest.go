package idleisv

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleIsvSpuSearchAPIRequest
spu搜索接口 API请求
alibaba.idle.isv.spu.search

搜索的品牌和型号，供服务商进行选择 */
type AlibabaIdleIsvSpuSearchAPIRequest struct {
	model.Params
	// 闲鱼渠道类目的id
	_channelCatId string
	// 搜索的文本
	_searchText string
}

// New
