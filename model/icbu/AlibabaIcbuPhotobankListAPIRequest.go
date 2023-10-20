package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaicbuphotobanklistAPIRequest 国际站图片银行查询接口 API请求
// alibaba.icbu.photobank.list
//
// 国际站图片银行查询接口
type AlibabaicbuphotobanklistAPIRequest struct {
	model.Params
	// 额外的上下文信息. 例如:icvId
	_extraContext string
	// 图片组id
	_groupId string
	// 存放位置 必要条件, 包括ALL_GROUP(所有目录), SUB_GROUP(指定图片组下),UNGROUP(未分组), TEMP(disable)四个值
	_locationType string
	// 当前翻页数
	_currentPage int64
	// 每页显示数
	_pageSize int64
}

// NewAlibabaicbuphotobanklistRequest 初始化AlibabaicbuphotobanklistAPIRequest对象
func NewAlibabaicbuphotobanklistRequest() *AlibabaicbuphotobanklistAPIRequest {
	return &AlibabaicbuphotobanklistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaicbuphotobanklistAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.photobank.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaicbuphotobanklistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaicbuphotobanklistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtraContext is ExtraContext Setter
// 额外的上下文信息. 例如:icvId
func (r *AlibabaicbuphotobanklistAPIRequest) SetExtraContext(_extraContext string) error {
	r._extraContext = _extraContext
	r.Set("extra_context", _extraContext)
	return nil
}

// GetExtraContext ExtraContext Getter
func (r AlibabaicbuphotobanklistAPIRequest) GetExtraContext() string {
	return r._extraContext
}

// SetGroupId is GroupId Setter
// 图片组id
func (r *AlibabaicbuphotobanklistAPIRequest) SetGroupId(_groupId string) error {
	r._groupId = _groupId
	r.Set("group_id", _groupId)
	return nil
}

// GetGroupId GroupId Getter
func (r AlibabaicbuphotobanklistAPIRequest) GetGroupId() string {
	return r._groupId
}

// SetLocationType is LocationType Setter
// 存放位置 必要条件, 包括ALL_GROUP(所有目录), SUB_GROUP(指定图片组下),UNGROUP(未分组), TEMP(disable)四个值
func (r *AlibabaicbuphotobanklistAPIRequest) SetLocationType(_locationType string) error {
	r._locationType = _locationType
	r.Set("location_type", _locationType)
	return nil
}

// GetLocationType LocationType Getter
func (r AlibabaicbuphotobanklistAPIRequest) GetLocationType() string {
	return r._locationType
}

// SetCurrentPage is CurrentPage Setter
// 当前翻页数
func (r *AlibabaicbuphotobanklistAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaicbuphotobanklistAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页显示数
func (r *AlibabaicbuphotobanklistAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaicbuphotobanklistAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
