package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitCourttimePush 开庭时间推送（带附件）
// alibaba.legal.suit.courttime.push
//
// 开庭时间推送（带附件）
func AlibabaLegalSuitCourttimePush(clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourttimePushAPIRequest, session string) (*legalsuit.AlibabaLegalSuitCourttimePushAPIResponse, error) {
	var resp legalsuit.AlibabaLegalSuitCourttimePushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
