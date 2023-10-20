package usergrowth

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/usergrowth"
)

// TaobaoUsergrowthTaskConfigGet 用户增长营销玩法配置查询
// taobao.usergrowth.task.config.get
//
// 用户增长营销玩法配置查询
func TaobaoUsergrowthTaskConfigGet(clt *core.SDKClient, req *usergrowth.TaobaoUsergrowthTaskConfigGetAPIRequest, resp *usergrowth.TaobaoUsergrowthTaskConfigGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
