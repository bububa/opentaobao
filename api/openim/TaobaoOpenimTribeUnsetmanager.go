package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribeunsetmanager OPENIM群取消管理员
// taobao.openim.tribe.unsetmanager
//
// OPENIM群取消管理员
func Taobaoopenimtribeunsetmanager(clt *core.SDKClient, req *openim.TaobaoopenimtribeunsetmanagerAPIRequest, session string) (*openim.TaobaoopenimtribeunsetmanagerAPIResponse, error) {
	var resp openim.TaobaoopenimtribeunsetmanagerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
