package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓库批量扫码回传接口 API请求
alibaba.alihealth.drug.kyt.uploadinsign

连锁总部仓库在采购入库或者销售出库环节，批量采集追溯码之后回传到码上放心平台。
*/
type AlibabaAlihealthDrugKytUploadinsignRequest struct {
    model.Params
    // 单据编号（小于20位字符串，唯一）
    _billCode   string
    // 单据生成时间
    _billTime   string
    // 码上放心平台企业编码（仓库所有者）
    _refUserId   string
    // 仓库名称（企业自定义）
    _warehouseId   string
    // 药品ID
    _drugId   string
    // 追溯码[多个时用逗号分开]
    _traceCodes   []string
}

// 初始化AlibabaAlihealthDrugKytUploadinsignRequest对象
func NewAlibabaAlihealthDrugKytUploadinsignRequest() *AlibabaAlihealthDrugKytUploadinsignRequest{
    return &AlibabaAlihealthDrugKytUploadinsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.uploadinsign"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据编号（小于20位字符串，唯一）
func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetBillCode() string {
    return r._billCode
}
// BillTime Setter
// 单据生成时间
func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetBillTime(_billTime string) error {
    r._billTime = _billTime
    r.Set("bill_time", _billTime)
    return nil
}

// BillTime Getter
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetBillTime() string {
    return r._billTime
}
// RefUserId Setter
// 码上放心平台企业编码（仓库所有者）
func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetRefUserId(_refUserId string) error {
    r._refUserId = _refUserId
    r.Set("ref_user_id", _refUserId)
    return nil
}

// RefUserId Getter
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetRefUserId() string {
    return r._refUserId
}
// WarehouseId Setter
// 仓库名称（企业自定义）
func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetWarehouseId(_warehouseId string) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetWarehouseId() string {
    return r._warehouseId
}
// DrugId Setter
// 药品ID
func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetDrugId(_drugId string) error {
    r._drugId = _drugId
    r.Set("drug_id", _drugId)
    return nil
}

// DrugId Getter
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetDrugId() string {
    return r._drugId
}
// TraceCodes Setter
// 追溯码[多个时用逗号分开]
func (r *AlibabaAlihealthDrugKytUploadinsignRequest) SetTraceCodes(_traceCodes []string) error {
    r._traceCodes = _traceCodes
    r.Set("trace_codes", _traceCodes)
    return nil
}

// TraceCodes Getter
func (r AlibabaAlihealthDrugKytUploadinsignRequest) GetTraceCodes() []string {
    return r._traceCodes
}
