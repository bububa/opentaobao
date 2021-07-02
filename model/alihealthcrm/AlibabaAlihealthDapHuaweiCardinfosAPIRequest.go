package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDapHuaweiCardinfosAPIRequest 华为负一屏卡片查询 API请求
// alibaba.alihealth.dap.huawei.cardinfos
//
// 医疗健康频道卡片华为负一屏
type AlibabaAlihealthDapHuaweiCardinfosAPIRequest struct {
	model.Params
	// source     HUAWEI_HAG,OPPO_OAG
	_param string
}

// NewAlibabaAlihealthDapHuaweiCardinfosRequest 初始化AlibabaAlihealthDapHuaweiCardinfosAPIRequest对象
func NewAlibabaAlihealthDapHuaweiCardinfosRequest() *AlibabaAlihealthDapHuaweiCardinfosAPIRequest {
	return &AlibabaAlihealthDapHuaweiCardinfosAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDapHuaweiCardinfosAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dap.huawei.cardinfos"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDapHuaweiCardinfosAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// source     HUAWEI_HAG,OPPO_OAG
func (r *AlibabaAlihealthDapHuaweiCardinfosAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaAlihealthDapHuaweiCardinfosAPIRequest) GetParam() string {
	return r._param
}
