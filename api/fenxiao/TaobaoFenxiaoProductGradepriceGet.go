package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductGradepriceGet 等级折扣查询
// taobao.fenxiao.product.gradeprice.get
//
// 等级折扣查询
func TaobaoFenxiaoProductGradepriceGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductGradepriceGetAPIRequest, resp *fenxiao.TaobaoFenxiaoProductGradepriceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
