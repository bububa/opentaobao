package media

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaoPictureCategoryAdd 新增图片分类信息
// taobao.picture.category.add
//
// 同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符
func TaobaoPictureCategoryAdd(ctx context.Context, clt *core.SDKClient, req *media.TaobaoPictureCategoryAddAPIRequest, resp *media.TaobaoPictureCategoryAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
