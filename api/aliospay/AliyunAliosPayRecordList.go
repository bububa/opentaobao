package aliospay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliospay"
)

// AliyunAliosPayRecordList 支付记录批量查询接口
// aliyun.alios.pay.record.list
//
// 商户用来对账的接口
func AliyunAliosPayRecordList(clt *core.SDKClient, req *aliospay.AliyunAliosPayRecordListAPIRequest, resp *aliospay.AliyunAliosPayRecordListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
