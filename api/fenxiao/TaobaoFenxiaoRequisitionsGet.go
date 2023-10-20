package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoRequisitionsGet 合作申请查询
// taobao.fenxiao.requisitions.get
//
// 合作申请查询
func TaobaoFenxiaoRequisitionsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoRequisitionsGetAPIRequest, resp *fenxiao.TaobaoFenxiaoRequisitionsGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
