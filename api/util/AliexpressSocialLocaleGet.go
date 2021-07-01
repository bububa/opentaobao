package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

/* AliexpressSocialLocaleGet
Locale获取接口
aliexpress.social.locale.get

新增Locale获取接口 */
func AliexpressSocialLocaleGet(clt *core.SDKClient, req *util.AliexpressSocialLocaleGetAPIRequest, session string) (*util.AliexpressSocialLocaleGetAPIResponse, error) {
	var resp util.AliexpressSocialLocaleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
