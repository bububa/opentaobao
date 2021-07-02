package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripGrouptoursProductUpload 跟团游 产品维护接口
// alitrip.grouptours.product.upload
//
// 跟团游 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
func AlitripGrouptoursProductUpload(clt *core.SDKClient, req *travel.AlitripGrouptoursProductUploadAPIRequest, session string) (*travel.AlitripGrouptoursProductUploadAPIResponse, error) {
	var resp travel.AlitripGrouptoursProductUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
