package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminmanagetopiccontentlistAPIRequest 查看专题内容列表 API请求
// yunos.tvpubadmin.manage.topic.contentlist
//
// 查看专题内容列表
type YunostvpubadminmanagetopiccontentlistAPIRequest struct {
	model.Params
	// 节目查询参数
	_programQuery string
}

// NewYunostvpubadminmanagetopiccontentlistRequest 初始化YunostvpubadminmanagetopiccontentlistAPIRequest对象
func NewYunostvpubadminmanagetopiccontentlistRequest() *YunostvpubadminmanagetopiccontentlistAPIRequest {
	return &YunostvpubadminmanagetopiccontentlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminmanagetopiccontentlistAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.manage.topic.contentlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminmanagetopiccontentlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminmanagetopiccontentlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProgramQuery is ProgramQuery Setter
// 节目查询参数
func (r *YunostvpubadminmanagetopiccontentlistAPIRequest) SetProgramQuery(_programQuery string) error {
	r._programQuery = _programQuery
	r.Set("program_query", _programQuery)
	return nil
}

// GetProgramQuery ProgramQuery Getter
func (r YunostvpubadminmanagetopiccontentlistAPIRequest) GetProgramQuery() string {
	return r._programQuery
}
