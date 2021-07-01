package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDapHuaweiCardinfosAPIRequest
华为负一屏卡片查询 API请求
alibaba.alihealth.dap.huawei.cardinfos

医疗健康频道卡片华为负一屏 */
type AlibabaAlihealthDapHuaweiCardinfosAPIRequest struct {
	model.Params
	// source     HUAWEI_HAG,OPPO_OAG
	_param string
}

// New
