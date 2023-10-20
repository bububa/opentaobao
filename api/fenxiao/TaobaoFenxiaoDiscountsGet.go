package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoDiscountsGet 获取折扣信息
// taobao.fenxiao.discounts.get
//
// 查询折扣信息
func TaobaoFenxiaoDiscountsGet(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoDiscountsGetAPIRequest, session string) (*fenxiao.TaobaoFenxiaoDiscountsGetAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoDiscountsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
