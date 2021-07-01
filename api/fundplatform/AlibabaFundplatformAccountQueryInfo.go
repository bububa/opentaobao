package fundplatform

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fundplatform"
)

/* AlibabaFundplatformAccountQueryInfo
查询账户信息
alibaba.fundplatform.account.query.info

外部查询资金平台用户账户信息 */
func AlibabaFundplatformAccountQueryInfo(clt *core.SDKClient, req *fundplatform.AlibabaFundplatformAccountQueryInfoAPIRequest, session string) (*fundplatform.AlibabaFundplatformAccountQueryInfoAPIResponse, error) {
	var resp fundplatform.AlibabaFundplatformAccountQueryInfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
