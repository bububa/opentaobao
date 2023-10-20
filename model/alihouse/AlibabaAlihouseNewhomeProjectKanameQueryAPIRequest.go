package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest 查询KA楼盘名称 API请求
// alibaba.alihouse.newhome.project.kaname.query
//
// 查询KA楼盘名称
type AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest struct {
	model.Params
	// KA楼盘ID
	_kaProjectMid int64
}

// NewAlibabaAlihouseNewhomeProjectKanameQueryRequest 初始化AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectKanameQueryRequest() *AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest {
	return &AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) Reset() {
	r._kaProjectMid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.kaname.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKaProjectMid is KaProjectMid Setter
// KA楼盘ID
func (r *AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) SetKaProjectMid(_kaProjectMid int64) error {
	r._kaProjectMid = _kaProjectMid
	r.Set("ka_project_mid", _kaProjectMid)
	return nil
}

// GetKaProjectMid KaProjectMid Getter
func (r AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) GetKaProjectMid() int64 {
	return r._kaProjectMid
}

var poolAlibabaAlihouseNewhomeProjectKanameQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectKanameQueryRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectKanameQueryRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest
func GetAlibabaAlihouseNewhomeProjectKanameQueryAPIRequest() *AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectKanameQueryAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectKanameQueryAPIRequest 将 AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectKanameQueryAPIRequest(v *AlibabaAlihouseNewhomeProjectKanameQueryAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectKanameQueryAPIRequest.Put(v)
}
