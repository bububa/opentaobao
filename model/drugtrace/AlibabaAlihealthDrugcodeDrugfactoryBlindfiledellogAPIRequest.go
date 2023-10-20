package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest 接收盲底文件删除日志 API请求
// alibaba.alihealth.drugcode.drugfactory.blindfiledellog
//
// 临床用药试验-接收盲底文件删除日志
type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest struct {
	model.Params
	// 药厂企业id
	_refEntId string
	// 盲底文件名称，多个盲底文件用,分隔
	_blindFileNames string
	// 操作人
	_operator string
	// 盲底文件删除时间
	_blindFileDeleteTime string
}

// NewAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest 初始化AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest对象
func NewAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest() *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest {
	return &AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) Reset() {
	r._refEntId = ""
	r._blindFileNames = ""
	r._operator = ""
	r._blindFileDeleteTime = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.blindfiledellog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 药厂企业id
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBlindFileNames is BlindFileNames Setter
// 盲底文件名称，多个盲底文件用,分隔
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetBlindFileNames(_blindFileNames string) error {
	r._blindFileNames = _blindFileNames
	r.Set("blind_file_names", _blindFileNames)
	return nil
}

// GetBlindFileNames BlindFileNames Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetBlindFileNames() string {
	return r._blindFileNames
}

// SetOperator is Operator Setter
// 操作人
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetOperator() string {
	return r._operator
}

// SetBlindFileDeleteTime is BlindFileDeleteTime Setter
// 盲底文件删除时间
func (r *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) SetBlindFileDeleteTime(_blindFileDeleteTime string) error {
	r._blindFileDeleteTime = _blindFileDeleteTime
	r.Set("blind_file_delete_time", _blindFileDeleteTime)
	return nil
}

// GetBlindFileDeleteTime BlindFileDeleteTime Getter
func (r AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) GetBlindFileDeleteTime() string {
	return r._blindFileDeleteTime
}

var poolAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest()
	},
}

// GetAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogRequest 从 sync.Pool 获取 AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest
func GetAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest() *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest {
	return poolAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest.Get().(*AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest)
}

// ReleaseAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest 将 AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest(v *AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest.Put(v)
}
