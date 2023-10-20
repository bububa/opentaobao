package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriphotelhstdfshotelmatchsroomselfAPIRequest 匹配标准房型以及卖家房型 API请求
// alitrip.hotel.hstdf.shotel.matchsroomself
//
// 匹配卖家房型以及标准房型，返回匹配结果
type AlitriphotelhstdfshotelmatchsroomselfAPIRequest struct {
	model.Params
	// SroomTypeMatchParam
	_param0 *SroomTypeMatchParam
}

// NewAlitriphotelhstdfshotelmatchsroomselfRequest 初始化AlitriphotelhstdfshotelmatchsroomselfAPIRequest对象
func NewAlitriphotelhstdfshotelmatchsroomselfRequest() *AlitriphotelhstdfshotelmatchsroomselfAPIRequest {
	return &AlitriphotelhstdfshotelmatchsroomselfAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriphotelhstdfshotelmatchsroomselfAPIRequest) GetApiMethodName() string {
	return "alitrip.hotel.hstdf.shotel.matchsroomself"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriphotelhstdfshotelmatchsroomselfAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriphotelhstdfshotelmatchsroomselfAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// SroomTypeMatchParam
func (r *AlitriphotelhstdfshotelmatchsroomselfAPIRequest) SetParam0(_param0 *SroomTypeMatchParam) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlitriphotelhstdfshotelmatchsroomselfAPIRequest) GetParam0() *SroomTypeMatchParam {
	return r._param0
}
