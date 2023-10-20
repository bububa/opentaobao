package alimember

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberPointOperate 积分操作
// alibaba.member.point.operate
//
// 消费会员积分
func AlibabaMemberPointOperate(clt *core.SDKClient, req *alimember.AlibabaMemberPointOperateAPIRequest, resp *alimember.AlibabaMemberPointOperateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
