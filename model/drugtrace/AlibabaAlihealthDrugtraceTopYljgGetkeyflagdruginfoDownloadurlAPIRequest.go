package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest 获取重点追溯品种明细下载URL API请求
// alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl
//
// 获取重点追溯品种明细下载URL
type AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
}

// NewAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest 初始化AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest() *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest {
	return &AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest) Reset() {
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugtrace.top.yljg.getkeyflagdruginfo.downloadurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest()
	},
}

// GetAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlRequest 从 sync.Pool 获取 AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest
func GetAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest() *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest {
	return poolAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest.Get().(*AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest)
}

// ReleaseAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest 将 AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest(v *AlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugtraceTopYljgGetkeyflagdruginfoDownloadurlAPIRequest.Put(v)
}
