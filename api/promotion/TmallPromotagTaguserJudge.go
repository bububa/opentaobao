package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Tmallpromotagtaguserjudge 用户标签判断接口
// tmall.promotag.taguser.judge
//
// 查询用户是否有标签
func Tmallpromotagtaguserjudge(clt *core.SDKClient, req *promotion.TmallpromotagtaguserjudgeAPIRequest, session string) (*promotion.TmallpromotagtaguserjudgeAPIResponse, error) {
	var resp promotion.TmallpromotagtaguserjudgeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
