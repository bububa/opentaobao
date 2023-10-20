package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulequerytaglist 查询标签列表
// alibaba.alsc.crm.rule.querytaglist
//
// 查询标签列表
func Alibabaalsccrmrulequerytaglist(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulequerytaglistAPIRequest, session string) (*alsc.AlibabaalsccrmrulequerytaglistAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulequerytaglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
