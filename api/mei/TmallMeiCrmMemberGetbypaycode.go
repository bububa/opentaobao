package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// Tmallmeicrmmembergetbypaycode 支付码获取会员信息
// tmall.mei.crm.member.getbypaycode
//
// 通过支付码获取会员信息
func Tmallmeicrmmembergetbypaycode(clt *core.SDKClient, req *mei.TmallmeicrmmembergetbypaycodeAPIRequest, session string) (*mei.TmallmeicrmmembergetbypaycodeAPIResponse, error) {
	var resp mei.TmallmeicrmmembergetbypaycodeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
