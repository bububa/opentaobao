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

// New
