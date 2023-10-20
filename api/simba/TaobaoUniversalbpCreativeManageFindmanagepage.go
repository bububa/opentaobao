package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoUniversalbpCreativeManageFindmanagepage 创意库查询创意列表
// taobao.universalbp.creative.manage.findmanagepage
//
// 创意库查询创意列表
func TaobaoUniversalbpCreativeManageFindmanagepage(clt *core.SDKClient, req *simba.TaobaoUniversalbpCreativeManageFindmanagepageAPIRequest, session string) (*simba.TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse, error) {
	var resp simba.TaobaoUniversalbpCreativeManageFindmanagepageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
