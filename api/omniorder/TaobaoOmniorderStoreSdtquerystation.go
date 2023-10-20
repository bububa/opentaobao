package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderStoreSdtquerystation 速店通查询站点信息
// taobao.omniorder.store.sdtquerystation
//
// 速店通查询站点信息
func TaobaoOmniorderStoreSdtquerystation(clt *core.SDKClient, req *omniorder.TaobaoOmniorderStoreSdtquerystationAPIRequest, resp *omniorder.TaobaoOmniorderStoreSdtquerystationAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
