package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenStoreQuery 门店信息查询接口
// taobao.qimen.store.query
//
// 商家在ERP等系统中调用该接口，查询门店相关信息
func TaobaoQimenStoreQuery(clt *core.SDKClient, req *qimen.TaobaoQimenStoreQueryAPIRequest, resp *qimen.TaobaoQimenStoreQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
