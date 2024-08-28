package customizemarket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/customizemarket"
)

// TmallIndustryBabyAuthprofileBackflow 孕校云回流档案
// tmall.industry.baby.authprofile.backflow
//
// 孕校云回流档案
func TmallIndustryBabyAuthprofileBackflow(ctx context.Context, clt *core.SDKClient, req *customizemarket.TmallIndustryBabyAuthprofileBackflowAPIRequest, resp *customizemarket.TmallIndustryBabyAuthprofileBackflowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
