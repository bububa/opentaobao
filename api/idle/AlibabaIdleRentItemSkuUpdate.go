package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentitemskuupdate 更新/增加sku信息
// alibaba.idle.rent.item.sku.update
//
// 更新/增加sku信息
func Alibabaidlerentitemskuupdate(clt *core.SDKClient, req *idle.AlibabaidlerentitemskuupdateAPIRequest, session string) (*idle.AlibabaidlerentitemskuupdateAPIResponse, error) {
	var resp idle.AlibabaidlerentitemskuupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
