package dengta

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest 天下秀回传订单执行状态变动 API请求
// alibaba.pictures.dengta.ims.order.status.change
//
// 天下秀回传订单执行状态变动
type AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest struct {
	model.Params
	// 状态发生的时间 2020-01-02 10:02:03
	_changeTime string
	// 描述，如ims关单，返回关单原因。
	_comments string
	// 天下秀订单id
	_imsOrderId string
	// 扩展字段
	_extJson string
	// 3=抖音，1-微博 2-微信
	_accountType int64
	// 1:待执行  2:执行中  3:发布  4:完成  5:取消
	_status int64
}

// NewAlibabaPicturesDengtaImsOrderStatusChangeRequest 初始化AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest对象
func NewAlibabaPicturesDengtaImsOrderStatusChangeRequest() *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest {
	return &AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) Reset() {
	r._changeTime = ""
	r._comments = ""
	r._imsOrderId = ""
	r._extJson = ""
	r._accountType = 0
	r._status = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.ims.order.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeTime is ChangeTime Setter
// 状态发生的时间 2020-01-02 10:02:03
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetComments is Comments Setter
// 描述，如ims关单，返回关单原因。
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetComments(_comments string) error {
	r._comments = _comments
	r.Set("comments", _comments)
	return nil
}

// GetComments Comments Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetComments() string {
	return r._comments
}

// SetImsOrderId is ImsOrderId Setter
// 天下秀订单id
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetImsOrderId(_imsOrderId string) error {
	r._imsOrderId = _imsOrderId
	r.Set("ims_order_id", _imsOrderId)
	return nil
}

// GetImsOrderId ImsOrderId Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetImsOrderId() string {
	return r._imsOrderId
}

// SetExtJson is ExtJson Setter
// 扩展字段
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetAccountType is AccountType Setter
// 3=抖音，1-微博 2-微信
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetAccountType() int64 {
	return r._accountType
}

// SetStatus is Status Setter
// 1:待执行  2:执行中  3:发布  4:完成  5:取消
func (r *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) GetStatus() int64 {
	return r._status
}

var poolAlibabaPicturesDengtaImsOrderStatusChangeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaPicturesDengtaImsOrderStatusChangeRequest()
	},
}

// GetAlibabaPicturesDengtaImsOrderStatusChangeRequest 从 sync.Pool 获取 AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest
func GetAlibabaPicturesDengtaImsOrderStatusChangeAPIRequest() *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest {
	return poolAlibabaPicturesDengtaImsOrderStatusChangeAPIRequest.Get().(*AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest)
}

// ReleaseAlibabaPicturesDengtaImsOrderStatusChangeAPIRequest 将 AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest 放入 sync.Pool
func ReleaseAlibabaPicturesDengtaImsOrderStatusChangeAPIRequest(v *AlibabaPicturesDengtaImsOrderStatusChangeAPIRequest) {
	v.Reset()
	poolAlibabaPicturesDengtaImsOrderStatusChangeAPIRequest.Put(v)
}
