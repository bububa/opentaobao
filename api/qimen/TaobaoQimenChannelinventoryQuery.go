package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenchannelinventoryquery 渠道库存查询接口
// taobao.qimen.channelinventory.query
//
// 渠道库存查询
func Taobaoqimenchannelinventoryquery(clt *core.SDKClient, req *qimen.TaobaoqimenchannelinventoryqueryAPIRequest, session string) (*qimen.TaobaoqimenchannelinventoryqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenchannelinventoryqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
