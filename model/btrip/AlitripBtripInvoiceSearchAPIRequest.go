package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripinvoicesearchAPIRequest 根据发票抬头搜索发票 API请求
// alitrip.btrip.invoice.search
//
// 用户根据发票抬头搜索发票信息
type AlitripbtripinvoicesearchAPIRequest struct {
	model.Params
	// 企业id
	_corpId string
	// 用户id
	_userId string
	// 发票抬头
	_title string
}

// NewAlitripbtripinvoicesearchRequest 初始化AlitripbtripinvoicesearchAPIRequest对象
func NewAlitripbtripinvoicesearchRequest() *AlitripbtripinvoicesearchAPIRequest {
	return &AlitripbtripinvoicesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripbtripinvoicesearchAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.invoice.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripbtripinvoicesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripbtripinvoicesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCorpId is CorpId Setter
// 企业id
func (r *AlitripbtripinvoicesearchAPIRequest) SetCorpId(_corpId string) error {
	r._corpId = _corpId
	r.Set("corp_id", _corpId)
	return nil
}

// GetCorpId CorpId Getter
func (r AlitripbtripinvoicesearchAPIRequest) GetCorpId() string {
	return r._corpId
}

// SetUserId is UserId Setter
// 用户id
func (r *AlitripbtripinvoicesearchAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlitripbtripinvoicesearchAPIRequest) GetUserId() string {
	return r._userId
}

// SetTitle is Title Setter
// 发票抬头
func (r *AlitripbtripinvoicesearchAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AlitripbtripinvoicesearchAPIRequest) GetTitle() string {
	return r._title
}
