package qianniu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQianniuNumberPutAPIRequest ISV上传数据接口 API请求
// taobao.qianniu.number.put
//
// ISV提供给卖家使用的业务数据，需要通过这个接口上传到千牛数据中心。
type TaobaoQianniuNumberPutAPIRequest struct {
	model.Params
	// 考虑到稳定性，建议一次卖家最多为200个。标准json格式的数组构成的字符串。每个元素为{user_id:****,field:"****",value:"****"}分别是用户的userid，数据的名称，以及数据的值。
	_data string
}

// NewTaobaoQianniuNumberPutRequest 初始化TaobaoQianniuNumberPutAPIRequest对象
func NewTaobaoQianniuNumberPutRequest() *TaobaoQianniuNumberPutAPIRequest {
	return &TaobaoQianniuNumberPutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQianniuNumberPutAPIRequest) GetApiMethodName() string {
	return "taobao.qianniu.number.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQianniuNumberPutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQianniuNumberPutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetData is Data Setter
// 考虑到稳定性，建议一次卖家最多为200个。标准json格式的数组构成的字符串。每个元素为{user_id:****,field:&#34;****&#34;,value:&#34;****&#34;}分别是用户的userid，数据的名称，以及数据的值。
func (r *TaobaoQianniuNumberPutAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaoQianniuNumberPutAPIRequest) GetData() string {
	return r._data
}
