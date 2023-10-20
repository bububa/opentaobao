package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenitemmappingquery 前后端商品映射查询接口
// taobao.qimen.itemmapping.query
//
// 前后端商品映射查询接口
func Taobaoqimenitemmappingquery(clt *core.SDKClient, req *qimen.TaobaoqimenitemmappingqueryAPIRequest, session string) (*qimen.TaobaoqimenitemmappingqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenitemmappingqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
