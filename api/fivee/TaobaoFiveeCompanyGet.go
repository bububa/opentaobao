package fivee

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fivee"
)

// TaobaoFiveeCompanyGet 查询商信息
// taobao.fivee.company.get
//
// 资质共享平台查询商信息
func TaobaoFiveeCompanyGet(clt *core.SDKClient, req *fivee.TaobaoFiveeCompanyGetAPIRequest, resp *fivee.TaobaoFiveeCompanyGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
