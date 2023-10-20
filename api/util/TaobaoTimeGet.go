package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTimeGet 获取淘宝系统当前时间
// taobao.time.get
//
// 获取淘宝系统当前时间
func TaobaoTimeGet(clt *core.SDKClient, req *util.TaobaoTimeGetAPIRequest, resp *util.TaobaoTimeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
