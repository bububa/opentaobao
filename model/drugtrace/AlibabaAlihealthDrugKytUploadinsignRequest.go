package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓库批量扫码回传接口 APIRequest
alibaba.alihealth.drug.kyt.uploadinsign

连锁总部仓库在采购入库或者销售出库环节，批量采集追溯码之后回传到码上放心平台。
*/
type AlibabaAlihealthDrugKytUploadinsignRequest struct {
    model.Params

    // 单据编号（小于20位字符串，唯一）
    billCode   string 

    // 单据生成时间
    billTime   string 

    // 码上放心平台企业编码（仓库所有者）
    refUserId   string 

    // 仓库名称（企业自定义）
    warehouseId   string 

    // 药品ID
    drugId   string 

    // 追溯码[多个时用逗号分开]
    traceCodes   []string 

}

func NewAlibabaAlihealthDrugKytUploadinsignRequest() *AlibabaAlihealthDrugKytUploadinsignRequest{
    return &AlibabaAlihealthDrugKytUploadinsignRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadinsign"
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetBillTime(billTime string) error {
    r.billTime = billTime
    r.Set("bill_time", billTime)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetBillTime() string {
    return r.billTime
}

func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetRefUserId(refUserId string) error {
    r.refUserId = refUserId
    r.Set("ref_user_id", refUserId)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetRefUserId() string {
    return r.refUserId
}

func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetWarehouseId(warehouseId string) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetWarehouseId() string {
    return r.warehouseId
}

func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetDrugId() string {
    return r.drugId
}

func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetTraceCodes(traceCodes []string) error {
    r.traceCodes = traceCodes
    r.Set("trace_codes", traceCodes)
    return nil
}

func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetTraceCodes() []string {
    return r.traceCodes
}

