package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoitemupdatedelistingAPIRequest 商品下架 API请求
// taobao.item.update.delisting
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
type TaobaoitemupdatedelistingAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoitemupdatedelistingRequest 初始化TaobaoitemupdatedelistingAPIRequest对象
func NewTaobaoitemupdatedelistingRequest() *TaobaoitemupdatedelistingAPIRequest {
	return &TaobaoitemupdatedelistingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoitemupdatedelistingAPIRequest) GetApiMethodName() string {
	return "taobao.item.update.delisting"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoitemupdatedelistingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoitemupdatedelistingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoitemupdatedelistingAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoitemupdatedelistingAPIRequest) GetNumIid() int64 {
	return r._numIid
}
