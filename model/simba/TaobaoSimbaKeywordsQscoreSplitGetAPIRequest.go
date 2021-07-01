package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsQscoreSplitGetAPIRequest
新质量分服务 API请求
taobao.simba.keywords.qscore.split.get

获取关键词新的质量分 */
type TaobaoSimbaKeywordsQscoreSplitGetAPIRequest struct {
	model.Params
	// 账号昵称
	_nick string
	// 推广组id
	_adGroupId int64
	// 词id数组（最多批量获取20个）
	_bidwordIds []int64
}

// NewTaobaoSimbaKeywordsQscoreSplitGetRequest 初始化TaobaoSimbaKeywordsQscoreSplitGetAPIRequest对象
func NewTaobaoSimbaKeywordsQscoreSplitGetRequest() *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest {
	return &TaobaoSimbaKeywordsQscoreSplitGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.qscore.split.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 账号昵称
func (r *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetNick() string {
	return r._nick
}

// Set is AdGroupId Setter
// 推广组id
func (r *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) SetAdGroupId(_adGroupId int64) error {
	r._adGroupId = _adGroupId
	r.Set("ad_group_id", _adGroupId)
	return nil
}

// Get AdGroupId Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetAdGroupId() int64 {
	return r._adGroupId
}

// Set is BidwordIds Setter
// 词id数组（最多批量获取20个）
func (r *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) SetBidwordIds(_bidwordIds []int64) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// Get BidwordIds Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetBidwordIds() []int64 {
	return r._bidwordIds
}
