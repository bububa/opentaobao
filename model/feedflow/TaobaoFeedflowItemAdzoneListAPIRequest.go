package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadzonelistAPIRequest 批量查询可用广告位列表 API请求
// taobao.feedflow.item.adzone.list
//
// 批量查询可用广告位列表
type TaobaofeedflowitemadzonelistAPIRequest struct {
	model.Params
	// 广告位查询条件
	_adzoneQuery *AdzoneQueryDto
}

// NewTaobaofeedflowitemadzonelistRequest 初始化TaobaofeedflowitemadzonelistAPIRequest对象
func NewTaobaofeedflowitemadzonelistRequest() *TaobaofeedflowitemadzonelistAPIRequest {
	return &TaobaofeedflowitemadzonelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadzonelistAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adzone.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadzonelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadzonelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzoneQuery is AdzoneQuery Setter
// 广告位查询条件
func (r *TaobaofeedflowitemadzonelistAPIRequest) SetAdzoneQuery(_adzoneQuery *AdzoneQueryDto) error {
	r._adzoneQuery = _adzoneQuery
	r.Set("adzone_query", _adzoneQuery)
	return nil
}

// GetAdzoneQuery AdzoneQuery Getter
func (r TaobaofeedflowitemadzonelistAPIRequest) GetAdzoneQuery() *AdzoneQueryDto {
	return r._adzoneQuery
}
