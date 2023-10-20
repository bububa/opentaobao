package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSearchcrowdBatchAdd 推广单元增加搜索人群
// taobao.simba.searchcrowd.batch.add
//
// 推广单元新增搜索人群
func TaobaoSimbaSearchcrowdBatchAdd(clt *core.SDKClient, req *simba.TaobaoSimbaSearchcrowdBatchAddAPIRequest, resp *simba.TaobaoSimbaSearchcrowdBatchAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
