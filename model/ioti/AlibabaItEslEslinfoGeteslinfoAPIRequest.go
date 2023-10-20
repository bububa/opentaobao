package ioti

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaItEslEslinfoGeteslinfoAPIRequest 厂测查询价签当前信息 API请求
// alibaba.it.esl.eslinfo.geteslinfo
//
// 工厂对生产出的电子价签进行全流程功能测试，查询价签当前上报的信息
type AlibabaItEslEslinfoGeteslinfoAPIRequest struct {
	model.Params
	// mac
	_mac string
}

// NewAlibabaItEslEslinfoGeteslinfoRequest 初始化AlibabaItEslEslinfoGeteslinfoAPIRequest对象
func NewAlibabaItEslEslinfoGeteslinfoRequest() *AlibabaItEslEslinfoGeteslinfoAPIRequest {
	return &AlibabaItEslEslinfoGeteslinfoAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaItEslEslinfoGeteslinfoAPIRequest) Reset() {
	r._mac = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaItEslEslinfoGeteslinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.it.esl.eslinfo.geteslinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaItEslEslinfoGeteslinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaItEslEslinfoGeteslinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMac is Mac Setter
// mac
func (r *AlibabaItEslEslinfoGeteslinfoAPIRequest) SetMac(_mac string) error {
	r._mac = _mac
	r.Set("mac", _mac)
	return nil
}

// GetMac Mac Getter
func (r AlibabaItEslEslinfoGeteslinfoAPIRequest) GetMac() string {
	return r._mac
}

var poolAlibabaItEslEslinfoGeteslinfoAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaItEslEslinfoGeteslinfoRequest()
	},
}

// GetAlibabaItEslEslinfoGeteslinfoRequest 从 sync.Pool 获取 AlibabaItEslEslinfoGeteslinfoAPIRequest
func GetAlibabaItEslEslinfoGeteslinfoAPIRequest() *AlibabaItEslEslinfoGeteslinfoAPIRequest {
	return poolAlibabaItEslEslinfoGeteslinfoAPIRequest.Get().(*AlibabaItEslEslinfoGeteslinfoAPIRequest)
}

// ReleaseAlibabaItEslEslinfoGeteslinfoAPIRequest 将 AlibabaItEslEslinfoGeteslinfoAPIRequest 放入 sync.Pool
func ReleaseAlibabaItEslEslinfoGeteslinfoAPIRequest(v *AlibabaItEslEslinfoGeteslinfoAPIRequest) {
	v.Reset()
	poolAlibabaItEslEslinfoGeteslinfoAPIRequest.Put(v)
}
