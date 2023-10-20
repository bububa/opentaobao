package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacsrgamedatasynccheck 公益互动 外部游戏数据同步-校验
// alibaba.csr.game.data.sync.check
//
// 公益互动 外部游戏数据同步-校验
func Alibabacsrgamedatasynccheck(clt *core.SDKClient, req *charity.AlibabacsrgamedatasynccheckAPIRequest, session string) (*charity.AlibabacsrgamedatasynccheckAPIResponse, error) {
	var resp charity.AlibabacsrgamedatasynccheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
