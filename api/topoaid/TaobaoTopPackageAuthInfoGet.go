package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// Taobaotoppackageauthinfoget 淘宝末端包裹信息获取
// taobao.top.package.auth.info.get
//
// 淘宝末端包裹信息获取
func Taobaotoppackageauthinfoget(clt *core.SDKClient, req *topoaid.TaobaotoppackageauthinfogetAPIRequest, session string) (*topoaid.TaobaotoppackageauthinfogetAPIResponse, error) {
	var resp topoaid.TaobaotoppackageauthinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
