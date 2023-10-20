package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// Taobaocrmhistoryouidget 根据buyerNick获取ouid
// taobao.crm.history.ouid.get
//
// 根据buyerNick获取ouid
func Taobaocrmhistoryouidget(clt *core.SDKClient, req *topoaid.TaobaocrmhistoryouidgetAPIRequest, session string) (*topoaid.TaobaocrmhistoryouidgetAPIResponse, error) {
	var resp topoaid.TaobaocrmhistoryouidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
