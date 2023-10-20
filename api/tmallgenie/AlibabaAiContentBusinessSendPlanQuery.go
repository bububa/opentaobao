package tmallgenie

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// Alibabaaicontentbusinesssendplanquery 内容商业化发放权益查询
// alibaba.ai.content.business.send.plan.query
//
// 天猫精灵内容生态的权益查询
func Alibabaaicontentbusinesssendplanquery(clt *core.SDKClient, req *tmallgenie.AlibabaaicontentbusinesssendplanqueryAPIRequest, session string) (*tmallgenie.AlibabaaicontentbusinesssendplanqueryAPIResponse, error) {
	var resp tmallgenie.AlibabaaicontentbusinesssendplanqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
