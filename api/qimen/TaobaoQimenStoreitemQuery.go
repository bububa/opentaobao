package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstoreitemquery 门店关联商品查询接口
// taobao.qimen.storeitem.query
//
// 商家在ERP等系统中调用该接口，查询某门店所关联的线上商品列表
func Taobaoqimenstoreitemquery(clt *core.SDKClient, req *qimen.TaobaoqimenstoreitemqueryAPIRequest, session string) (*qimen.TaobaoqimenstoreitemqueryAPIResponse, error) {
	var resp qimen.TaobaoqimenstoreitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
