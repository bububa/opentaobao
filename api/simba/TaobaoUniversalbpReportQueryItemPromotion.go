package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportqueryitempromotion 宝贝主体报表查询
// taobao.universalbp.report.query.item.promotion
//
// 宝贝主体报表查询
func Taobaouniversalbpreportqueryitempromotion(clt *core.SDKClient, req *simba.TaobaouniversalbpreportqueryitempromotionAPIRequest, session string) (*simba.TaobaouniversalbpreportqueryitempromotionAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportqueryitempromotionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
