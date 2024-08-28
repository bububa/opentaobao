package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBaikeImportZhubaoData 导入数据到商品百科服务
// taobao.baike.import.zhubao.data
//
// 用于接入外部数据录入到商品百科中
func TaobaoBaikeImportZhubaoData(ctx context.Context, clt *core.SDKClient, req *product.TaobaoBaikeImportZhubaoDataAPIRequest, resp *product.TaobaoBaikeImportZhubaoDataAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
