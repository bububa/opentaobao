package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpmbbslistgetAPIRequest 通过ids列表获取营销积木块列表 API请求
// taobao.ump.mbbs.list.get
//
// 通过营销积木id列表来获取营销积木块列表。
type TaobaoumpmbbslistgetAPIRequest struct {
	model.Params
	// 营销积木块id组成的字符串。
	_ids []int64
}

// NewTaobaoumpmbbslistgetRequest 初始化TaobaoumpmbbslistgetAPIRequest对象
func NewTaobaoumpmbbslistgetRequest() *TaobaoumpmbbslistgetAPIRequest {
	return &TaobaoumpmbbslistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpmbbslistgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbbs.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpmbbslistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpmbbslistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 营销积木块id组成的字符串。
func (r *TaobaoumpmbbslistgetAPIRequest) SetIds(_ids []int64) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoumpmbbslistgetAPIRequest) GetIds() []int64 {
	return r._ids
}
