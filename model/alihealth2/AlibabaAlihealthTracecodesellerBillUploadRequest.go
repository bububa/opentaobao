package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传入出库单api API请求
alibaba.alihealth.tracecodeseller.bill.upload

上传入出库单api
*/
type AlibabaAlihealthTracecodesellerBillUploadRequest struct {
    model.Params
    // 身份认证
    skeyCode   string
    // 商家id
    entInfoId   int64
    // 单据编号
    billCode   string
    // 出入库类型
    type   string
    // 出入库时间 精确到 年 月 日 时 分 秒
    time   string
    // 自己仓库id
    warehouseId   int64
    // 渠道商id
    entMerchantId   int64
    // 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
    codeInfo   string
}

// 初始化AlibabaAlihealthTracecodesellerBillUploadRequest对象
func NewAlibabaAlihealthTracecodesellerBillUploadRequest() *AlibabaAlihealthTracecodesellerBillUploadRequest{
    return &AlibabaAlihealthTracecodesellerBillUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.bill.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SkeyCode Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetSkeyCode(skeyCode string) error {
    r.skeyCode = skeyCode
    r.Set("skey_code", skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetSkeyCode() string {
    return r.skeyCode
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetEntInfoId(entInfoId int64) error {
    r.entInfoId = entInfoId
    r.Set("ent_info_id", entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetEntInfoId() int64 {
    return r.entInfoId
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetBillCode() string {
    return r.billCode
}
// Type Setter
// 出入库类型
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetType(type string) error {
    r.type = type
    r.Set("type", type)
    return nil
}

// Type Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetType() string {
    return r.type
}
// Time Setter
// 出入库时间 精确到 年 月 日 时 分 秒
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetTime(time string) error {
    r.time = time
    r.Set("time", time)
    return nil
}

// Time Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetTime() string {
    return r.time
}
// WarehouseId Setter
// 自己仓库id
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetWarehouseId(warehouseId int64) error {
    r.warehouseId = warehouseId
    r.Set("warehouse_id", warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetWarehouseId() int64 {
    return r.warehouseId
}
// EntMerchantId Setter
// 渠道商id
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetEntMerchantId(entMerchantId int64) error {
    r.entMerchantId = entMerchantId
    r.Set("ent_merchant_id", entMerchantId)
    return nil
}

// EntMerchantId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetEntMerchantId() int64 {
    return r.entMerchantId
}
// CodeInfo Setter
// 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetCodeInfo(codeInfo string) error {
    r.codeInfo = codeInfo
    r.Set("code_info", codeInfo)
    return nil
}

// CodeInfo Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetCodeInfo() string {
    return r.codeInfo
}
