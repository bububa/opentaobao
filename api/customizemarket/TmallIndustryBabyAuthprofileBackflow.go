package customizemarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// TmallIndustryBabyAuthprofileBackflow 孕校云回流档案
// tmall.industry.baby.authprofile.backflow
//
// 孕校云回流档案
func TmallIndustryBabyAuthprofileBackflow(clt *core.SDKClient, req *customizemarket.TmallIndustryBabyAuthprofileBackflowAPIRequest, session string) (*customizemarket.TmallIndustryBabyAuthprofileBackflowAPIResponse, error) {
	var resp customizemarket.TmallIndustryBabyAuthprofileBackflowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
