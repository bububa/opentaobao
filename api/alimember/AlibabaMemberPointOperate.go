package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// Alibabamemberpointoperate 积分操作
// alibaba.member.point.operate
//
// 消费会员积分
func Alibabamemberpointoperate(clt *core.SDKClient, req *alimember.AlibabamemberpointoperateAPIRequest, session string) (*alimember.AlibabamemberpointoperateAPIResponse, error) {
	var resp alimember.AlibabamemberpointoperateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
