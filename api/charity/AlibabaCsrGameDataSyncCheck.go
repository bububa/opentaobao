package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrGameDataSyncCheck 公益互动 外部游戏数据同步-校验
// alibaba.csr.game.data.sync.check
//
// 公益互动 外部游戏数据同步-校验
func AlibabaCsrGameDataSyncCheck(clt *core.SDKClient, req *charity.AlibabaCsrGameDataSyncCheckAPIRequest, resp *charity.AlibabaCsrGameDataSyncCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
