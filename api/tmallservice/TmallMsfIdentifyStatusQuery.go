package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallmsfidentifystatusquery 喵师傅定案核销状态查询
// tmall.msf.identify.status.query
//
// 喵师傅定案核销状态查询，供服务商erp系统调用
func Tmallmsfidentifystatusquery(clt *core.SDKClient, req *tmallservice.TmallmsfidentifystatusqueryAPIRequest, session string) (*tmallservice.TmallmsfidentifystatusqueryAPIResponse, error) {
	var resp tmallservice.TmallmsfidentifystatusqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
