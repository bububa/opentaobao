package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoCrmHistoryOmidGet 根据buyerNick获取omid
// taobao.crm.history.omid.get
//
// 根据buyerNick获取ouid
func TaobaoCrmHistoryOmidGet(clt *core.SDKClient, req *topoaid.TaobaoCrmHistoryOmidGetAPIRequest, session string) (*topoaid.TaobaoCrmHistoryOmidGetAPIResponse, error) {
	var resp topoaid.TaobaoCrmHistoryOmidGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
