package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest 免费儿童列表查询 API请求
// taobao.train.agent.freechildrenlist.query.vtwo
//
// 免费儿童列表查询
type TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest struct {
	model.Params
}

// NewTaobaotrainagentfreechildrenlistqueryvtwoRequest 初始化TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest对象
func NewTaobaotrainagentfreechildrenlistqueryvtwoRequest() *TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest {
	return &TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.freechildrenlist.query.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentfreechildrenlistqueryvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}
