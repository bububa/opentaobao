package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSerchcrowdBatchDelete 单品搜索人群批量取消溢价
// taobao.simba.serchcrowd.batch.delete
//
// 删除单品搜索人群溢价功能
func TaobaoSimbaSerchcrowdBatchDelete(clt *core.SDKClient, req *simba.TaobaoSimbaSerchcrowdBatchDeleteAPIRequest, resp *simba.TaobaoSimbaSerchcrowdBatchDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
