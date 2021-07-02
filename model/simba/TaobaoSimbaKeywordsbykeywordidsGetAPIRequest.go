package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsbykeywordidsGetAPIRequest 根据一个关键词Id列表取得一组关键词 API请求
// taobao.simba.keywordsbykeywordids.get
//
// 根据一个关键词Id列表取得一组关键词
type TaobaoSimbaKeywordsbykeywordidsGetAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 关键词Id数组，最多200个；
	_keywordIds []int64
}

// NewTaobaoSimbaKeywordsbykeywordidsGetRequest 初始化TaobaoSimbaKeywordsbykeywordidsGetAPIRequest对象
func NewTaobaoSimbaKeywordsbykeywordidsGetRequest() *TaobaoSimbaKeywordsbykeywordidsGetAPIRequest {
	return &TaobaoSimbaKeywordsbykeywordidsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywordsbykeywordids.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is KeywordIds Setter
// 关键词Id数组，最多200个；
func (r *TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) SetKeywordIds(_keywordIds []int64) error {
	r._keywordIds = _keywordIds
	r.Set("keyword_ids", _keywordIds)
	return nil
}

// Get KeywordIds Getter
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetKeywordIds() []int64 {
	return r._keywordIds
}
