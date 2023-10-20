package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Aliexpresssocialcurrencyget 币种获取接口
// aliexpress.social.currency.get
//
// 获取目前AE社交支持的币种
func Aliexpresssocialcurrencyget(clt *core.SDKClient, req *util.AliexpresssocialcurrencygetAPIRequest, session string) (*util.AliexpresssocialcurrencygetAPIResponse, error) {
	var resp util.AliexpresssocialcurrencygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
