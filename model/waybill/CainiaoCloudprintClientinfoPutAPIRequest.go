package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoCloudprintClientinfoPutAPIRequest
云打印客户端监控信息收集 API请求
cainiao.cloudprint.clientinfo.put

云打印客户端监控信息收集 */
type CainiaoCloudprintClientinfoPutAPIRequest struct {
	model.Params
	// 客户端上传json数据
	_jsonData string
}

// New
