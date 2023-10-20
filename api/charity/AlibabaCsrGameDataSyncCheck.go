package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrGameDataSyncCheck 公益互动 外部游戏数据同步-校验
// alibaba.csr.game.data.sync.check
//
// 公益互动 外部游戏数据同步-校验
func AlibabaCsrGameDataSyncCheck(clt *core.SDKClient, req *charity.AlibabaCsrGameDataSyncCheckAPIRequest, session string) (*charity.AlibabaCsrGameDataSyncCheckAPIResponse, error) {
	var resp charity.AlibabaCsrGameDataSyncCheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
