package iot

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/iot"
)

// Tmalldevicestorefollowurlget 获取店铺关注链接
// tmall.device.store.followurl.get
//
// 获取智能硬件上的关注店铺的URL
func Tmalldevicestorefollowurlget(clt *core.SDKClient, req *iot.TmalldevicestorefollowurlgetAPIRequest, session string) (*iot.TmalldevicestorefollowurlgetAPIResponse, error) {
	var resp iot.TmalldevicestorefollowurlgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
