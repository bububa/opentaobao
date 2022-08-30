package customizemarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// AlibabaTaobaoIndustryPetProfileQuery 用户宠物列表查询
// alibaba.taobao.industry.pet.profile.query
//
// 用户宠物列表查询
func AlibabaTaobaoIndustryPetProfileQuery(clt *core.SDKClient, req *customizemarket.AlibabaTaobaoIndustryPetProfileQueryAPIRequest, session string) (*customizemarket.AlibabaTaobaoIndustryPetProfileQueryAPIResponse, error) {
	var resp customizemarket.AlibabaTaobaoIndustryPetProfileQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
