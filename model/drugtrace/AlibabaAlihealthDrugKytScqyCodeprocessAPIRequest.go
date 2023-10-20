package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest 关联关系处理查询 API请求
// alibaba.alihealth.drug.kyt.scqy.codeprocess
//
// 关联关系处理查询
type AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest struct {
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

// NewAlibabaAlihealthDrugKytScqyCodeprocessRequest 初始化AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest对象
func NewAlibabaAlihealthDrugKytScqyCodeprocessRequest() *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest {
	return &AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest{
		Params: model.NewParams(14),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) Reset() {
	r._startDate = ""
	r._endDate = ""
	r._uploadFlag = ""
	r._processFlag = ""
	r._produceBatchNo = ""
	r._queryFlag = ""
	r._physicType = ""
	r._prodSeqNo = ""
	r._drugEntBaseInfoId = ""
	r._pkgSpec = ""
	r._refEntId = ""
	r._clientType = ""
	r._page = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.scqy.codeprocess"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStartDate is StartDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetStartDate(_startDate string) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetStartDate() string {
	return r._startDate
}

// SetEndDate is EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetEndDate(_endDate string) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetEndDate() string {
	return r._endDate
}

// SetUploadFlag is UploadFlag Setter
// 上传标识
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetUploadFlag(_uploadFlag string) error {
	r._uploadFlag = _uploadFlag
	r.Set("upload_flag", _uploadFlag)
	return nil
}

// GetUploadFlag UploadFlag Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetUploadFlag() string {
	return r._uploadFlag
}

// SetProcessFlag is ProcessFlag Setter
// 处理状态(所有状态 A ,待处理  1 ,处理成功  3 ,处理失败  4)
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetProcessFlag(_processFlag string) error {
	r._processFlag = _processFlag
	r.Set("process_flag", _processFlag)
	return nil
}

// GetProcessFlag ProcessFlag Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetProcessFlag() string {
	return r._processFlag
}

// SetProduceBatchNo is ProduceBatchNo Setter
// 批次号
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetProduceBatchNo(_produceBatchNo string) error {
	r._produceBatchNo = _produceBatchNo
	r.Set("produce_batch_no", _produceBatchNo)
	return nil
}

// GetProduceBatchNo ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetProduceBatchNo() string {
	return r._produceBatchNo
}

// SetQueryFlag is QueryFlag Setter
// 查询标识(处理状态查询 传0 , 关联关系个修改 传1)
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetQueryFlag(_queryFlag string) error {
	r._queryFlag = _queryFlag
	r.Set("query_flag", _queryFlag)
	return nil
}

// GetQueryFlag QueryFlag Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetQueryFlag() string {
	return r._queryFlag
}

// SetPhysicType is PhysicType Setter
// 药品类型（所有药品  A ，未分类 9， 特殊药品原料药  1， 特殊药品制  2， 普通药品	3）
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetPhysicType(_physicType string) error {
	r._physicType = _physicType
	r.Set("physic_type", _physicType)
	return nil
}

// GetPhysicType PhysicType Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetPhysicType() string {
	return r._physicType
}

// SetProdSeqNo is ProdSeqNo Setter
// 生产企业ID
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetProdSeqNo(_prodSeqNo string) error {
	r._prodSeqNo = _prodSeqNo
	r.Set("prod_seq_no", _prodSeqNo)
	return nil
}

// GetProdSeqNo ProdSeqNo Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetProdSeqNo() string {
	return r._prodSeqNo
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetPkgSpec is PkgSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetPkgSpec(_pkgSpec string) error {
	r._pkgSpec = _pkgSpec
	r.Set("pkg_spec", _pkgSpec)
	return nil
}

// GetPkgSpec PkgSpec Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetPkgSpec() string {
	return r._pkgSpec
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetClientType is ClientType Setter
// 客户端
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetClientType(_clientType string) error {
	r._clientType = _clientType
	r.Set("client_type", _clientType)
	return nil
}

// GetClientType ClientType Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetClientType() string {
	return r._clientType
}

// SetPage is Page Setter
// 页数
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 条数
func (r *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAlibabaAlihealthDrugKytScqyCodeprocessAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytScqyCodeprocessRequest()
	},
}

// GetAlibabaAlihealthDrugKytScqyCodeprocessRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest
func GetAlibabaAlihealthDrugKytScqyCodeprocessAPIRequest() *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest {
	return poolAlibabaAlihealthDrugKytScqyCodeprocessAPIRequest.Get().(*AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytScqyCodeprocessAPIRequest 将 AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytScqyCodeprocessAPIRequest(v *AlibabaAlihealthDrugKytScqyCodeprocessAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytScqyCodeprocessAPIRequest.Put(v)
}
