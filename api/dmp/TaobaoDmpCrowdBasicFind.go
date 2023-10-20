package dmp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/dmp"
)

// Taobaodmpcrowdbasicfind DMP_BP版人群列表查询
// taobao.dmp.crowd.basic.find
//
// DMP_BP版人群列表查询
func Taobaodmpcrowdbasicfind(clt *core.SDKClient, req *dmp.TaobaodmpcrowdbasicfindAPIRequest, session string) (*dmp.TaobaodmpcrowdbasicfindAPIResponse, error) {
	var resp dmp.TaobaodmpcrowdbasicfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
