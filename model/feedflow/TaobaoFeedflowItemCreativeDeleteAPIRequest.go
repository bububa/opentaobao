package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcreativedeleteAPIRequest 信息流删除创意 API请求
// taobao.feedflow.item.creative.delete
//
// 信息流删除创意
type TaobaofeedflowitemcreativedeleteAPIRequest struct {
	model.Params
	// 创意id列表
	_creativeIdList []string
}

// NewTaobaofeedflowitemcreativedeleteRequest 初始化TaobaofeedflowitemcreativedeleteAPIRequest对象
func NewTaobaofeedflowitemcreativedeleteRequest() *TaobaofeedflowitemcreativedeleteAPIRequest {
	return &TaobaofeedflowitemcreativedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemcreativedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.creative.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemcreativedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemcreativedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeIdList is CreativeIdList Setter
// 创意id列表
func (r *TaobaofeedflowitemcreativedeleteAPIRequest) SetCreativeIdList(_creativeIdList []string) error {
	r._creativeIdList = _creativeIdList
	r.Set("creative_id_list", _creativeIdList)
	return nil
}

// GetCreativeIdList CreativeIdList Getter
func (r TaobaofeedflowitemcreativedeleteAPIRequest) GetCreativeIdList() []string {
	return r._creativeIdList
}
