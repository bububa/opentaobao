package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

/* TaobaoTrainAgentBookordersGet
代理商获取待出票订单列表
taobao.train.agent.bookorders.get

代理商获取待出票订单列表，只返回订单号 */
func TaobaoTrainAgentBookordersGet(clt *core.SDKClient, req *train.TaobaoTrainAgentBookordersGetAPIRequest, session string) (*train.TaobaoTrainAgentBookordersGetAPIResponse, error) {
	var resp train.TaobaoTrainAgentBookordersGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
