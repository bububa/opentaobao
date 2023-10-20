package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsvonaddAPIRequest 创建一批关键词 API请求
// taobao.simba.keywordsvon.add
//
// 创建一批关键词
type TaobaosimbakeywordsvonaddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词、计算机出价、无线出价和匹配方式json字符串，word:词，不能有一些特殊字符。maxPrice：计算机出价，是整数，以“分”为单位，不能小于5，不能大于日限额, maxMobilePrice：代表无线出价，规则同maxPice 当matchscope只能是1,4（1代表精确匹配，4代表广泛匹配）。
	_keywordPrices string
	// 推广组id
	_adgroupId int64
}

// NewTaobaosimbakeywordsvonaddRequest 初始化TaobaosimbakeywordsvonaddAPIRequest对象
func NewTaobaosimbakeywordsvonaddRequest() *TaobaosimbakeywordsvonaddAPIRequest {
	return &TaobaosimbakeywordsvonaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordsvonaddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordsvon.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordsvonaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordsvonaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbakeywordsvonaddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordsvonaddAPIRequest) GetNick() string {
	return r._nick
}

// SetKeywordPrices is KeywordPrices Setter
// 关键词、计算机出价、无线出价和匹配方式json字符串，word:词，不能有一些特殊字符。maxPrice：计算机出价，是整数，以“分”为单位，不能小于5，不能大于日限额, maxMobilePrice：代表无线出价，规则同maxPice 当matchscope只能是1,4（1代表精确匹配，4代表广泛匹配）。
func (r *TaobaosimbakeywordsvonaddAPIRequest) SetKeywordPrices(_keywordPrices string) error {
	r._keywordPrices = _keywordPrices
	r.Set("keyword_prices", _keywordPrices)
	return nil
}

// GetKeywordPrices KeywordPrices Getter
func (r TaobaosimbakeywordsvonaddAPIRequest) GetKeywordPrices() string {
	return r._keywordPrices
}

// SetAdgroupId is AdgroupId Setter
// 推广组id
func (r *TaobaosimbakeywordsvonaddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbakeywordsvonaddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
