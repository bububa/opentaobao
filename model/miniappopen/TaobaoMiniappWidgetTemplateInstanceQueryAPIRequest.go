package miniappopen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest 小部件实例化版本查询 API请求
// taobao.miniapp.widget.template.instance.query
//
// 小部件实例化版本查询
type TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest struct {
	model.Params
	// 入参
	_param0 *MiniAppInstantiateAppOpenQuery
}

// NewTaobaoMiniappWidgetTemplateInstanceQueryRequest 初始化TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest对象
func NewTaobaoMiniappWidgetTemplateInstanceQueryRequest() *TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest {
	return &TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.widget.template.instance.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 入参
func (r *TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest) SetParam0(_param0 *MiniAppInstantiateAppOpenQuery) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest) GetParam0() *MiniAppInstantiateAppOpenQuery {
	return r._param0
}

var poolTaobaoMiniappWidgetTemplateInstanceQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMiniappWidgetTemplateInstanceQueryRequest()
	},
}

// GetTaobaoMiniappWidgetTemplateInstanceQueryRequest 从 sync.Pool 获取 TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest
func GetTaobaoMiniappWidgetTemplateInstanceQueryAPIRequest() *TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest {
	return poolTaobaoMiniappWidgetTemplateInstanceQueryAPIRequest.Get().(*TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest)
}

// ReleaseTaobaoMiniappWidgetTemplateInstanceQueryAPIRequest 将 TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoMiniappWidgetTemplateInstanceQueryAPIRequest(v *TaobaoMiniappWidgetTemplateInstanceQueryAPIRequest) {
	v.Reset()
	poolTaobaoMiniappWidgetTemplateInstanceQueryAPIRequest.Put(v)
}
