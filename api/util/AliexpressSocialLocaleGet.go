package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Aliexpresssociallocaleget Locale获取接口
// aliexpress.social.locale.get
//
// 新增Locale获取接口
func Aliexpresssociallocaleget(clt *core.SDKClient, req *util.AliexpresssociallocalegetAPIRequest, session string) (*util.AliexpresssociallocalegetAPIResponse, error) {
	var resp util.AliexpresssociallocalegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
