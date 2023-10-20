package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallPromotagTagRemovetagAPIRequest 删除标签定义 API请求
// tmall.promotag.tag.removetag
//
// 用于删除标签定义，但是要确保目前该标签没有人群在使用。
type TmallPromotagTagRemovetagAPIRequest struct {
	model.Params
	// 需要删除的标签id
	_tagId int64
}

// NewTmallPromotagTagRemovetagRequest 初始化TmallPromotagTagRemovetagAPIRequest对象
func NewTmallPromotagTagRemovetagRequest() *TmallPromotagTagRemovetagAPIRequest {
	return &TmallPromotagTagRemovetagAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallPromotagTagRemovetagAPIRequest) Reset() {
	r._tagId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallPromotagTagRemovetagAPIRequest) GetApiMethodName() string {
	return "tmall.promotag.tag.removetag"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallPromotagTagRemovetagAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallPromotagTagRemovetagAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagId is TagId Setter
// 需要删除的标签id
func (r *TmallPromotagTagRemovetagAPIRequest) SetTagId(_tagId int64) error {
	r._tagId = _tagId
	r.Set("tag_id", _tagId)
	return nil
}

// GetTagId TagId Getter
func (r TmallPromotagTagRemovetagAPIRequest) GetTagId() int64 {
	return r._tagId
}

var poolTmallPromotagTagRemovetagAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallPromotagTagRemovetagRequest()
	},
}

// GetTmallPromotagTagRemovetagRequest 从 sync.Pool 获取 TmallPromotagTagRemovetagAPIRequest
func GetTmallPromotagTagRemovetagAPIRequest() *TmallPromotagTagRemovetagAPIRequest {
	return poolTmallPromotagTagRemovetagAPIRequest.Get().(*TmallPromotagTagRemovetagAPIRequest)
}

// ReleaseTmallPromotagTagRemovetagAPIRequest 将 TmallPromotagTagRemovetagAPIRequest 放入 sync.Pool
func ReleaseTmallPromotagTagRemovetagAPIRequest(v *TmallPromotagTagRemovetagAPIRequest) {
	v.Reset()
	poolTmallPromotagTagRemovetagAPIRequest.Put(v)
}
