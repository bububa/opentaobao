package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupcreativeaddbind 信息流新增并且绑定创意
// taobao.feedflow.item.adgroup.creative.add.bind
//
// 信息流新增并且绑定创意
func Taobaofeedflowitemadgroupcreativeaddbind(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupcreativeaddbindAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupcreativeaddbindAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupcreativeaddbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
