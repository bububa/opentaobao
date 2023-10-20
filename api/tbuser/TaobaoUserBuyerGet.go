package tbuser

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbuser"
)

// Taobaouserbuyerget 查询买家信息API
// taobao.user.buyer.get
//
// 查询买家信息API，只能买家类应用调用。
func Taobaouserbuyerget(clt *core.SDKClient, req *tbuser.TaobaouserbuyergetAPIRequest, session string) (*tbuser.TaobaouserbuyergetAPIResponse, error) {
	var resp tbuser.TaobaouserbuyergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
