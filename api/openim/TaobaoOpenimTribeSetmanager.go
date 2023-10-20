package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribesetmanager OPENIM群设置管理员
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
func Taobaoopenimtribesetmanager(clt *core.SDKClient, req *openim.TaobaoopenimtribesetmanagerAPIRequest, session string) (*openim.TaobaoopenimtribesetmanagerAPIResponse, error) {
	var resp openim.TaobaoopenimtribesetmanagerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
