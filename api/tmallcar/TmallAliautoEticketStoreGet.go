package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoEticketStoreGet 查询电子凭证对应门店信息
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
func TmallAliautoEticketStoreGet(clt *core.SDKClient, req *tmallcar.TmallAliautoEticketStoreGetAPIRequest, resp *tmallcar.TmallAliautoEticketStoreGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
