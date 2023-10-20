package tmc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTmcUserPermitAPIRequest 为已授权的用户开通消息服务 API请求
// taobao.tmc.user.permit
//
// 为已授权的用户开通消息服务，授权消息使用。&lt;br/&gt;&lt;span style=&#34;color:red&#34;&gt;注意：topic覆盖更新,务必传入全量topic，或者不传topics，使用appkey订阅的所有topic&lt;/span&gt;
type TaobaoTmcUserPermitAPIRequest struct {
	model.Params
	// 消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。
	_topics []string
}

// NewTaobaoTmcUserPermitRequest 初始化TaobaoTmcUserPermitAPIRequest对象
func NewTaobaoTmcUserPermitRequest() *TaobaoTmcUserPermitAPIRequest {
	return &TaobaoTmcUserPermitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTmcUserPermitAPIRequest) Reset() {
	r._topics = r._topics[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTmcUserPermitAPIRequest) GetApiMethodName() string {
	return "taobao.tmc.user.permit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTmcUserPermitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTmcUserPermitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTopics is Topics Setter
// 消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。
func (r *TaobaoTmcUserPermitAPIRequest) SetTopics(_topics []string) error {
	r._topics = _topics
	r.Set("topics", _topics)
	return nil
}

// GetTopics Topics Getter
func (r TaobaoTmcUserPermitAPIRequest) GetTopics() []string {
	return r._topics
}

var poolTaobaoTmcUserPermitAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTmcUserPermitRequest()
	},
}

// GetTaobaoTmcUserPermitRequest 从 sync.Pool 获取 TaobaoTmcUserPermitAPIRequest
func GetTaobaoTmcUserPermitAPIRequest() *TaobaoTmcUserPermitAPIRequest {
	return poolTaobaoTmcUserPermitAPIRequest.Get().(*TaobaoTmcUserPermitAPIRequest)
}

// ReleaseTaobaoTmcUserPermitAPIRequest 将 TaobaoTmcUserPermitAPIRequest 放入 sync.Pool
func ReleaseTaobaoTmcUserPermitAPIRequest(v *TaobaoTmcUserPermitAPIRequest) {
	v.Reset()
	poolTaobaoTmcUserPermitAPIRequest.Put(v)
}
