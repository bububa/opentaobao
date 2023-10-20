package legalsuit

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalsuit"
)

// Alibabalegalsuitdominationpush 更新或者保存管辖信息
// alibaba.legal.suit.domination.push
//
// ISV推送管辖信息到诉讼平台
func Alibabalegalsuitdominationpush(clt *core.SDKClient, req *legalsuit.AlibabalegalsuitdominationpushAPIRequest, session string) (*legalsuit.AlibabalegalsuitdominationpushAPIResponse, error) {
	var resp legalsuit.AlibabalegalsuitdominationpushAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
