package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenitemstorequery 商品关联门店查询接口
// taobao.qimen.itemstore.query
//
// 商家在ERP等系统中调用该接口，查询线上商品所关联的门店列表
func Taobaoqimenitemstorequery(clt *core.SDKClient, req *qimen.TaobaoqimenitemstorequeryAPIRequest, session string) (*qimen.TaobaoqimenitemstorequeryAPIResponse, error) {
	var resp qimen.TaobaoqimenitemstorequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
