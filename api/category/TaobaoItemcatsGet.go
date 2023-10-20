package category

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/category"
)

// Taobaoitemcatsget 获取后台供卖家发布商品的标准商品类目
// taobao.itemcats.get
//
// 获取后台供卖家发布商品的标准商品类目。
func Taobaoitemcatsget(clt *core.SDKClient, req *category.TaobaoitemcatsgetAPIRequest, session string) (*category.TaobaoitemcatsgetAPIResponse, error) {
	var resp category.TaobaoitemcatsgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
