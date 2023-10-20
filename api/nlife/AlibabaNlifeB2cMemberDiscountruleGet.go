package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeB2cMemberDiscountruleGet 会员抵扣规则
// alibaba.nlife.b2c.member.discountrule.get
//
// 获取企业会员抵扣规则
func AlibabaNlifeB2cMemberDiscountruleGet(clt *core.SDKClient, req *nlife.AlibabaNlifeB2cMemberDiscountruleGetAPIRequest, session string) (*nlife.AlibabaNlifeB2cMemberDiscountruleGetAPIResponse, error) {
	var resp nlife.AlibabaNlifeB2cMemberDiscountruleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
