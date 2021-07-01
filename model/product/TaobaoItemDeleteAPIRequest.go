package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoItemDeleteAPIRequest
删除单条商品 API请求
taobao.item.delete

删除单条商品 */
type TaobaoItemDeleteAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoItemDeleteRequest 初始化TaobaoItemDeleteAPIRequest对象
func NewTaobaoItemDeleteRequest() *TaobaoItemDeleteAPIRequest {
	return &TaobaoItemDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemDeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// Get NumIid Getter
func (r TaobaoItemDeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}
