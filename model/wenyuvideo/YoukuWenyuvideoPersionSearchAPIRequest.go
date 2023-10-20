package wenyuvideo

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuwenyuvideopersionsearchAPIRequest 根据人物名称查询人物列表 API请求
// youku.wenyuvideo.persion.search
//
// 根据人物名称查询人物列表
type YoukuwenyuvideopersionsearchAPIRequest struct {
	model.Params
	// 人物名字，搜索规则是完全匹配，即只返回同名人物列表
	_personName string
}

// NewYoukuwenyuvideopersionsearchRequest 初始化YoukuwenyuvideopersionsearchAPIRequest对象
func NewYoukuwenyuvideopersionsearchRequest() *YoukuwenyuvideopersionsearchAPIRequest {
	return &YoukuwenyuvideopersionsearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuwenyuvideopersionsearchAPIRequest) GetApiMethodName() string {
	return "youku.wenyuvideo.persion.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuwenyuvideopersionsearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuwenyuvideopersionsearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPersonName is PersonName Setter
// 人物名字，搜索规则是完全匹配，即只返回同名人物列表
func (r *YoukuwenyuvideopersionsearchAPIRequest) SetPersonName(_personName string) error {
	r._personName = _personName
	r.Set("person_name", _personName)
	return nil
}

// GetPersonName PersonName Getter
func (r YoukuwenyuvideopersionsearchAPIRequest) GetPersonName() string {
	return r._personName
}
