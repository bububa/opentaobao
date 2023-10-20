package openim

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openim"
)

// Taobaoopenimtribemodifytribeinfo OPENIM群信息修改
// taobao.openim.tribe.modifytribeinfo
//
// OPENIM群信息修改
func Taobaoopenimtribemodifytribeinfo(clt *core.SDKClient, req *openim.TaobaoopenimtribemodifytribeinfoAPIRequest, session string) (*openim.TaobaoopenimtribemodifytribeinfoAPIResponse, error) {
	var resp openim.TaobaoopenimtribemodifytribeinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
