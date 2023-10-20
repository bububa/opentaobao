package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest 获取关键词的新版实时排名 API请求
// taobao.simba.keywords.realtime.ranking.batch.get
//
// 根据关键词ID获取关键词的新版实时排名
type TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest struct {
	model.Params
	// 关键词列表集合,id用半角逗号分割，一次最多20个
	_bidwordIds []string
	// 旺旺名称
	_nick string
	// adgroupId
	_adGroupId int64
}

// NewTaobaosimbakeywordsrealtimerankingbatchgetRequest 初始化TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest对象
func NewTaobaosimbakeywordsrealtimerankingbatchgetRequest() *TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest {
	return &TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) GetApiMethodName() string {
	return "taobao.simba.keywords.realtime.ranking.batch.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBidwordIds is BidwordIds Setter
// 关键词列表集合,id用半角逗号分割，一次最多20个
func (r *TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) SetBidwordIds(_bidwordIds []string) error {
	r._bidwordIds = _bidwordIds
	r.Set("bidword_ids", _bidwordIds)
	return nil
}

// GetBidwordIds BidwordIds Getter
func (r TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) GetBidwordIds() []string {
	return r._bidwordIds
}

// SetNick is Nick Setter
// 旺旺名称
func (r *TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) GetNick() string {
	return r._nick
}

// SetAdGroupId is AdGroupId Setter
// adgroupId
func (r *TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) SetAdGroupId(_adGroupId int64) error {
	r._adGroupId = _adGroupId
	r.Set("ad_group_id", _adGroupId)
	return nil
}

// GetAdGroupId AdGroupId Getter
func (r TaobaosimbakeywordsrealtimerankingbatchgetAPIRequest) GetAdGroupId() int64 {
	return r._adGroupId
}
