package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytCodeprocessAPIRequest 关联关系处理查询 API请求
// alibaba.alihealth.drug.kyt.codeprocess
//
// 关联关系处理查询
type AlibabaAlihealthDrugKytCodeprocessAPIRequest struct {
	model.Params
	// 开始时间
	_startDate string
	// 结束时间
	_endDate string
	// 上传标识
	_uploadFlag string
	// 处理状态(所有状态 A ,待处理  1 ,处理成功  3 ,处理失败  4)
	_processFlag string
	// 批次号
	_produceBatchNo string
	// 查询标识(处理状态查询 传0 , 关联关系个修改 传1)
	_queryFlag string
	// 药品类型（所有药品  A ，未分类 9， 特殊药品原料药  1， 特殊药品制  2， 普通药品	3）
	_physicType string
	// 生产企业ID
	_prodSeqNo string
	// 药品ID
	_drugEntBaseInfoId string
	// 包装规格
	_pkgSpec string
	// 企业ID
	_refEntId string
	// 客户端
	_clientType string
	// 页数
	_page int64
	// 条数
	_pageSize int64
}

// NewAlibabaAlihealthDrugKytCodeprocessRequest 初始化AlibabaAlihealthDrugKytCodeprocessAPIRequest对象
func NewAlibabaAlihealthDrugKytCodeprocessRequest() *AlibabaAlihealthDrugKytCodeprocessAPIRequest {
	return &AlibabaAlihealthDrugKytCodeprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codeprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetUploadFlag is UploadFlag Setter
// 上传标识
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetUploadFlag(_uploadFlag string) error {
	r._uploadFlag = _uploadFlag
	r.Set("upload_flag", _uploadFlag)
	return nil
}

// GetUploadFlag UploadFlag Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetUploadFlag() string {
	return r._uploadFlag
}

// SetProcessFlag is ProcessFlag Setter
// 处理状态(所有状态 A ,待处理  1 ,处理成功  3 ,处理失败  4)
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetProcessFlag(_processFlag string) error {
	r._processFlag = _processFlag
	r.Set("process_flag", _processFlag)
	return nil
}

// GetProcessFlag ProcessFlag Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetProcessFlag() string {
	return r._processFlag
}

// SetProduceBatchNo is ProduceBatchNo Setter
// 批次号
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// GetProduceBatchNo ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// SetQueryFlag is QueryFlag Setter
// 查询标识(处理状态查询 传0 , 关联关系个修改 传1)
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetQueryFlag(_queryFlag string) error {
	r._queryFlag = _queryFlag
	r.Set("query_flag", _queryFlag)
	return nil
}

// GetQueryFlag QueryFlag Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetQueryFlag() string {
	return r._queryFlag
}

// SetPhysicType is PhysicType Setter
// 药品类型（所有药品  A ，未分类 9， 特殊药品原料药  1， 特殊药品制  2， 普通药品	3）
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetPhysicType() string {
	return r._physicType
}

// SetProdSeqNo is ProdSeqNo Setter
// 生产企业ID
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetProdSeqNo(_prodSeqNo string) error {
	r._prodSeqNo = _prodSeqNo
	r.Set("prod_seq_no", _prodSeqNo)
	return nil
}

// GetProdSeqNo ProdSeqNo Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetProdSeqNo() string {
	return r._prodSeqNo
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetPkgSpec is PkgSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetPkgSpec(_pkgSpec string) error {
	r._pkgSpec = _pkgSpec
	r.Set("pkg_spec", _pkgSpec)
	return nil
}

// GetPkgSpec PkgSpec Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetPkgSpec() string {
	return r._pkgSpec
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetClientType is ClientType Setter
// 客户端
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetClientType() string {
	return r._clientType
}

// SetPage is Page Setter
// 页数
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 条数
func (r *AlibabaAlihealthDrugKytCodeprocessAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytCodeprocessAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
