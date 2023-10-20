package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// Taobaovmarketeticketoplogsget 电子凭证操作日志查询
// taobao.vmarket.eticket.oplogs.get
//
// 电子凭证核销日志查询
func Taobaovmarketeticketoplogsget(clt *core.SDKClient, req *eticket.TaobaovmarketeticketoplogsgetAPIRequest, session string) (*eticket.TaobaovmarketeticketoplogsgetAPIResponse, error) {
	var resp eticket.TaobaovmarketeticketoplogsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
