package scs

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScsImageMatteAPIRequest 阿里妈妈智能创意平台在线抠图 API请求
// alibaba.scs.image.matte
//
// 该API对外输出一个在线抠图(Deep Image Matting)接口，合作方可以通过该接口利用深度学习抠图算法从图片中抠出目标对象(比如商品或者人物轮廓)
type AlibabaScsImageMatteAPIRequest struct {
	model.Params
	// 资源位ID，接入前由智能创意平台分配
	_pid string
	// 服务名称，可选值: scs
	_name string
	// 场景名称，可选值: image_cutout
	_scenes string
	// 抠图上下文信息，json字符串格式，json中matting_type字段可选值: external_matting，url: 需要抠图的目标图片url
	_obExt string
	// 32位uuid
	_sessionid string
	// 当前秒级时间戳
	_ts string
}

// NewAlibabaScsImageMatteRequest 初始化AlibabaScsImageMatteAPIRequest对象
func NewAlibabaScsImageMatteRequest() *AlibabaScsImageMatteAPIRequest {
	return &AlibabaScsImageMatteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaScsImageMatteAPIRequest) GetApiMethodName() string {
	return "alibaba.scs.image.matte"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaScsImageMatteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPid is Pid Setter
// 资源位ID，接入前由智能创意平台分配
func (r *AlibabaScsImageMatteAPIRequest) SetPid(_pid string) error {
	r._pid = _pid
	r.Set("pid", _pid)
	return nil
}

// GetPid Pid Getter
func (r AlibabaScsImageMatteAPIRequest) GetPid() string {
	return r._pid
}

// SetName is Name Setter
// 服务名称，可选值: scs
func (r *AlibabaScsImageMatteAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaScsImageMatteAPIRequest) GetName() string {
	return r._name
}

// SetScenes is Scenes Setter
// 场景名称，可选值: image_cutout
func (r *AlibabaScsImageMatteAPIRequest) SetScenes(_scenes string) error {
	r._scenes = _scenes
	r.Set("scenes", _scenes)
	return nil
}

// GetScenes Scenes Getter
func (r AlibabaScsImageMatteAPIRequest) GetScenes() string {
	return r._scenes
}

// SetObExt is ObExt Setter
// 抠图上下文信息，json字符串格式，json中matting_type字段可选值: external_matting，url: 需要抠图的目标图片url
func (r *AlibabaScsImageMatteAPIRequest) SetObExt(_obExt string) error {
	r._obExt = _obExt
	r.Set("ob_ext", _obExt)
	return nil
}

// GetObExt ObExt Getter
func (r AlibabaScsImageMatteAPIRequest) GetObExt() string {
	return r._obExt
}

// SetSessionid is Sessionid Setter
// 32位uuid
func (r *AlibabaScsImageMatteAPIRequest) SetSessionid(_sessionid string) error {
	r._sessionid = _sessionid
	r.Set("sessionid", _sessionid)
	return nil
}

// GetSessionid Sessionid Getter
func (r AlibabaScsImageMatteAPIRequest) GetSessionid() string {
	return r._sessionid
}

// SetTs is Ts Setter
// 当前秒级时间戳
func (r *AlibabaScsImageMatteAPIRequest) SetTs(_ts string) error {
	r._ts = _ts
	r.Set("ts", _ts)
	return nil
}

// GetTs Ts Getter
func (r AlibabaScsImageMatteAPIRequest) GetTs() string {
	return r._ts
}
