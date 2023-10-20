package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSmyxListpartsAPIRequest 药店查询往来单位 API请求
// alibaba.alihealth.drug.kyt.smyx.listparts
//
// 查询往来单位列表
type AlibabaAlihealthDrugKytSmyxListpartsAPIRequest struct {
	model.Params
	// 企业唯一标识
	_refEntId string
	// 企业名称
	_entName string
	// 企业自定义编号
	_refPartnerId string
	// 开始时间
	_beginDate string
	// 结束时间
	_endDate string
	// 页大小
	_pageSize int64
	// 页码
	_page int64
}

// NewAlibabaAlihealthDrugKytSmyxListpartsRequest 初始化AlibabaAlihealthDrugKytSmyxListpartsAPIRequest对象
func NewAlibabaAlihealthDrugKytSmyxListpartsRequest() *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest {
	return &AlibabaAlihealthDrugKytSmyxListpartsAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) Reset() {
	r._refEntId = ""
	r._entName = ""
	r._refPartnerId = ""
	r._beginDate = ""
	r._endDate = ""
	r._pageSize = 0
	r._page = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.smyx.listparts"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业唯一标识
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetEntName() string {
	return r._entName
}

// SetRefPartnerId is RefPartnerId Setter
// 企业自定义编号
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) SetRefPartnerId(_refPartnerId string) error {
	r._refPartnerId = _refPartnerId
	r.Set("ref_partner_id", _refPartnerId)
	return nil
}

// GetRefPartnerId RefPartnerId Getter
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetRefPartnerId() string {
	return r._refPartnerId
}

// SetBeginDate is BeginDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) SetBeginDate(_beginDate string) error {
	r._beginDate = _beginDate
	r.Set("begin_date", _beginDate)
	return nil
}

// GetBeginDate BeginDate Getter
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetBeginDate() string {
	return r._beginDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) GetPage() int64 {
	return r._page
}

var poolAlibabaAlihealthDrugKytSmyxListpartsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytSmyxListpartsRequest()
	},
}

// GetAlibabaAlihealthDrugKytSmyxListpartsRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytSmyxListpartsAPIRequest
func GetAlibabaAlihealthDrugKytSmyxListpartsAPIRequest() *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest {
	return poolAlibabaAlihealthDrugKytSmyxListpartsAPIRequest.Get().(*AlibabaAlihealthDrugKytSmyxListpartsAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytSmyxListpartsAPIRequest 将 AlibabaAlihealthDrugKytSmyxListpartsAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSmyxListpartsAPIRequest(v *AlibabaAlihealthDrugKytSmyxListpartsAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSmyxListpartsAPIRequest.Put(v)
}
