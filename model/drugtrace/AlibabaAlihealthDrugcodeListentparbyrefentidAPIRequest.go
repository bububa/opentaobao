package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest 根据企业id获取往来单位 API请求
// alibaba.alihealth.drugcode.listentparbyrefentid
//
// 根据企业id获取往来单位
type AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest struct {
	model.Params
	// 登录企业refEntId
	_refEntId string
	// 企业名称
	_entName string
	// 往来单位id
	_partnerIdb string
	// 页数
	_page int64
	// 每页条数
	_pageSize int64
}

// NewAlibabaAlihealthDrugcodeListentparbyrefentidRequest 初始化AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest对象
func NewAlibabaAlihealthDrugcodeListentparbyrefentidRequest() *AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest {
	return &AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.listentparbyrefentid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 登录企业refEntId
func (r *AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetEntName() string {
	return r._entName
}

// SetPartnerIdb is PartnerIdb Setter
// 往来单位id
func (r *AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) SetPartnerIdb(_partnerIdb string) error {
	r._partnerIdb = _partnerIdb
	r.Set("partner_idb", _partnerIdb)
	return nil
}

// GetPartnerIdb PartnerIdb Getter
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetPartnerIdb() string {
	return r._partnerIdb
}

// SetPage is Page Setter
// 页数
func (r *AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugcodeListentparbyrefentidAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
