package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportquerynotitempromotion 其他主体报表查询
// taobao.universalbp.report.query.not.item.promotion
//
// 其他主体报表查询
func Taobaouniversalbpreportquerynotitempromotion(clt *core.SDKClient, req *simba.TaobaouniversalbpreportquerynotitempromotionAPIRequest, session string) (*simba.TaobaouniversalbpreportquerynotitempromotionAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportquerynotitempromotionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
