package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFcMallxInteractionAiPicListAPIRequest 花园ai作画定制查询 API请求
// alibaba.fc.mallx.interaction.ai.pic.list
//
// 花园ai作画定制查询
type AlibabaFcMallxInteractionAiPicListAPIRequest struct {
	model.Params
	// openid
	_openid string
	// openid
	_openId string
	// appKey
	_appId string
	// 商品信息
	_param *InteractiveTopItemParam
}

// NewAlibabaFcMallxInteractionAiPicListRequest 初始化AlibabaFcMallxInteractionAiPicListAPIRequest对象
func NewAlibabaFcMallxInteractionAiPicListRequest() *AlibabaFcMallxInteractionAiPicListAPIRequest {
	return &AlibabaFcMallxInteractionAiPicListAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFcMallxInteractionAiPicListAPIRequest) Reset() {
	r._openid = ""
	r._openId = ""
	r._appId = ""
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFcMallxInteractionAiPicListAPIRequest) GetApiMethodName() string {
	return "alibaba.fc.mallx.interaction.ai.pic.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFcMallxInteractionAiPicListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFcMallxInteractionAiPicListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOpenid is Openid Setter
// openid
func (r *AlibabaFcMallxInteractionAiPicListAPIRequest) SetOpenid(_openid string) error {
	r._openid = _openid
	r.Set("openid", _openid)
	return nil
}

// GetOpenid Openid Getter
func (r AlibabaFcMallxInteractionAiPicListAPIRequest) GetOpenid() string {
	return r._openid
}

// SetOpenId is OpenId Setter
// openid
func (r *AlibabaFcMallxInteractionAiPicListAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r AlibabaFcMallxInteractionAiPicListAPIRequest) GetOpenId() string {
	return r._openId
}

// SetAppId is AppId Setter
// appKey
func (r *AlibabaFcMallxInteractionAiPicListAPIRequest) SetAppId(_appId string) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r AlibabaFcMallxInteractionAiPicListAPIRequest) GetAppId() string {
	return r._appId
}

// SetParam is Param Setter
// 商品信息
func (r *AlibabaFcMallxInteractionAiPicListAPIRequest) SetParam(_param *InteractiveTopItemParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaFcMallxInteractionAiPicListAPIRequest) GetParam() *InteractiveTopItemParam {
	return r._param
}

var poolAlibabaFcMallxInteractionAiPicListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFcMallxInteractionAiPicListRequest()
	},
}

// GetAlibabaFcMallxInteractionAiPicListRequest 从 sync.Pool 获取 AlibabaFcMallxInteractionAiPicListAPIRequest
func GetAlibabaFcMallxInteractionAiPicListAPIRequest() *AlibabaFcMallxInteractionAiPicListAPIRequest {
	return poolAlibabaFcMallxInteractionAiPicListAPIRequest.Get().(*AlibabaFcMallxInteractionAiPicListAPIRequest)
}

// ReleaseAlibabaFcMallxInteractionAiPicListAPIRequest 将 AlibabaFcMallxInteractionAiPicListAPIRequest 放入 sync.Pool
func ReleaseAlibabaFcMallxInteractionAiPicListAPIRequest(v *AlibabaFcMallxInteractionAiPicListAPIRequest) {
	v.Reset()
	poolAlibabaFcMallxInteractionAiPicListAPIRequest.Put(v)
}
