package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTagAddAPIRequest
创建关键词分组 API请求
alibaba.scbp.tag.add

创建关键词分组 */
type AlibabaScbpTagAddAPIRequest struct {
	model.Params
	// 分组名称，最多允许创建100个
	_tagName string
}

// NewAlibabaScbpTagAddRequest 初始化AlibabaScbpTagAddAPIRequest对象
func NewAlibabaScbpTagAddRequest() *AlibabaScbpTagAddAPIRequest {
	return &AlibabaScbpTagAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScbpTagAddAPIRequest) GetApiMethodName() string {
	return "alibaba.scbp.tag.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScbpTagAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TagName Setter
// 分组名称，最多允许创建100个
func (r *AlibabaScbpTagAddAPIRequest) SetTagName(_tagName string) error {
	r._tagName = _tagName
	r.Set("tag_name", _tagName)
	return nil
}

// Get TagName Getter
func (r AlibabaScbpTagAddAPIRequest) GetTagName() string {
	return r._tagName
}
