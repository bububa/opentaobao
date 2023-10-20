package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// Alitriplocalplayproductupload 当地玩乐 产品维护接口
// alitrip.localplay.product.upload
//
// 当地玩乐（境内当地玩乐/境外玩乐套餐） 产品维护接口。
// 接口同时支持新商品发布 和 现有商品编辑：
// 1）只上传out_product_id的情况：如果out_product_id已经关联过某个商品id，则认为是编辑现有商品，否则认为是新发布一个商品。
// 2）同时上传out_product_id和item_id，则认为是将out_product_id与item_id进行关联，同时对该商品进行编辑。
func Alitriplocalplayproductupload(clt *core.SDKClient, req *travel.AlitriplocalplayproductuploadAPIRequest, session string) (*travel.AlitriplocalplayproductuploadAPIResponse, error) {
	var resp travel.AlitriplocalplayproductuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
