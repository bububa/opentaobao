package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribequit OPENIM群成员退出
// taobao.openim.tribe.quit
//
// OPENIM群成员退出
func Taobaoopenimtribequit(clt *core.SDKClient, req *openim.TaobaoopenimtribequitAPIRequest, session string) (*openim.TaobaoopenimtribequitAPIResponse, error) {
	var resp openim.TaobaoopenimtribequitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
