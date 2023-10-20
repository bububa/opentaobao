package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapicturesdengtaimsorderstatuschangeAPIRequest 天下秀回传订单执行状态变动 API请求
// alibaba.pictures.dengta.ims.order.status.change
//
// 天下秀回传订单执行状态变动
type AlibabapicturesdengtaimsorderstatuschangeAPIRequest struct {
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

// NewAlibabapicturesdengtaimsorderstatuschangeRequest 初始化AlibabapicturesdengtaimsorderstatuschangeAPIRequest对象
func NewAlibabapicturesdengtaimsorderstatuschangeRequest() *AlibabapicturesdengtaimsorderstatuschangeAPIRequest {
	return &AlibabapicturesdengtaimsorderstatuschangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.ims.order.status.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetChangeTime is ChangeTime Setter
// 状态发生的时间 2020-01-02 10:02:03
func (r *AlibabapicturesdengtaimsorderstatuschangeAPIRequest) SetChangeTime(_changeTime string) error {
	r._changeTime = _changeTime
	r.Set("change_time", _changeTime)
	return nil
}

// GetChangeTime ChangeTime Getter
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetChangeTime() string {
	return r._changeTime
}

// SetComments is Comments Setter
// 描述，如ims关单，返回关单原因。
func (r *AlibabapicturesdengtaimsorderstatuschangeAPIRequest) SetComments(_comments string) error {
	r._comments = _comments
	r.Set("comments", _comments)
	return nil
}

// GetComments Comments Getter
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetComments() string {
	return r._comments
}

// SetImsOrderId is ImsOrderId Setter
// 天下秀订单id
func (r *AlibabapicturesdengtaimsorderstatuschangeAPIRequest) SetImsOrderId(_imsOrderId string) error {
	r._imsOrderId = _imsOrderId
	r.Set("ims_order_id", _imsOrderId)
	return nil
}

// GetImsOrderId ImsOrderId Getter
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetImsOrderId() string {
	return r._imsOrderId
}

// SetExtJson is ExtJson Setter
// 扩展字段
func (r *AlibabapicturesdengtaimsorderstatuschangeAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetAccountType is AccountType Setter
// 3=抖音，1-微博 2-微信
func (r *AlibabapicturesdengtaimsorderstatuschangeAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetAccountType() int64 {
	return r._accountType
}

// SetStatus is Status Setter
// 1:待执行  2:执行中  3:发布  4:完成  5:取消
func (r *AlibabapicturesdengtaimsorderstatuschangeAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r AlibabapicturesdengtaimsorderstatuschangeAPIRequest) GetStatus() int64 {
	return r._status
}
