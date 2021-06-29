package dengta

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
接收发生变化的抖音帐号 API请求
alibaba.pictures.dengta.ims.douyin.account.changed

接收发生变化的抖音帐号
*/
type AlibabaPicturesDengtaImsDouyinAccountChangedRequest struct {
    model.Params
    // 天下秀账号ID列表，多个用逗号分隔
    _accountIds   string
    // 3=抖音，1-微博 2-微信
    _accountType   int64
    // 1  下架 2  账号变更（包括账号上架）
    _changeType   int64
}

// 初始化AlibabaPicturesDengtaImsDouyinAccountChangedRequest对象
func NewAlibabaPicturesDengtaImsDouyinAccountChangedRequest() *AlibabaPicturesDengtaImsDouyinAccountChangedRequest{
    return &AlibabaPicturesDengtaImsDouyinAccountChangedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaPicturesDengtaImsDouyinAccountChangedRequest) GetApiMethodName() string {
    return "alibaba.pictures.dengta.ims.douyin.account.changed"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaPicturesDengtaImsDouyinAccountChangedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AccountIds Setter
// 天下秀账号ID列表，多个用逗号分隔
func (r *AlibabaPicturesDengtaImsDouyinAccountChangedRequest) SetAccountIds(_accountIds string) error {
    r._accountIds = _accountIds
    r.Set("account_ids", _accountIds)
    return nil
}

// AccountIds Getter
func (r AlibabaPicturesDengtaImsDouyinAccountChangedRequest) GetAccountIds() string {
    return r._accountIds
}
// AccountType Setter
// 3=抖音，1-微博 2-微信
func (r *AlibabaPicturesDengtaImsDouyinAccountChangedRequest) SetAccountType(_accountType int64) error {
    r._accountType = _accountType
    r.Set("account_type", _accountType)
    return nil
}

// AccountType Getter
func (r AlibabaPicturesDengtaImsDouyinAccountChangedRequest) GetAccountType() int64 {
    return r._accountType
}
// ChangeType Setter
// 1  下架 2  账号变更（包括账号上架）
func (r *AlibabaPicturesDengtaImsDouyinAccountChangedRequest) SetChangeType(_changeType int64) error {
    r._changeType = _changeType
    r.Set("change_type", _changeType)
    return nil
}

// ChangeType Getter
func (r AlibabaPicturesDengtaImsDouyinAccountChangedRequest) GetChangeType() int64 {
    return r._changeType
}
