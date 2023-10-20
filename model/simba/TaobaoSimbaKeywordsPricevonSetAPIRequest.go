package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsPricevonSetAPIRequest 设置一批关键词的信息 API请求
// taobao.simba.keywords.pricevon.set
//
// 设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
type TaobaoSimbaKeywordsPricevonSetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词ID，出价和匹配方式json字符串，keywordId:词ID，整数。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,也不能大于99.99元,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0。matchscope只能是1,4（1代表精确匹配，4代表广泛匹配），maxMobilePrice：代表无线出价（如果是0，则代表无线出价=PC*无线溢价），mobileIsDefaultPrice代表无线出价是否采用pc*无线溢价，只能传0
	_keywordidPrices string
}

// NewTaobaoSimbaKeywordsPricevonSetRequest 初始化TaobaoSimbaKeywordsPricevonSetAPIRequest对象
func NewTaobaoSimbaKeywordsPricevonSetRequest() *TaobaoSimbaKeywordsPricevonSetAPIRequest {
	return &TaobaoSimbaKeywordsPricevonSetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordsPricevonSetAPIRequest) Reset() {
	r._nick = ""
	r._keywordidPrices = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsPricevonSetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.pricevon.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsPricevonSetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordsPricevonSetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsPricevonSetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordsPricevonSetAPIRequest) GetNick() string {
	return r._nick
}

// SetKeywordidPrices is KeywordidPrices Setter
// 关键词ID，出价和匹配方式json字符串，keywordId:词ID，整数。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,也不能大于99.99元,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0。matchscope只能是1,4（1代表精确匹配，4代表广泛匹配），maxMobilePrice：代表无线出价（如果是0，则代表无线出价=PC*无线溢价），mobileIsDefaultPrice代表无线出价是否采用pc*无线溢价，只能传0
func (r *TaobaoSimbaKeywordsPricevonSetAPIRequest) SetKeywordidPrices(_keywordidPrices string) error {
	r._keywordidPrices = _keywordidPrices
	r.Set("keywordid_prices", _keywordidPrices)
	return nil
}

// GetKeywordidPrices KeywordidPrices Getter
func (r TaobaoSimbaKeywordsPricevonSetAPIRequest) GetKeywordidPrices() string {
	return r._keywordidPrices
}

var poolTaobaoSimbaKeywordsPricevonSetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordsPricevonSetRequest()
	},
}

// GetTaobaoSimbaKeywordsPricevonSetRequest 从 sync.Pool 获取 TaobaoSimbaKeywordsPricevonSetAPIRequest
func GetTaobaoSimbaKeywordsPricevonSetAPIRequest() *TaobaoSimbaKeywordsPricevonSetAPIRequest {
	return poolTaobaoSimbaKeywordsPricevonSetAPIRequest.Get().(*TaobaoSimbaKeywordsPricevonSetAPIRequest)
}

// ReleaseTaobaoSimbaKeywordsPricevonSetAPIRequest 将 TaobaoSimbaKeywordsPricevonSetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordsPricevonSetAPIRequest(v *TaobaoSimbaKeywordsPricevonSetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordsPricevonSetAPIRequest.Put(v)
}
