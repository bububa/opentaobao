package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytSynonymauthsAPIRequest 物流企业查询货主企业信息 API请求
// alibaba.alihealth.drug.kyt.synonymauths
//
// 物流企业查询货主企业信息
type AlibabaAlihealthDrugKytSynonymauthsAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业名称
	_entName string
	// 货主自定义编号
	_synOwnEntId string
	// 页码
	_pageSize int64
	// 页面大小
	_page int64
}

// NewAlibabaAlihealthDrugKytSynonymauthsRequest 初始化AlibabaAlihealthDrugKytSynonymauthsAPIRequest对象
func NewAlibabaAlihealthDrugKytSynonymauthsRequest() *AlibabaAlihealthDrugKytSynonymauthsAPIRequest {
	return &AlibabaAlihealthDrugKytSynonymauthsAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytSynonymauthsAPIRequest) Reset() {
	r._refEntId = ""
	r._entName = ""
	r._synOwnEntId = ""
	r._pageSize = 0
	r._page = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.synonymauths"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytSynonymauthsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytSynonymauthsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetEntName() string {
	return r._entName
}

// SetSynOwnEntId is SynOwnEntId Setter
// 货主自定义编号
func (r *AlibabaAlihealthDrugKytSynonymauthsAPIRequest) SetSynOwnEntId(_synOwnEntId string) error {
	r._synOwnEntId = _synOwnEntId
	r.Set("syn_own_ent_id", _synOwnEntId)
	return nil
}

// GetSynOwnEntId SynOwnEntId Getter
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetSynOwnEntId() string {
	return r._synOwnEntId
}

// SetPageSize is PageSize Setter
// 页码
func (r *AlibabaAlihealthDrugKytSynonymauthsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPage is Page Setter
// 页面大小
func (r *AlibabaAlihealthDrugKytSynonymauthsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytSynonymauthsAPIRequest) GetPage() int64 {
	return r._page
}

var poolAlibabaAlihealthDrugKytSynonymauthsAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytSynonymauthsRequest()
	},
}

// GetAlibabaAlihealthDrugKytSynonymauthsRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytSynonymauthsAPIRequest
func GetAlibabaAlihealthDrugKytSynonymauthsAPIRequest() *AlibabaAlihealthDrugKytSynonymauthsAPIRequest {
	return poolAlibabaAlihealthDrugKytSynonymauthsAPIRequest.Get().(*AlibabaAlihealthDrugKytSynonymauthsAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytSynonymauthsAPIRequest 将 AlibabaAlihealthDrugKytSynonymauthsAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytSynonymauthsAPIRequest(v *AlibabaAlihealthDrugKytSynonymauthsAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytSynonymauthsAPIRequest.Put(v)
}
