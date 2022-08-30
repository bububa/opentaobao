package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoCrmHistoryOuidGet 根据buyerNick获取ouid
// taobao.crm.history.ouid.get
//
// 根据buyerNick获取ouid
func TaobaoCrmHistoryOuidGet(clt *core.SDKClient, req *topoaid.TaobaoCrmHistoryOuidGetAPIRequest, session string) (*topoaid.TaobaoCrmHistoryOuidGetAPIResponse, error) {
	var resp topoaid.TaobaoCrmHistoryOuidGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
