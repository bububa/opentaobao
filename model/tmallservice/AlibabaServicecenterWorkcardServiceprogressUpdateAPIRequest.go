package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaservicecenterworkcardserviceprogressupdateAPIRequest 更新服务进度 API请求
// alibaba.servicecenter.workcard.serviceprogress.update
//
// 提供给外部合作服务商更新服务进度的接口
type AlibabaservicecenterworkcardserviceprogressupdateAPIRequest struct {
	model.Params
	// 图片列表
	_picUrlList []string
	// 扩展参数
	_extendInfo string
	// 请求节点的动作描述，唯一标识一个节点
	_action string
	// 真实服务商nick，仅限isv服务商调用时使用
	_realTpNick string
	// 工单id
	_workcardId int64
}

// NewAlibabaservicecenterworkcardserviceprogressupdateRequest 初始化AlibabaservicecenterworkcardserviceprogressupdateAPIRequest对象
func NewAlibabaservicecenterworkcardserviceprogressupdateRequest() *AlibabaservicecenterworkcardserviceprogressupdateAPIRequest {
	return &AlibabaservicecenterworkcardserviceprogressupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.workcard.serviceprogress.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicUrlList is PicUrlList Setter
// 图片列表
func (r *AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) SetPicUrlList(_picUrlList []string) error {
	r._picUrlList = _picUrlList
	r.Set("pic_url_list", _picUrlList)
	return nil
}

// GetPicUrlList PicUrlList Getter
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetPicUrlList() []string {
	return r._picUrlList
}

// SetExtendInfo is ExtendInfo Setter
// 扩展参数
func (r *AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) SetExtendInfo(_extendInfo string) error {
	r._extendInfo = _extendInfo
	r.Set("extend_info", _extendInfo)
	return nil
}

// GetExtendInfo ExtendInfo Getter
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetExtendInfo() string {
	return r._extendInfo
}

// SetAction is Action Setter
// 请求节点的动作描述，唯一标识一个节点
func (r *AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetAction() string {
	return r._action
}

// SetRealTpNick is RealTpNick Setter
// 真实服务商nick，仅限isv服务商调用时使用
func (r *AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) SetRealTpNick(_realTpNick string) error {
	r._realTpNick = _realTpNick
	r.Set("real_tp_nick", _realTpNick)
	return nil
}

// GetRealTpNick RealTpNick Getter
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetRealTpNick() string {
	return r._realTpNick
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r AlibabaservicecenterworkcardserviceprogressupdateAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}
