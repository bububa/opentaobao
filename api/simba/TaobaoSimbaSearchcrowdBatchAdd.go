package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSearchcrowdBatchAdd 推广单元增加搜索人群
// taobao.simba.searchcrowd.batch.add
//
// 推广单元新增搜索人群
func TaobaoSimbaSearchcrowdBatchAdd(clt *core.SDKClient, req *simba.TaobaoSimbaSearchcrowdBatchAddAPIRequest, session string) (*simba.TaobaoSimbaSearchcrowdBatchAddAPIResponse, error) {
	var resp simba.TaobaoSimbaSearchcrowdBatchAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
