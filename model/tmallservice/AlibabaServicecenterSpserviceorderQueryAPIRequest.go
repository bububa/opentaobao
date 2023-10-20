package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterspserviceorderqueryAPIRequest 服务供应链服务单查询 API请求
// alibaba.servicecenter.spserviceorder.query
//
// 服务供应链服务单查询，服务商通过此接口拉取用户的购买的服务信息，以此为依据为用户提供安装维修等服务
type AlibabaservicecenterspserviceorderqueryAPIRequest struct {
	model.Params
	// 状态码，可传多个
	_statusCodes string
	// 服务单修改时间(时间段15分钟以内)(包含时分秒)
	_gmtModifiedEnd string
	// 服务单修改时间(包含时分秒)
	_gmtModifiedStart string
	// 查询第几页，默认1
	_currentPage int64
	// 每页大小，默认50，最大50
	_pageSize int64
	// 实物主订单id(消费者在淘宝订单里看到的订单号)
	_masterParentBizOrderId int64
	// 服务单id
	_spServiceOrderId int64
}

// NewAlibabaservicecenterspserviceorderqueryRequest 初始化AlibabaservicecenterspserviceorderqueryAPIRequest对象
func NewAlibabaservicecenterspserviceorderqueryRequest() *AlibabaservicecenterspserviceorderqueryAPIRequest {
	return &AlibabaservicecenterspserviceorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.spserviceorder.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStatusCodes is StatusCodes Setter
// 状态码，可传多个
func (r *AlibabaservicecenterspserviceorderqueryAPIRequest) SetStatusCodes(_statusCodes string) error {
	r._statusCodes = _statusCodes
	r.Set("status_codes", _statusCodes)
	return nil
}

// GetStatusCodes StatusCodes Getter
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetStatusCodes() string {
	return r._statusCodes
}

// SetGmtModifiedEnd is GmtModifiedEnd Setter
// 服务单修改时间(时间段15分钟以内)(包含时分秒)
func (r *AlibabaservicecenterspserviceorderqueryAPIRequest) SetGmtModifiedEnd(_gmtModifiedEnd string) error {
	r._gmtModifiedEnd = _gmtModifiedEnd
	r.Set("gmt_modified_end", _gmtModifiedEnd)
	return nil
}

// GetGmtModifiedEnd GmtModifiedEnd Getter
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetGmtModifiedEnd() string {
	return r._gmtModifiedEnd
}

// SetGmtModifiedStart is GmtModifiedStart Setter
// 服务单修改时间(包含时分秒)
func (r *AlibabaservicecenterspserviceorderqueryAPIRequest) SetGmtModifiedStart(_gmtModifiedStart string) error {
	r._gmtModifiedStart = _gmtModifiedStart
	r.Set("gmt_modified_start", _gmtModifiedStart)
	return nil
}

// GetGmtModifiedStart GmtModifiedStart Getter
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetGmtModifiedStart() string {
	return r._gmtModifiedStart
}

// SetCurrentPage is CurrentPage Setter
// 查询第几页，默认1
func (r *AlibabaservicecenterspserviceorderqueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}

// SetPageSize is PageSize Setter
// 每页大小，默认50，最大50
func (r *AlibabaservicecenterspserviceorderqueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetMasterParentBizOrderId is MasterParentBizOrderId Setter
// 实物主订单id(消费者在淘宝订单里看到的订单号)
func (r *AlibabaservicecenterspserviceorderqueryAPIRequest) SetMasterParentBizOrderId(_masterParentBizOrderId int64) error {
	r._masterParentBizOrderId = _masterParentBizOrderId
	r.Set("master_parent_biz_order_id", _masterParentBizOrderId)
	return nil
}

// GetMasterParentBizOrderId MasterParentBizOrderId Getter
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetMasterParentBizOrderId() int64 {
	return r._masterParentBizOrderId
}

// SetSpServiceOrderId is SpServiceOrderId Setter
// 服务单id
func (r *AlibabaservicecenterspserviceorderqueryAPIRequest) SetSpServiceOrderId(_spServiceOrderId int64) error {
	r._spServiceOrderId = _spServiceOrderId
	r.Set("sp_service_order_id", _spServiceOrderId)
	return nil
}

// GetSpServiceOrderId SpServiceOrderId Getter
func (r AlibabaservicecenterspserviceorderqueryAPIRequest) GetSpServiceOrderId() int64 {
	return r._spServiceOrderId
}
