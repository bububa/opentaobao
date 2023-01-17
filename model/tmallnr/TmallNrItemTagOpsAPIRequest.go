package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrItemTagOpsAPIRequest 区域零售商品打标去标 API请求
// tmall.nr.item.tag.ops
//
// 参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
type TmallNrItemTagOpsAPIRequest struct {
	model.Params
	// 请求入参
	_tagReqDTO *TagReqDto
}

// NewTmallNrItemTagOpsRequest 初始化TmallNrItemTagOpsAPIRequest对象
func NewTmallNrItemTagOpsRequest() *TmallNrItemTagOpsAPIRequest {
	return &TmallNrItemTagOpsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrItemTagOpsAPIRequest) GetApiMethodName() string {
	return "tmall.nr.item.tag.ops"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrItemTagOpsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrItemTagOpsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTagReqDTO is TagReqDTO Setter
// 请求入参
func (r *TmallNrItemTagOpsAPIRequest) SetTagReqDTO(_tagReqDTO *TagReqDto) error {
	r._tagReqDTO = _tagReqDTO
	r.Set("tag_req_d_t_o", _tagReqDTO)
	return nil
}

// GetTagReqDTO TagReqDTO Getter
func (r TmallNrItemTagOpsAPIRequest) GetTagReqDTO() *TagReqDto {
	return r._tagReqDTO
}
