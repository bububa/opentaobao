package lstlogistics2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsttradesellerofflineorderuploadAPIRequest 供应商-线下订单-导入接口 API请求
// alibaba.lst.trade.seller.offline.order.upload
//
// 供应商线下订单数据上传、实现和零售通本地云仓订单的共配
type AlibabalsttradesellerofflineorderuploadAPIRequest struct {
	model.Params
	// 入参
	_offlineOrderUploadParam *LstOffLineOrderUploadParam
}

// NewAlibabalsttradesellerofflineorderuploadRequest 初始化AlibabalsttradesellerofflineorderuploadAPIRequest对象
func NewAlibabalsttradesellerofflineorderuploadRequest() *AlibabalsttradesellerofflineorderuploadAPIRequest {
	return &AlibabalsttradesellerofflineorderuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsttradesellerofflineorderuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.trade.seller.offline.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsttradesellerofflineorderuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsttradesellerofflineorderuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOfflineOrderUploadParam is OfflineOrderUploadParam Setter
// 入参
func (r *AlibabalsttradesellerofflineorderuploadAPIRequest) SetOfflineOrderUploadParam(_offlineOrderUploadParam *LstOffLineOrderUploadParam) error {
	r._offlineOrderUploadParam = _offlineOrderUploadParam
	r.Set("offline_order_upload_param", _offlineOrderUploadParam)
	return nil
}

// GetOfflineOrderUploadParam OfflineOrderUploadParam Getter
func (r AlibabalsttradesellerofflineorderuploadAPIRequest) GetOfflineOrderUploadParam() *LstOffLineOrderUploadParam {
	return r._offlineOrderUploadParam
}
