package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLsyCrmActivityStoreGetstorelistAPIRequest
ISV查询门店 API请求
alibaba.lsy.crm.activity.store.getstorelist

ISV查询门店 */
type AlibabaLsyCrmActivityStoreGetstorelistAPIRequest struct {
	model.Params
	// 系统自动生成
	_queryStoreReq *NrtQueryStoreReq
}

// New
