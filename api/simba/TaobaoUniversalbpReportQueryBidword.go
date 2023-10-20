package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// Taobaouniversalbpreportquerybidword 关键词报表查询
// taobao.universalbp.report.query.bidword
//
// 关键词报表查询
func Taobaouniversalbpreportquerybidword(clt *core.SDKClient, req *simba.TaobaouniversalbpreportquerybidwordAPIRequest, session string) (*simba.TaobaouniversalbpreportquerybidwordAPIResponse, error) {
	var resp simba.TaobaouniversalbpreportquerybidwordAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
