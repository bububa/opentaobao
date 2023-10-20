package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallFuwuServiceitemList 获取服务商品扩展信息
// tmall.fuwu.serviceitem.list
//
// 获取服务商品扩展信息
func TmallFuwuServiceitemList(clt *core.SDKClient, req *tmallservice.TmallFuwuServiceitemListAPIRequest, session string) (*tmallservice.TmallFuwuServiceitemListAPIResponse, error) {
	var resp tmallservice.TmallFuwuServiceitemListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
