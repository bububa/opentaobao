package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsqscoresplitgetAPIRequest 新质量分服务 API请求
// taobao.simba.keywords.qscore.split.get
//
// 获取关键词新的质量分
type TaobaosimbakeywordsqscoresplitgetAPIRequest struct {
	model.Params
	// 词id数组（最多批量获取20个）
	_bidwordIds []string
	// 账号昵称
	_nick string
	// 推广组id
	_adGroupId int64
}

// NewTaobaosimbakeywordsqscoresplitgetRequest 初始化TaobaosimbakeywordsqscoresplitgetAPIRequest对象
func NewTaobaosimbakeywordsqscoresplitgetRequest() *TaobaosimbakeywordsqscoresplitgetAPIRequest {
	return &TaobaosimbakeywordsqscoresplitgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordsqscoresplitgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.qscore.split.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordsqscoresplitgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordsqscoresplitgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordIds is BidwordIds Setter
// 词id数组（最多批量获取20个）
func (r *TaobaosimbakeywordsqscoresplitgetAPIRequest) SetBidwordIds(_bidwordIds []string) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// GetBidwordIds BidwordIds Getter
func (r TaobaosimbakeywordsqscoresplitgetAPIRequest) GetBidwordIds() []string {
	return r._bidwordIds
}

// SetNick is Nick Setter
// 账号昵称
func (r *TaobaosimbakeywordsqscoresplitgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordsqscoresplitgetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdGroupId is AdGroupId Setter
// 推广组id
func (r *TaobaosimbakeywordsqscoresplitgetAPIRequest) SetAdGroupId(_adGroupId int64) error {
	r._adGroupId = _adGroupId
	r.Set("ad_group_id", _adGroupId)
	return nil
}

// GetAdGroupId AdGroupId Getter
func (r TaobaosimbakeywordsqscoresplitgetAPIRequest) GetAdGroupId() int64 {
	return r._adGroupId
}
