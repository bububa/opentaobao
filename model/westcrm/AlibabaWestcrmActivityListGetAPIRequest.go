package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmActivityListGetAPIRequest 获取活动列表接口 API请求
// alibaba.westcrm.activity.list.get
//
// 获取活动列表提供给友盟&互动屏
type AlibabaWestcrmActivityListGetAPIRequest struct {
	model.Params
	// 活动状态
	_status int64
	// 园区id
	_campusId int64
	// 排序方向
	_sord string
	// 页,默认第一页
	_page int64
	// 排序字段
	_sidx string
	// 分页偏移量 eq . limit offset ,rows
	_offset int64
	// 页大小,默认每页查询10条数据
	_rows int64
}

// NewAlibabaWestcrmActivityListGetRequest 初始化AlibabaWestcrmActivityListGetAPIRequest对象
func NewAlibabaWestcrmActivityListGetRequest() *AlibabaWestcrmActivityListGetAPIRequest {
	return &AlibabaWestcrmActivityListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmActivityListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.activity.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmActivityListGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Status Setter
// 活动状态
func (r *AlibabaWestcrmActivityListGetAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// Get Status Getter
func (r AlibabaWestcrmActivityListGetAPIRequest) GetStatus() int64 {
	return r._status
}

// Set is CampusId Setter
// 园区id
func (r *AlibabaWestcrmActivityListGetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// Get CampusId Getter
func (r AlibabaWestcrmActivityListGetAPIRequest) GetCampusId() int64 {
	return r._campusId
}

// Set is Sord Setter
// 排序方向
func (r *AlibabaWestcrmActivityListGetAPIRequest) SetSord(_sord string) error {
	r._sord = _sord
	r.Set("sord", _sord)
	return nil
}

// Get Sord Getter
func (r AlibabaWestcrmActivityListGetAPIRequest) GetSord() string {
	return r._sord
}

// Set is Page Setter
// 页,默认第一页
func (r *AlibabaWestcrmActivityListGetAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// Get Page Getter
func (r AlibabaWestcrmActivityListGetAPIRequest) GetPage() int64 {
	return r._page
}

// Set is Sidx Setter
// 排序字段
func (r *AlibabaWestcrmActivityListGetAPIRequest) SetSidx(_sidx string) error {
	r._sidx = _sidx
	r.Set("sidx", _sidx)
	return nil
}

// Get Sidx Getter
func (r AlibabaWestcrmActivityListGetAPIRequest) GetSidx() string {
	return r._sidx
}

// Set is Offset Setter
// 分页偏移量 eq . limit offset ,rows
func (r *AlibabaWestcrmActivityListGetAPIRequest) SetOffset(_offset int64) error {
	r._offset = _offset
	r.Set("offset", _offset)
	return nil
}

// Get Offset Getter
func (r AlibabaWestcrmActivityListGetAPIRequest) GetOffset() int64 {
	return r._offset
}

// Set is Rows Setter
// 页大小,默认每页查询10条数据
func (r *AlibabaWestcrmActivityListGetAPIRequest) SetRows(_rows int64) error {
	r._rows = _rows
	r.Set("rows", _rows)
	return nil
}

// Get Rows Getter
func (r AlibabaWestcrmActivityListGetAPIRequest) GetRows() int64 {
	return r._rows
}
