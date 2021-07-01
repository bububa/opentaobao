package ascpffo

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ascpffo"
)

/* AliexpressAscpFfoItemQuery
AliExpress发货单明细分页查询API
aliexpress.ascp.ffo.item.query

AE履约发货单明细分页查询 */
func AliexpressAscpFfoItemQuery(clt *core.SDKClient, req *ascpffo.AliexpressAscpFfoItemQueryAPIRequest, session string) (*ascpffo.AliexpressAscpFfoItemQueryAPIResponse, error) {
	var resp ascpffo.AliexpressAscpFfoItemQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
