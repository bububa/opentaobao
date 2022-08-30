package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest (新)销量明星质量分相关接口 API请求
// taobao.simba.salestar.keywords.qscore.split.get
//
// 获取关键词新的质量分
type TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest struct {
	model.Params
	// 词id数组（最多批量获取20个）
	_bidwordIds []string
	// 账号昵称
	_nick string
	// 推广组id
	_adGroupId int64
}

// NewTaobaoSimbaSalestarKeywordsQscoreSplitGetRequest 初始化TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest对象
func NewTaobaoSimbaSalestarKeywordsQscoreSplitGetRequest() *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest {
	return &TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.keywords.qscore.split.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBidwordIds is BidwordIds Setter
// 词id数组（最多批量获取20个）
func (r *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) SetBidwordIds(_bidwordIds []string) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// GetBidwordIds BidwordIds Getter
func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) GetBidwordIds() []string {
	return r._bidwordIds
}

// SetNick is Nick Setter
// 账号昵称
func (r *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdGroupId is AdGroupId Setter
// 推广组id
func (r *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) SetAdGroupId(_adGroupId int64) error {
	r._adGroupId = _adGroupId
	r.Set("ad_group_id", _adGroupId)
	return nil
}

// GetAdGroupId AdGroupId Getter
func (r TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIRequest) GetAdGroupId() int64 {
	return r._adGroupId
}
