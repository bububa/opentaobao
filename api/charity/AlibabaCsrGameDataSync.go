package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrGameDataSync 公益互动 外部游戏数据同步
// alibaba.csr.game.data.sync
//
// 公益互动 外部游戏数据同步
func AlibabaCsrGameDataSync(clt *core.SDKClient, req *charity.AlibabaCsrGameDataSyncAPIRequest, session string) (*charity.AlibabaCsrGameDataSyncAPIResponse, error) {
	var resp charity.AlibabaCsrGameDataSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
