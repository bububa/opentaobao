package tmalltrend

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTrendStyleBasicinfoUploadAPIRequest 3D款式基本信息同步API API请求
// tmall.trend.style.basicinfo.upload
//
// 3D款式基本信息同步至天猫趋势中心
type TmallTrendStyleBasicinfoUploadAPIRequest struct {
	model.Params
	// 款式基本信息列表，单次同步最多1000条
	_styleBasicInfoBoList []StyleBasicInfoBo
}

// NewTmallTrendStyleBasicinfoUploadRequest 初始化TmallTrendStyleBasicinfoUploadAPIRequest对象
func NewTmallTrendStyleBasicinfoUploadRequest() *TmallTrendStyleBasicinfoUploadAPIRequest {
	return &TmallTrendStyleBasicinfoUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTrendStyleBasicinfoUploadAPIRequest) Reset() {
	r._styleBasicInfoBoList = r._styleBasicInfoBoList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTrendStyleBasicinfoUploadAPIRequest) GetApiMethodName() string {
	return "tmall.trend.style.basicinfo.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTrendStyleBasicinfoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTrendStyleBasicinfoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStyleBasicInfoBoList is StyleBasicInfoBoList Setter
// 款式基本信息列表，单次同步最多1000条
func (r *TmallTrendStyleBasicinfoUploadAPIRequest) SetStyleBasicInfoBoList(_styleBasicInfoBoList []StyleBasicInfoBo) error {
	r._styleBasicInfoBoList = _styleBasicInfoBoList
	r.Set("style_basic_info_bo_list", _styleBasicInfoBoList)
	return nil
}

// GetStyleBasicInfoBoList StyleBasicInfoBoList Getter
func (r TmallTrendStyleBasicinfoUploadAPIRequest) GetStyleBasicInfoBoList() []StyleBasicInfoBo {
	return r._styleBasicInfoBoList
}

var poolTmallTrendStyleBasicinfoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTrendStyleBasicinfoUploadRequest()
	},
}

// GetTmallTrendStyleBasicinfoUploadRequest 从 sync.Pool 获取 TmallTrendStyleBasicinfoUploadAPIRequest
func GetTmallTrendStyleBasicinfoUploadAPIRequest() *TmallTrendStyleBasicinfoUploadAPIRequest {
	return poolTmallTrendStyleBasicinfoUploadAPIRequest.Get().(*TmallTrendStyleBasicinfoUploadAPIRequest)
}

// ReleaseTmallTrendStyleBasicinfoUploadAPIRequest 将 TmallTrendStyleBasicinfoUploadAPIRequest 放入 sync.Pool
func ReleaseTmallTrendStyleBasicinfoUploadAPIRequest(v *TmallTrendStyleBasicinfoUploadAPIRequest) {
	v.Reset()
	poolTmallTrendStyleBasicinfoUploadAPIRequest.Put(v)
}
