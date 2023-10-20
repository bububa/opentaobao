package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportquerycreative 创意报表查询
// taobao.universalbp.report.query.creative
//
// 创意报表查询
func Taobaouniversalbpreportquerycreative(clt *core.SDKClient, req *simba.TaobaouniversalbpreportquerycreativeAPIRequest, session string) (*simba.TaobaouniversalbpreportquerycreativeAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportquerycreativeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
