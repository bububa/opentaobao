package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportqueryarea 地域报表查询
// taobao.universalbp.report.query.area
//
// 地域报表查询
func Taobaouniversalbpreportqueryarea(clt *core.SDKClient, req *simba.TaobaouniversalbpreportqueryareaAPIRequest, session string) (*simba.TaobaouniversalbpreportqueryareaAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportqueryareaAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
