package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatUpdate 修改产品线
// taobao.fenxiao.productcat.update
//
// 修改产品线
func TaobaoFenxiaoProductcatUpdate(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoProductcatUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
