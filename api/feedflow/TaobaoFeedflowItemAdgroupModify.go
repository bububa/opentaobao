package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemadgroupmodify 信息流单元修改
// taobao.feedflow.item.adgroup.modify
//
// 信息流单元修改
func Taobaofeedflowitemadgroupmodify(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemadgroupmodifyAPIRequest, session string) (*feedflow.TaobaofeedflowitemadgroupmodifyAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemadgroupmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
