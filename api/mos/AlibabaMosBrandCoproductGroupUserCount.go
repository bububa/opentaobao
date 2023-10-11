package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosBrandCoproductGroupUserCount 按照查询条件统计总数
// alibaba.mos.brand.coproduct.group.user.count
//
// 按照查询条件统计总数
func AlibabaMosBrandCoproductGroupUserCount(clt *core.SDKClient, req *mos.AlibabaMosBrandCoproductGroupUserCountAPIRequest, session string) (*mos.AlibabaMosBrandCoproductGroupUserCountAPIResponse, error) {
	var resp mos.AlibabaMosBrandCoproductGroupUserCountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
