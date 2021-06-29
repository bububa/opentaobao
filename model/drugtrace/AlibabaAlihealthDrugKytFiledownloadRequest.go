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
type AlibabaAlihealthDrugKytFiledownloadRequest struct {
    model.Params
    // 企业ID
    refUserId   string
    // 文件地址
    url   string
    // 单据类型
    billType   string
    // 单据队列ID
    billQueueId   string
}

// 初始化AlibabaAlihealthDrugKytFiledownloadRequest对象
func NewAlibabaAlihealthDrugKytFiledownloadRequest() *AlibabaAlihealthDrugKytFiledownloadRequest{
    return &AlibabaAlihealthDrugKytFiledownloadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.filedownload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefUserId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetRefUserId() string {
    return r.refUserId
}
// Url Setter
// 文件地址
func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

// Url Getter
func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetUrl() string {
    return r.url
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetBillType() string {
    return r.billType
}
// BillQueueId Setter
// 单据队列ID
func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetBillQueueId(billQueueId string) error {
    r.billQueueId = billQueueId
    r.Set("bill_queue_id", billQueueId)
    return nil
}

// BillQueueId Getter
func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetBillQueueId() string {
    return r.billQueueId
}
