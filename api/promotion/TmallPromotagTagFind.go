package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// Tmallpromotagtagfind 查询标签接口
// tmall.promotag.tag.find
//
// 查询用户创建的所有标签
func Tmallpromotagtagfind(clt *core.SDKClient, req *promotion.TmallpromotagtagfindAPIRequest, session string) (*promotion.TmallpromotagtagfindAPIResponse, error) {
	var resp promotion.TmallpromotagtagfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
