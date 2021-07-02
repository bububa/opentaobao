package ihome

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ihome"
)

// AlibabaIhomeCtomPostsaleStatusSync C2M售后状态同步
// alibaba.ihome.ctom.postsale.status.sync
//
// 供给三维家同步定制、成品商品售后进度状态
func AlibabaIhomeCtomPostsaleStatusSync(clt *core.SDKClient, req *ihome.AlibabaIhomeCtomPostsaleStatusSyncAPIRequest, session string) (*ihome.AlibabaIhomeCtomPostsaleStatusSyncAPIResponse, error) {
	var resp ihome.AlibabaIhomeCtomPostsaleStatusSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
