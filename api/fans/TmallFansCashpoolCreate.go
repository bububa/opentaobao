package fans

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// TmallFansCashpoolCreate 创建资金池
// tmall.fans.cashpool.create
//
// 商家创建资金池接口
func TmallFansCashpoolCreate(clt *core.SDKClient, req *fans.TmallFansCashpoolCreateAPIRequest, resp *fans.TmallFansCashpoolCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
