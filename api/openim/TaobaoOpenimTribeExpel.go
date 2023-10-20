package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribeexpel OPENIM群踢出成员
// taobao.openim.tribe.expel
//
// OPENIM群踢出成员
func Taobaoopenimtribeexpel(clt *core.SDKClient, req *openim.TaobaoopenimtribeexpelAPIRequest, session string) (*openim.TaobaoopenimtribeexpelAPIResponse, error) {
	var resp openim.TaobaoopenimtribeexpelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
