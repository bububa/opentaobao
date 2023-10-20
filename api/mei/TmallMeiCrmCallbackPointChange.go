package mei

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mei"
)

// TmallMeiCrmCallbackPointChange 品牌积分变更回调API
// tmall.mei.crm.callback.point.change
//
// 线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
func TmallMeiCrmCallbackPointChange(clt *core.SDKClient, req *mei.TmallMeiCrmCallbackPointChangeAPIRequest, session string) (*mei.TmallMeiCrmCallbackPointChangeAPIResponse, error) {
	var resp mei.TmallMeiCrmCallbackPointChangeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
