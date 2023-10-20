package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaseakingtitlerewritesubmitAPIRequest 提交标题改写任务 API请求
// alibaba.seaking.titlerewrite.submit
//
// 提交标题改写任务
type AlibabaseakingtitlerewritesubmitAPIRequest struct {
	model.Params
	// 任务详情列表
	_titleRewriteDetailList []TitleRewriteDetailDto
	// token来源站点
	_tokenFrom string
	// 用户token
	_token string
}

// NewAlibabaseakingtitlerewritesubmitRequest 初始化AlibabaseakingtitlerewritesubmitAPIRequest对象
func NewAlibabaseakingtitlerewritesubmitRequest() *AlibabaseakingtitlerewritesubmitAPIRequest {
	return &AlibabaseakingtitlerewritesubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaseakingtitlerewritesubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.titlerewrite.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaseakingtitlerewritesubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaseakingtitlerewritesubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitleRewriteDetailList is TitleRewriteDetailList Setter
// 任务详情列表
func (r *AlibabaseakingtitlerewritesubmitAPIRequest) SetTitleRewriteDetailList(_titleRewriteDetailList []TitleRewriteDetailDto) error {
	r._titleRewriteDetailList = _titleRewriteDetailList
	r.Set("title_rewrite_detail_list", _titleRewriteDetailList)
	return nil
}

// GetTitleRewriteDetailList TitleRewriteDetailList Getter
func (r AlibabaseakingtitlerewritesubmitAPIRequest) GetTitleRewriteDetailList() []TitleRewriteDetailDto {
	return r._titleRewriteDetailList
}

// SetTokenFrom is TokenFrom Setter
// token来源站点
func (r *AlibabaseakingtitlerewritesubmitAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// GetTokenFrom TokenFrom Getter
func (r AlibabaseakingtitlerewritesubmitAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}

// SetToken is Token Setter
// 用户token
func (r *AlibabaseakingtitlerewritesubmitAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaseakingtitlerewritesubmitAPIRequest) GetToken() string {
	return r._token
}
