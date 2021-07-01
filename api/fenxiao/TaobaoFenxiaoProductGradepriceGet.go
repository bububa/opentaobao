package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoFenxiaoProductGradepriceGet
等级折扣查询
taobao.fenxiao.product.gradeprice.get

等级折扣查询 */
func TaobaoFenxiaoProductGradepriceGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductGradepriceGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductGradepriceGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductGradepriceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
