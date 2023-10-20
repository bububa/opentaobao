package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotimeget 获取淘宝系统当前时间
// taobao.time.get
//
// 获取淘宝系统当前时间
func Taobaotimeget(clt *core.SDKClient, req *util.TaobaotimegetAPIRequest, session string) (*util.TaobaotimegetAPIResponse, error) {
	var resp util.TaobaotimegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
