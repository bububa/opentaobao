package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
处理失败单据下载 APIRequest
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

func NewAlibabaAlihealthDrugKytFiledownloadRequest() *AlibabaAlihealthDrugKytFiledownloadRequest{
    return &AlibabaAlihealthDrugKytFiledownloadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.filedownload"
}

func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetUrl(url string) error {
    r.url = url
    r.Set("url", url)
    return nil
}

func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetUrl() string {
    return r.url
}

func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetBillType() string {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytFiledownloadRequest) SetBillQueueId(billQueueId string) error {
    r.billQueueId = billQueueId
    r.Set("bill_queue_id", billQueueId)
    return nil
}

func (r AlibabaAlihealthDrugKytFiledownloadRequest) GetBillQueueId() string {
    return r.billQueueId
}

