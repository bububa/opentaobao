package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcrowddelete 删除单品人群
// taobao.feedflow.item.crowd.delete
//
// 删除单品人群
func Taobaofeedflowitemcrowddelete(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcrowddeleteAPIRequest, session string) (*feedflow.TaobaofeedflowitemcrowddeleteAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcrowddeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
