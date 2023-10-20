package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtbrandinfoquery 品牌数据查询
// tmall.nrt.brandinfo.query
//
// 商家获取自己旗舰店授权的品牌id列表
func Tmallnrtbrandinfoquery(clt *core.SDKClient, req *nrt.TmallnrtbrandinfoqueryAPIRequest, session string) (*nrt.TmallnrtbrandinfoqueryAPIResponse, error) {
	var resp nrt.TmallnrtbrandinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
