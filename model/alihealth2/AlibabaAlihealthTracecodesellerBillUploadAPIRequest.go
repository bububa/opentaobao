package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerBillUploadAPIRequest 上传入出库单api API请求
// alibaba.alihealth.tracecodeseller.bill.upload
//
// 上传入出库单api
type AlibabaAlihealthTracecodesellerBillUploadAPIRequest struct {
	model.Params
	// 身份认证
	_skeyCode string
	// 单据编号
	_billCode string
	// 出入库类型
	_type string
	// 出入库时间 精确到 年 月 日 时 分 秒
	_time string
	// 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
	_codeInfo string
	// 商家id
	_entInfoId int64
	// 自己仓库id
	_warehouseId int64
	// 渠道商id
	_entMerchantId int64
}

// NewAlibabaAlihealthTracecodesellerBillUploadRequest 初始化AlibabaAlihealthTracecodesellerBillUploadAPIRequest对象
func NewAlibabaAlihealthTracecodesellerBillUploadRequest() *AlibabaAlihealthTracecodesellerBillUploadAPIRequest {
	return &AlibabaAlihealthTracecodesellerBillUploadAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) Reset() {
	r._skeyCode = ""
	r._billCode = ""
	r._type = ""
	r._time = ""
	r._codeInfo = ""
	r._entInfoId = 0
	r._warehouseId = 0
	r._entMerchantId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.bill.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkeyCode is SkeyCode Setter
// 身份认证
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetSkeyCode(_skeyCode string) error {
	r._skeyCode = _skeyCode
	r.Set("skey_code", _skeyCode)
	return nil
}

// GetSkeyCode SkeyCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetSkeyCode() string {
	return r._skeyCode
}

// SetBillCode is BillCode Setter
// 单据编号
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetBillCode(_billCode string) error {
	r._billCode = _billCode
	r.Set("bill_code", _billCode)
	return nil
}

// GetBillCode BillCode Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetBillCode() string {
	return r._billCode
}

// SetType is Type Setter
// 出入库类型
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetType() string {
	return r._type
}

// SetTime is Time Setter
// 出入库时间 精确到 年 月 日 时 分 秒
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetTime(_time string) error {
	r._time = _time
	r.Set("time", _time)
	return nil
}

// GetTime Time Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetTime() string {
	return r._time
}

// SetCodeInfo is CodeInfo Setter
// 把txt格式的文件转成16进制的字符串，txt文件是每个码一行
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetCodeInfo(_codeInfo string) error {
	r._codeInfo = _codeInfo
	r.Set("code_info", _codeInfo)
	return nil
}

// GetCodeInfo CodeInfo Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetCodeInfo() string {
	return r._codeInfo
}

// SetEntInfoId is EntInfoId Setter
// 商家id
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetEntInfoId(_entInfoId int64) error {
	r._entInfoId = _entInfoId
	r.Set("ent_info_id", _entInfoId)
	return nil
}

// GetEntInfoId EntInfoId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetEntInfoId() int64 {
	return r._entInfoId
}

// SetWarehouseId is WarehouseId Setter
// 自己仓库id
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetWarehouseId(_warehouseId int64) error {
	r._warehouseId = _warehouseId
	r.Set("warehouse_id", _warehouseId)
	return nil
}

// GetWarehouseId WarehouseId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetWarehouseId() int64 {
	return r._warehouseId
}

// SetEntMerchantId is EntMerchantId Setter
// 渠道商id
func (r *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) SetEntMerchantId(_entMerchantId int64) error {
	r._entMerchantId = _entMerchantId
	r.Set("ent_merchant_id", _entMerchantId)
	return nil
}

// GetEntMerchantId EntMerchantId Getter
func (r AlibabaAlihealthTracecodesellerBillUploadAPIRequest) GetEntMerchantId() int64 {
	return r._entMerchantId
}

var poolAlibabaAlihealthTracecodesellerBillUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthTracecodesellerBillUploadRequest()
	},
}

// GetAlibabaAlihealthTracecodesellerBillUploadRequest 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerBillUploadAPIRequest
func GetAlibabaAlihealthTracecodesellerBillUploadAPIRequest() *AlibabaAlihealthTracecodesellerBillUploadAPIRequest {
	return poolAlibabaAlihealthTracecodesellerBillUploadAPIRequest.Get().(*AlibabaAlihealthTracecodesellerBillUploadAPIRequest)
}

// ReleaseAlibabaAlihealthTracecodesellerBillUploadAPIRequest 将 AlibabaAlihealthTracecodesellerBillUploadAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerBillUploadAPIRequest(v *AlibabaAlihealthTracecodesellerBillUploadAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerBillUploadAPIRequest.Put(v)
}
