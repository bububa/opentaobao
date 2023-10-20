package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenitemmappingcreate 前后端商品映射接口
// taobao.qimen.itemmapping.create
//
// 前后端商品映射
func Taobaoqimenitemmappingcreate(clt *core.SDKClient, req *qimen.TaobaoqimenitemmappingcreateAPIRequest, session string) (*qimen.TaobaoqimenitemmappingcreateAPIResponse, error) {
	var resp qimen.TaobaoqimenitemmappingcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
