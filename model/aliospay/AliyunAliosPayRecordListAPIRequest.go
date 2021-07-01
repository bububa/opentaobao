package aliospay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliyunAliosPayRecordListAPIRequest
支付记录批量查询接口 API请求
aliyun.alios.pay.record.list

商户用来对账的接口 */
type AliyunAliosPayRecordListAPIRequest struct {
	model.Params
	// 请求参数
	_searchRecordRequest *SearchRecordRequest
}

// New
