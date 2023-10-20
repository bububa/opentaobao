package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest 接收发生变化的抖音帐号 API请求
// alibaba.pictures.dengta.ims.douyin.account.changed
//
// 接收发生变化的抖音帐号
type AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest struct {
	model.Params
	// 天下秀账号ID列表，多个用逗号分隔
	_accountIds string
	// 3=抖音，1-微博 2-微信
	_accountType int64
	// 1  下架 2  账号变更（包括账号上架）
	_changeType int64
}

// NewAlibabapicturesdengtaimsdouyinaccountchangedRequest 初始化AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest对象
func NewAlibabapicturesdengtaimsdouyinaccountchangedRequest() *AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest {
	return &AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.ims.douyin.account.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAccountIds is AccountIds Setter
// 天下秀账号ID列表，多个用逗号分隔
func (r *AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) SetAccountIds(_accountIds string) error {
	r._accountIds = _accountIds
	r.Set("account_ids", _accountIds)
	return nil
}

// GetAccountIds AccountIds Getter
func (r AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) GetAccountIds() string {
	return r._accountIds
}

// SetAccountType is AccountType Setter
// 3=抖音，1-微博 2-微信
func (r *AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// GetAccountType AccountType Getter
func (r AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) GetAccountType() int64 {
	return r._accountType
}

// SetChangeType is ChangeType Setter
// 1  下架 2  账号变更（包括账号上架）
func (r *AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) SetChangeType(_changeType int64) error {
	r._changeType = _changeType
	r.Set("change_type", _changeType)
	return nil
}

// GetChangeType ChangeType Getter
func (r AlibabapicturesdengtaimsdouyinaccountchangedAPIRequest) GetChangeType() int64 {
	return r._changeType
}
