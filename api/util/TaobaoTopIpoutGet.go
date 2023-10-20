package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopipoutget 获取开放平台出口IP段
// taobao.top.ipout.get
//
// 获取开放平台出口IP段
func Taobaotopipoutget(clt *core.SDKClient, req *util.TaobaotopipoutgetAPIRequest, session string) (*util.TaobaotopipoutgetAPIResponse, error) {
	var resp util.TaobaotopipoutgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
