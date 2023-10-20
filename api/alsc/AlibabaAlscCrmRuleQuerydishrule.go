package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulequerydishrule 查询品牌下的入会菜品规则
// alibaba.alsc.crm.rule.querydishrule
//
// 查询品牌下的入会菜品规则
func Alibabaalsccrmrulequerydishrule(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulequerydishruleAPIRequest, session string) (*alsc.AlibabaalsccrmrulequerydishruleAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulequerydishruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
