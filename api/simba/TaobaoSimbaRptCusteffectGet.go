package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaRptCusteffectGet 用户账户报表效果数据查询（只有汇总数据，无分类数据）
// taobao.simba.rpt.custeffect.get
//
// 用户账户报表效果数据查询（只有汇总数据，无分类数据）
func TaobaoSimbaRptCusteffectGet(clt *core.SDKClient, req *simba.TaobaoSimbaRptCusteffectGetAPIRequest, session string) (*simba.TaobaoSimbaRptCusteffectGetAPIResponse, error) {
	var resp simba.TaobaoSimbaRptCusteffectGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
