package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// TaobaoAlitripItFareBatchdelete 【国际机票自有政策】批量删除
// taobao.alitrip.it.fare.batchdelete
//
// 批量删除自有政策，单次删除最大5万，大于5万时候提示失败，需要缩小删除条件。此接口同步返回任务id，异步执行操作。每个接入方最多同时只能有10个处理中的任务，超过后直接返回失败。
func TaobaoAlitripItFareBatchdelete(clt *core.SDKClient, req *itpolicy.TaobaoAlitripItFareBatchdeleteAPIRequest, resp *itpolicy.TaobaoAlitripItFareBatchdeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
