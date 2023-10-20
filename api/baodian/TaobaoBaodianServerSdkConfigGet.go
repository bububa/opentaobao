package baodian

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baodian"
)

// Taobaobaodianserversdkconfigget 获取宝点SDK的配置项（已迁移）
// taobao.baodian.server.sdk.config.get
//
// 获取SDK各种配置项（已迁移）
func Taobaobaodianserversdkconfigget(clt *core.SDKClient, req *baodian.TaobaobaodianserversdkconfiggetAPIRequest, session string) (*baodian.TaobaobaodianserversdkconfiggetAPIResponse, error) {
	var resp baodian.TaobaobaodianserversdkconfiggetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
