package dengta

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest
接收发生变化的抖音帐号 API请求
alibaba.pictures.dengta.ims.douyin.account.changed

接收发生变化的抖音帐号 */
type AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest struct {
	model.Params
	// 天下秀账号ID列表，多个用逗号分隔
	_accountIds string
	// 3=抖音，1-微博 2-微信
	_accountType int64
	// 1  下架 2  账号变更（包括账号上架）
	_changeType int64
}

// NewAlibabaPicturesDengtaImsDouyinAccountChangedRequest 初始化AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest对象
func NewAlibabaPicturesDengtaImsDouyinAccountChangedRequest() *AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest {
	return &AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) GetApiMethodName() string {
	return "alibaba.pictures.dengta.ims.douyin.account.changed"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AccountIds Setter
// 天下秀账号ID列表，多个用逗号分隔
func (r *AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) SetAccountIds(_accountIds string) error {
	r._accountIds = _accountIds
	r.Set("account_ids", _accountIds)
	return nil
}

// Get AccountIds Getter
func (r AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) GetAccountIds() string {
	return r._accountIds
}

// Set is AccountType Setter
// 3=抖音，1-微博 2-微信
func (r *AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) SetAccountType(_accountType int64) error {
	r._accountType = _accountType
	r.Set("account_type", _accountType)
	return nil
}

// Get AccountType Getter
func (r AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) GetAccountType() int64 {
	return r._accountType
}

// Set is ChangeType Setter
// 1  下架 2  账号变更（包括账号上架）
func (r *AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) SetChangeType(_changeType int64) error {
	r._changeType = _changeType
	r.Set("change_type", _changeType)
	return nil
}

// Get ChangeType Getter
func (r AlibabaPicturesDengtaImsDouyinAccountChangedAPIRequest) GetChangeType() int64 {
	return r._changeType
}
