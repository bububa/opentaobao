package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSerchcrowdPriceBatchUpdate 单品推广搜索人群修改溢价
// taobao.simba.serchcrowd.price.batch.update
//
// 单品推广搜索人群修改溢价, 不支持跨推广单元修改
func TaobaoSimbaSerchcrowdPriceBatchUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaSerchcrowdPriceBatchUpdateAPIRequest, resp *simba.TaobaoSimbaSerchcrowdPriceBatchUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
