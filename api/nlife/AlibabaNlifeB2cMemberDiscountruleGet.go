package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cMemberDiscountruleGet 会员抵扣规则
// alibaba.nlife.b2c.member.discountrule.get
//
// 获取企业会员抵扣规则
func AlibabaNlifeB2cMemberDiscountruleGet(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cMemberDiscountruleGetAPIRequest, resp *nlife.AlibabaNlifeB2cMemberDiscountruleGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
