package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivitiesgetAPIRequest 查询活动列表 API请求
// taobao.ump.activities.get
//
// 查询活动列表
type TaobaoumpactivitiesgetAPIRequest struct {
	model.Params
	// 工具id
	_toolId int64
	// 分页的页码
	_pageNo int64
	// 每页的最大条数
	_pageSize int64
}

// NewTaobaoumpactivitiesgetRequest 初始化TaobaoumpactivitiesgetAPIRequest对象
func NewTaobaoumpactivitiesgetRequest() *TaobaoumpactivitiesgetAPIRequest {
	return &TaobaoumpactivitiesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpactivitiesgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activities.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpactivitiesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpactivitiesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToolId is ToolId Setter
// 工具id
func (r *TaobaoumpactivitiesgetAPIRequest) SetToolId(_toolId int64) error {
	r._toolId = _toolId
	r.Set("tool_id", _toolId)
	return nil
}

// GetToolId ToolId Getter
func (r TaobaoumpactivitiesgetAPIRequest) GetToolId() int64 {
	return r._toolId
}

// SetPageNo is PageNo Setter
// 分页的页码
func (r *TaobaoumpactivitiesgetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoumpactivitiesgetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页的最大条数
func (r *TaobaoumpactivitiesgetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoumpactivitiesgetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
