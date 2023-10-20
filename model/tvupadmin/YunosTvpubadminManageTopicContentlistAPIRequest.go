package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminManageTopicContentlistAPIRequest 查看专题内容列表 API请求
// yunos.tvpubadmin.manage.topic.contentlist
//
// 查看专题内容列表
type YunosTvpubadminManageTopicContentlistAPIRequest struct {
	model.Params
	// 节目查询参数
	_programQuery string
}

// NewYunosTvpubadminManageTopicContentlistRequest 初始化YunosTvpubadminManageTopicContentlistAPIRequest对象
func NewYunosTvpubadminManageTopicContentlistRequest() *YunosTvpubadminManageTopicContentlistAPIRequest {
	return &YunosTvpubadminManageTopicContentlistAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminManageTopicContentlistAPIRequest) Reset() {
	r._programQuery = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminManageTopicContentlistAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminManageTopicContentlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminManageTopicContentlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProgramQuery is ProgramQuery Setter
// 节目查询参数
func (r *YunosTvpubadminManageTopicContentlistAPIRequest) SetProgramQuery(_programQuery string) error {
	r._programQuery = _programQuery
	r.Set("program_query", _programQuery)
	return nil
}

// GetProgramQuery ProgramQuery Getter
func (r YunosTvpubadminManageTopicContentlistAPIRequest) GetProgramQuery() string {
	return r._programQuery
}

var poolYunosTvpubadminManageTopicContentlistAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminManageTopicContentlistRequest()
	},
}

// GetYunosTvpubadminManageTopicContentlistRequest 从 sync.Pool 获取 YunosTvpubadminManageTopicContentlistAPIRequest
func GetYunosTvpubadminManageTopicContentlistAPIRequest() *YunosTvpubadminManageTopicContentlistAPIRequest {
	return poolYunosTvpubadminManageTopicContentlistAPIRequest.Get().(*YunosTvpubadminManageTopicContentlistAPIRequest)
}

// ReleaseYunosTvpubadminManageTopicContentlistAPIRequest 将 YunosTvpubadminManageTopicContentlistAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminManageTopicContentlistAPIRequest(v *YunosTvpubadminManageTopicContentlistAPIRequest) {
	v.Reset()
	poolYunosTvpubadminManageTopicContentlistAPIRequest.Put(v)
}
