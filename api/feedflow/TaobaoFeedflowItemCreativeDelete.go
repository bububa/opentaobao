package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcreativedelete 信息流删除创意
// taobao.feedflow.item.creative.delete
//
// 信息流删除创意
func Taobaofeedflowitemcreativedelete(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcreativedeleteAPIRequest, session string) (*feedflow.TaobaofeedflowitemcreativedeleteAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcreativedeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
