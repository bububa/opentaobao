package einvoice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaeinvoiceserialnobatchgenerateAPIRequest 开票流水号批量生成接口 API请求
// alibaba.einvoice.serialno.batch.generate
//
// 批量获取开票流水号接口。此接口1次返回1000条开票流水号，每个应用每天限流1000次调用。
// 优先使用alibaba.einvoice.serial.generate。
type AlibabaeinvoiceserialnobatchgenerateAPIRequest struct {
	model.Params
}

// NewAlibabaeinvoiceserialnobatchgenerateRequest 初始化AlibabaeinvoiceserialnobatchgenerateAPIRequest对象
func NewAlibabaeinvoiceserialnobatchgenerateRequest() *AlibabaeinvoiceserialnobatchgenerateAPIRequest {
	return &AlibabaeinvoiceserialnobatchgenerateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaeinvoiceserialnobatchgenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.einvoice.serialno.batch.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaeinvoiceserialnobatchgenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaeinvoiceserialnobatchgenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}
