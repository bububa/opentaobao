package taotv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/taotv"
)

// Taobaotaotvcarouselcategorylist 获取轮播分类列表
// taobao.taotv.carousel.category.list
//
// 获取轮播分类列表
func Taobaotaotvcarouselcategorylist(clt *core.SDKClient, req *taotv.TaobaotaotvcarouselcategorylistAPIRequest, session string) (*taotv.TaobaotaotvcarouselcategorylistAPIResponse, error) {
	var resp taotv.TaobaotaotvcarouselcategorylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
