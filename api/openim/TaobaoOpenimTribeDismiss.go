package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribedismiss OPENIM群解散
// taobao.openim.tribe.dismiss
//
// OPENIM群解散
func Taobaoopenimtribedismiss(clt *core.SDKClient, req *openim.TaobaoopenimtribedismissAPIRequest, session string) (*openim.TaobaoopenimtribedismissAPIResponse, error) {
	var resp openim.TaobaoopenimtribedismissAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
