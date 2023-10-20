package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentreturnticketinfogetvtwo 代理商获取退票详情回调
// taobao.train.agent.returnticketinfo.get.vtwo
//
// 代理商获取退票详情回调
func Taobaotrainagentreturnticketinfogetvtwo(clt *core.SDKClient, req *train.TaobaotrainagentreturnticketinfogetvtwoAPIRequest, session string) (*train.TaobaotrainagentreturnticketinfogetvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentreturnticketinfogetvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
