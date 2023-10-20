package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoeticketstoreget 查询电子凭证对应门店信息
// tmall.aliauto.eticket.store.get
//
// 查询电子凭证对应门店信息
func Tmallaliautoeticketstoreget(clt *core.SDKClient, req *tmallcar.TmallaliautoeticketstoregetAPIRequest, session string) (*tmallcar.TmallaliautoeticketstoregetAPIResponse, error) {
	var resp tmallcar.TmallaliautoeticketstoregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
