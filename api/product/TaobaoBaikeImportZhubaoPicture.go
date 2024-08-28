package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoBaikeImportZhubaoPicture 百科图片数据导入
// taobao.baike.import.zhubao.picture
//
// 用于接入外部--图片--录入到商品百科中
func TaobaoBaikeImportZhubaoPicture(ctx context.Context, clt *core.SDKClient, req *product.TaobaoBaikeImportZhubaoPictureAPIRequest, resp *product.TaobaoBaikeImportZhubaoPictureAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
