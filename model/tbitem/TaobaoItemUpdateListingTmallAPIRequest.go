package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateListingTmallAPIRequest taobao.item.update.listing天猫分流 API请求
// taobao.item.update.listing.tmall
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateListingTmallAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
	// 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
	_num int64
}

// NewTaobaoItemUpdateListingTmallRequest 初始化TaobaoItemUpdateListingTmallAPIRequest对象
func NewTaobaoItemUpdateListingTmallRequest() *TaobaoItemUpdateListingTmallAPIRequest {
	return &TaobaoItemUpdateListingTmallAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemUpdateListingTmallAPIRequest) Reset() {
	r._numIid = 0
	r._num = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateListingTmallAPIRequest) GetApiMethodName() string {
	return "taobao.item.update.listing.tmall"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateListingTmallAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemUpdateListingTmallAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemUpdateListingTmallAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemUpdateListingTmallAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetNum is Num Setter
// 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
func (r *TaobaoItemUpdateListingTmallAPIRequest) SetNum(_num int64) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// GetNum Num Getter
func (r TaobaoItemUpdateListingTmallAPIRequest) GetNum() int64 {
	return r._num
}

var poolTaobaoItemUpdateListingTmallAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemUpdateListingTmallRequest()
	},
}

// GetTaobaoItemUpdateListingTmallRequest 从 sync.Pool 获取 TaobaoItemUpdateListingTmallAPIRequest
func GetTaobaoItemUpdateListingTmallAPIRequest() *TaobaoItemUpdateListingTmallAPIRequest {
	return poolTaobaoItemUpdateListingTmallAPIRequest.Get().(*TaobaoItemUpdateListingTmallAPIRequest)
}

// ReleaseTaobaoItemUpdateListingTmallAPIRequest 将 TaobaoItemUpdateListingTmallAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemUpdateListingTmallAPIRequest(v *TaobaoItemUpdateListingTmallAPIRequest) {
	v.Reset()
	poolTaobaoItemUpdateListingTmallAPIRequest.Put(v)
}
