package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribecreate 创建群
// taobao.openim.tribe.create
//
// 创建一个openim的群
func Taobaoopenimtribecreate(clt *core.SDKClient, req *openim.TaobaoopenimtribecreateAPIRequest, session string) (*openim.TaobaoopenimtribecreateAPIResponse, error) {
	var resp openim.TaobaoopenimtribecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
