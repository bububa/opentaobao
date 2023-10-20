package pentraprism

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pentraprism"
)

// Taobaopentaprismtasktriggerfrom 任务进度推进（根据fromtoken）
// taobao.pentaprism.task.trigger.from
//
// 外网用户推进单条五棱镜任务进度
func Taobaopentaprismtasktriggerfrom(clt *core.SDKClient, req *pentraprism.TaobaopentaprismtasktriggerfromAPIRequest, session string) (*pentraprism.TaobaopentaprismtasktriggerfromAPIResponse, error) {
	var resp pentraprism.TaobaopentaprismtasktriggerfromAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
