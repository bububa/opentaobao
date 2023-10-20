package tanx

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tanx"
)

// Taobaotanxqualificationfind 资质查询接口
// taobao.tanx.qualification.find
//
// 资质查询接口
func Taobaotanxqualificationfind(clt *core.SDKClient, req *tanx.TaobaotanxqualificationfindAPIRequest, session string) (*tanx.TaobaotanxqualificationfindAPIResponse, error) {
	var resp tanx.TaobaotanxqualificationfindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
