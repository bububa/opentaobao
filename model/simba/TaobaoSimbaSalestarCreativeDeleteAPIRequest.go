package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarCreativeDeleteAPIRequest (新)销量明星删除创意相关接口 API请求
// taobao.simba.salestar.creative.delete
//
// 删除一个创意
type TaobaoSimbaSalestarCreativeDeleteAPIRequest struct {
	model.Params
	// 创意Id
	_creativeId int64
}

// NewTaobaoSimbaSalestarCreativeDeleteRequest 初始化TaobaoSimbaSalestarCreativeDeleteAPIRequest对象
func NewTaobaoSimbaSalestarCreativeDeleteRequest() *TaobaoSimbaSalestarCreativeDeleteAPIRequest {
	return &TaobaoSimbaSalestarCreativeDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarCreativeDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.creative.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarCreativeDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSalestarCreativeDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeId is CreativeId Setter
// 创意Id
func (r *TaobaoSimbaSalestarCreativeDeleteAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaoSimbaSalestarCreativeDeleteAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}
