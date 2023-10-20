package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmrulequerymdayerule 查询品牌下的会员日规则
// alibaba.alsc.crm.rule.querymdayerule
//
// 查询品牌下的会员日规则
func Alibabaalsccrmrulequerymdayerule(clt *core.SDKClient, req *alsc.AlibabaalsccrmrulequerymdayeruleAPIRequest, session string) (*alsc.AlibabaalsccrmrulequerymdayeruleAPIResponse, error) {
	var resp alsc.AlibabaalsccrmrulequerymdayeruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
