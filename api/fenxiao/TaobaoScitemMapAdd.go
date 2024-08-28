package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemMapAdd 创建IC商品与后端商品的映射关系
// taobao.scitem.map.add
//
// 创建IC商品或分销商品与后端商品的映射关系
func TaobaoScitemMapAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoScitemMapAddAPIRequest, resp *fenxiao.TaobaoScitemMapAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
