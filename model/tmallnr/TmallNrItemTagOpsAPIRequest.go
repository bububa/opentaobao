package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnritemtagopsAPIRequest 区域零售商品打标去标 API请求
// tmall.nr.item.tag.ops
//
// 参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
type TmallnritemtagopsAPIRequest struct {
	model.Params
	// 请求入参
	_tagReqDTO *TagReqDto
}

// NewTmallnritemtagopsRequest 初始化TmallnritemtagopsAPIRequest对象
func NewTmallnritemtagopsRequest() *TmallnritemtagopsAPIRequest {
	return &TmallnritemtagopsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnritemtagopsAPIRequest) GetApiMethodName() string {
	return "tmall.nr.item.tag.ops"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnritemtagopsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnritemtagopsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagReqDTO is TagReqDTO Setter
// 请求入参
func (r *TmallnritemtagopsAPIRequest) SetTagReqDTO(_tagReqDTO *TagReqDto) error {
	r._tagReqDTO = _tagReqDTO
	r.Set("tag_req_d_t_o", _tagReqDTO)
	return nil
}

// GetTagReqDTO TagReqDTO Getter
func (r TmallnritemtagopsAPIRequest) GetTagReqDTO() *TagReqDto {
	return r._tagReqDTO
}
