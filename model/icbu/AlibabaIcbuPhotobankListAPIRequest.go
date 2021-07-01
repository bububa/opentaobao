package icbu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际站图片银行查询接口 API请求
alibaba.icbu.photobank.list

国际站图片银行查询接口
*/
type AlibabaIcbuPhotobankListAPIRequest struct {
    model.Params
    // 当前翻页数
    _currentPage   int64
    // 图片组id
    _groupId   string
    // 存放位置 必要条件, 包括ALL_GROUP(所有目录), SUB_GROUP(指定图片组下),UNGROUP(未分组), TEMP(disable)四个值
    _locationType   string
    // 每页显示数
    _pageSize   int64
    // 额外的上下文信息. 例如:icvId
    _extraContext   string
}

// 初始化AlibabaIcbuPhotobankListAPIRequest对象
func NewAlibabaIcbuPhotobankListRequest() *AlibabaIcbuPhotobankListAPIRequest{
    return &AlibabaIcbuPhotobankListAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIcbuPhotobankListAPIRequest) GetApiMethodName() string {
    return "alibaba.icbu.photobank.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIcbuPhotobankListAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CurrentPage Setter
// 当前翻页数
func (r *AlibabaIcbuPhotobankListAPIRequest) SetCurrentPage(_currentPage int64) error {
    r._currentPage = _currentPage
    r.Set("current_page", _currentPage)
    return nil
}

// CurrentPage Getter
func (r AlibabaIcbuPhotobankListAPIRequest) GetCurrentPage() int64 {
    return r._currentPage
}
// GroupId Setter
// 图片组id
func (r *AlibabaIcbuPhotobankListAPIRequest) SetGroupId(_groupId string) error {
    r._groupId = _groupId
    r.Set("group_id", _groupId)
    return nil
}

// GroupId Getter
func (r AlibabaIcbuPhotobankListAPIRequest) GetGroupId() string {
    return r._groupId
}
// LocationType Setter
// 存放位置 必要条件, 包括ALL_GROUP(所有目录), SUB_GROUP(指定图片组下),UNGROUP(未分组), TEMP(disable)四个值
func (r *AlibabaIcbuPhotobankListAPIRequest) SetLocationType(_locationType string) error {
    r._locationType = _locationType
    r.Set("location_type", _locationType)
    return nil
}

// LocationType Getter
func (r AlibabaIcbuPhotobankListAPIRequest) GetLocationType() string {
    return r._locationType
}
// PageSize Setter
// 每页显示数
func (r *AlibabaIcbuPhotobankListAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaIcbuPhotobankListAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// ExtraContext Setter
// 额外的上下文信息. 例如:icvId
func (r *AlibabaIcbuPhotobankListAPIRequest) SetExtraContext(_extraContext string) error {
    r._extraContext = _extraContext
    r.Set("extra_context", _extraContext)
    return nil
}

// ExtraContext Getter
func (r AlibabaIcbuPhotobankListAPIRequest) GetExtraContext() string {
    return r._extraContext
}
