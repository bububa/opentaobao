package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosbrandcoproductgroupuserquery 按照条件查询分页数据
// alibaba.mos.brand.coproduct.group.user.query
//
// 按照条件查询分页数据
func Alibabamosbrandcoproductgroupuserquery(clt *core.SDKClient, req *mos.AlibabamosbrandcoproductgroupuserqueryAPIRequest, session string) (*mos.AlibabamosbrandcoproductgroupuserqueryAPIResponse, error) {
	var resp mos.AlibabamosbrandcoproductgroupuserqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
