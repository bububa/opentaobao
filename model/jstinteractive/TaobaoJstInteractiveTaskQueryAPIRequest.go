package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveTaskQueryAPIRequest
互动任务列表查询接口 API请求
taobao.jst.interactive.task.query

查询用户的互动任务列表 */
type TaobaoJstInteractiveTaskQueryAPIRequest struct {
	model.Params
}

// NewTaobaoJstInteractiveTaskQueryRequest 初始化TaobaoJstInteractiveTaskQueryAPIRequest对象
func NewTaobaoJstInteractiveTaskQueryRequest() *TaobaoJstInteractiveTaskQueryAPIRequest {
	return &TaobaoJstInteractiveTaskQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstInteractiveTaskQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.interactive.task.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstInteractiveTaskQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
