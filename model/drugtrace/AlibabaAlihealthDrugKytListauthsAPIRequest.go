package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytListauthsAPIRequest 企业搜索自己授权的物流企业 API请求
// alibaba.alihealth.drug.kyt.listauths
//
// 企业搜索自己授权的物流企业
type AlibabaAlihealthDrugKytListauthsAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业名称
	_entName string
	// 页码
	_page int64
	// 页大小
	_pageSize int64
}

// NewAlibabaAlihealthDrugKytListauthsRequest 初始化AlibabaAlihealthDrugKytListauthsAPIRequest对象
func NewAlibabaAlihealthDrugKytListauthsRequest() *AlibabaAlihealthDrugKytListauthsAPIRequest {
	return &AlibabaAlihealthDrugKytListauthsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytListauthsAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.listauths"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytListauthsAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytListauthsAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytListauthsAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetEntName is EntName Setter
// 企业名称
func (r *AlibabaAlihealthDrugKytListauthsAPIRequest) SetEntName(_entName string) error {
	r._entName = _entName
	r.Set("ent_name", _entName)
	return nil
}

// GetEntName EntName Getter
func (r AlibabaAlihealthDrugKytListauthsAPIRequest) GetEntName() string {
	return r._entName
}

// SetPage is Page Setter
// 页码
func (r *AlibabaAlihealthDrugKytListauthsAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaAlihealthDrugKytListauthsAPIRequest) GetPage() int64 {
	return r._page
}

// SetPageSize is PageSize Setter
// 页大小
func (r *AlibabaAlihealthDrugKytListauthsAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaAlihealthDrugKytListauthsAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
