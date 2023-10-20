package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaosimbarptcusteffectget 用户账户报表效果数据查询（只有汇总数据，无分类数据）
// taobao.simba.rpt.custeffect.get
//
// 用户账户报表效果数据查询（只有汇总数据，无分类数据）
func Taobaosimbarptcusteffectget(clt *core.SDKClient, req *simba.TaobaosimbarptcusteffectgetAPIRequest, session string) (*simba.TaobaosimbarptcusteffectgetAPIResponse, error) {
	var resp simba.TaobaosimbarptcusteffectgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
