package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopidlistconverterAPIRequest 将淘宝openId或者设备id/用户id转换为openId API请求
// taobao.ailab.aicloud.top.id.list.converter
//
// 将淘宝openId或者设备id/用户id转换为openId
type TaobaoailabaicloudtopidlistconverterAPIRequest struct {
	model.Params
	// 入参数据
	_convertData *ConverterIdRequest
}

// NewTaobaoailabaicloudtopidlistconverterRequest 初始化TaobaoailabaicloudtopidlistconverterAPIRequest对象
func NewTaobaoailabaicloudtopidlistconverterRequest() *TaobaoailabaicloudtopidlistconverterAPIRequest {
	return &TaobaoailabaicloudtopidlistconverterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopidlistconverterAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.id.list.converter"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopidlistconverterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopidlistconverterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetConvertData is ConvertData Setter
// 入参数据
func (r *TaobaoailabaicloudtopidlistconverterAPIRequest) SetConvertData(_convertData *ConverterIdRequest) error {
	r._convertData = _convertData
	r.Set("convert_data", _convertData)
	return nil
}

// GetConvertData ConvertData Getter
func (r TaobaoailabaicloudtopidlistconverterAPIRequest) GetConvertData() *ConverterIdRequest {
	return r._convertData
}
