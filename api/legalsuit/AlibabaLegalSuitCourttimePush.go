package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitcourttimepush 开庭时间推送（带附件）
// alibaba.legal.suit.courttime.push
//
// 开庭时间推送（带附件）
func Alibabalegalsuitcourttimepush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitcourttimepushAPIRequest, session string) (*legalsuit.AlibabalegalsuitcourttimepushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitcourttimepushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
