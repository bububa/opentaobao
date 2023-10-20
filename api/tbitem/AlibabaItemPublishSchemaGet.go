package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Alibabaitempublishschemaget 获取商品发布规则信息
// alibaba.item.publish.schema.get
//
// 新商品发布，获取商品发布规则信息
func Alibabaitempublishschemaget(clt *core.SDKClient, req *tbitem.AlibabaitempublishschemagetAPIRequest, session string) (*tbitem.AlibabaitempublishschemagetAPIResponse, error) {
	var resp tbitem.AlibabaitempublishschemagetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
