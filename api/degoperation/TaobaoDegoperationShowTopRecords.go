package degoperation

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/degoperation"
)

// Taobaodegoperationshowtoprecords 活动中奖记录
// taobao.degoperation.show.top.records
//
// 活动中奖记录
func Taobaodegoperationshowtoprecords(clt *core.SDKClient, req *degoperation.TaobaodegoperationshowtoprecordsAPIRequest, session string) (*degoperation.TaobaodegoperationshowtoprecordsAPIResponse, error) {
	var resp degoperation.TaobaodegoperationshowtoprecordsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
