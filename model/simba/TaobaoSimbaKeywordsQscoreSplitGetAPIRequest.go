package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsQscoreSplitGetAPIRequest 新质量分服务 API请求
// taobao.simba.keywords.qscore.split.get
//
// 获取关键词新的质量分
type TaobaoSimbaKeywordsQscoreSplitGetAPIRequest struct {
	model.Params
	// 词id数组（最多批量获取20个）
	_bidwordIds []string
	// 账号昵称
	_nick string
	// 推广组id
	_adGroupId int64
}

// NewTaobaoSimbaKeywordsQscoreSplitGetRequest 初始化TaobaoSimbaKeywordsQscoreSplitGetAPIRequest对象
func NewTaobaoSimbaKeywordsQscoreSplitGetRequest() *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest {
	return &TaobaoSimbaKeywordsQscoreSplitGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) Reset() {
	r._bidwordIds = r._bidwordIds[:0]
	r._nick = ""
	r._adGroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.qscore.split.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordIds is BidwordIds Setter
// 词id数组（最多批量获取20个）
func (r *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) SetBidwordIds(_bidwordIds []string) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// GetBidwordIds BidwordIds Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetBidwordIds() []string {
	return r._bidwordIds
}

// SetNick is Nick Setter
// 账号昵称
func (r *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdGroupId is AdGroupId Setter
// 推广组id
func (r *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) SetAdGroupId(_adGroupId int64) error {
	r._adGroupId = _adGroupId
	r.Set("ad_group_id", _adGroupId)
	return nil
}

// GetAdGroupId AdGroupId Getter
func (r TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) GetAdGroupId() int64 {
	return r._adGroupId
}

var poolTaobaoSimbaKeywordsQscoreSplitGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaKeywordsQscoreSplitGetRequest()
	},
}

// GetTaobaoSimbaKeywordsQscoreSplitGetRequest 从 sync.Pool 获取 TaobaoSimbaKeywordsQscoreSplitGetAPIRequest
func GetTaobaoSimbaKeywordsQscoreSplitGetAPIRequest() *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest {
	return poolTaobaoSimbaKeywordsQscoreSplitGetAPIRequest.Get().(*TaobaoSimbaKeywordsQscoreSplitGetAPIRequest)
}

// ReleaseTaobaoSimbaKeywordsQscoreSplitGetAPIRequest 将 TaobaoSimbaKeywordsQscoreSplitGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaKeywordsQscoreSplitGetAPIRequest(v *TaobaoSimbaKeywordsQscoreSplitGetAPIRequest) {
	v.Reset()
	poolTaobaoSimbaKeywordsQscoreSplitGetAPIRequest.Put(v)
}
