package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateDelistingAPIRequest 商品下架 API请求
// taobao.item.update.delisting
//
// * 单个商品下架<br/>    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoItemUpdateDelistingRequest 初始化TaobaoItemUpdateDelistingAPIRequest对象
func NewTaobaoItemUpdateDelistingRequest() *TaobaoItemUpdateDelistingAPIRequest {
	return &TaobaoItemUpdateDelistingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateDelistingAPIRequest) GetApiMethodName() string {
	return "taobao.item.update.delisting"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateDelistingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemUpdateDelistingAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemUpdateDelistingAPIRequest) GetNumIid() int64 {
	return r._numIid
}
