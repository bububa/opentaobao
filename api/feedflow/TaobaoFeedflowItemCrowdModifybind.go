package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// Taobaofeedflowitemcrowdmodifybind 修改人群出价或状态
// taobao.feedflow.item.crowd.modifybind
//
// 修改人群出价或状态
func Taobaofeedflowitemcrowdmodifybind(clt *core.SDKClient, req *feedflow.TaobaofeedflowitemcrowdmodifybindAPIRequest, session string) (*feedflow.TaobaofeedflowitemcrowdmodifybindAPIResponse, error) {
	var resp feedflow.TaobaofeedflowitemcrowdmodifybindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
