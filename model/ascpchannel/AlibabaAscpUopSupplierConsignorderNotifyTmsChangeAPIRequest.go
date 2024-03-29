package ascpchannel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest 商家修改运单号 API请求
// alibaba.ascp.uop.supplier.consignorder.notify.tms.change
//
// 供应商可以通过此接口，对出库回告上报的运单号进行修改，目前一次调用只能支持一个运单号的修改
type AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest struct {
	model.Params
	// 修改运单号请求模型
	_modifyMailNoRequest *Modifymailnorequest
}

// NewAlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest 初始化AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest对象
func NewAlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest() *AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest {
	return &AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest) Reset() {
	r._modifyMailNoRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.ascp.uop.supplier.consignorder.notify.tms.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetModifyMailNoRequest is ModifyMailNoRequest Setter
// 修改运单号请求模型
func (r *AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest) SetModifyMailNoRequest(_modifyMailNoRequest *Modifymailnorequest) error {
	r._modifyMailNoRequest = _modifyMailNoRequest
	r.Set("modify_mail_no_request", _modifyMailNoRequest)
	return nil
}

// GetModifyMailNoRequest ModifyMailNoRequest Getter
func (r AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest) GetModifyMailNoRequest() *Modifymailnorequest {
	return r._modifyMailNoRequest
}

var poolAlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest()
	},
}

// GetAlibabaAscpUopSupplierConsignorderNotifyTmsChangeRequest 从 sync.Pool 获取 AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest
func GetAlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest() *AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest {
	return poolAlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest.Get().(*AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest)
}

// ReleaseAlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest 将 AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest(v *AlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest) {
	v.Reset()
	poolAlibabaAscpUopSupplierConsignorderNotifyTmsChangeAPIRequest.Put(v)
}
