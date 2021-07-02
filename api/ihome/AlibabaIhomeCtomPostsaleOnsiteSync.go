package ihome

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ihome"
)

// AlibabaIhomeCtomPostsaleOnsiteSync 售后上门信息同步
// alibaba.ihome.ctom.postsale.onsite.sync
//
// 用于三维家同步售后单上门人员和时间信息
func AlibabaIhomeCtomPostsaleOnsiteSync(clt *core.SDKClient, req *ihome.AlibabaIhomeCtomPostsaleOnsiteSyncAPIRequest, session string) (*ihome.AlibabaIhomeCtomPostsaleOnsiteSyncAPIResponse, error) {
	var resp ihome.AlibabaIhomeCtomPostsaleOnsiteSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
