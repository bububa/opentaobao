package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAdzoneListAPIRequest
批量查询可用广告位列表 API请求
taobao.feedflow.item.adzone.list

批量查询可用广告位列表 */
type TaobaoFeedflowItemAdzoneListAPIRequest struct {
	model.Params
	// 广告位查询条件
	_adzoneQuery *AdzoneQueryDto
}

// NewTaobaoFeedflowItemAdzoneListRequest 初始化TaobaoFeedflowItemAdzoneListAPIRequest对象
func NewTaobaoFeedflowItemAdzoneListRequest() *TaobaoFeedflowItemAdzoneListAPIRequest {
	return &TaobaoFeedflowItemAdzoneListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdzoneListAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adzone.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdzoneListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AdzoneQuery Setter
// 广告位查询条件
func (r *TaobaoFeedflowItemAdzoneListAPIRequest) SetAdzoneQuery(_adzoneQuery *AdzoneQueryDto) error {
	r._adzoneQuery = _adzoneQuery
	r.Set("adzone_query", _adzoneQuery)
	return nil
}

// Get AdzoneQuery Getter
func (r TaobaoFeedflowItemAdzoneListAPIRequest) GetAdzoneQuery() *AdzoneQueryDto {
	return r._adzoneQuery
}
