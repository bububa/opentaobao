package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerentitemquery 查询租赁商品信息
// alibaba.idle.rent.item.query
//
// 查询租赁商品信息
func Alibabaidlerentitemquery(clt *core.SDKClient, req *idle.AlibabaidlerentitemqueryAPIRequest, session string) (*idle.AlibabaidlerentitemqueryAPIResponse, error) {
	var resp idle.AlibabaidlerentitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
