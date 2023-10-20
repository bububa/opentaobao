package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// Taobaodegoperationshowuserrecords 用户中奖记录
// taobao.degoperation.show.user.records
//
// 用户中奖记录
func Taobaodegoperationshowuserrecords(clt *core.SDKClient, req *degoperation.TaobaodegoperationshowuserrecordsAPIRequest, session string) (*degoperation.TaobaodegoperationshowuserrecordsAPIResponse, error) {
	var resp degoperation.TaobaodegoperationshowuserrecordsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
