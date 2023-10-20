package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniitemcategoryget 全渠道商品轻发布类目信息
// taobao.omniitem.category.get
//
// 全渠道商品轻发布类目信息
func Taobaoomniitemcategoryget(clt *core.SDKClient, req *omniorder.TaobaoomniitemcategorygetAPIRequest, session string) (*omniorder.TaobaoomniitemcategorygetAPIResponse, error) {
	var resp omniorder.TaobaoomniitemcategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
