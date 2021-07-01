package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

/* AliexpressSocialCurrencyGet
币种获取接口
aliexpress.social.currency.get

获取目前AE社交支持的币种 */
func AliexpressSocialCurrencyGet(clt *core.SDKClient, req *util.AliexpressSocialCurrencyGetAPIRequest, session string) (*util.AliexpressSocialCurrencyGetAPIResponse, error) {
	var resp util.AliexpressSocialCurrencyGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
