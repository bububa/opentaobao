package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// AlibabaRetailShorturlGet 短链接获取
// alibaba.retail.shorturl.get
//
// 短链接获取
func AlibabaRetailShorturlGet(clt *core.SDKClient, req *util.AlibabaRetailShorturlGetAPIRequest, resp *util.AlibabaRetailShorturlGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
