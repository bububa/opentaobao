package seaking

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSeakingTitlerewriteSubmitAPIRequest 提交标题改写任务 API请求
// alibaba.seaking.titlerewrite.submit
//
// 提交标题改写任务
type AlibabaSeakingTitlerewriteSubmitAPIRequest struct {
	model.Params
	// 任务详情列表
	_titleRewriteDetailList []TitleRewriteDetailDto
	// token来源站点
	_tokenFrom string
	// 用户token
	_token string
}

// NewAlibabaSeakingTitlerewriteSubmitRequest 初始化AlibabaSeakingTitlerewriteSubmitAPIRequest对象
func NewAlibabaSeakingTitlerewriteSubmitRequest() *AlibabaSeakingTitlerewriteSubmitAPIRequest {
	return &AlibabaSeakingTitlerewriteSubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaSeakingTitlerewriteSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.seaking.titlerewrite.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaSeakingTitlerewriteSubmitAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTitleRewriteDetailList is TitleRewriteDetailList Setter
// 任务详情列表
func (r *AlibabaSeakingTitlerewriteSubmitAPIRequest) SetTitleRewriteDetailList(_titleRewriteDetailList []TitleRewriteDetailDto) error {
	r._titleRewriteDetailList = _titleRewriteDetailList
	r.Set("title_rewrite_detail_list", _titleRewriteDetailList)
	return nil
}

// GetTitleRewriteDetailList TitleRewriteDetailList Getter
func (r AlibabaSeakingTitlerewriteSubmitAPIRequest) GetTitleRewriteDetailList() []TitleRewriteDetailDto {
	return r._titleRewriteDetailList
}

// SetTokenFrom is TokenFrom Setter
// token来源站点
func (r *AlibabaSeakingTitlerewriteSubmitAPIRequest) SetTokenFrom(_tokenFrom string) error {
	r._tokenFrom = _tokenFrom
	r.Set("token_from", _tokenFrom)
	return nil
}

// GetTokenFrom TokenFrom Getter
func (r AlibabaSeakingTitlerewriteSubmitAPIRequest) GetTokenFrom() string {
	return r._tokenFrom
}

// SetToken is Token Setter
// 用户token
func (r *AlibabaSeakingTitlerewriteSubmitAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaSeakingTitlerewriteSubmitAPIRequest) GetToken() string {
	return r._token
}
