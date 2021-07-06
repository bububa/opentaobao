package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelOrderFutureInfoPutAPIRequest 订单信息上传更新 API请求
// taobao.xhotel.order.future.info.put
//
// 商家调用推送信息给飞猪平台。 支持如下操作类型：21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
type TaobaoXhotelOrderFutureInfoPutAPIRequest struct {
	model.Params
	// 商家请求流水号
	_outUuid string
	// 酒店编码
	_hotelCode string
	// 字段详细介绍参见 https://open.alitrip.com/docs/doc.htm?docType=1&articleId=106153
	_context string
	// 商家vendor信息。具体值咨询淘宝技术
	_vendor string
	// 请求流水号
	_requestId string
	// 操作类型 21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
	_operateType int64
}

// NewTaobaoXhotelOrderFutureInfoPutRequest 初始化TaobaoXhotelOrderFutureInfoPutAPIRequest对象
func NewTaobaoXhotelOrderFutureInfoPutRequest() *TaobaoXhotelOrderFutureInfoPutAPIRequest {
	return &TaobaoXhotelOrderFutureInfoPutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.order.future.info.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOutUuid is OutUuid Setter
// 商家请求流水号
func (r *TaobaoXhotelOrderFutureInfoPutAPIRequest) SetOutUuid(_outUuid string) error {
	r._outUuid = _outUuid
	r.Set("out_uuid", _outUuid)
	return nil
}

// GetOutUuid OutUuid Getter
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetOutUuid() string {
	return r._outUuid
}

// SetHotelCode is HotelCode Setter
// 酒店编码
func (r *TaobaoXhotelOrderFutureInfoPutAPIRequest) SetHotelCode(_hotelCode string) error {
	r._hotelCode = _hotelCode
	r.Set("hotel_code", _hotelCode)
	return nil
}

// GetHotelCode HotelCode Getter
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetHotelCode() string {
	return r._hotelCode
}

// SetContext is Context Setter
// 字段详细介绍参见 https://open.alitrip.com/docs/doc.htm?docType=1&articleId=106153
func (r *TaobaoXhotelOrderFutureInfoPutAPIRequest) SetContext(_context string) error {
	r._context = _context
	r.Set("context", _context)
	return nil
}

// GetContext Context Getter
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetContext() string {
	return r._context
}

// SetVendor is Vendor Setter
// 商家vendor信息。具体值咨询淘宝技术
func (r *TaobaoXhotelOrderFutureInfoPutAPIRequest) SetVendor(_vendor string) error {
	r._vendor = _vendor
	r.Set("vendor", _vendor)
	return nil
}

// GetVendor Vendor Getter
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetVendor() string {
	return r._vendor
}

// SetRequestId is RequestId Setter
// 请求流水号
func (r *TaobaoXhotelOrderFutureInfoPutAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetOperateType is OperateType Setter
// 操作类型 21: 订单状态更新（商家推送订单状态变更）23：酒店房态信息上传（上传一段时间内的酒店房态）25：在线开发票请求确认 26：自助选房请求进行请求确认   27：自助checkIn请求进行请求确认 32: 扫脸入住入住信息回传 （飞猪将登记至公安系统）
func (r *TaobaoXhotelOrderFutureInfoPutAPIRequest) SetOperateType(_operateType int64) error {
	r._operateType = _operateType
	r.Set("operate_type", _operateType)
	return nil
}

// GetOperateType OperateType Getter
func (r TaobaoXhotelOrderFutureInfoPutAPIRequest) GetOperateType() int64 {
	return r._operateType
}
