package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupadd 信息流增加单元
// taobao.feedflow.item.adgroup.add
//
// 信息流增加单元
func Taobaofeedflowitemadgroupadd(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupaddAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupaddAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
