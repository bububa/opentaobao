package jstinteractive

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractiveAssetsConfig 任务素材配置接口
// taobao.jst.interactive.assets.config
//
// 任务素材配置接口
func TaobaoJstInteractiveAssetsConfig(clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractiveAssetsConfigAPIRequest, session string) (*jstinteractive.TaobaoJstInteractiveAssetsConfigAPIResponse, error) {
	var resp jstinteractive.TaobaoJstInteractiveAssetsConfigAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
