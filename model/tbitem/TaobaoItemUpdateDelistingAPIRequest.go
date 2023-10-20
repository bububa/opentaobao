package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateDelistingAPIRequest 商品下架 API请求
// taobao.item.update.delisting
//
// * 单个商品下架&lt;br/&gt;    * 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateDelistingAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoItemUpdateDelistingRequest 初始化TaobaoItemUpdateDelistingAPIRequest对象
func NewTaobaoItemUpdateDelistingRequest() *TaobaoItemUpdateDelistingAPIRequest {
	return &TaobaoItemUpdateDelistingAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemUpdateDelistingAPIRequest) Reset() {
	r._numIid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateDelistingAPIRequest) GetApiMethodName() string {
	return "taobao.item.update.delisting"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateDelistingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemUpdateDelistingAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTaobaoItemUpdateDelistingAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemUpdateDelistingRequest()
	},
}

// GetTaobaoItemUpdateDelistingRequest 从 sync.Pool 获取 TaobaoItemUpdateDelistingAPIRequest
func GetTaobaoItemUpdateDelistingAPIRequest() *TaobaoItemUpdateDelistingAPIRequest {
	return poolTaobaoItemUpdateDelistingAPIRequest.Get().(*TaobaoItemUpdateDelistingAPIRequest)
}

// ReleaseTaobaoItemUpdateDelistingAPIRequest 将 TaobaoItemUpdateDelistingAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemUpdateDelistingAPIRequest(v *TaobaoItemUpdateDelistingAPIRequest) {
	v.Reset()
	poolTaobaoItemUpdateDelistingAPIRequest.Put(v)
}
