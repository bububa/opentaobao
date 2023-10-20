package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Aliexpresssocialcountryget 获取国家列表
// aliexpress.social.country.get
//
// 获取目前AE支持的国家列表
func Aliexpresssocialcountryget(clt *core.SDKClient, req *util.AliexpresssocialcountrygetAPIRequest, session string) (*util.AliexpresssocialcountrygetAPIResponse, error) {
	var resp util.AliexpresssocialcountrygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
