package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitcourtafterpush 更新或者新增庭后信息
// alibaba.legal.suit.court.after.push
//
// 供外部ISV供应商 推送庭后信息给集团诉讼
func Alibabalegalsuitcourtafterpush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitcourtafterpushAPIRequest, session string) (*legalsuit.AlibabalegalsuitcourtafterpushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitcourtafterpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
