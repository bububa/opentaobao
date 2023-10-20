package train

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/train"
)

// Taobaotrainagentfreechildrendealconfirmvtwo 免费儿童处理
// taobao.train.agent.freechildrendeal.confirm.vtwo
//
// 免费儿童列表查询
func Taobaotrainagentfreechildrendealconfirmvtwo(clt *core.SDKClient, req *train.TaobaotrainagentfreechildrendealconfirmvtwoAPIRequest, session string) (*train.TaobaotrainagentfreechildrendealconfirmvtwoAPIResponse, error) {
	var resp train.TaobaotrainagentfreechildrendealconfirmvtwoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
