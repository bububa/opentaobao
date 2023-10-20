package elife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/elife"
)

// Taobaoelifelifecardquery 查询交易结果
// taobao.elife.lifecard.query
//
// 卖家在交易状态不明的情况下, 查询交易结果.
func Taobaoelifelifecardquery(clt *core.SDKClient, req *elife.TaobaoelifelifecardqueryAPIRequest, session string) (*elife.TaobaoelifelifecardqueryAPIResponse, error) {
	var resp elife.TaobaoelifelifecardqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
