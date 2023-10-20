package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribejoin OPENIM群主动加入
// taobao.openim.tribe.join
//
// OPENIM群主动加入
func Taobaoopenimtribejoin(clt *core.SDKClient, req *openim.TaobaoopenimtribejoinAPIRequest, session string) (*openim.TaobaoopenimtribejoinAPIResponse, error) {
	var resp openim.TaobaoopenimtribejoinAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
