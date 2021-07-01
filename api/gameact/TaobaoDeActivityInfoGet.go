package gameact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/gameact"
)

/* TaobaoDeActivityInfoGet
获取活动信息
taobao.de.activity.info.get

根据appKey和活动id获取活动 */
func TaobaoDeActivityInfoGet(clt *core.SDKClient, req *gameact.TaobaoDeActivityInfoGetAPIRequest, session string) (*gameact.TaobaoDeActivityInfoGetAPIResponse, error) {
	var resp gameact.TaobaoDeActivityInfoGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
