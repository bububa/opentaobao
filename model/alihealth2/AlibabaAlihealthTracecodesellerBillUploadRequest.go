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
    _skeyCode   string
    // 商家id
    _entInfoId   int64
    // 单据编号
    _billCode   string
    // 出入库类型
    _type   string
    // 出入库时间 精确到 年 月 日 时 分 秒
    _time   string
    // 自己仓库id
    _warehouseId   int64
    // 渠道商id
    _entMerchantId   int64
    // 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
    _codeInfo   string
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
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetSkeyCode(_skeyCode string) error {
    r._skeyCode = _skeyCode
    r.Set("skey_code", _skeyCode)
    return nil
}

// SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetSkeyCode() string {
    return r._skeyCode
}
// EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetEntInfoId(_entInfoId int64) error {
    r._entInfoId = _entInfoId
    r.Set("ent_info_id", _entInfoId)
    return nil
}

// EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetEntInfoId() int64 {
    return r._entInfoId
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetBillCode() string {
    return r._billCode
}
// Type Setter
// 出入库类型
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetType(_type string) error {
    r._type = _type
    r.Set("type", _type)
    return nil
}

// Type Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetType() string {
    return r._type
}
// Time Setter
// 出入库时间 精确到 年 月 日 时 分 秒
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetTime(_time string) error {
    r._time = _time
    r.Set("time", _time)
    return nil
}

// Time Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetTime() string {
    return r._time
}
// WarehouseId Setter
// 自己仓库id
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetWarehouseId(_warehouseId int64) error {
    r._warehouseId = _warehouseId
    r.Set("warehouse_id", _warehouseId)
    return nil
}

// WarehouseId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetWarehouseId() int64 {
    return r._warehouseId
}
// EntMerchantId Setter
// 渠道商id
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetEntMerchantId(_entMerchantId int64) error {
    r._entMerchantId = _entMerchantId
    r.Set("ent_merchant_id", _entMerchantId)
    return nil
}

// EntMerchantId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetEntMerchantId() int64 {
    return r._entMerchantId
}
// CodeInfo Setter
// 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
func (r *AlibabaAlihealthTracecodesellerBillUploadRequest) SetCodeInfo(_codeInfo string) error {
    r._codeInfo = _codeInfo
    r.Set("code_info", _codeInfo)
    return nil
}

// CodeInfo Getter
func (r AlibabaAlihealthTracecodesellerBillUploadRequest) GetCodeInfo() string {
    return r._codeInfo
}
