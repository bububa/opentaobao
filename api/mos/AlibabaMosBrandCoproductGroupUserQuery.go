package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosBrandCoproductGroupUserQuery 按照条件查询分页数据
// alibaba.mos.brand.coproduct.group.user.query
//
// 按照条件查询分页数据
func AlibabaMosBrandCoproductGroupUserQuery(clt *core.SDKClient, req *mos.AlibabaMosBrandCoproductGroupUserQueryAPIRequest, session string) (*mos.AlibabaMosBrandCoproductGroupUserQueryAPIResponse, error) {
	var resp mos.AlibabaMosBrandCoproductGroupUserQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
