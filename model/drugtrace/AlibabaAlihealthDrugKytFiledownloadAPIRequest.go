package drugtrace

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.filedownload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
