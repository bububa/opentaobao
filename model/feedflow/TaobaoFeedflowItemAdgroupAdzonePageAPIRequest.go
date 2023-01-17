package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAdzonePageAPIRequest 信息流单元下查看绑定资源位 API请求
// taobao.feedflow.item.adgroup.adzone.page
//
// 信息流单元下查看绑定资源位
type TaobaoFeedflowItemAdgroupAdzonePageAPIRequest struct {
	model.Params
	// 查询条件
	_adzoneBindQuery *AdzoneBindQueryDto
}

// NewTaobaoFeedflowItemAdgroupAdzonePageRequest 初始化TaobaoFeedflowItemAdgroupAdzonePageAPIRequest对象
func NewTaobaoFeedflowItemAdgroupAdzonePageRequest() *TaobaoFeedflowItemAdgroupAdzonePageAPIRequest {
	return &TaobaoFeedflowItemAdgroupAdzonePageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.adzone.page"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzoneBindQuery is AdzoneBindQuery Setter
// 查询条件
func (r *TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) SetAdzoneBindQuery(_adzoneBindQuery *AdzoneBindQueryDto) error {
	r._adzoneBindQuery = _adzoneBindQuery
	r.Set("adzone_bind_query", _adzoneBindQuery)
	return nil
}

// GetAdzoneBindQuery AdzoneBindQuery Getter
func (r TaobaoFeedflowItemAdgroupAdzonePageAPIRequest) GetAdzoneBindQuery() *AdzoneBindQueryDto {
	return r._adzoneBindQuery
}
