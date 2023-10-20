package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamosbrandcoproductgroupusercount 按照查询条件统计总数
// alibaba.mos.brand.coproduct.group.user.count
//
// 按照查询条件统计总数
func Alibabamosbrandcoproductgroupusercount(clt *core.SDKClient, req *mos.AlibabamosbrandcoproductgroupusercountAPIRequest, session string) (*mos.AlibabamosbrandcoproductgroupusercountAPIResponse, error) {
	var resp mos.AlibabamosbrandcoproductgroupusercountAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
