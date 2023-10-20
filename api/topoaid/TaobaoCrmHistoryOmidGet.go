package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// Taobaocrmhistoryomidget 根据buyerNick获取omid
// taobao.crm.history.omid.get
//
// 根据buyerNick获取ouid
func Taobaocrmhistoryomidget(clt *core.SDKClient, req *topoaid.TaobaocrmhistoryomidgetAPIRequest, session string) (*topoaid.TaobaocrmhistoryomidgetAPIResponse, error) {
	var resp topoaid.TaobaocrmhistoryomidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
