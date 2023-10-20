package customizemarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// AlibabaTaobaoIndustryPetProfileQuery 用户宠物列表查询
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
func AlibabaTaobaoIndustryPetProfileQuery(clt *core.SDKClient, req *customizemarket.AlibabaTaobaoIndustryPetProfileQueryAPIRequest, resp *customizemarket.AlibabaTaobaoIndustryPetProfileQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
