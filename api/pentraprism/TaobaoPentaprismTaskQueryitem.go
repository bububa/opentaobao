package pentraprism

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pentraprism"
)

// TaobaoPentaprismTaskQueryitem 查询任务当前进度
// taobao.pentaprism.task.queryitem
//
// 外网用户查询五棱镜任务系统当前进度
func TaobaoPentaprismTaskQueryitem(clt *core.SDKClient, req *pentraprism.TaobaoPentaprismTaskQueryitemAPIRequest, session string) (*pentraprism.TaobaoPentaprismTaskQueryitemAPIResponse, error) {
	var resp pentraprism.TaobaoPentaprismTaskQueryitemAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
