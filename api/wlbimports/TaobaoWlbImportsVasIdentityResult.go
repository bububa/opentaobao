package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// Taobaowlbimportsvasidentityresult 集货鉴定结果
// taobao.wlb.imports.vas.identity.result
//
// 集货鉴定结果查询
func Taobaowlbimportsvasidentityresult(clt *core.SDKClient, req *wlbimports.TaobaowlbimportsvasidentityresultAPIRequest, session string) (*wlbimports.TaobaowlbimportsvasidentityresultAPIResponse, error) {
	var resp wlbimports.TaobaowlbimportsvasidentityresultAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
