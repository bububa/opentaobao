package wenyuvideo

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YoukuWenyuvideoPersionSearchAPIRequest 根据人物名称查询人物列表 API请求
// youku.wenyuvideo.persion.search
//
// 根据人物名称查询人物列表
type YoukuWenyuvideoPersionSearchAPIRequest struct {
	model.Params
	// 人物名字，搜索规则是完全匹配，即只返回同名人物列表
	_personName string
}

// NewYoukuWenyuvideoPersionSearchRequest 初始化YoukuWenyuvideoPersionSearchAPIRequest对象
func NewYoukuWenyuvideoPersionSearchRequest() *YoukuWenyuvideoPersionSearchAPIRequest {
	return &YoukuWenyuvideoPersionSearchAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuWenyuvideoPersionSearchAPIRequest) Reset() {
	r._personName = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuWenyuvideoPersionSearchAPIRequest) GetApiMethodName() string {
	return "youku.wenyuvideo.persion.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuWenyuvideoPersionSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuWenyuvideoPersionSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPersonName is PersonName Setter
// 人物名字，搜索规则是完全匹配，即只返回同名人物列表
func (r *YoukuWenyuvideoPersionSearchAPIRequest) SetPersonName(_personName string) error {
	r._personName = _personName
	r.Set("person_name", _personName)
	return nil
}

// GetPersonName PersonName Getter
func (r YoukuWenyuvideoPersionSearchAPIRequest) GetPersonName() string {
	return r._personName
}

var poolYoukuWenyuvideoPersionSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuWenyuvideoPersionSearchRequest()
	},
}

// GetYoukuWenyuvideoPersionSearchRequest 从 sync.Pool 获取 YoukuWenyuvideoPersionSearchAPIRequest
func GetYoukuWenyuvideoPersionSearchAPIRequest() *YoukuWenyuvideoPersionSearchAPIRequest {
	return poolYoukuWenyuvideoPersionSearchAPIRequest.Get().(*YoukuWenyuvideoPersionSearchAPIRequest)
}

// ReleaseYoukuWenyuvideoPersionSearchAPIRequest 将 YoukuWenyuvideoPersionSearchAPIRequest 放入 sync.Pool
func ReleaseYoukuWenyuvideoPersionSearchAPIRequest(v *YoukuWenyuvideoPersionSearchAPIRequest) {
	v.Reset()
	poolYoukuWenyuvideoPersionSearchAPIRequest.Put(v)
}
