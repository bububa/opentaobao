package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytfiledownloadAPIRequest 处理失败单据下载 API请求
// alibaba.alihealth.drug.kyt.filedownload
//
// 处理失败单据下载
type AlibabaalihealthdrugkytfiledownloadAPIRequest struct {
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

// NewAlibabaalihealthdrugkytfiledownloadRequest 初始化AlibabaalihealthdrugkytfiledownloadAPIRequest对象
func NewAlibabaalihealthdrugkytfiledownloadRequest() *AlibabaalihealthdrugkytfiledownloadAPIRequest {
	return &AlibabaalihealthdrugkytfiledownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytfiledownloadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.filedownload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytfiledownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytfiledownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefUserId is RefUserId Setter
// 企业ID
func (r *AlibabaalihealthdrugkytfiledownloadAPIRequest) SetRefUserId(_refUserId string) error {
	r._refUserId = _refUserId
	r.Set("ref_user_id", _refUserId)
	return nil
}

// GetRefUserId RefUserId Getter
func (r AlibabaalihealthdrugkytfiledownloadAPIRequest) GetRefUserId() string {
	return r._refUserId
}

// SetUrl is Url Setter
// 文件地址
func (r *AlibabaalihealthdrugkytfiledownloadAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r AlibabaalihealthdrugkytfiledownloadAPIRequest) GetUrl() string {
	return r._url
}

// SetBillType is BillType Setter
// 单据类型
func (r *AlibabaalihealthdrugkytfiledownloadAPIRequest) SetBillType(_billType string) error {
	r._billType = _billType
	r.Set("bill_type", _billType)
	return nil
}

// GetBillType BillType Getter
func (r AlibabaalihealthdrugkytfiledownloadAPIRequest) GetBillType() string {
	return r._billType
}

// SetBillQueueId is BillQueueId Setter
// 单据队列ID
func (r *AlibabaalihealthdrugkytfiledownloadAPIRequest) SetBillQueueId(_billQueueId string) error {
	r._billQueueId = _billQueueId
	r.Set("bill_queue_id", _billQueueId)
	return nil
}

// GetBillQueueId BillQueueId Getter
func (r AlibabaalihealthdrugkytfiledownloadAPIRequest) GetBillQueueId() string {
	return r._billQueueId
}
