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
	// 关键词Id数组，最多200个；
	_keywordIds []string
	// 主人昵称
	_nick string
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
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywordIds is KeywordIds Setter
// 关键词Id数组，最多200个；
func (r *TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) SetKeywordIds(_keywordIds []string) error {
	r._keywordIds = _keywordIds
	r.Set("keyword_ids", _keywordIds)
	return nil
}

// GetKeywordIds KeywordIds Getter
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetKeywordIds() []string {
	return r._keywordIds
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordsbykeywordidsGetAPIRequest) GetNick() string {
	return r._nick
}
