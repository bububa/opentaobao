package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// Tmallcrmmemberpointchange 会员积分变更
// tmall.crm.member.point.change
//
// 品牌CRM项目中，会员积分变更接口。
func Tmallcrmmemberpointchange(clt *core.SDKClient, req *mei.TmallcrmmemberpointchangeAPIRequest, session string) (*mei.TmallcrmmemberpointchangeAPIResponse, error) {
	var resp mei.TmallcrmmemberpointchangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
