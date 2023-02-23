package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateDelistingTmallAPIRequest taobao.item.update.delisting.tmall API请求
// taobao.item.update.delisting.tmall
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingTmallAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoItemUpdateDelistingTmallRequest 初始化TaobaoItemUpdateDelistingTmallAPIRequest对象
func NewTaobaoItemUpdateDelistingTmallRequest() *TaobaoItemUpdateDelistingTmallAPIRequest {
	return &TaobaoItemUpdateDelistingTmallAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateDelistingTmallAPIRequest) GetApiMethodName() string {
	return "taobao.item.update.delisting.tmall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateDelistingTmallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemUpdateDelistingTmallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemUpdateDelistingTmallAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemUpdateDelistingTmallAPIRequest) GetNumIid() int64 {
	return r._numIid
}
