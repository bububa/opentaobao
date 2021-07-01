package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpMbbGetbyidAPIRequest
获取营销积木块 API请求
taobao.ump.mbb.getbyid

根据积木块id获取营销积木块。 */
type TaobaoUmpMbbGetbyidAPIRequest struct {
	model.Params
	// 积木块的id
	_id int64
}

// NewTaobaoUmpMbbGetbyidRequest 初始化TaobaoUmpMbbGetbyidAPIRequest对象
func NewTaobaoUmpMbbGetbyidRequest() *TaobaoUmpMbbGetbyidAPIRequest {
	return &TaobaoUmpMbbGetbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbGetbyidAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbb.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbGetbyidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Id Setter
// 积木块的id
func (r *TaobaoUmpMbbGetbyidAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// Get Id Getter
func (r TaobaoUmpMbbGetbyidAPIRequest) GetId() int64 {
	return r._id
}
