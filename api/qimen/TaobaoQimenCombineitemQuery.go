package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimencombineitemquery 组合货品关系查询接口
// taobao.qimen.combineitem.query
//
// 组合货品关系查询
func Taobaoqimencombineitemquery(clt *core.SDKClient, req *qimen.TaobaoqimencombineitemqueryAPIRequest, session string) (*qimen.TaobaoqimencombineitemqueryAPIResponse, error) {
	var resp qimen.TaobaoqimencombineitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
