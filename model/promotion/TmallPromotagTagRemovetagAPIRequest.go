package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallpromotagtagremovetagAPIRequest 删除标签定义 API请求
// tmall.promotag.tag.removetag
//
// 用于删除标签定义，但是要确保目前该标签没有人群在使用。
type TmallpromotagtagremovetagAPIRequest struct {
	model.Params
	// 需要删除的标签id
	_tagId int64
}

// NewTmallpromotagtagremovetagRequest 初始化TmallpromotagtagremovetagAPIRequest对象
func NewTmallpromotagtagremovetagRequest() *TmallpromotagtagremovetagAPIRequest {
	return &TmallpromotagtagremovetagAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallpromotagtagremovetagAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.tag.removetag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallpromotagtagremovetagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallpromotagtagremovetagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagId is TagId Setter
// 需要删除的标签id
func (r *TmallpromotagtagremovetagAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallpromotagtagremovetagAPIRequest) GetTagId() int64 {
	return r._tagId
}
