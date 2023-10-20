package dmp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// Taobaodmpcrowdtemplateapply 人群模版采纳并生成人群API
// taobao.dmp.crowd.template.apply
//
// 人群模版采纳并生成人群API
func Taobaodmpcrowdtemplateapply(clt *core.SDKClient, req *dmp.TaobaodmpcrowdtemplateapplyAPIRequest, session string) (*dmp.TaobaodmpcrowdtemplateapplyAPIResponse, error) {
	var resp dmp.TaobaodmpcrowdtemplateapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
