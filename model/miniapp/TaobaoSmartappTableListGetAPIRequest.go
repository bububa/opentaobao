package miniapp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosmartapptablelistgetAPIRequest 智能应用服务登记工作表列表查询 API请求
// taobao.smartapp.table.list.get
//
// 智能应用服务登记工作表列表查询
type TaobaosmartapptablelistgetAPIRequest struct {
	model.Params
}

// NewTaobaosmartapptablelistgetRequest 初始化TaobaosmartapptablelistgetAPIRequest对象
func NewTaobaosmartapptablelistgetRequest() *TaobaosmartapptablelistgetAPIRequest {
	return &TaobaosmartapptablelistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosmartapptablelistgetAPIRequest) GetApiMethodName() string {
	return "taobao.smartapp.table.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosmartapptablelistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosmartapptablelistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}
