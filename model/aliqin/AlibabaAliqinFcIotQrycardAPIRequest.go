package aliqin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFcIotQrycardAPIRequest 查询终端信息 API请求
// alibaba.aliqin.fc.iot.qrycard
//
// 查询终端信息
type AlibabaAliqinFcIotQrycardAPIRequest struct {
	model.Params
	// 外部计费来源
	_billSource string
	// 外部计费号
	_billReal string
	// ICCID
	_iccid string
}

// NewAlibabaAliqinFcIotQrycardRequest 初始化AlibabaAliqinFcIotQrycardAPIRequest对象
func NewAlibabaAliqinFcIotQrycardRequest() *AlibabaAliqinFcIotQrycardAPIRequest {
	return &AlibabaAliqinFcIotQrycardAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAliqinFcIotQrycardAPIRequest) Reset() {
	r._billSource = ""
	r._billReal = ""
	r._iccid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.qrycard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBillSource is BillSource Setter
// 外部计费来源
func (r *AlibabaAliqinFcIotQrycardAPIRequest) SetBillSource(_billSource string) error {
	r._billSource = _billSource
	r.Set("bill_source", _billSource)
	return nil
}

// GetBillSource BillSource Getter
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetBillSource() string {
	return r._billSource
}

// SetBillReal is BillReal Setter
// 外部计费号
func (r *AlibabaAliqinFcIotQrycardAPIRequest) SetBillReal(_billReal string) error {
	r._billReal = _billReal
	r.Set("bill_real", _billReal)
	return nil
}

// GetBillReal BillReal Getter
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetBillReal() string {
	return r._billReal
}

// SetIccid is Iccid Setter
// ICCID
func (r *AlibabaAliqinFcIotQrycardAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// GetIccid Iccid Getter
func (r AlibabaAliqinFcIotQrycardAPIRequest) GetIccid() string {
	return r._iccid
}

var poolAlibabaAliqinFcIotQrycardAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAliqinFcIotQrycardRequest()
	},
}

// GetAlibabaAliqinFcIotQrycardRequest 从 sync.Pool 获取 AlibabaAliqinFcIotQrycardAPIRequest
func GetAlibabaAliqinFcIotQrycardAPIRequest() *AlibabaAliqinFcIotQrycardAPIRequest {
	return poolAlibabaAliqinFcIotQrycardAPIRequest.Get().(*AlibabaAliqinFcIotQrycardAPIRequest)
}

// ReleaseAlibabaAliqinFcIotQrycardAPIRequest 将 AlibabaAliqinFcIotQrycardAPIRequest 放入 sync.Pool
func ReleaseAlibabaAliqinFcIotQrycardAPIRequest(v *AlibabaAliqinFcIotQrycardAPIRequest) {
	v.Reset()
	poolAlibabaAliqinFcIotQrycardAPIRequest.Put(v)
}
