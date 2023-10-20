package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytweslistupoutAPIRequest 查询货主/本企业上游企业出库单据信息 API请求
// alibaba.alihealth.drug.kyt.wes.listupout
//
// 查询货主/本企业上游企业出库单据信息
type AlibabaalihealthdrugkytweslistupoutAPIRequest struct {
	model.Params
	// 货主企业唯一标识（一般情况下填写自已企业）
	_refEntId string
	// licenseToken
	_licenseToken string
	// 单据时间的开始日期（不写时分秒）
	_beginDate string
	// 单据时间的结束日期（不写时分秒）
	_endDate string
	// 发货单位
	_fromUserId string
	// 生产批号
	_produceBatchNo string
	// 药品ID
	_drugEntBaseInfoId string
	// 单据类型
	_billType string
	// 药品类型
	_physicType string
	// 状态
	_status string
	// 单据号
	_billCode string
	// 第三方物流企业唯一标识（只有代查其它企业数据时填写）
	_agentRefEntId string
	// 单据上传的开始日期（不写时分秒）
	_uploadTimeBegin string
	// 单据上传的结束日期（不写时分秒）
	_uploadTimeEnd string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaalihealthdrugkytweslistupoutRequest 初始化AlibabaalihealthdrugkytweslistupoutAPIRequest对象
func NewAlibabaalihealthdrugkytweslistupoutRequest() *AlibabaalihealthdrugkytweslistupoutAPIRequest {
	return &AlibabaalihealthdrugkytweslistupoutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.wes.listupout"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 货主企业唯一标识（一般情况下填写自已企业）
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetLicenseToken is LicenseToken Setter
// licenseToken
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetLicenseToken(_licenseToken string) error {
	r._licenseToken = _licenseToken
	r.Set("license_token", _licenseToken)
	return nil
}

// GetLicenseToken LicenseToken Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetLicenseToken() string {
	return r._licenseToken
}

// SetBeginDate is BeginDate Setter
// 单据时间的开始日期（不写时分秒）
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 单据时间的结束日期（不写时分秒）
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetFromUserId is FromUserId Setter
// 发货单位
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetFromUserId(_fromUserId string) error {
	r._fromUserId = _fromUserId
	r.Set("from_user_id", _fromUserId)
	return nil
}

// GetFromUserId FromUserId Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetFromUserId() string {
	return r._fromUserId
}

// SetProduceBatchNo is ProduceBatchNo Setter
// 生产批号
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// GetProduceBatchNo ProduceBatchNo Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetBillType() string {
	return r._billType
}

// SetPhysicType is PhysicType Setter
// 药品类型
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetPhysicType() string {
	return r._physicType
}

// SetStatus is Status Setter
// 状态
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetStatus() string {
	return r._status
}

// SetBillCode is BillCode Setter
// 单据号
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetAgentRefEntId is AgentRefEntId Setter
// 第三方物流企业唯一标识（只有代查其它企业数据时填写）
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetAgentRefEntId(_agentRefEntId string) error {
	r._agentRefEntId = _agentRefEntId
	r.Set("agent_ref_ent_id", _agentRefEntId)
	return nil
}

// GetAgentRefEntId AgentRefEntId Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetAgentRefEntId() string {
	return r._agentRefEntId
}

// SetUploadTimeBegin is UploadTimeBegin Setter
// 单据上传的开始日期（不写时分秒）
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetUploadTimeBegin(_uploadTimeBegin string) error {
	r._uploadTimeBegin = _uploadTimeBegin
	r.Set("upload_time_begin", _uploadTimeBegin)
	return nil
}

// GetUploadTimeBegin UploadTimeBegin Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetUploadTimeBegin() string {
	return r._uploadTimeBegin
}

// SetUploadTimeEnd is UploadTimeEnd Setter
// 单据上传的结束日期（不写时分秒）
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetUploadTimeEnd(_uploadTimeEnd string) error {
	r._uploadTimeEnd = _uploadTimeEnd
	r.Set("upload_time_end", _uploadTimeEnd)
	return nil
}

// GetUploadTimeEnd UploadTimeEnd Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetUploadTimeEnd() string {
	return r._uploadTimeEnd
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaalihealthdrugkytweslistupoutAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthdrugkytweslistupoutAPIRequest) GetPage() int64 {
	return r._page
}
