package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacsrgamedatasync 公益互动 外部游戏数据同步
// alibaba.csr.game.data.sync
//
// 公益互动 外部游戏数据同步
func Alibabacsrgamedatasync(clt *core.SDKClient, req *charity.AlibabacsrgamedatasyncAPIRequest, session string) (*charity.AlibabacsrgamedatasyncAPIResponse, error) {
	var resp charity.AlibabacsrgamedatasyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
