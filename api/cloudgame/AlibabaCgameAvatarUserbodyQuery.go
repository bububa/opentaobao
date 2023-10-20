package cloudgame

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cloudgame"
)

// Alibabacgameavataruserbodyquery 用户Avatar body查询
// alibaba.cgame.avatar.userbody.query
//
// Avatar用户body数据查询
func Alibabacgameavataruserbodyquery(clt *core.SDKClient, req *cloudgame.AlibabacgameavataruserbodyqueryAPIRequest, session string) (*cloudgame.AlibabacgameavataruserbodyqueryAPIResponse, error) {
	var resp cloudgame.AlibabacgameavataruserbodyqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
