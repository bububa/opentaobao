package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpmbbgetbyidAPIRequest 获取营销积木块 API请求
// taobao.ump.mbb.getbyid
//
// 根据积木块id获取营销积木块。
type TaobaoumpmbbgetbyidAPIRequest struct {
	model.Params
	// 积木块的id
	_id int64
}

// NewTaobaoumpmbbgetbyidRequest 初始化TaobaoumpmbbgetbyidAPIRequest对象
func NewTaobaoumpmbbgetbyidRequest() *TaobaoumpmbbgetbyidAPIRequest {
	return &TaobaoumpmbbgetbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpmbbgetbyidAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbb.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpmbbgetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpmbbgetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 积木块的id
func (r *TaobaoumpmbbgetbyidAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoumpmbbgetbyidAPIRequest) GetId() int64 {
	return r._id
}
