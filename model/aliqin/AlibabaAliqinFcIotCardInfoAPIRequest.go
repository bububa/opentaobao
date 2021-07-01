package aliqin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFcIotCardInfoAPIRequest
物联卡信息查询 API请求
alibaba.aliqin.fc.iot.cardInfo

物联卡信息查询 */
type AlibabaAliqinFcIotCardInfoAPIRequest struct {
	model.Params
	// SIM卡号
	_iccid string
}

// NewAlibabaAliqinFcIotCardInfoRequest 初始化AlibabaAliqinFcIotCardInfoAPIRequest对象
func NewAlibabaAliqinFcIotCardInfoRequest() *AlibabaAliqinFcIotCardInfoAPIRequest {
	return &AlibabaAliqinFcIotCardInfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFcIotCardInfoAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.fc.iot.cardInfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFcIotCardInfoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Iccid Setter
// SIM卡号
func (r *AlibabaAliqinFcIotCardInfoAPIRequest) SetIccid(_iccid string) error {
	r._iccid = _iccid
	r.Set("iccid", _iccid)
	return nil
}

// Get Iccid Getter
func (r AlibabaAliqinFcIotCardInfoAPIRequest) GetIccid() string {
	return r._iccid
}
