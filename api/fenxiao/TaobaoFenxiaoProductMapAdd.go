package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductMapAdd 创建分销和后端商品映射关系
// taobao.fenxiao.product.map.add
//
// 创建分销和供应链商品映射关系。
func TaobaoFenxiaoProductMapAdd(clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductMapAddAPIRequest, resp *fenxiao.TaobaoFenxiaoProductMapAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
