package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemTemplatesGetAPIRequest 获取用户宝贝详情页模板名称 API请求
// taobao.item.templates.get
//
// 查询当前登录用户的店铺的宝贝详情页的模板名称
type TaobaoItemTemplatesGetAPIRequest struct {
	model.Params
}

// NewTaobaoItemTemplatesGetRequest 初始化TaobaoItemTemplatesGetAPIRequest对象
func NewTaobaoItemTemplatesGetRequest() *TaobaoItemTemplatesGetAPIRequest {
	return &TaobaoItemTemplatesGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemTemplatesGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemTemplatesGetAPIRequest) GetApiMethodName() string {
	return "taobao.item.templates.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemTemplatesGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemTemplatesGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoItemTemplatesGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemTemplatesGetRequest()
	},
}

// GetTaobaoItemTemplatesGetRequest 从 sync.Pool 获取 TaobaoItemTemplatesGetAPIRequest
func GetTaobaoItemTemplatesGetAPIRequest() *TaobaoItemTemplatesGetAPIRequest {
	return poolTaobaoItemTemplatesGetAPIRequest.Get().(*TaobaoItemTemplatesGetAPIRequest)
}

// ReleaseTaobaoItemTemplatesGetAPIRequest 将 TaobaoItemTemplatesGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemTemplatesGetAPIRequest(v *TaobaoItemTemplatesGetAPIRequest) {
	v.Reset()
	poolTaobaoItemTemplatesGetAPIRequest.Put(v)
}
