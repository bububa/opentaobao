package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Tmallccfcrowdactivityuserupload 品牌营销活动用户上传
// tmall.ccf.crowd.activityuser.upload
//
// 搜集ISV的活动用户信息，将其沉淀为活动人群数据
func Tmallccfcrowdactivityuserupload(clt *core.SDKClient, req *promotion.TmallccfcrowdactivityuseruploadAPIRequest, session string) (*promotion.TmallccfcrowdactivityuseruploadAPIResponse, error) {
	var resp promotion.TmallccfcrowdactivityuseruploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
