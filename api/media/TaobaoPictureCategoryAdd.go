package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

/* TaobaoPictureCategoryAdd
新增图片分类信息
taobao.picture.category.add

同一卖家最多添加500个图片分类，图片分类名称长度最大为20个字符 */
func TaobaoPictureCategoryAdd(clt *core.SDKClient, req *media.TaobaoPictureCategoryAddAPIRequest, session string) (*media.TaobaoPictureCategoryAddAPIResponse, error) {
	var resp media.TaobaoPictureCategoryAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
