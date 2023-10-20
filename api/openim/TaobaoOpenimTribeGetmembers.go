package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribegetmembers OPENIM群成员获取
// taobao.openim.tribe.getmembers
//
// OPENIM群成员获取
func Taobaoopenimtribegetmembers(clt *core.SDKClient, req *openim.TaobaoopenimtribegetmembersAPIRequest, session string) (*openim.TaobaoopenimtribegetmembersAPIResponse, error) {
	var resp openim.TaobaoopenimtribegetmembersAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
