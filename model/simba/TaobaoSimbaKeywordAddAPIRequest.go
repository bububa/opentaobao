package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordAddAPIRequest （新）关键词新增接口 API请求
// taobao.simba.keyword.add
//
// （新）关键词更新相关接口
type TaobaoSimbaKeywordAddAPIRequest struct {
	model.Params
	// 关键词相关信息
	_bidwords []SiriusBidwordDto
	// 推广单元id
	_adgroupId int64
}

// NewTaobaoSimbaKeywordAddRequest 初始化TaobaoSimbaKeywordAddAPIRequest对象
func NewTaobaoSimbaKeywordAddRequest() *TaobaoSimbaKeywordAddAPIRequest {
	return &TaobaoSimbaKeywordAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordAddAPIRequest) Reset() {
	r._bidwords = r._bidwords[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keyword.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwords is Bidwords Setter
// 关键词相关信息
func (r *TaobaoSimbaKeywordAddAPIRequest) SetBidwords(_bidwords []SiriusBidwordDto) error {
	r._bidwords = _bidwords
	r.Set("bidwords", _bidwords)
	return nil
}

// GetBidwords Bidwords Getter
func (r TaobaoSimbaKeywordAddAPIRequest) GetBidwords() []SiriusBidwordDto {
	return r._bidwords
}

// SetAdgroupId is AdgroupId Setter
// 推广单元id
func (r *TaobaoSimbaKeywordAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaKeywordAddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaKeywordAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordAddRequest()
	},
}

// GetTaobaoSimbaKeywordAddRequest 从 sync.Pool 获取 TaobaoSimbaKeywordAddAPIRequest
func GetTaobaoSimbaKeywordAddAPIRequest() *TaobaoSimbaKeywordAddAPIRequest {
	return poolTaobaoSimbaKeywordAddAPIRequest.Get().(*TaobaoSimbaKeywordAddAPIRequest)
}

// ReleaseTaobaoSimbaKeywordAddAPIRequest 将 TaobaoSimbaKeywordAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordAddAPIRequest(v *TaobaoSimbaKeywordAddAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordAddAPIRequest.Put(v)
}
