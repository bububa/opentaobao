package xhotelonlineorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelpmsguestbillgetvtwoAPIRequest 客人PMS账单信息查询 API请求
// taobao.xhotel.pms.guestbill.get.vtwo
//
// 从pms获取客人账单信息
type TaobaoxhotelpmsguestbillgetvtwoAPIRequest struct {
	model.Params
	// 开票点税号
	_taxNum string
	// 身份证后4位
	_shortIdNum string
	// 房间号
	_roomNum string
	// 请求id (32位唯一值)
	_requestId string
	// 用户渠道(0:未知,1:淘宝)
	_userChannel int64
}

// NewTaobaoxhotelpmsguestbillgetvtwoRequest 初始化TaobaoxhotelpmsguestbillgetvtwoAPIRequest对象
func NewTaobaoxhotelpmsguestbillgetvtwoRequest() *TaobaoxhotelpmsguestbillgetvtwoAPIRequest {
	return &TaobaoxhotelpmsguestbillgetvtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.pms.guestbill.get.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaxNum is TaxNum Setter
// 开票点税号
func (r *TaobaoxhotelpmsguestbillgetvtwoAPIRequest) SetTaxNum(_taxNum string) error {
	r._taxNum = _taxNum
	r.Set("tax_num", _taxNum)
	return nil
}

// GetTaxNum TaxNum Getter
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetTaxNum() string {
	return r._taxNum
}

// SetShortIdNum is ShortIdNum Setter
// 身份证后4位
func (r *TaobaoxhotelpmsguestbillgetvtwoAPIRequest) SetShortIdNum(_shortIdNum string) error {
	r._shortIdNum = _shortIdNum
	r.Set("short_id_num", _shortIdNum)
	return nil
}

// GetShortIdNum ShortIdNum Getter
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetShortIdNum() string {
	return r._shortIdNum
}

// SetRoomNum is RoomNum Setter
// 房间号
func (r *TaobaoxhotelpmsguestbillgetvtwoAPIRequest) SetRoomNum(_roomNum string) error {
	r._roomNum = _roomNum
	r.Set("room_num", _roomNum)
	return nil
}

// GetRoomNum RoomNum Getter
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetRoomNum() string {
	return r._roomNum
}

// SetRequestId is RequestId Setter
// 请求id (32位唯一值)
func (r *TaobaoxhotelpmsguestbillgetvtwoAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetUserChannel is UserChannel Setter
// 用户渠道(0:未知,1:淘宝)
func (r *TaobaoxhotelpmsguestbillgetvtwoAPIRequest) SetUserChannel(_userChannel int64) error {
	r._userChannel = _userChannel
	r.Set("user_channel", _userChannel)
	return nil
}

// GetUserChannel UserChannel Getter
func (r TaobaoxhotelpmsguestbillgetvtwoAPIRequest) GetUserChannel() int64 {
	return r._userChannel
}
