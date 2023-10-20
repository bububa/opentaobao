package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuServiceitemList 获取服务商品扩展信息
// tmall.fuwu.serviceitem.list
//
// 获取服务商品扩展信息
func TmallFuwuServiceitemList(clt *core.SDKClient, req *tmallservice.TmallFuwuServiceitemListAPIRequest, resp *tmallservice.TmallFuwuServiceitemListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
