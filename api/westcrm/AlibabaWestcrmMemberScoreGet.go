package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

// AlibabaWestcrmMemberScoreGet 获取会员某时间段积分
// alibaba.westcrm.member.score.get
//
// 获取会员某时间段积分
func AlibabaWestcrmMemberScoreGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmMemberScoreGetAPIRequest, session string) (*westcrm.AlibabaWestcrmMemberScoreGetAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmMemberScoreGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
