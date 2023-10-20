package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// TmallAliautoEticketStoreGet 查询电子凭证对应门店信息
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
func TmallAliautoEticketStoreGet(clt *core.SDKClient, req *tmallcar.TmallAliautoEticketStoreGetAPIRequest, session string) (*tmallcar.TmallAliautoEticketStoreGetAPIResponse, error) {
	var resp tmallcar.TmallAliautoEticketStoreGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
