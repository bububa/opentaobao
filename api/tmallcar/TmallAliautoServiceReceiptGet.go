package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallaliautoservicereceiptget isv查询服务工单详情
// tmall.aliauto.service.receipt.get
//
// isv查询服务工单详情
func Tmallaliautoservicereceiptget(clt *core.SDKClient, req *tmallcar.TmallaliautoservicereceiptgetAPIRequest, session string) (*tmallcar.TmallaliautoservicereceiptgetAPIResponse, error) {
	var resp tmallcar.TmallaliautoservicereceiptgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
