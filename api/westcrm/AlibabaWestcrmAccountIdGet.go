package westcrm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/westcrm"
)

/* AlibabaWestcrmAccountIdGet
根据支付宝id查询IB的id
alibaba.westcrm.account.id.get

根据支付宝id查询IB的id */
func AlibabaWestcrmAccountIdGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmAccountIdGetAPIRequest, session string) (*westcrm.AlibabaWestcrmAccountIdGetAPIResponse, error) {
	var resp westcrm.AlibabaWestcrmAccountIdGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
