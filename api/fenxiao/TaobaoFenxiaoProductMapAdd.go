package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

/* TaobaoFenxiaoProductMapAdd
创建分销和后端商品映射关系
taobao.fenxiao.product.map.add

创建分销和供应链商品映射关系。 */
func TaobaoFenxiaoProductMapAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductMapAddAPIRequest, session string) (*fenxiao.TaobaoFenxiaoProductMapAddAPIResponse, error) {
	var resp fenxiao.TaobaoFenxiaoProductMapAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
