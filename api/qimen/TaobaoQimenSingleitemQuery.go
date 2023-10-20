package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimensingleitemquery 商品查询接口
// taobao.qimen.singleitem.query
//
// 商品查询接口
func Taobaoqimensingleitemquery(clt *core.SDKClient, req *qimen.TaobaoqimensingleitemqueryAPIRequest, session string) (*qimen.TaobaoqimensingleitemqueryAPIResponse, error) {
	var resp qimen.TaobaoqimensingleitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
