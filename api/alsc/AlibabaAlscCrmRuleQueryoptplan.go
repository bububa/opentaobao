package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulequeryoptplan 查询运营计划
// alibaba.alsc.crm.rule.queryoptplan
//
// 查询运营计划
func Alibabaalsccrmrulequeryoptplan(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulequeryoptplanAPIRequest, session string) (*alsc.AlibabaalsccrmrulequeryoptplanAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulequeryoptplanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
