package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemUpdateListingAPIRequest 一口价商品上架 API请求
// taobao.item.update.listing
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
type TaobaoItemUpdateListingAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
	// 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
	_num int64
}

// NewTaobaoItemUpdateListingRequest 初始化TaobaoItemUpdateListingAPIRequest对象
func NewTaobaoItemUpdateListingRequest() *TaobaoItemUpdateListingAPIRequest {
	return &TaobaoItemUpdateListingAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemUpdateListingAPIRequest) Reset() {
	r._numIid = 0
	r._num = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemUpdateListingAPIRequest) GetApiMethodName() string {
	return "taobao.item.update.listing"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemUpdateListingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemUpdateListingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemUpdateListingAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemUpdateListingAPIRequest) GetNumIid() int64 {
	return r._numIid
}

// SetNum is Num Setter
// 需要上架的商品的数量。取值范围:大于零的整数。如果商品有sku，则上架数量默认为所有sku数量总和，不可修改。否则商品数量根据设置数量调整为num
func (r *TaobaoItemUpdateListingAPIRequest) SetNum(_num int64) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// GetNum Num Getter
func (r TaobaoItemUpdateListingAPIRequest) GetNum() int64 {
	return r._num
}

var poolTaobaoItemUpdateListingAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemUpdateListingRequest()
	},
}

// GetTaobaoItemUpdateListingRequest 从 sync.Pool 获取 TaobaoItemUpdateListingAPIRequest
func GetTaobaoItemUpdateListingAPIRequest() *TaobaoItemUpdateListingAPIRequest {
	return poolTaobaoItemUpdateListingAPIRequest.Get().(*TaobaoItemUpdateListingAPIRequest)
}

// ReleaseTaobaoItemUpdateListingAPIRequest 将 TaobaoItemUpdateListingAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemUpdateListingAPIRequest(v *TaobaoItemUpdateListingAPIRequest) {
	v.Reset()
	poolTaobaoItemUpdateListingAPIRequest.Put(v)
}
