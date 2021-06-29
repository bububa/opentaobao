package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
关联关系处理查询 API请求
alibaba.alihealth.drug.kyt.codeprocess

关联关系处理查询
*/
type AlibabaAlihealthDrugKytCodeprocessRequest struct {
    model.Params
    // 开始时间
    _startDate   string
    // 结束时间
    _endDate   string
    // 上传标识
    _uploadFlag   string
    // 处理状态
    _processFlag   string
    // 批次号
    _produceBatchNo   string
    // 查询标识
    _queryFlag   string
    // 药品类型
    _physicType   string
    // 生产企业ID
    _prodSeqNo   string
    // 药品ID
    _drugEntBaseInfoId   string
    // 包装规格
    _pkgSpec   string
    // 页数
    _page   int64
    // 条数
    _pageSize   int64
    // 客户端
    _clientType   string
    // 企业ID
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugKytCodeprocessRequest对象
func NewAlibabaAlihealthDrugKytCodeprocessRequest() *AlibabaAlihealthDrugKytCodeprocessRequest{
    return &AlibabaAlihealthDrugKytCodeprocessRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.codeprocess"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartDate Setter
// 开始时间
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetStartDate(_startDate string) error {
    r._startDate = _startDate
    r.Set("start_date", _startDate)
    return nil
}

// StartDate Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetStartDate() string {
    return r._startDate
}
// EndDate Setter
// 结束时间
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetEndDate(_endDate string) error {
    r._endDate = _endDate
    r.Set("end_date", _endDate)
    return nil
}

// EndDate Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetEndDate() string {
    return r._endDate
}
// UploadFlag Setter
// 上传标识
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetUploadFlag(_uploadFlag string) error {
    r._uploadFlag = _uploadFlag
    r.Set("upload_flag", _uploadFlag)
    return nil
}

// UploadFlag Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetUploadFlag() string {
    return r._uploadFlag
}
// ProcessFlag Setter
// 处理状态
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetProcessFlag(_processFlag string) error {
    r._processFlag = _processFlag
    r.Set("process_flag", _processFlag)
    return nil
}

// ProcessFlag Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetProcessFlag() string {
    return r._processFlag
}
// ProduceBatchNo Setter
// 批次号
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetProduceBatchNo(_produceBatchNo string) error {
    r._produceBatchNo = _produceBatchNo
    r.Set("produce_batch_no", _produceBatchNo)
    return nil
}

// ProduceBatchNo Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetProduceBatchNo() string {
    return r._produceBatchNo
}
// QueryFlag Setter
// 查询标识
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetQueryFlag(_queryFlag string) error {
    r._queryFlag = _queryFlag
    r.Set("query_flag", _queryFlag)
    return nil
}

// QueryFlag Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetQueryFlag() string {
    return r._queryFlag
}
// PhysicType Setter
// 药品类型
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPhysicType(_physicType string) error {
    r._physicType = _physicType
    r.Set("physic_type", _physicType)
    return nil
}

// PhysicType Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPhysicType() string {
    return r._physicType
}
// ProdSeqNo Setter
// 生产企业ID
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetProdSeqNo(_prodSeqNo string) error {
    r._prodSeqNo = _prodSeqNo
    r.Set("prod_seq_no", _prodSeqNo)
    return nil
}

// ProdSeqNo Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetProdSeqNo() string {
    return r._prodSeqNo
}
// DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
    r._drugEntBaseInfoId = _drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
    return nil
}

// DrugEntBaseInfoId Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetDrugEntBaseInfoId() string {
    return r._drugEntBaseInfoId
}
// PkgSpec Setter
// 包装规格
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPkgSpec(_pkgSpec string) error {
    r._pkgSpec = _pkgSpec
    r.Set("pkg_spec", _pkgSpec)
    return nil
}

// PkgSpec Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPkgSpec() string {
    return r._pkgSpec
}
// Page Setter
// 页数
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPage(_page int64) error {
    r._page = _page
    r.Set("page", _page)
    return nil
}

// Page Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPage() int64 {
    return r._page
}
// PageSize Setter
// 条数
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetPageSize() int64 {
    return r._pageSize
}
// ClientType Setter
// 客户端
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetClientType(_clientType string) error {
    r._clientType = _clientType
    r.Set("client_type", _clientType)
    return nil
}

// ClientType Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetClientType() string {
    return r._clientType
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytCodeprocessRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytCodeprocessRequest) GetRefEntId() string {
    return r._refEntId
}
