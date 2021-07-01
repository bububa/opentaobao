package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

/* AlibabaRetailShorturlGet
短链接获取
alibaba.retail.shorturl.get

短链接获取 */
func AlibabaRetailShorturlGet(clt *core.SDKClient, req *util.AlibabaRetailShorturlGetAPIRequest, session string) (*util.AlibabaRetailShorturlGetAPIResponse, error) {
	var resp util.AlibabaRetailShorturlGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
