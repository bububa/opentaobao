package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处理失败单据下载 API请求
alibaba.alihealth.drug.kyt.filedownload

处理失败单据下载
*/
type AlibabaAlihealthDrugKytFiledownloadAPIRequest struct {
    model.Params
    // 企业ID
    _refUserId   string
    // 文件地址
    _url   string
    // 单据类型
    _billType   string
    // 单据队列ID
    _billQueueId   string
}

// 初始化AlibabaAlihealthDrugKytFiledownloadAPIRequest对象
func NewAlibabaAlihealthDrugKytFiledownloadRequest() *AlibabaAlihealthDrugKytFiledownloadAPIRequest{
    return &AlibabaAlihealthDrugKytFiledownloadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.filedownload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefUserId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetRefUserId() string {
    return r._refUserId
}
// Url Setter
// 文件地址
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetUrl(_url string) error {
    r._url = _url
    r.Set("url", _url)
    return nil
}

// Url Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetUrl() string {
    return r._url
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetBillType() string {
    return r._billType
}
// BillQueueId Setter
// 单据队列ID
func (r *AlibabaAlihealthDrugKytFiledownloadAPIRequest) SetBillQueueId(_billQueueId string) error {
    r._billQueueId = _billQueueId
    r.Set("bill_queue_id", _billQueueId)
    return nil
}

// BillQueueId Getter
func (r AlibabaAlihealthDrugKytFiledownloadAPIRequest) GetBillQueueId() string {
    return r._billQueueId
}
