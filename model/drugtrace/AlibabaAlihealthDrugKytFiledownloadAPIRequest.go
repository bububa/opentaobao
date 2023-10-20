package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugKytFiledownloadAPIRequest 处理失败单据下载 API请求
// alibaba.alihealth.drug.kyt.filedownload
//
// 处理失败单据下载
type AlibabaAlihealthDrugKytFiledownloadAPIRequest struct {
	model.Params
	// 企业ID
	_refUserId string
	// 文件地址
	_url string
	// 单据类型
	_billType string
	// 单据队列ID
	_billQueueId string
}

// NewAlibabaAlihealthDrugKytFiledownloadRequest 初始化AlibabaAlihealthDrugKytFiledownloadAPIRequest对象
func NewAlibabaAlihealthDrugKytFiledownloadRequest() *AlibabaAlihealthDrugKytFiledownloadAPIRequest {
	return &AlibabaAlihealthDrugKytFiledownloadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) Reset() {
	r._refUserId = ""
	r._url = ""
	r._billType = ""
	r._billQueueId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.filedownload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefUserId is RefUserId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetUrl is Url Setter
// 文件地址
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetUrl() string {
	return r._url
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetBillType() string {
	return r._billType
}

// SetBillQueueId is BillQueueId Setter
// 单据队列ID
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetBillQueueId(_billQueueId string) error {
	r._billQueueId = _billQueueId
	r.Set("bill_queue_id", _billQueueId)
	return nil
}

// GetBillQueueId BillQueueId Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetBillQueueId() string {
	return r._billQueueId
}

var poolAlibabaAlihealthDrugKytFiledownloadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugKytFiledownloadRequest()
	},
}

// GetAlibabaAlihealthDrugKytFiledownloadRequest 从 sync.Pool 获取 AlibabaAlihealthDrugKytFiledownloadAPIRequest
func GetAlibabaAlihealthDrugKytFiledownloadAPIRequest() *AlibabaAlihealthDrugKytFiledownloadAPIRequest {
	return poolAlibabaAlihealthDrugKytFiledownloadAPIRequest.Get().(*AlibabaAlihealthDrugKytFiledownloadAPIRequest)
}

// ReleaseAlibabaAlihealthDrugKytFiledownloadAPIRequest 将 AlibabaAlihealthDrugKytFiledownloadAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugKytFiledownloadAPIRequest(v *AlibabaAlihealthDrugKytFiledownloadAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugKytFiledownloadAPIRequest.Put(v)
}
