package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordspricevonsetAPIRequest 设置一批关键词的信息 API请求
// taobao.simba.keywords.pricevon.set
//
// 设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
type TaobaosimbakeywordspricevonsetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词ID，出价和匹配方式json字符串，keywordId:词ID，整数。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,也不能大于99.99元,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0。matchscope只能是1,4（1代表精确匹配，4代表广泛匹配），maxMobilePrice：代表无线出价（如果是0，则代表无线出价=PC*无线溢价），mobileIsDefaultPrice代表无线出价是否采用pc*无线溢价，只能传0
	_keywordidPrices string
}

// NewTaobaosimbakeywordspricevonsetRequest 初始化TaobaosimbakeywordspricevonsetAPIRequest对象
func NewTaobaosimbakeywordspricevonsetRequest() *TaobaosimbakeywordspricevonsetAPIRequest {
	return &TaobaosimbakeywordspricevonsetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordspricevonsetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.pricevon.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordspricevonsetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordspricevonsetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordspricevonsetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordspricevonsetAPIRequest) GetNick() string {
	return r._nick
}

// SetKeywordidPrices is KeywordidPrices Setter
// 关键词ID，出价和匹配方式json字符串，keywordId:词ID，整数。maxPrice：价格，是整数，以“分”为单位，不能小于5，不能大于日限额,也不能大于99.99元,当使用默认出价时必须将这个值设置为0。; isDefaultPrice：是否使用默认出价，只能是0。matchscope只能是1,4（1代表精确匹配，4代表广泛匹配），maxMobilePrice：代表无线出价（如果是0，则代表无线出价=PC*无线溢价），mobileIsDefaultPrice代表无线出价是否采用pc*无线溢价，只能传0
func (r *TaobaosimbakeywordspricevonsetAPIRequest) SetKeywordidPrices(_keywordidPrices string) error {
	r._keywordidPrices = _keywordidPrices
	r.Set("keywordid_prices", _keywordidPrices)
	return nil
}

// GetKeywordidPrices KeywordidPrices Getter
func (r TaobaosimbakeywordspricevonsetAPIRequest) GetKeywordidPrices() string {
	return r._keywordidPrices
}
