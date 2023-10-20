package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemdeleteAPIRequest 删除单条商品 API请求
// taobao.item.delete
//
// 删除单条商品
type TaobaoitemdeleteAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoitemdeleteRequest 初始化TaobaoitemdeleteAPIRequest对象
func NewTaobaoitemdeleteRequest() *TaobaoitemdeleteAPIRequest {
	return &TaobaoitemdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoitemdeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemdeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}
