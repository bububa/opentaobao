package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportqueryadgroup 单元报表查询
// taobao.universalbp.report.query.adgroup
//
// 单元报表查询
func Taobaouniversalbpreportqueryadgroup(clt *core.SDKClient, req *simba.TaobaouniversalbpreportqueryadgroupAPIRequest, session string) (*simba.TaobaouniversalbpreportqueryadgroupAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportqueryadgroupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
