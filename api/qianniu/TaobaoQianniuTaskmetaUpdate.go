package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskmetaUpdate 更新任务元数据
// taobao.qianniu.taskmeta.update
//
// 由任务发起者调用
func TaobaoQianniuTaskmetaUpdate(clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskmetaUpdateAPIRequest, resp *qianniu.TaobaoQianniuTaskmetaUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
