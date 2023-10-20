package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AliexpressSocialCountryGet 获取国家列表
// aliexpress.social.country.get
//
// 获取目前AE支持的国家列表
func AliexpressSocialCountryGet(clt *core.SDKClient, req *util.AliexpressSocialCountryGetAPIRequest, resp *util.AliexpressSocialCountryGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
