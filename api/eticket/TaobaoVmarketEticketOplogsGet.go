package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

/* TaobaoVmarketEticketOplogsGet
电子凭证操作日志查询
taobao.vmarket.eticket.oplogs.get

电子凭证核销日志查询 */
func TaobaoVmarketEticketOplogsGet(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketOplogsGetAPIRequest, session string) (*eticket.TaobaoVmarketEticketOplogsGetAPIResponse, error) {
	var resp eticket.TaobaoVmarketEticketOplogsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
