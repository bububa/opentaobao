package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimencombineitemdelete 组合货品删除接口
// taobao.qimen.combineitem.delete
//
// 组合货品删除
func Taobaoqimencombineitemdelete(clt *core.SDKClient, req *qimen.TaobaoqimencombineitemdeleteAPIRequest, session string) (*qimen.TaobaoqimencombineitemdeleteAPIResponse, error) {
	var resp qimen.TaobaoqimencombineitemdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
