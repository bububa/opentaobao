package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest 药品全量数据下载 API请求
// alibaba.alihealth.drug.kyt.getdruginfo.downloadurl
//
// 药品全量数据下载
type AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest struct {
	model.Params
	// 调用接口的企业ID
	_refEntId string
}

// NewAlibabaAlihealthDrugKytGetdruginfoDownloadurlRequest 初始化AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest对象
func NewAlibabaAlihealthDrugKytGetdruginfoDownloadurlRequest() *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest {
	return &AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest) Reset() {
	r._refEntId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.getdruginfo.downloadurl"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 调用接口的企业ID
func (r *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest) GetRefEntId() string {
	return r._refEntId
}

var poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytGetdruginfoDownloadurlRequest()
	},
}

// GetAlibabaAlihealthDrugKytGetdruginfoDownloadurlRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest
func GetAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest() *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest {
	return poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest.Get().(*AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest 将 AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest(v *AlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlAPIRequest.Put(v)
}
